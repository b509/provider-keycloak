//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProvider) DeepCopyInto(out *IdentityProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProvider.
func (in *IdentityProvider) DeepCopy() *IdentityProvider {
	if in == nil {
		return nil
	}
	out := new(IdentityProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderInitParameters) DeepCopyInto(out *IdentityProviderInitParameters) {
	*out = *in
	if in.AcceptsPromptNoneForwardFromClient != nil {
		in, out := &in.AcceptsPromptNoneForwardFromClient, &out.AcceptsPromptNoneForwardFromClient
		*out = new(bool)
		**out = **in
	}
	if in.AddReadTokenRoleOnCreate != nil {
		in, out := &in.AddReadTokenRoleOnCreate, &out.AddReadTokenRoleOnCreate
		*out = new(bool)
		**out = **in
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.AuthenticateByDefault != nil {
		in, out := &in.AuthenticateByDefault, &out.AuthenticateByDefault
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizationURL != nil {
		in, out := &in.AuthorizationURL, &out.AuthorizationURL
		*out = new(string)
		**out = **in
	}
	if in.BackchannelSupported != nil {
		in, out := &in.BackchannelSupported, &out.BackchannelSupported
		*out = new(bool)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientIDRef != nil {
		in, out := &in.ClientIDRef, &out.ClientIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientIDSelector != nil {
		in, out := &in.ClientIDSelector, &out.ClientIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	out.ClientSecretSecretRef = in.ClientSecretSecretRef
	if in.DefaultScopes != nil {
		in, out := &in.DefaultScopes, &out.DefaultScopes
		*out = new(string)
		**out = **in
	}
	if in.DisableUserInfo != nil {
		in, out := &in.DisableUserInfo, &out.DisableUserInfo
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.FirstBrokerLoginFlowAlias != nil {
		in, out := &in.FirstBrokerLoginFlowAlias, &out.FirstBrokerLoginFlowAlias
		*out = new(string)
		**out = **in
	}
	if in.FirstBrokerLoginFlowAliasRef != nil {
		in, out := &in.FirstBrokerLoginFlowAliasRef, &out.FirstBrokerLoginFlowAliasRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FirstBrokerLoginFlowAliasSelector != nil {
		in, out := &in.FirstBrokerLoginFlowAliasSelector, &out.FirstBrokerLoginFlowAliasSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.GuiOrder != nil {
		in, out := &in.GuiOrder, &out.GuiOrder
		*out = new(string)
		**out = **in
	}
	if in.HideOnLoginPage != nil {
		in, out := &in.HideOnLoginPage, &out.HideOnLoginPage
		*out = new(bool)
		**out = **in
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	if in.JwksURL != nil {
		in, out := &in.JwksURL, &out.JwksURL
		*out = new(string)
		**out = **in
	}
	if in.LinkOnly != nil {
		in, out := &in.LinkOnly, &out.LinkOnly
		*out = new(bool)
		**out = **in
	}
	if in.LoginHint != nil {
		in, out := &in.LoginHint, &out.LoginHint
		*out = new(string)
		**out = **in
	}
	if in.LogoutURL != nil {
		in, out := &in.LogoutURL, &out.LogoutURL
		*out = new(string)
		**out = **in
	}
	if in.PostBrokerLoginFlowAlias != nil {
		in, out := &in.PostBrokerLoginFlowAlias, &out.PostBrokerLoginFlowAlias
		*out = new(string)
		**out = **in
	}
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(string)
		**out = **in
	}
	if in.RealmRef != nil {
		in, out := &in.RealmRef, &out.RealmRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmSelector != nil {
		in, out := &in.RealmSelector, &out.RealmSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.StoreToken != nil {
		in, out := &in.StoreToken, &out.StoreToken
		*out = new(bool)
		**out = **in
	}
	if in.SyncMode != nil {
		in, out := &in.SyncMode, &out.SyncMode
		*out = new(string)
		**out = **in
	}
	if in.TokenURL != nil {
		in, out := &in.TokenURL, &out.TokenURL
		*out = new(string)
		**out = **in
	}
	if in.TrustEmail != nil {
		in, out := &in.TrustEmail, &out.TrustEmail
		*out = new(bool)
		**out = **in
	}
	if in.UILocales != nil {
		in, out := &in.UILocales, &out.UILocales
		*out = new(bool)
		**out = **in
	}
	if in.UserInfoURL != nil {
		in, out := &in.UserInfoURL, &out.UserInfoURL
		*out = new(string)
		**out = **in
	}
	if in.ValidateSignature != nil {
		in, out := &in.ValidateSignature, &out.ValidateSignature
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderInitParameters.
func (in *IdentityProviderInitParameters) DeepCopy() *IdentityProviderInitParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderList) DeepCopyInto(out *IdentityProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdentityProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderList.
func (in *IdentityProviderList) DeepCopy() *IdentityProviderList {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderObservation) DeepCopyInto(out *IdentityProviderObservation) {
	*out = *in
	if in.AcceptsPromptNoneForwardFromClient != nil {
		in, out := &in.AcceptsPromptNoneForwardFromClient, &out.AcceptsPromptNoneForwardFromClient
		*out = new(bool)
		**out = **in
	}
	if in.AddReadTokenRoleOnCreate != nil {
		in, out := &in.AddReadTokenRoleOnCreate, &out.AddReadTokenRoleOnCreate
		*out = new(bool)
		**out = **in
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.AuthenticateByDefault != nil {
		in, out := &in.AuthenticateByDefault, &out.AuthenticateByDefault
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizationURL != nil {
		in, out := &in.AuthorizationURL, &out.AuthorizationURL
		*out = new(string)
		**out = **in
	}
	if in.BackchannelSupported != nil {
		in, out := &in.BackchannelSupported, &out.BackchannelSupported
		*out = new(bool)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.DefaultScopes != nil {
		in, out := &in.DefaultScopes, &out.DefaultScopes
		*out = new(string)
		**out = **in
	}
	if in.DisableUserInfo != nil {
		in, out := &in.DisableUserInfo, &out.DisableUserInfo
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.FirstBrokerLoginFlowAlias != nil {
		in, out := &in.FirstBrokerLoginFlowAlias, &out.FirstBrokerLoginFlowAlias
		*out = new(string)
		**out = **in
	}
	if in.GuiOrder != nil {
		in, out := &in.GuiOrder, &out.GuiOrder
		*out = new(string)
		**out = **in
	}
	if in.HideOnLoginPage != nil {
		in, out := &in.HideOnLoginPage, &out.HideOnLoginPage
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InternalID != nil {
		in, out := &in.InternalID, &out.InternalID
		*out = new(string)
		**out = **in
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	if in.JwksURL != nil {
		in, out := &in.JwksURL, &out.JwksURL
		*out = new(string)
		**out = **in
	}
	if in.LinkOnly != nil {
		in, out := &in.LinkOnly, &out.LinkOnly
		*out = new(bool)
		**out = **in
	}
	if in.LoginHint != nil {
		in, out := &in.LoginHint, &out.LoginHint
		*out = new(string)
		**out = **in
	}
	if in.LogoutURL != nil {
		in, out := &in.LogoutURL, &out.LogoutURL
		*out = new(string)
		**out = **in
	}
	if in.PostBrokerLoginFlowAlias != nil {
		in, out := &in.PostBrokerLoginFlowAlias, &out.PostBrokerLoginFlowAlias
		*out = new(string)
		**out = **in
	}
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(string)
		**out = **in
	}
	if in.StoreToken != nil {
		in, out := &in.StoreToken, &out.StoreToken
		*out = new(bool)
		**out = **in
	}
	if in.SyncMode != nil {
		in, out := &in.SyncMode, &out.SyncMode
		*out = new(string)
		**out = **in
	}
	if in.TokenURL != nil {
		in, out := &in.TokenURL, &out.TokenURL
		*out = new(string)
		**out = **in
	}
	if in.TrustEmail != nil {
		in, out := &in.TrustEmail, &out.TrustEmail
		*out = new(bool)
		**out = **in
	}
	if in.UILocales != nil {
		in, out := &in.UILocales, &out.UILocales
		*out = new(bool)
		**out = **in
	}
	if in.UserInfoURL != nil {
		in, out := &in.UserInfoURL, &out.UserInfoURL
		*out = new(string)
		**out = **in
	}
	if in.ValidateSignature != nil {
		in, out := &in.ValidateSignature, &out.ValidateSignature
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderObservation.
func (in *IdentityProviderObservation) DeepCopy() *IdentityProviderObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderParameters) DeepCopyInto(out *IdentityProviderParameters) {
	*out = *in
	if in.AcceptsPromptNoneForwardFromClient != nil {
		in, out := &in.AcceptsPromptNoneForwardFromClient, &out.AcceptsPromptNoneForwardFromClient
		*out = new(bool)
		**out = **in
	}
	if in.AddReadTokenRoleOnCreate != nil {
		in, out := &in.AddReadTokenRoleOnCreate, &out.AddReadTokenRoleOnCreate
		*out = new(bool)
		**out = **in
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.AuthenticateByDefault != nil {
		in, out := &in.AuthenticateByDefault, &out.AuthenticateByDefault
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizationURL != nil {
		in, out := &in.AuthorizationURL, &out.AuthorizationURL
		*out = new(string)
		**out = **in
	}
	if in.BackchannelSupported != nil {
		in, out := &in.BackchannelSupported, &out.BackchannelSupported
		*out = new(bool)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientIDRef != nil {
		in, out := &in.ClientIDRef, &out.ClientIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientIDSelector != nil {
		in, out := &in.ClientIDSelector, &out.ClientIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	out.ClientSecretSecretRef = in.ClientSecretSecretRef
	if in.DefaultScopes != nil {
		in, out := &in.DefaultScopes, &out.DefaultScopes
		*out = new(string)
		**out = **in
	}
	if in.DisableUserInfo != nil {
		in, out := &in.DisableUserInfo, &out.DisableUserInfo
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.FirstBrokerLoginFlowAlias != nil {
		in, out := &in.FirstBrokerLoginFlowAlias, &out.FirstBrokerLoginFlowAlias
		*out = new(string)
		**out = **in
	}
	if in.FirstBrokerLoginFlowAliasRef != nil {
		in, out := &in.FirstBrokerLoginFlowAliasRef, &out.FirstBrokerLoginFlowAliasRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FirstBrokerLoginFlowAliasSelector != nil {
		in, out := &in.FirstBrokerLoginFlowAliasSelector, &out.FirstBrokerLoginFlowAliasSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.GuiOrder != nil {
		in, out := &in.GuiOrder, &out.GuiOrder
		*out = new(string)
		**out = **in
	}
	if in.HideOnLoginPage != nil {
		in, out := &in.HideOnLoginPage, &out.HideOnLoginPage
		*out = new(bool)
		**out = **in
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	if in.JwksURL != nil {
		in, out := &in.JwksURL, &out.JwksURL
		*out = new(string)
		**out = **in
	}
	if in.LinkOnly != nil {
		in, out := &in.LinkOnly, &out.LinkOnly
		*out = new(bool)
		**out = **in
	}
	if in.LoginHint != nil {
		in, out := &in.LoginHint, &out.LoginHint
		*out = new(string)
		**out = **in
	}
	if in.LogoutURL != nil {
		in, out := &in.LogoutURL, &out.LogoutURL
		*out = new(string)
		**out = **in
	}
	if in.PostBrokerLoginFlowAlias != nil {
		in, out := &in.PostBrokerLoginFlowAlias, &out.PostBrokerLoginFlowAlias
		*out = new(string)
		**out = **in
	}
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(string)
		**out = **in
	}
	if in.RealmRef != nil {
		in, out := &in.RealmRef, &out.RealmRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmSelector != nil {
		in, out := &in.RealmSelector, &out.RealmSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.StoreToken != nil {
		in, out := &in.StoreToken, &out.StoreToken
		*out = new(bool)
		**out = **in
	}
	if in.SyncMode != nil {
		in, out := &in.SyncMode, &out.SyncMode
		*out = new(string)
		**out = **in
	}
	if in.TokenURL != nil {
		in, out := &in.TokenURL, &out.TokenURL
		*out = new(string)
		**out = **in
	}
	if in.TrustEmail != nil {
		in, out := &in.TrustEmail, &out.TrustEmail
		*out = new(bool)
		**out = **in
	}
	if in.UILocales != nil {
		in, out := &in.UILocales, &out.UILocales
		*out = new(bool)
		**out = **in
	}
	if in.UserInfoURL != nil {
		in, out := &in.UserInfoURL, &out.UserInfoURL
		*out = new(string)
		**out = **in
	}
	if in.ValidateSignature != nil {
		in, out := &in.ValidateSignature, &out.ValidateSignature
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderParameters.
func (in *IdentityProviderParameters) DeepCopy() *IdentityProviderParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderSpec) DeepCopyInto(out *IdentityProviderSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderSpec.
func (in *IdentityProviderSpec) DeepCopy() *IdentityProviderSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderStatus) DeepCopyInto(out *IdentityProviderStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderStatus.
func (in *IdentityProviderStatus) DeepCopy() *IdentityProviderStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderStatus)
	in.DeepCopyInto(out)
	return out
}
