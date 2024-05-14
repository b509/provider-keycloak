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
func (in *AuthenticationFlowBindingOverridesInitParameters) DeepCopyInto(out *AuthenticationFlowBindingOverridesInitParameters) {
	*out = *in
	if in.BrowserID != nil {
		in, out := &in.BrowserID, &out.BrowserID
		*out = new(string)
		**out = **in
	}
	if in.DirectGrantID != nil {
		in, out := &in.DirectGrantID, &out.DirectGrantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationFlowBindingOverridesInitParameters.
func (in *AuthenticationFlowBindingOverridesInitParameters) DeepCopy() *AuthenticationFlowBindingOverridesInitParameters {
	if in == nil {
		return nil
	}
	out := new(AuthenticationFlowBindingOverridesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationFlowBindingOverridesObservation) DeepCopyInto(out *AuthenticationFlowBindingOverridesObservation) {
	*out = *in
	if in.BrowserID != nil {
		in, out := &in.BrowserID, &out.BrowserID
		*out = new(string)
		**out = **in
	}
	if in.DirectGrantID != nil {
		in, out := &in.DirectGrantID, &out.DirectGrantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationFlowBindingOverridesObservation.
func (in *AuthenticationFlowBindingOverridesObservation) DeepCopy() *AuthenticationFlowBindingOverridesObservation {
	if in == nil {
		return nil
	}
	out := new(AuthenticationFlowBindingOverridesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationFlowBindingOverridesParameters) DeepCopyInto(out *AuthenticationFlowBindingOverridesParameters) {
	*out = *in
	if in.BrowserID != nil {
		in, out := &in.BrowserID, &out.BrowserID
		*out = new(string)
		**out = **in
	}
	if in.DirectGrantID != nil {
		in, out := &in.DirectGrantID, &out.DirectGrantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationFlowBindingOverridesParameters.
func (in *AuthenticationFlowBindingOverridesParameters) DeepCopy() *AuthenticationFlowBindingOverridesParameters {
	if in == nil {
		return nil
	}
	out := new(AuthenticationFlowBindingOverridesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Client) DeepCopyInto(out *Client) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Client.
func (in *Client) DeepCopy() *Client {
	if in == nil {
		return nil
	}
	out := new(Client)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Client) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientDefaultScopes) DeepCopyInto(out *ClientDefaultScopes) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientDefaultScopes.
func (in *ClientDefaultScopes) DeepCopy() *ClientDefaultScopes {
	if in == nil {
		return nil
	}
	out := new(ClientDefaultScopes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientDefaultScopes) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientDefaultScopesInitParameters) DeepCopyInto(out *ClientDefaultScopesInitParameters) {
	*out = *in
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
	if in.DefaultScopes != nil {
		in, out := &in.DefaultScopes, &out.DefaultScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientDefaultScopesInitParameters.
func (in *ClientDefaultScopesInitParameters) DeepCopy() *ClientDefaultScopesInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClientDefaultScopesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientDefaultScopesList) DeepCopyInto(out *ClientDefaultScopesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClientDefaultScopes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientDefaultScopesList.
func (in *ClientDefaultScopesList) DeepCopy() *ClientDefaultScopesList {
	if in == nil {
		return nil
	}
	out := new(ClientDefaultScopesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientDefaultScopesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientDefaultScopesObservation) DeepCopyInto(out *ClientDefaultScopesObservation) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.DefaultScopes != nil {
		in, out := &in.DefaultScopes, &out.DefaultScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientDefaultScopesObservation.
func (in *ClientDefaultScopesObservation) DeepCopy() *ClientDefaultScopesObservation {
	if in == nil {
		return nil
	}
	out := new(ClientDefaultScopesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientDefaultScopesParameters) DeepCopyInto(out *ClientDefaultScopesParameters) {
	*out = *in
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
	if in.DefaultScopes != nil {
		in, out := &in.DefaultScopes, &out.DefaultScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientDefaultScopesParameters.
func (in *ClientDefaultScopesParameters) DeepCopy() *ClientDefaultScopesParameters {
	if in == nil {
		return nil
	}
	out := new(ClientDefaultScopesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientDefaultScopesSpec) DeepCopyInto(out *ClientDefaultScopesSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientDefaultScopesSpec.
func (in *ClientDefaultScopesSpec) DeepCopy() *ClientDefaultScopesSpec {
	if in == nil {
		return nil
	}
	out := new(ClientDefaultScopesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientDefaultScopesStatus) DeepCopyInto(out *ClientDefaultScopesStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientDefaultScopesStatus.
func (in *ClientDefaultScopesStatus) DeepCopy() *ClientDefaultScopesStatus {
	if in == nil {
		return nil
	}
	out := new(ClientDefaultScopesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientInitParameters) DeepCopyInto(out *ClientInitParameters) {
	*out = *in
	if in.AssertionConsumerPostURL != nil {
		in, out := &in.AssertionConsumerPostURL, &out.AssertionConsumerPostURL
		*out = new(string)
		**out = **in
	}
	if in.AssertionConsumerRedirectURL != nil {
		in, out := &in.AssertionConsumerRedirectURL, &out.AssertionConsumerRedirectURL
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationFlowBindingOverrides != nil {
		in, out := &in.AuthenticationFlowBindingOverrides, &out.AuthenticationFlowBindingOverrides
		*out = make([]AuthenticationFlowBindingOverridesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BaseURL != nil {
		in, out := &in.BaseURL, &out.BaseURL
		*out = new(string)
		**out = **in
	}
	if in.CanonicalizationMethod != nil {
		in, out := &in.CanonicalizationMethod, &out.CanonicalizationMethod
		*out = new(string)
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
	if in.ClientSignatureRequired != nil {
		in, out := &in.ClientSignatureRequired, &out.ClientSignatureRequired
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptAssertions != nil {
		in, out := &in.EncryptAssertions, &out.EncryptAssertions
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionCertificate != nil {
		in, out := &in.EncryptionCertificate, &out.EncryptionCertificate
		*out = new(string)
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
	if in.ForceNameIDFormat != nil {
		in, out := &in.ForceNameIDFormat, &out.ForceNameIDFormat
		*out = new(bool)
		**out = **in
	}
	if in.ForcePostBinding != nil {
		in, out := &in.ForcePostBinding, &out.ForcePostBinding
		*out = new(bool)
		**out = **in
	}
	if in.FrontChannelLogout != nil {
		in, out := &in.FrontChannelLogout, &out.FrontChannelLogout
		*out = new(bool)
		**out = **in
	}
	if in.FullScopeAllowed != nil {
		in, out := &in.FullScopeAllowed, &out.FullScopeAllowed
		*out = new(bool)
		**out = **in
	}
	if in.IdpInitiatedSsoRelayState != nil {
		in, out := &in.IdpInitiatedSsoRelayState, &out.IdpInitiatedSsoRelayState
		*out = new(string)
		**out = **in
	}
	if in.IdpInitiatedSsoURLName != nil {
		in, out := &in.IdpInitiatedSsoURLName, &out.IdpInitiatedSsoURLName
		*out = new(string)
		**out = **in
	}
	if in.IncludeAuthnStatement != nil {
		in, out := &in.IncludeAuthnStatement, &out.IncludeAuthnStatement
		*out = new(bool)
		**out = **in
	}
	if in.LoginTheme != nil {
		in, out := &in.LoginTheme, &out.LoginTheme
		*out = new(string)
		**out = **in
	}
	if in.LogoutServicePostBindingURL != nil {
		in, out := &in.LogoutServicePostBindingURL, &out.LogoutServicePostBindingURL
		*out = new(string)
		**out = **in
	}
	if in.LogoutServiceRedirectBindingURL != nil {
		in, out := &in.LogoutServiceRedirectBindingURL, &out.LogoutServiceRedirectBindingURL
		*out = new(string)
		**out = **in
	}
	if in.MasterSAMLProcessingURL != nil {
		in, out := &in.MasterSAMLProcessingURL, &out.MasterSAMLProcessingURL
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameIDFormat != nil {
		in, out := &in.NameIDFormat, &out.NameIDFormat
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RootURL != nil {
		in, out := &in.RootURL, &out.RootURL
		*out = new(string)
		**out = **in
	}
	if in.SignAssertions != nil {
		in, out := &in.SignAssertions, &out.SignAssertions
		*out = new(bool)
		**out = **in
	}
	if in.SignDocuments != nil {
		in, out := &in.SignDocuments, &out.SignDocuments
		*out = new(bool)
		**out = **in
	}
	if in.SignatureAlgorithm != nil {
		in, out := &in.SignatureAlgorithm, &out.SignatureAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.SignatureKeyName != nil {
		in, out := &in.SignatureKeyName, &out.SignatureKeyName
		*out = new(string)
		**out = **in
	}
	if in.SigningCertificate != nil {
		in, out := &in.SigningCertificate, &out.SigningCertificate
		*out = new(string)
		**out = **in
	}
	if in.SigningPrivateKey != nil {
		in, out := &in.SigningPrivateKey, &out.SigningPrivateKey
		*out = new(string)
		**out = **in
	}
	if in.ValidRedirectUris != nil {
		in, out := &in.ValidRedirectUris, &out.ValidRedirectUris
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientInitParameters.
func (in *ClientInitParameters) DeepCopy() *ClientInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClientInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientList) DeepCopyInto(out *ClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Client, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientList.
func (in *ClientList) DeepCopy() *ClientList {
	if in == nil {
		return nil
	}
	out := new(ClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientObservation) DeepCopyInto(out *ClientObservation) {
	*out = *in
	if in.AssertionConsumerPostURL != nil {
		in, out := &in.AssertionConsumerPostURL, &out.AssertionConsumerPostURL
		*out = new(string)
		**out = **in
	}
	if in.AssertionConsumerRedirectURL != nil {
		in, out := &in.AssertionConsumerRedirectURL, &out.AssertionConsumerRedirectURL
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationFlowBindingOverrides != nil {
		in, out := &in.AuthenticationFlowBindingOverrides, &out.AuthenticationFlowBindingOverrides
		*out = make([]AuthenticationFlowBindingOverridesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BaseURL != nil {
		in, out := &in.BaseURL, &out.BaseURL
		*out = new(string)
		**out = **in
	}
	if in.CanonicalizationMethod != nil {
		in, out := &in.CanonicalizationMethod, &out.CanonicalizationMethod
		*out = new(string)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientSignatureRequired != nil {
		in, out := &in.ClientSignatureRequired, &out.ClientSignatureRequired
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptAssertions != nil {
		in, out := &in.EncryptAssertions, &out.EncryptAssertions
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionCertificate != nil {
		in, out := &in.EncryptionCertificate, &out.EncryptionCertificate
		*out = new(string)
		**out = **in
	}
	if in.EncryptionCertificateSha1 != nil {
		in, out := &in.EncryptionCertificateSha1, &out.EncryptionCertificateSha1
		*out = new(string)
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
	if in.ForceNameIDFormat != nil {
		in, out := &in.ForceNameIDFormat, &out.ForceNameIDFormat
		*out = new(bool)
		**out = **in
	}
	if in.ForcePostBinding != nil {
		in, out := &in.ForcePostBinding, &out.ForcePostBinding
		*out = new(bool)
		**out = **in
	}
	if in.FrontChannelLogout != nil {
		in, out := &in.FrontChannelLogout, &out.FrontChannelLogout
		*out = new(bool)
		**out = **in
	}
	if in.FullScopeAllowed != nil {
		in, out := &in.FullScopeAllowed, &out.FullScopeAllowed
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IdpInitiatedSsoRelayState != nil {
		in, out := &in.IdpInitiatedSsoRelayState, &out.IdpInitiatedSsoRelayState
		*out = new(string)
		**out = **in
	}
	if in.IdpInitiatedSsoURLName != nil {
		in, out := &in.IdpInitiatedSsoURLName, &out.IdpInitiatedSsoURLName
		*out = new(string)
		**out = **in
	}
	if in.IncludeAuthnStatement != nil {
		in, out := &in.IncludeAuthnStatement, &out.IncludeAuthnStatement
		*out = new(bool)
		**out = **in
	}
	if in.LoginTheme != nil {
		in, out := &in.LoginTheme, &out.LoginTheme
		*out = new(string)
		**out = **in
	}
	if in.LogoutServicePostBindingURL != nil {
		in, out := &in.LogoutServicePostBindingURL, &out.LogoutServicePostBindingURL
		*out = new(string)
		**out = **in
	}
	if in.LogoutServiceRedirectBindingURL != nil {
		in, out := &in.LogoutServiceRedirectBindingURL, &out.LogoutServiceRedirectBindingURL
		*out = new(string)
		**out = **in
	}
	if in.MasterSAMLProcessingURL != nil {
		in, out := &in.MasterSAMLProcessingURL, &out.MasterSAMLProcessingURL
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameIDFormat != nil {
		in, out := &in.NameIDFormat, &out.NameIDFormat
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RootURL != nil {
		in, out := &in.RootURL, &out.RootURL
		*out = new(string)
		**out = **in
	}
	if in.SignAssertions != nil {
		in, out := &in.SignAssertions, &out.SignAssertions
		*out = new(bool)
		**out = **in
	}
	if in.SignDocuments != nil {
		in, out := &in.SignDocuments, &out.SignDocuments
		*out = new(bool)
		**out = **in
	}
	if in.SignatureAlgorithm != nil {
		in, out := &in.SignatureAlgorithm, &out.SignatureAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.SignatureKeyName != nil {
		in, out := &in.SignatureKeyName, &out.SignatureKeyName
		*out = new(string)
		**out = **in
	}
	if in.SigningCertificate != nil {
		in, out := &in.SigningCertificate, &out.SigningCertificate
		*out = new(string)
		**out = **in
	}
	if in.SigningCertificateSha1 != nil {
		in, out := &in.SigningCertificateSha1, &out.SigningCertificateSha1
		*out = new(string)
		**out = **in
	}
	if in.SigningPrivateKey != nil {
		in, out := &in.SigningPrivateKey, &out.SigningPrivateKey
		*out = new(string)
		**out = **in
	}
	if in.SigningPrivateKeySha1 != nil {
		in, out := &in.SigningPrivateKeySha1, &out.SigningPrivateKeySha1
		*out = new(string)
		**out = **in
	}
	if in.ValidRedirectUris != nil {
		in, out := &in.ValidRedirectUris, &out.ValidRedirectUris
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientObservation.
func (in *ClientObservation) DeepCopy() *ClientObservation {
	if in == nil {
		return nil
	}
	out := new(ClientObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientParameters) DeepCopyInto(out *ClientParameters) {
	*out = *in
	if in.AssertionConsumerPostURL != nil {
		in, out := &in.AssertionConsumerPostURL, &out.AssertionConsumerPostURL
		*out = new(string)
		**out = **in
	}
	if in.AssertionConsumerRedirectURL != nil {
		in, out := &in.AssertionConsumerRedirectURL, &out.AssertionConsumerRedirectURL
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationFlowBindingOverrides != nil {
		in, out := &in.AuthenticationFlowBindingOverrides, &out.AuthenticationFlowBindingOverrides
		*out = make([]AuthenticationFlowBindingOverridesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BaseURL != nil {
		in, out := &in.BaseURL, &out.BaseURL
		*out = new(string)
		**out = **in
	}
	if in.CanonicalizationMethod != nil {
		in, out := &in.CanonicalizationMethod, &out.CanonicalizationMethod
		*out = new(string)
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
	if in.ClientSignatureRequired != nil {
		in, out := &in.ClientSignatureRequired, &out.ClientSignatureRequired
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptAssertions != nil {
		in, out := &in.EncryptAssertions, &out.EncryptAssertions
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionCertificate != nil {
		in, out := &in.EncryptionCertificate, &out.EncryptionCertificate
		*out = new(string)
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
	if in.ForceNameIDFormat != nil {
		in, out := &in.ForceNameIDFormat, &out.ForceNameIDFormat
		*out = new(bool)
		**out = **in
	}
	if in.ForcePostBinding != nil {
		in, out := &in.ForcePostBinding, &out.ForcePostBinding
		*out = new(bool)
		**out = **in
	}
	if in.FrontChannelLogout != nil {
		in, out := &in.FrontChannelLogout, &out.FrontChannelLogout
		*out = new(bool)
		**out = **in
	}
	if in.FullScopeAllowed != nil {
		in, out := &in.FullScopeAllowed, &out.FullScopeAllowed
		*out = new(bool)
		**out = **in
	}
	if in.IdpInitiatedSsoRelayState != nil {
		in, out := &in.IdpInitiatedSsoRelayState, &out.IdpInitiatedSsoRelayState
		*out = new(string)
		**out = **in
	}
	if in.IdpInitiatedSsoURLName != nil {
		in, out := &in.IdpInitiatedSsoURLName, &out.IdpInitiatedSsoURLName
		*out = new(string)
		**out = **in
	}
	if in.IncludeAuthnStatement != nil {
		in, out := &in.IncludeAuthnStatement, &out.IncludeAuthnStatement
		*out = new(bool)
		**out = **in
	}
	if in.LoginTheme != nil {
		in, out := &in.LoginTheme, &out.LoginTheme
		*out = new(string)
		**out = **in
	}
	if in.LogoutServicePostBindingURL != nil {
		in, out := &in.LogoutServicePostBindingURL, &out.LogoutServicePostBindingURL
		*out = new(string)
		**out = **in
	}
	if in.LogoutServiceRedirectBindingURL != nil {
		in, out := &in.LogoutServiceRedirectBindingURL, &out.LogoutServiceRedirectBindingURL
		*out = new(string)
		**out = **in
	}
	if in.MasterSAMLProcessingURL != nil {
		in, out := &in.MasterSAMLProcessingURL, &out.MasterSAMLProcessingURL
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameIDFormat != nil {
		in, out := &in.NameIDFormat, &out.NameIDFormat
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RootURL != nil {
		in, out := &in.RootURL, &out.RootURL
		*out = new(string)
		**out = **in
	}
	if in.SignAssertions != nil {
		in, out := &in.SignAssertions, &out.SignAssertions
		*out = new(bool)
		**out = **in
	}
	if in.SignDocuments != nil {
		in, out := &in.SignDocuments, &out.SignDocuments
		*out = new(bool)
		**out = **in
	}
	if in.SignatureAlgorithm != nil {
		in, out := &in.SignatureAlgorithm, &out.SignatureAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.SignatureKeyName != nil {
		in, out := &in.SignatureKeyName, &out.SignatureKeyName
		*out = new(string)
		**out = **in
	}
	if in.SigningCertificate != nil {
		in, out := &in.SigningCertificate, &out.SigningCertificate
		*out = new(string)
		**out = **in
	}
	if in.SigningPrivateKey != nil {
		in, out := &in.SigningPrivateKey, &out.SigningPrivateKey
		*out = new(string)
		**out = **in
	}
	if in.ValidRedirectUris != nil {
		in, out := &in.ValidRedirectUris, &out.ValidRedirectUris
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientParameters.
func (in *ClientParameters) DeepCopy() *ClientParameters {
	if in == nil {
		return nil
	}
	out := new(ClientParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientScope) DeepCopyInto(out *ClientScope) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientScope.
func (in *ClientScope) DeepCopy() *ClientScope {
	if in == nil {
		return nil
	}
	out := new(ClientScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientScope) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientScopeInitParameters) DeepCopyInto(out *ClientScopeInitParameters) {
	*out = *in
	if in.ConsentScreenText != nil {
		in, out := &in.ConsentScreenText, &out.ConsentScreenText
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.GuiOrder != nil {
		in, out := &in.GuiOrder, &out.GuiOrder
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientScopeInitParameters.
func (in *ClientScopeInitParameters) DeepCopy() *ClientScopeInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClientScopeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientScopeList) DeepCopyInto(out *ClientScopeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClientScope, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientScopeList.
func (in *ClientScopeList) DeepCopy() *ClientScopeList {
	if in == nil {
		return nil
	}
	out := new(ClientScopeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientScopeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientScopeObservation) DeepCopyInto(out *ClientScopeObservation) {
	*out = *in
	if in.ConsentScreenText != nil {
		in, out := &in.ConsentScreenText, &out.ConsentScreenText
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.GuiOrder != nil {
		in, out := &in.GuiOrder, &out.GuiOrder
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientScopeObservation.
func (in *ClientScopeObservation) DeepCopy() *ClientScopeObservation {
	if in == nil {
		return nil
	}
	out := new(ClientScopeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientScopeParameters) DeepCopyInto(out *ClientScopeParameters) {
	*out = *in
	if in.ConsentScreenText != nil {
		in, out := &in.ConsentScreenText, &out.ConsentScreenText
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.GuiOrder != nil {
		in, out := &in.GuiOrder, &out.GuiOrder
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientScopeParameters.
func (in *ClientScopeParameters) DeepCopy() *ClientScopeParameters {
	if in == nil {
		return nil
	}
	out := new(ClientScopeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientScopeSpec) DeepCopyInto(out *ClientScopeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientScopeSpec.
func (in *ClientScopeSpec) DeepCopy() *ClientScopeSpec {
	if in == nil {
		return nil
	}
	out := new(ClientScopeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientScopeStatus) DeepCopyInto(out *ClientScopeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientScopeStatus.
func (in *ClientScopeStatus) DeepCopy() *ClientScopeStatus {
	if in == nil {
		return nil
	}
	out := new(ClientScopeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSpec) DeepCopyInto(out *ClientSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSpec.
func (in *ClientSpec) DeepCopy() *ClientSpec {
	if in == nil {
		return nil
	}
	out := new(ClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientStatus) DeepCopyInto(out *ClientStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientStatus.
func (in *ClientStatus) DeepCopy() *ClientStatus {
	if in == nil {
		return nil
	}
	out := new(ClientStatus)
	in.DeepCopyInto(out)
	return out
}
