/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/source"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/claimbinding"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/claimdefaulting"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/claimscheduling"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	cachev1alpha1 "github.com/crossplane/crossplane/apis/cache/v1alpha1"

	"github.com/crossplane/provider-aws/apis/cache/v1beta1"
	aws "github.com/crossplane/provider-aws/pkg/clients"
)

// SetupReplicationGroupClaimScheduling adds a controller that reconciles
// RedisCluster claims that include a class selector but omit their class and
// resource references by picking a random matching ReplicationGroupClass, if
// any.
func SetupReplicationGroupClaimScheduling(mgr ctrl.Manager, l logging.Logger) error {
	name := claimscheduling.ControllerName(cachev1alpha1.RedisClusterGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&cachev1alpha1.RedisCluster{}).
		WithEventFilter(resource.NewPredicates(resource.AllOf(
			resource.HasClassSelector(),
			resource.HasNoClassReference(),
			resource.HasNoManagedResourceReference(),
		))).
		Complete(claimscheduling.NewReconciler(mgr,
			resource.ClaimKind(cachev1alpha1.RedisClusterGroupVersionKind),
			resource.ClassKind(v1beta1.ReplicationGroupClassGroupVersionKind),
			claimscheduling.WithLogger(l.WithValues("controller", name)),
			claimscheduling.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		))
}

// SetupReplicationGroupClaimDefaulting adds a controller that reconciles
// RedisCluster claims that omit their resource ref, class ref, and class
// selector by choosing a default ReplicationGroupClass if one exists.
func SetupReplicationGroupClaimDefaulting(mgr ctrl.Manager, l logging.Logger) error {
	name := claimdefaulting.ControllerName(cachev1alpha1.RedisClusterGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&cachev1alpha1.RedisCluster{}).
		WithEventFilter(resource.NewPredicates(resource.AllOf(
			resource.HasNoClassSelector(),
			resource.HasNoClassReference(),
			resource.HasNoManagedResourceReference(),
		))).
		Complete(claimdefaulting.NewReconciler(mgr,
			resource.ClaimKind(cachev1alpha1.RedisClusterGroupVersionKind),
			resource.ClassKind(v1beta1.ReplicationGroupClassGroupVersionKind),
			claimdefaulting.WithLogger(l.WithValues("controller", name)),
			claimdefaulting.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		))
}

// SetupReplicationGroupClaimBinding adds a controller that reconciles
// RedisCluster claims with ReplicationGroups, dynamically provisioning them if
// needed.
func SetupReplicationGroupClaimBinding(mgr ctrl.Manager, l logging.Logger) error {
	name := claimbinding.ControllerName(cachev1alpha1.RedisClusterGroupKind)

	r := claimbinding.NewReconciler(mgr,
		resource.ClaimKind(cachev1alpha1.RedisClusterGroupVersionKind),
		resource.ClassKind(v1beta1.ReplicationGroupClassGroupVersionKind),
		resource.ManagedKind(v1beta1.ReplicationGroupGroupVersionKind),
		claimbinding.WithManagedConfigurators(
			claimbinding.ManagedConfiguratorFn(ConfigureReplicationGroup),
			claimbinding.ManagedConfiguratorFn(claimbinding.ConfigureReclaimPolicy),
			claimbinding.ManagedConfiguratorFn(claimbinding.ConfigureNames)),
		claimbinding.WithLogger(l.WithValues("controller", name)),
		claimbinding.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
	)

	p := resource.NewPredicates(resource.AnyOf(
		resource.HasClassReferenceKind(resource.ClassKind(v1beta1.ReplicationGroupClassGroupVersionKind)),
		resource.HasManagedResourceReferenceKind(resource.ManagedKind(v1beta1.ReplicationGroupGroupVersionKind)),
		resource.IsManagedKind(resource.ManagedKind(v1beta1.ReplicationGroupGroupVersionKind), mgr.GetScheme()),
	))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		Watches(&source.Kind{Type: &v1beta1.ReplicationGroup{}}, &resource.EnqueueRequestForClaim{}).
		For(&cachev1alpha1.RedisCluster{}).
		WithEventFilter(p).
		Complete(r)
}

// ConfigureReplicationGroup configures the supplied resource (presumed to be a
// ReplicationGroup) using the supplied resource claim (presumed to be a
// RedisCluster) and resource class.
func ConfigureReplicationGroup(_ context.Context, cm resource.Claim, cs resource.Class, mg resource.Managed) error {
	rc, cmok := cm.(*cachev1alpha1.RedisCluster)
	if !cmok {
		return errors.Errorf("expected resource claim %s to be %s", cm.GetName(), cachev1alpha1.RedisClusterGroupVersionKind)
	}

	rgc, csok := cs.(*v1beta1.ReplicationGroupClass)
	if !csok {
		return errors.Errorf("expected resource class %s to be %s", cs.GetName(), v1beta1.ReplicationGroupClassGroupVersionKind)
	}

	rg, mgok := mg.(*v1beta1.ReplicationGroup)
	if !mgok {
		return errors.Errorf("expected managed resource %s to be %s", mg.GetName(), v1beta1.ReplicationGroupGroupVersionKind)
	}

	spec := &v1beta1.ReplicationGroupSpec{
		ResourceSpec: runtimev1alpha1.ResourceSpec{
			ReclaimPolicy: runtimev1alpha1.ReclaimRetain,
		},
		ForProvider: rgc.SpecTemplate.ForProvider,
	}

	if err := resolveAWSClassInstanceValues(&spec.ForProvider, rc); err != nil {
		return errors.Wrap(err, "cannot resolve AWS class instance values")
	}

	spec.WriteConnectionSecretToReference = &runtimev1alpha1.SecretReference{
		Namespace: rgc.SpecTemplate.WriteConnectionSecretsToNamespace,
		Name:      string(cm.GetUID()),
	}
	spec.ReclaimPolicy = rgc.SpecTemplate.ReclaimPolicy
	rg.Spec = *spec

	return nil
}

func resolveAWSClassInstanceValues(spec *v1beta1.ReplicationGroupParameters, rc *cachev1alpha1.RedisCluster) error {
	var err error
	switch {
	case aws.StringValue(spec.EngineVersion) == "" && rc.Spec.EngineVersion == "":
	// Neither the claim nor its class specified a version. Let AWS pick.

	case aws.StringValue(spec.EngineVersion) == "" && rc.Spec.EngineVersion != "":
		// Only the claim specified a version. Use the latest supported patch
		// version for said claim (minor) version.
		spec.EngineVersion, err = latestSupportedPatchVersion(rc.Spec.EngineVersion)

	case aws.StringValue(spec.EngineVersion) != "" && rc.Spec.EngineVersion == "":
		// Only the class specified a version. Use it.

	case !strings.HasPrefix(aws.StringValue(spec.EngineVersion), rc.Spec.EngineVersion+"."):
		// Both the claim and its class specified a version, but the class
		// version is not a patch of the claim version.
		err = errors.Errorf("class version %s is not a patch of claim version %s", aws.StringValue(spec.EngineVersion), rc.Spec.EngineVersion)

	default:
		// Both the claim and its class specified a version, and the class
		// version is a patch of the claim version. Use the class version.
	}

	return errors.Wrap(err, "cannot resolve class claim values")
}

func latestSupportedPatchVersion(minorVersion string) (*string, error) {
	p := v1beta1.LatestSupportedPatchVersion[v1beta1.MinorVersion(minorVersion)]
	if p == v1beta1.UnsupportedVersion {
		return nil, errors.Errorf("minor version %s is not currently supported", minorVersion)
	}
	s := string(p)
	return &s, nil
}
