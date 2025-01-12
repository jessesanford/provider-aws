/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package userpoolclient

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/cognitoidentityprovider/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeUserPoolClientInput returns input for read
// operation.
func GenerateDescribeUserPoolClientInput(cr *svcapitypes.UserPoolClient) *svcsdk.DescribeUserPoolClientInput {
	res := &svcsdk.DescribeUserPoolClientInput{}

	if cr.Status.AtProvider.ClientID != nil {
		res.SetClientId(*cr.Status.AtProvider.ClientID)
	}
	if cr.Status.AtProvider.UserPoolID != nil {
		res.SetUserPoolId(*cr.Status.AtProvider.UserPoolID)
	}

	return res
}

// GenerateUserPoolClient returns the current state in the form of *svcapitypes.UserPoolClient.
func GenerateUserPoolClient(resp *svcsdk.DescribeUserPoolClientOutput) *svcapitypes.UserPoolClient {
	cr := &svcapitypes.UserPoolClient{}

	if resp.UserPoolClient.AccessTokenValidity != nil {
		cr.Spec.ForProvider.AccessTokenValidity = resp.UserPoolClient.AccessTokenValidity
	} else {
		cr.Spec.ForProvider.AccessTokenValidity = nil
	}
	if resp.UserPoolClient.AllowedOAuthFlows != nil {
		f1 := []*string{}
		for _, f1iter := range resp.UserPoolClient.AllowedOAuthFlows {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		cr.Spec.ForProvider.AllowedOAuthFlows = f1
	} else {
		cr.Spec.ForProvider.AllowedOAuthFlows = nil
	}
	if resp.UserPoolClient.AllowedOAuthFlowsUserPoolClient != nil {
		cr.Spec.ForProvider.AllowedOAuthFlowsUserPoolClient = resp.UserPoolClient.AllowedOAuthFlowsUserPoolClient
	} else {
		cr.Spec.ForProvider.AllowedOAuthFlowsUserPoolClient = nil
	}
	if resp.UserPoolClient.AllowedOAuthScopes != nil {
		f3 := []*string{}
		for _, f3iter := range resp.UserPoolClient.AllowedOAuthScopes {
			var f3elem string
			f3elem = *f3iter
			f3 = append(f3, &f3elem)
		}
		cr.Spec.ForProvider.AllowedOAuthScopes = f3
	} else {
		cr.Spec.ForProvider.AllowedOAuthScopes = nil
	}
	if resp.UserPoolClient.AnalyticsConfiguration != nil {
		f4 := &svcapitypes.AnalyticsConfigurationType{}
		if resp.UserPoolClient.AnalyticsConfiguration.ApplicationArn != nil {
			f4.ApplicationARN = resp.UserPoolClient.AnalyticsConfiguration.ApplicationArn
		}
		if resp.UserPoolClient.AnalyticsConfiguration.ApplicationId != nil {
			f4.ApplicationID = resp.UserPoolClient.AnalyticsConfiguration.ApplicationId
		}
		if resp.UserPoolClient.AnalyticsConfiguration.ExternalId != nil {
			f4.ExternalID = resp.UserPoolClient.AnalyticsConfiguration.ExternalId
		}
		if resp.UserPoolClient.AnalyticsConfiguration.RoleArn != nil {
			f4.RoleARN = resp.UserPoolClient.AnalyticsConfiguration.RoleArn
		}
		if resp.UserPoolClient.AnalyticsConfiguration.UserDataShared != nil {
			f4.UserDataShared = resp.UserPoolClient.AnalyticsConfiguration.UserDataShared
		}
		cr.Spec.ForProvider.AnalyticsConfiguration = f4
	} else {
		cr.Spec.ForProvider.AnalyticsConfiguration = nil
	}
	if resp.UserPoolClient.CallbackURLs != nil {
		f5 := []*string{}
		for _, f5iter := range resp.UserPoolClient.CallbackURLs {
			var f5elem string
			f5elem = *f5iter
			f5 = append(f5, &f5elem)
		}
		cr.Spec.ForProvider.CallbackURLs = f5
	} else {
		cr.Spec.ForProvider.CallbackURLs = nil
	}
	if resp.UserPoolClient.ClientId != nil {
		cr.Status.AtProvider.ClientID = resp.UserPoolClient.ClientId
	} else {
		cr.Status.AtProvider.ClientID = nil
	}
	if resp.UserPoolClient.ClientName != nil {
		cr.Spec.ForProvider.ClientName = resp.UserPoolClient.ClientName
	} else {
		cr.Spec.ForProvider.ClientName = nil
	}
	if resp.UserPoolClient.ClientSecret != nil {
		cr.Status.AtProvider.ClientSecret = resp.UserPoolClient.ClientSecret
	} else {
		cr.Status.AtProvider.ClientSecret = nil
	}
	if resp.UserPoolClient.CreationDate != nil {
		cr.Status.AtProvider.CreationDate = &metav1.Time{*resp.UserPoolClient.CreationDate}
	} else {
		cr.Status.AtProvider.CreationDate = nil
	}
	if resp.UserPoolClient.DefaultRedirectURI != nil {
		cr.Spec.ForProvider.DefaultRedirectURI = resp.UserPoolClient.DefaultRedirectURI
	} else {
		cr.Spec.ForProvider.DefaultRedirectURI = nil
	}
	if resp.UserPoolClient.EnableTokenRevocation != nil {
		cr.Spec.ForProvider.EnableTokenRevocation = resp.UserPoolClient.EnableTokenRevocation
	} else {
		cr.Spec.ForProvider.EnableTokenRevocation = nil
	}
	if resp.UserPoolClient.ExplicitAuthFlows != nil {
		f12 := []*string{}
		for _, f12iter := range resp.UserPoolClient.ExplicitAuthFlows {
			var f12elem string
			f12elem = *f12iter
			f12 = append(f12, &f12elem)
		}
		cr.Spec.ForProvider.ExplicitAuthFlows = f12
	} else {
		cr.Spec.ForProvider.ExplicitAuthFlows = nil
	}
	if resp.UserPoolClient.IdTokenValidity != nil {
		cr.Spec.ForProvider.IDTokenValidity = resp.UserPoolClient.IdTokenValidity
	} else {
		cr.Spec.ForProvider.IDTokenValidity = nil
	}
	if resp.UserPoolClient.LastModifiedDate != nil {
		cr.Status.AtProvider.LastModifiedDate = &metav1.Time{*resp.UserPoolClient.LastModifiedDate}
	} else {
		cr.Status.AtProvider.LastModifiedDate = nil
	}
	if resp.UserPoolClient.LogoutURLs != nil {
		f15 := []*string{}
		for _, f15iter := range resp.UserPoolClient.LogoutURLs {
			var f15elem string
			f15elem = *f15iter
			f15 = append(f15, &f15elem)
		}
		cr.Spec.ForProvider.LogoutURLs = f15
	} else {
		cr.Spec.ForProvider.LogoutURLs = nil
	}
	if resp.UserPoolClient.PreventUserExistenceErrors != nil {
		cr.Spec.ForProvider.PreventUserExistenceErrors = resp.UserPoolClient.PreventUserExistenceErrors
	} else {
		cr.Spec.ForProvider.PreventUserExistenceErrors = nil
	}
	if resp.UserPoolClient.ReadAttributes != nil {
		f17 := []*string{}
		for _, f17iter := range resp.UserPoolClient.ReadAttributes {
			var f17elem string
			f17elem = *f17iter
			f17 = append(f17, &f17elem)
		}
		cr.Spec.ForProvider.ReadAttributes = f17
	} else {
		cr.Spec.ForProvider.ReadAttributes = nil
	}
	if resp.UserPoolClient.RefreshTokenValidity != nil {
		cr.Spec.ForProvider.RefreshTokenValidity = resp.UserPoolClient.RefreshTokenValidity
	} else {
		cr.Spec.ForProvider.RefreshTokenValidity = nil
	}
	if resp.UserPoolClient.SupportedIdentityProviders != nil {
		f19 := []*string{}
		for _, f19iter := range resp.UserPoolClient.SupportedIdentityProviders {
			var f19elem string
			f19elem = *f19iter
			f19 = append(f19, &f19elem)
		}
		cr.Spec.ForProvider.SupportedIdentityProviders = f19
	} else {
		cr.Spec.ForProvider.SupportedIdentityProviders = nil
	}
	if resp.UserPoolClient.TokenValidityUnits != nil {
		f20 := &svcapitypes.TokenValidityUnitsType{}
		if resp.UserPoolClient.TokenValidityUnits.AccessToken != nil {
			f20.AccessToken = resp.UserPoolClient.TokenValidityUnits.AccessToken
		}
		if resp.UserPoolClient.TokenValidityUnits.IdToken != nil {
			f20.IDToken = resp.UserPoolClient.TokenValidityUnits.IdToken
		}
		if resp.UserPoolClient.TokenValidityUnits.RefreshToken != nil {
			f20.RefreshToken = resp.UserPoolClient.TokenValidityUnits.RefreshToken
		}
		cr.Spec.ForProvider.TokenValidityUnits = f20
	} else {
		cr.Spec.ForProvider.TokenValidityUnits = nil
	}
	if resp.UserPoolClient.UserPoolId != nil {
		cr.Status.AtProvider.UserPoolID = resp.UserPoolClient.UserPoolId
	} else {
		cr.Status.AtProvider.UserPoolID = nil
	}
	if resp.UserPoolClient.WriteAttributes != nil {
		f22 := []*string{}
		for _, f22iter := range resp.UserPoolClient.WriteAttributes {
			var f22elem string
			f22elem = *f22iter
			f22 = append(f22, &f22elem)
		}
		cr.Spec.ForProvider.WriteAttributes = f22
	} else {
		cr.Spec.ForProvider.WriteAttributes = nil
	}

	return cr
}

// GenerateCreateUserPoolClientInput returns a create input.
func GenerateCreateUserPoolClientInput(cr *svcapitypes.UserPoolClient) *svcsdk.CreateUserPoolClientInput {
	res := &svcsdk.CreateUserPoolClientInput{}

	if cr.Spec.ForProvider.AccessTokenValidity != nil {
		res.SetAccessTokenValidity(*cr.Spec.ForProvider.AccessTokenValidity)
	}
	if cr.Spec.ForProvider.AllowedOAuthFlows != nil {
		f1 := []*string{}
		for _, f1iter := range cr.Spec.ForProvider.AllowedOAuthFlows {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetAllowedOAuthFlows(f1)
	}
	if cr.Spec.ForProvider.AllowedOAuthFlowsUserPoolClient != nil {
		res.SetAllowedOAuthFlowsUserPoolClient(*cr.Spec.ForProvider.AllowedOAuthFlowsUserPoolClient)
	}
	if cr.Spec.ForProvider.AllowedOAuthScopes != nil {
		f3 := []*string{}
		for _, f3iter := range cr.Spec.ForProvider.AllowedOAuthScopes {
			var f3elem string
			f3elem = *f3iter
			f3 = append(f3, &f3elem)
		}
		res.SetAllowedOAuthScopes(f3)
	}
	if cr.Spec.ForProvider.AnalyticsConfiguration != nil {
		f4 := &svcsdk.AnalyticsConfigurationType{}
		if cr.Spec.ForProvider.AnalyticsConfiguration.ApplicationARN != nil {
			f4.SetApplicationArn(*cr.Spec.ForProvider.AnalyticsConfiguration.ApplicationARN)
		}
		if cr.Spec.ForProvider.AnalyticsConfiguration.ApplicationID != nil {
			f4.SetApplicationId(*cr.Spec.ForProvider.AnalyticsConfiguration.ApplicationID)
		}
		if cr.Spec.ForProvider.AnalyticsConfiguration.ExternalID != nil {
			f4.SetExternalId(*cr.Spec.ForProvider.AnalyticsConfiguration.ExternalID)
		}
		if cr.Spec.ForProvider.AnalyticsConfiguration.RoleARN != nil {
			f4.SetRoleArn(*cr.Spec.ForProvider.AnalyticsConfiguration.RoleARN)
		}
		if cr.Spec.ForProvider.AnalyticsConfiguration.UserDataShared != nil {
			f4.SetUserDataShared(*cr.Spec.ForProvider.AnalyticsConfiguration.UserDataShared)
		}
		res.SetAnalyticsConfiguration(f4)
	}
	if cr.Spec.ForProvider.CallbackURLs != nil {
		f5 := []*string{}
		for _, f5iter := range cr.Spec.ForProvider.CallbackURLs {
			var f5elem string
			f5elem = *f5iter
			f5 = append(f5, &f5elem)
		}
		res.SetCallbackURLs(f5)
	}
	if cr.Spec.ForProvider.ClientName != nil {
		res.SetClientName(*cr.Spec.ForProvider.ClientName)
	}
	if cr.Spec.ForProvider.DefaultRedirectURI != nil {
		res.SetDefaultRedirectURI(*cr.Spec.ForProvider.DefaultRedirectURI)
	}
	if cr.Spec.ForProvider.EnableTokenRevocation != nil {
		res.SetEnableTokenRevocation(*cr.Spec.ForProvider.EnableTokenRevocation)
	}
	if cr.Spec.ForProvider.ExplicitAuthFlows != nil {
		f9 := []*string{}
		for _, f9iter := range cr.Spec.ForProvider.ExplicitAuthFlows {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		res.SetExplicitAuthFlows(f9)
	}
	if cr.Spec.ForProvider.GenerateSecret != nil {
		res.SetGenerateSecret(*cr.Spec.ForProvider.GenerateSecret)
	}
	if cr.Spec.ForProvider.IDTokenValidity != nil {
		res.SetIdTokenValidity(*cr.Spec.ForProvider.IDTokenValidity)
	}
	if cr.Spec.ForProvider.LogoutURLs != nil {
		f12 := []*string{}
		for _, f12iter := range cr.Spec.ForProvider.LogoutURLs {
			var f12elem string
			f12elem = *f12iter
			f12 = append(f12, &f12elem)
		}
		res.SetLogoutURLs(f12)
	}
	if cr.Spec.ForProvider.PreventUserExistenceErrors != nil {
		res.SetPreventUserExistenceErrors(*cr.Spec.ForProvider.PreventUserExistenceErrors)
	}
	if cr.Spec.ForProvider.ReadAttributes != nil {
		f14 := []*string{}
		for _, f14iter := range cr.Spec.ForProvider.ReadAttributes {
			var f14elem string
			f14elem = *f14iter
			f14 = append(f14, &f14elem)
		}
		res.SetReadAttributes(f14)
	}
	if cr.Spec.ForProvider.RefreshTokenValidity != nil {
		res.SetRefreshTokenValidity(*cr.Spec.ForProvider.RefreshTokenValidity)
	}
	if cr.Spec.ForProvider.SupportedIdentityProviders != nil {
		f16 := []*string{}
		for _, f16iter := range cr.Spec.ForProvider.SupportedIdentityProviders {
			var f16elem string
			f16elem = *f16iter
			f16 = append(f16, &f16elem)
		}
		res.SetSupportedIdentityProviders(f16)
	}
	if cr.Spec.ForProvider.TokenValidityUnits != nil {
		f17 := &svcsdk.TokenValidityUnitsType{}
		if cr.Spec.ForProvider.TokenValidityUnits.AccessToken != nil {
			f17.SetAccessToken(*cr.Spec.ForProvider.TokenValidityUnits.AccessToken)
		}
		if cr.Spec.ForProvider.TokenValidityUnits.IDToken != nil {
			f17.SetIdToken(*cr.Spec.ForProvider.TokenValidityUnits.IDToken)
		}
		if cr.Spec.ForProvider.TokenValidityUnits.RefreshToken != nil {
			f17.SetRefreshToken(*cr.Spec.ForProvider.TokenValidityUnits.RefreshToken)
		}
		res.SetTokenValidityUnits(f17)
	}
	if cr.Spec.ForProvider.WriteAttributes != nil {
		f18 := []*string{}
		for _, f18iter := range cr.Spec.ForProvider.WriteAttributes {
			var f18elem string
			f18elem = *f18iter
			f18 = append(f18, &f18elem)
		}
		res.SetWriteAttributes(f18)
	}

	return res
}

// GenerateUpdateUserPoolClientInput returns an update input.
func GenerateUpdateUserPoolClientInput(cr *svcapitypes.UserPoolClient) *svcsdk.UpdateUserPoolClientInput {
	res := &svcsdk.UpdateUserPoolClientInput{}

	if cr.Spec.ForProvider.AccessTokenValidity != nil {
		res.SetAccessTokenValidity(*cr.Spec.ForProvider.AccessTokenValidity)
	}
	if cr.Spec.ForProvider.AllowedOAuthFlows != nil {
		f1 := []*string{}
		for _, f1iter := range cr.Spec.ForProvider.AllowedOAuthFlows {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetAllowedOAuthFlows(f1)
	}
	if cr.Spec.ForProvider.AllowedOAuthFlowsUserPoolClient != nil {
		res.SetAllowedOAuthFlowsUserPoolClient(*cr.Spec.ForProvider.AllowedOAuthFlowsUserPoolClient)
	}
	if cr.Spec.ForProvider.AllowedOAuthScopes != nil {
		f3 := []*string{}
		for _, f3iter := range cr.Spec.ForProvider.AllowedOAuthScopes {
			var f3elem string
			f3elem = *f3iter
			f3 = append(f3, &f3elem)
		}
		res.SetAllowedOAuthScopes(f3)
	}
	if cr.Spec.ForProvider.AnalyticsConfiguration != nil {
		f4 := &svcsdk.AnalyticsConfigurationType{}
		if cr.Spec.ForProvider.AnalyticsConfiguration.ApplicationARN != nil {
			f4.SetApplicationArn(*cr.Spec.ForProvider.AnalyticsConfiguration.ApplicationARN)
		}
		if cr.Spec.ForProvider.AnalyticsConfiguration.ApplicationID != nil {
			f4.SetApplicationId(*cr.Spec.ForProvider.AnalyticsConfiguration.ApplicationID)
		}
		if cr.Spec.ForProvider.AnalyticsConfiguration.ExternalID != nil {
			f4.SetExternalId(*cr.Spec.ForProvider.AnalyticsConfiguration.ExternalID)
		}
		if cr.Spec.ForProvider.AnalyticsConfiguration.RoleARN != nil {
			f4.SetRoleArn(*cr.Spec.ForProvider.AnalyticsConfiguration.RoleARN)
		}
		if cr.Spec.ForProvider.AnalyticsConfiguration.UserDataShared != nil {
			f4.SetUserDataShared(*cr.Spec.ForProvider.AnalyticsConfiguration.UserDataShared)
		}
		res.SetAnalyticsConfiguration(f4)
	}
	if cr.Spec.ForProvider.CallbackURLs != nil {
		f5 := []*string{}
		for _, f5iter := range cr.Spec.ForProvider.CallbackURLs {
			var f5elem string
			f5elem = *f5iter
			f5 = append(f5, &f5elem)
		}
		res.SetCallbackURLs(f5)
	}
	if cr.Status.AtProvider.ClientID != nil {
		res.SetClientId(*cr.Status.AtProvider.ClientID)
	}
	if cr.Spec.ForProvider.ClientName != nil {
		res.SetClientName(*cr.Spec.ForProvider.ClientName)
	}
	if cr.Spec.ForProvider.DefaultRedirectURI != nil {
		res.SetDefaultRedirectURI(*cr.Spec.ForProvider.DefaultRedirectURI)
	}
	if cr.Spec.ForProvider.EnableTokenRevocation != nil {
		res.SetEnableTokenRevocation(*cr.Spec.ForProvider.EnableTokenRevocation)
	}
	if cr.Spec.ForProvider.ExplicitAuthFlows != nil {
		f10 := []*string{}
		for _, f10iter := range cr.Spec.ForProvider.ExplicitAuthFlows {
			var f10elem string
			f10elem = *f10iter
			f10 = append(f10, &f10elem)
		}
		res.SetExplicitAuthFlows(f10)
	}
	if cr.Spec.ForProvider.IDTokenValidity != nil {
		res.SetIdTokenValidity(*cr.Spec.ForProvider.IDTokenValidity)
	}
	if cr.Spec.ForProvider.LogoutURLs != nil {
		f12 := []*string{}
		for _, f12iter := range cr.Spec.ForProvider.LogoutURLs {
			var f12elem string
			f12elem = *f12iter
			f12 = append(f12, &f12elem)
		}
		res.SetLogoutURLs(f12)
	}
	if cr.Spec.ForProvider.PreventUserExistenceErrors != nil {
		res.SetPreventUserExistenceErrors(*cr.Spec.ForProvider.PreventUserExistenceErrors)
	}
	if cr.Spec.ForProvider.ReadAttributes != nil {
		f14 := []*string{}
		for _, f14iter := range cr.Spec.ForProvider.ReadAttributes {
			var f14elem string
			f14elem = *f14iter
			f14 = append(f14, &f14elem)
		}
		res.SetReadAttributes(f14)
	}
	if cr.Spec.ForProvider.RefreshTokenValidity != nil {
		res.SetRefreshTokenValidity(*cr.Spec.ForProvider.RefreshTokenValidity)
	}
	if cr.Spec.ForProvider.SupportedIdentityProviders != nil {
		f16 := []*string{}
		for _, f16iter := range cr.Spec.ForProvider.SupportedIdentityProviders {
			var f16elem string
			f16elem = *f16iter
			f16 = append(f16, &f16elem)
		}
		res.SetSupportedIdentityProviders(f16)
	}
	if cr.Spec.ForProvider.TokenValidityUnits != nil {
		f17 := &svcsdk.TokenValidityUnitsType{}
		if cr.Spec.ForProvider.TokenValidityUnits.AccessToken != nil {
			f17.SetAccessToken(*cr.Spec.ForProvider.TokenValidityUnits.AccessToken)
		}
		if cr.Spec.ForProvider.TokenValidityUnits.IDToken != nil {
			f17.SetIdToken(*cr.Spec.ForProvider.TokenValidityUnits.IDToken)
		}
		if cr.Spec.ForProvider.TokenValidityUnits.RefreshToken != nil {
			f17.SetRefreshToken(*cr.Spec.ForProvider.TokenValidityUnits.RefreshToken)
		}
		res.SetTokenValidityUnits(f17)
	}
	if cr.Status.AtProvider.UserPoolID != nil {
		res.SetUserPoolId(*cr.Status.AtProvider.UserPoolID)
	}
	if cr.Spec.ForProvider.WriteAttributes != nil {
		f19 := []*string{}
		for _, f19iter := range cr.Spec.ForProvider.WriteAttributes {
			var f19elem string
			f19elem = *f19iter
			f19 = append(f19, &f19elem)
		}
		res.SetWriteAttributes(f19)
	}

	return res
}

// GenerateDeleteUserPoolClientInput returns a deletion input.
func GenerateDeleteUserPoolClientInput(cr *svcapitypes.UserPoolClient) *svcsdk.DeleteUserPoolClientInput {
	res := &svcsdk.DeleteUserPoolClientInput{}

	if cr.Status.AtProvider.ClientID != nil {
		res.SetClientId(*cr.Status.AtProvider.ClientID)
	}
	if cr.Status.AtProvider.UserPoolID != nil {
		res.SetUserPoolId(*cr.Status.AtProvider.UserPoolID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}
