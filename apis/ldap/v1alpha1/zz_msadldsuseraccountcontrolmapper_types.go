/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MsadLdsUserAccountControlMapperInitParameters struct {

	// The ID of the LDAP user federation provider to attach this mapper to.
	// The ldap user federation provider to attach this mapper to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/ldap/v1alpha1.UserFederation
	LdapUserFederationID *string `json:"ldapUserFederationId,omitempty" tf:"ldap_user_federation_id,omitempty"`

	// Reference to a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDRef *v1.Reference `json:"ldapUserFederationIdRef,omitempty" tf:"-"`

	// Selector for a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDSelector *v1.Selector `json:"ldapUserFederationIdSelector,omitempty" tf:"-"`

	// Display name of this mapper when displayed in the console.
	// Display name of the mapper when displayed in the console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm that this LDAP mapper will exist in.
	// The realm in which the ldap user federation provider exists.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type MsadLdsUserAccountControlMapperObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the LDAP user federation provider to attach this mapper to.
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationID *string `json:"ldapUserFederationId,omitempty" tf:"ldap_user_federation_id,omitempty"`

	// Display name of this mapper when displayed in the console.
	// Display name of the mapper when displayed in the console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm that this LDAP mapper will exist in.
	// The realm in which the ldap user federation provider exists.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type MsadLdsUserAccountControlMapperParameters struct {

	// The ID of the LDAP user federation provider to attach this mapper to.
	// The ldap user federation provider to attach this mapper to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/ldap/v1alpha1.UserFederation
	// +kubebuilder:validation:Optional
	LdapUserFederationID *string `json:"ldapUserFederationId,omitempty" tf:"ldap_user_federation_id,omitempty"`

	// Reference to a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDRef *v1.Reference `json:"ldapUserFederationIdRef,omitempty" tf:"-"`

	// Selector for a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDSelector *v1.Selector `json:"ldapUserFederationIdSelector,omitempty" tf:"-"`

	// Display name of this mapper when displayed in the console.
	// Display name of the mapper when displayed in the console.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm that this LDAP mapper will exist in.
	// The realm in which the ldap user federation provider exists.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// MsadLdsUserAccountControlMapperSpec defines the desired state of MsadLdsUserAccountControlMapper
type MsadLdsUserAccountControlMapperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MsadLdsUserAccountControlMapperParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MsadLdsUserAccountControlMapperInitParameters `json:"initProvider,omitempty"`
}

// MsadLdsUserAccountControlMapperStatus defines the observed state of MsadLdsUserAccountControlMapper.
type MsadLdsUserAccountControlMapperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MsadLdsUserAccountControlMapperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MsadLdsUserAccountControlMapper is the Schema for the MsadLdsUserAccountControlMappers API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type MsadLdsUserAccountControlMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MsadLdsUserAccountControlMapperSpec   `json:"spec"`
	Status MsadLdsUserAccountControlMapperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MsadLdsUserAccountControlMapperList contains a list of MsadLdsUserAccountControlMappers
type MsadLdsUserAccountControlMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MsadLdsUserAccountControlMapper `json:"items"`
}

// Repository type metadata.
var (
	MsadLdsUserAccountControlMapper_Kind             = "MsadLdsUserAccountControlMapper"
	MsadLdsUserAccountControlMapper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MsadLdsUserAccountControlMapper_Kind}.String()
	MsadLdsUserAccountControlMapper_KindAPIVersion   = MsadLdsUserAccountControlMapper_Kind + "." + CRDGroupVersion.String()
	MsadLdsUserAccountControlMapper_GroupVersionKind = CRDGroupVersion.WithKind(MsadLdsUserAccountControlMapper_Kind)
)

func init() {
	SchemeBuilder.Register(&MsadLdsUserAccountControlMapper{}, &MsadLdsUserAccountControlMapperList{})
}
