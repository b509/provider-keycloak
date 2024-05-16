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
func (in *CacheInitParameters) DeepCopyInto(out *CacheInitParameters) {
	*out = *in
	if in.EvictionDay != nil {
		in, out := &in.EvictionDay, &out.EvictionDay
		*out = new(float64)
		**out = **in
	}
	if in.EvictionHour != nil {
		in, out := &in.EvictionHour, &out.EvictionHour
		*out = new(float64)
		**out = **in
	}
	if in.EvictionMinute != nil {
		in, out := &in.EvictionMinute, &out.EvictionMinute
		*out = new(float64)
		**out = **in
	}
	if in.MaxLifespan != nil {
		in, out := &in.MaxLifespan, &out.MaxLifespan
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheInitParameters.
func (in *CacheInitParameters) DeepCopy() *CacheInitParameters {
	if in == nil {
		return nil
	}
	out := new(CacheInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheObservation) DeepCopyInto(out *CacheObservation) {
	*out = *in
	if in.EvictionDay != nil {
		in, out := &in.EvictionDay, &out.EvictionDay
		*out = new(float64)
		**out = **in
	}
	if in.EvictionHour != nil {
		in, out := &in.EvictionHour, &out.EvictionHour
		*out = new(float64)
		**out = **in
	}
	if in.EvictionMinute != nil {
		in, out := &in.EvictionMinute, &out.EvictionMinute
		*out = new(float64)
		**out = **in
	}
	if in.MaxLifespan != nil {
		in, out := &in.MaxLifespan, &out.MaxLifespan
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheObservation.
func (in *CacheObservation) DeepCopy() *CacheObservation {
	if in == nil {
		return nil
	}
	out := new(CacheObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheParameters) DeepCopyInto(out *CacheParameters) {
	*out = *in
	if in.EvictionDay != nil {
		in, out := &in.EvictionDay, &out.EvictionDay
		*out = new(float64)
		**out = **in
	}
	if in.EvictionHour != nil {
		in, out := &in.EvictionHour, &out.EvictionHour
		*out = new(float64)
		**out = **in
	}
	if in.EvictionMinute != nil {
		in, out := &in.EvictionMinute, &out.EvictionMinute
		*out = new(float64)
		**out = **in
	}
	if in.MaxLifespan != nil {
		in, out := &in.MaxLifespan, &out.MaxLifespan
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheParameters.
func (in *CacheParameters) DeepCopy() *CacheParameters {
	if in == nil {
		return nil
	}
	out := new(CacheParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KerberosInitParameters) DeepCopyInto(out *KerberosInitParameters) {
	*out = *in
	if in.KerberosRealm != nil {
		in, out := &in.KerberosRealm, &out.KerberosRealm
		*out = new(string)
		**out = **in
	}
	if in.KeyTab != nil {
		in, out := &in.KeyTab, &out.KeyTab
		*out = new(string)
		**out = **in
	}
	if in.ServerPrincipal != nil {
		in, out := &in.ServerPrincipal, &out.ServerPrincipal
		*out = new(string)
		**out = **in
	}
	if in.UseKerberosForPasswordAuthentication != nil {
		in, out := &in.UseKerberosForPasswordAuthentication, &out.UseKerberosForPasswordAuthentication
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KerberosInitParameters.
func (in *KerberosInitParameters) DeepCopy() *KerberosInitParameters {
	if in == nil {
		return nil
	}
	out := new(KerberosInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KerberosObservation) DeepCopyInto(out *KerberosObservation) {
	*out = *in
	if in.KerberosRealm != nil {
		in, out := &in.KerberosRealm, &out.KerberosRealm
		*out = new(string)
		**out = **in
	}
	if in.KeyTab != nil {
		in, out := &in.KeyTab, &out.KeyTab
		*out = new(string)
		**out = **in
	}
	if in.ServerPrincipal != nil {
		in, out := &in.ServerPrincipal, &out.ServerPrincipal
		*out = new(string)
		**out = **in
	}
	if in.UseKerberosForPasswordAuthentication != nil {
		in, out := &in.UseKerberosForPasswordAuthentication, &out.UseKerberosForPasswordAuthentication
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KerberosObservation.
func (in *KerberosObservation) DeepCopy() *KerberosObservation {
	if in == nil {
		return nil
	}
	out := new(KerberosObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KerberosParameters) DeepCopyInto(out *KerberosParameters) {
	*out = *in
	if in.KerberosRealm != nil {
		in, out := &in.KerberosRealm, &out.KerberosRealm
		*out = new(string)
		**out = **in
	}
	if in.KeyTab != nil {
		in, out := &in.KeyTab, &out.KeyTab
		*out = new(string)
		**out = **in
	}
	if in.ServerPrincipal != nil {
		in, out := &in.ServerPrincipal, &out.ServerPrincipal
		*out = new(string)
		**out = **in
	}
	if in.UseKerberosForPasswordAuthentication != nil {
		in, out := &in.UseKerberosForPasswordAuthentication, &out.UseKerberosForPasswordAuthentication
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KerberosParameters.
func (in *KerberosParameters) DeepCopy() *KerberosParameters {
	if in == nil {
		return nil
	}
	out := new(KerberosParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAttributeMapper) DeepCopyInto(out *UserAttributeMapper) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAttributeMapper.
func (in *UserAttributeMapper) DeepCopy() *UserAttributeMapper {
	if in == nil {
		return nil
	}
	out := new(UserAttributeMapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserAttributeMapper) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAttributeMapperInitParameters) DeepCopyInto(out *UserAttributeMapperInitParameters) {
	*out = *in
	if in.AlwaysReadValueFromLdap != nil {
		in, out := &in.AlwaysReadValueFromLdap, &out.AlwaysReadValueFromLdap
		*out = new(bool)
		**out = **in
	}
	if in.AttributeDefaultValue != nil {
		in, out := &in.AttributeDefaultValue, &out.AttributeDefaultValue
		*out = new(string)
		**out = **in
	}
	if in.IsBinaryAttribute != nil {
		in, out := &in.IsBinaryAttribute, &out.IsBinaryAttribute
		*out = new(bool)
		**out = **in
	}
	if in.IsMandatoryInLdap != nil {
		in, out := &in.IsMandatoryInLdap, &out.IsMandatoryInLdap
		*out = new(bool)
		**out = **in
	}
	if in.LdapAttribute != nil {
		in, out := &in.LdapAttribute, &out.LdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.LdapUserFederationID != nil {
		in, out := &in.LdapUserFederationID, &out.LdapUserFederationID
		*out = new(string)
		**out = **in
	}
	if in.LdapUserFederationIDRef != nil {
		in, out := &in.LdapUserFederationIDRef, &out.LdapUserFederationIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LdapUserFederationIDSelector != nil {
		in, out := &in.LdapUserFederationIDSelector, &out.LdapUserFederationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
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
	if in.UserModelAttribute != nil {
		in, out := &in.UserModelAttribute, &out.UserModelAttribute
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAttributeMapperInitParameters.
func (in *UserAttributeMapperInitParameters) DeepCopy() *UserAttributeMapperInitParameters {
	if in == nil {
		return nil
	}
	out := new(UserAttributeMapperInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAttributeMapperList) DeepCopyInto(out *UserAttributeMapperList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserAttributeMapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAttributeMapperList.
func (in *UserAttributeMapperList) DeepCopy() *UserAttributeMapperList {
	if in == nil {
		return nil
	}
	out := new(UserAttributeMapperList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserAttributeMapperList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAttributeMapperObservation) DeepCopyInto(out *UserAttributeMapperObservation) {
	*out = *in
	if in.AlwaysReadValueFromLdap != nil {
		in, out := &in.AlwaysReadValueFromLdap, &out.AlwaysReadValueFromLdap
		*out = new(bool)
		**out = **in
	}
	if in.AttributeDefaultValue != nil {
		in, out := &in.AttributeDefaultValue, &out.AttributeDefaultValue
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsBinaryAttribute != nil {
		in, out := &in.IsBinaryAttribute, &out.IsBinaryAttribute
		*out = new(bool)
		**out = **in
	}
	if in.IsMandatoryInLdap != nil {
		in, out := &in.IsMandatoryInLdap, &out.IsMandatoryInLdap
		*out = new(bool)
		**out = **in
	}
	if in.LdapAttribute != nil {
		in, out := &in.LdapAttribute, &out.LdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.LdapUserFederationID != nil {
		in, out := &in.LdapUserFederationID, &out.LdapUserFederationID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.UserModelAttribute != nil {
		in, out := &in.UserModelAttribute, &out.UserModelAttribute
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAttributeMapperObservation.
func (in *UserAttributeMapperObservation) DeepCopy() *UserAttributeMapperObservation {
	if in == nil {
		return nil
	}
	out := new(UserAttributeMapperObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAttributeMapperParameters) DeepCopyInto(out *UserAttributeMapperParameters) {
	*out = *in
	if in.AlwaysReadValueFromLdap != nil {
		in, out := &in.AlwaysReadValueFromLdap, &out.AlwaysReadValueFromLdap
		*out = new(bool)
		**out = **in
	}
	if in.AttributeDefaultValue != nil {
		in, out := &in.AttributeDefaultValue, &out.AttributeDefaultValue
		*out = new(string)
		**out = **in
	}
	if in.IsBinaryAttribute != nil {
		in, out := &in.IsBinaryAttribute, &out.IsBinaryAttribute
		*out = new(bool)
		**out = **in
	}
	if in.IsMandatoryInLdap != nil {
		in, out := &in.IsMandatoryInLdap, &out.IsMandatoryInLdap
		*out = new(bool)
		**out = **in
	}
	if in.LdapAttribute != nil {
		in, out := &in.LdapAttribute, &out.LdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.LdapUserFederationID != nil {
		in, out := &in.LdapUserFederationID, &out.LdapUserFederationID
		*out = new(string)
		**out = **in
	}
	if in.LdapUserFederationIDRef != nil {
		in, out := &in.LdapUserFederationIDRef, &out.LdapUserFederationIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LdapUserFederationIDSelector != nil {
		in, out := &in.LdapUserFederationIDSelector, &out.LdapUserFederationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
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
	if in.UserModelAttribute != nil {
		in, out := &in.UserModelAttribute, &out.UserModelAttribute
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAttributeMapperParameters.
func (in *UserAttributeMapperParameters) DeepCopy() *UserAttributeMapperParameters {
	if in == nil {
		return nil
	}
	out := new(UserAttributeMapperParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAttributeMapperSpec) DeepCopyInto(out *UserAttributeMapperSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAttributeMapperSpec.
func (in *UserAttributeMapperSpec) DeepCopy() *UserAttributeMapperSpec {
	if in == nil {
		return nil
	}
	out := new(UserAttributeMapperSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAttributeMapperStatus) DeepCopyInto(out *UserAttributeMapperStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAttributeMapperStatus.
func (in *UserAttributeMapperStatus) DeepCopy() *UserAttributeMapperStatus {
	if in == nil {
		return nil
	}
	out := new(UserAttributeMapperStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserFederation) DeepCopyInto(out *UserFederation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserFederation.
func (in *UserFederation) DeepCopy() *UserFederation {
	if in == nil {
		return nil
	}
	out := new(UserFederation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserFederation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserFederationInitParameters) DeepCopyInto(out *UserFederationInitParameters) {
	*out = *in
	if in.BatchSizeForSync != nil {
		in, out := &in.BatchSizeForSync, &out.BatchSizeForSync
		*out = new(float64)
		**out = **in
	}
	if in.BindDn != nil {
		in, out := &in.BindDn, &out.BindDn
		*out = new(string)
		**out = **in
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = make([]CacheInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChangedSyncPeriod != nil {
		in, out := &in.ChangedSyncPeriod, &out.ChangedSyncPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ConnectionTimeout != nil {
		in, out := &in.ConnectionTimeout, &out.ConnectionTimeout
		*out = new(string)
		**out = **in
	}
	if in.ConnectionURL != nil {
		in, out := &in.ConnectionURL, &out.ConnectionURL
		*out = new(string)
		**out = **in
	}
	if in.CustomUserSearchFilter != nil {
		in, out := &in.CustomUserSearchFilter, &out.CustomUserSearchFilter
		*out = new(string)
		**out = **in
	}
	if in.DeleteDefaultMappers != nil {
		in, out := &in.DeleteDefaultMappers, &out.DeleteDefaultMappers
		*out = new(bool)
		**out = **in
	}
	if in.EditMode != nil {
		in, out := &in.EditMode, &out.EditMode
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.FullSyncPeriod != nil {
		in, out := &in.FullSyncPeriod, &out.FullSyncPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ImportEnabled != nil {
		in, out := &in.ImportEnabled, &out.ImportEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Kerberos != nil {
		in, out := &in.Kerberos, &out.Kerberos
		*out = make([]KerberosInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Pagination != nil {
		in, out := &in.Pagination, &out.Pagination
		*out = new(bool)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.RdnLdapAttribute != nil {
		in, out := &in.RdnLdapAttribute, &out.RdnLdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.ReadTimeout != nil {
		in, out := &in.ReadTimeout, &out.ReadTimeout
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
	if in.SearchScope != nil {
		in, out := &in.SearchScope, &out.SearchScope
		*out = new(string)
		**out = **in
	}
	if in.StartTLS != nil {
		in, out := &in.StartTLS, &out.StartTLS
		*out = new(bool)
		**out = **in
	}
	if in.SyncRegistrations != nil {
		in, out := &in.SyncRegistrations, &out.SyncRegistrations
		*out = new(bool)
		**out = **in
	}
	if in.TrustEmail != nil {
		in, out := &in.TrustEmail, &out.TrustEmail
		*out = new(bool)
		**out = **in
	}
	if in.UUIDLdapAttribute != nil {
		in, out := &in.UUIDLdapAttribute, &out.UUIDLdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.UsePasswordModifyExtendedOp != nil {
		in, out := &in.UsePasswordModifyExtendedOp, &out.UsePasswordModifyExtendedOp
		*out = new(bool)
		**out = **in
	}
	if in.UseTruststoreSpi != nil {
		in, out := &in.UseTruststoreSpi, &out.UseTruststoreSpi
		*out = new(string)
		**out = **in
	}
	if in.UserObjectClasses != nil {
		in, out := &in.UserObjectClasses, &out.UserObjectClasses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UsernameLdapAttribute != nil {
		in, out := &in.UsernameLdapAttribute, &out.UsernameLdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.UsersDn != nil {
		in, out := &in.UsersDn, &out.UsersDn
		*out = new(string)
		**out = **in
	}
	if in.ValidatePasswordPolicy != nil {
		in, out := &in.ValidatePasswordPolicy, &out.ValidatePasswordPolicy
		*out = new(bool)
		**out = **in
	}
	if in.Vendor != nil {
		in, out := &in.Vendor, &out.Vendor
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserFederationInitParameters.
func (in *UserFederationInitParameters) DeepCopy() *UserFederationInitParameters {
	if in == nil {
		return nil
	}
	out := new(UserFederationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserFederationList) DeepCopyInto(out *UserFederationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserFederation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserFederationList.
func (in *UserFederationList) DeepCopy() *UserFederationList {
	if in == nil {
		return nil
	}
	out := new(UserFederationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserFederationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserFederationObservation) DeepCopyInto(out *UserFederationObservation) {
	*out = *in
	if in.BatchSizeForSync != nil {
		in, out := &in.BatchSizeForSync, &out.BatchSizeForSync
		*out = new(float64)
		**out = **in
	}
	if in.BindDn != nil {
		in, out := &in.BindDn, &out.BindDn
		*out = new(string)
		**out = **in
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = make([]CacheObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChangedSyncPeriod != nil {
		in, out := &in.ChangedSyncPeriod, &out.ChangedSyncPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ConnectionTimeout != nil {
		in, out := &in.ConnectionTimeout, &out.ConnectionTimeout
		*out = new(string)
		**out = **in
	}
	if in.ConnectionURL != nil {
		in, out := &in.ConnectionURL, &out.ConnectionURL
		*out = new(string)
		**out = **in
	}
	if in.CustomUserSearchFilter != nil {
		in, out := &in.CustomUserSearchFilter, &out.CustomUserSearchFilter
		*out = new(string)
		**out = **in
	}
	if in.DeleteDefaultMappers != nil {
		in, out := &in.DeleteDefaultMappers, &out.DeleteDefaultMappers
		*out = new(bool)
		**out = **in
	}
	if in.EditMode != nil {
		in, out := &in.EditMode, &out.EditMode
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.FullSyncPeriod != nil {
		in, out := &in.FullSyncPeriod, &out.FullSyncPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ImportEnabled != nil {
		in, out := &in.ImportEnabled, &out.ImportEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Kerberos != nil {
		in, out := &in.Kerberos, &out.Kerberos
		*out = make([]KerberosObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Pagination != nil {
		in, out := &in.Pagination, &out.Pagination
		*out = new(bool)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.RdnLdapAttribute != nil {
		in, out := &in.RdnLdapAttribute, &out.RdnLdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.ReadTimeout != nil {
		in, out := &in.ReadTimeout, &out.ReadTimeout
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.SearchScope != nil {
		in, out := &in.SearchScope, &out.SearchScope
		*out = new(string)
		**out = **in
	}
	if in.StartTLS != nil {
		in, out := &in.StartTLS, &out.StartTLS
		*out = new(bool)
		**out = **in
	}
	if in.SyncRegistrations != nil {
		in, out := &in.SyncRegistrations, &out.SyncRegistrations
		*out = new(bool)
		**out = **in
	}
	if in.TrustEmail != nil {
		in, out := &in.TrustEmail, &out.TrustEmail
		*out = new(bool)
		**out = **in
	}
	if in.UUIDLdapAttribute != nil {
		in, out := &in.UUIDLdapAttribute, &out.UUIDLdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.UsePasswordModifyExtendedOp != nil {
		in, out := &in.UsePasswordModifyExtendedOp, &out.UsePasswordModifyExtendedOp
		*out = new(bool)
		**out = **in
	}
	if in.UseTruststoreSpi != nil {
		in, out := &in.UseTruststoreSpi, &out.UseTruststoreSpi
		*out = new(string)
		**out = **in
	}
	if in.UserObjectClasses != nil {
		in, out := &in.UserObjectClasses, &out.UserObjectClasses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UsernameLdapAttribute != nil {
		in, out := &in.UsernameLdapAttribute, &out.UsernameLdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.UsersDn != nil {
		in, out := &in.UsersDn, &out.UsersDn
		*out = new(string)
		**out = **in
	}
	if in.ValidatePasswordPolicy != nil {
		in, out := &in.ValidatePasswordPolicy, &out.ValidatePasswordPolicy
		*out = new(bool)
		**out = **in
	}
	if in.Vendor != nil {
		in, out := &in.Vendor, &out.Vendor
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserFederationObservation.
func (in *UserFederationObservation) DeepCopy() *UserFederationObservation {
	if in == nil {
		return nil
	}
	out := new(UserFederationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserFederationParameters) DeepCopyInto(out *UserFederationParameters) {
	*out = *in
	if in.BatchSizeForSync != nil {
		in, out := &in.BatchSizeForSync, &out.BatchSizeForSync
		*out = new(float64)
		**out = **in
	}
	if in.BindCredentialSecretRef != nil {
		in, out := &in.BindCredentialSecretRef, &out.BindCredentialSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.BindDn != nil {
		in, out := &in.BindDn, &out.BindDn
		*out = new(string)
		**out = **in
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = make([]CacheParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChangedSyncPeriod != nil {
		in, out := &in.ChangedSyncPeriod, &out.ChangedSyncPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ConnectionTimeout != nil {
		in, out := &in.ConnectionTimeout, &out.ConnectionTimeout
		*out = new(string)
		**out = **in
	}
	if in.ConnectionURL != nil {
		in, out := &in.ConnectionURL, &out.ConnectionURL
		*out = new(string)
		**out = **in
	}
	if in.CustomUserSearchFilter != nil {
		in, out := &in.CustomUserSearchFilter, &out.CustomUserSearchFilter
		*out = new(string)
		**out = **in
	}
	if in.DeleteDefaultMappers != nil {
		in, out := &in.DeleteDefaultMappers, &out.DeleteDefaultMappers
		*out = new(bool)
		**out = **in
	}
	if in.EditMode != nil {
		in, out := &in.EditMode, &out.EditMode
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.FullSyncPeriod != nil {
		in, out := &in.FullSyncPeriod, &out.FullSyncPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ImportEnabled != nil {
		in, out := &in.ImportEnabled, &out.ImportEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Kerberos != nil {
		in, out := &in.Kerberos, &out.Kerberos
		*out = make([]KerberosParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Pagination != nil {
		in, out := &in.Pagination, &out.Pagination
		*out = new(bool)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.RdnLdapAttribute != nil {
		in, out := &in.RdnLdapAttribute, &out.RdnLdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.ReadTimeout != nil {
		in, out := &in.ReadTimeout, &out.ReadTimeout
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
	if in.SearchScope != nil {
		in, out := &in.SearchScope, &out.SearchScope
		*out = new(string)
		**out = **in
	}
	if in.StartTLS != nil {
		in, out := &in.StartTLS, &out.StartTLS
		*out = new(bool)
		**out = **in
	}
	if in.SyncRegistrations != nil {
		in, out := &in.SyncRegistrations, &out.SyncRegistrations
		*out = new(bool)
		**out = **in
	}
	if in.TrustEmail != nil {
		in, out := &in.TrustEmail, &out.TrustEmail
		*out = new(bool)
		**out = **in
	}
	if in.UUIDLdapAttribute != nil {
		in, out := &in.UUIDLdapAttribute, &out.UUIDLdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.UsePasswordModifyExtendedOp != nil {
		in, out := &in.UsePasswordModifyExtendedOp, &out.UsePasswordModifyExtendedOp
		*out = new(bool)
		**out = **in
	}
	if in.UseTruststoreSpi != nil {
		in, out := &in.UseTruststoreSpi, &out.UseTruststoreSpi
		*out = new(string)
		**out = **in
	}
	if in.UserObjectClasses != nil {
		in, out := &in.UserObjectClasses, &out.UserObjectClasses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UsernameLdapAttribute != nil {
		in, out := &in.UsernameLdapAttribute, &out.UsernameLdapAttribute
		*out = new(string)
		**out = **in
	}
	if in.UsersDn != nil {
		in, out := &in.UsersDn, &out.UsersDn
		*out = new(string)
		**out = **in
	}
	if in.ValidatePasswordPolicy != nil {
		in, out := &in.ValidatePasswordPolicy, &out.ValidatePasswordPolicy
		*out = new(bool)
		**out = **in
	}
	if in.Vendor != nil {
		in, out := &in.Vendor, &out.Vendor
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserFederationParameters.
func (in *UserFederationParameters) DeepCopy() *UserFederationParameters {
	if in == nil {
		return nil
	}
	out := new(UserFederationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserFederationSpec) DeepCopyInto(out *UserFederationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserFederationSpec.
func (in *UserFederationSpec) DeepCopy() *UserFederationSpec {
	if in == nil {
		return nil
	}
	out := new(UserFederationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserFederationStatus) DeepCopyInto(out *UserFederationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserFederationStatus.
func (in *UserFederationStatus) DeepCopy() *UserFederationStatus {
	if in == nil {
		return nil
	}
	out := new(UserFederationStatus)
	in.DeepCopyInto(out)
	return out
}
