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

type GroupsInitParameters struct {

	// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to true.
	Exhaustive *bool `json:"exhaustive,omitempty" tf:"exhaustive,omitempty"`

	// A list of group IDs that the user is member of.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/group/v1alpha1.Group
	// +listType=set
	GroupIds []*string `json:"groupIds,omitempty" tf:"group_ids,omitempty"`

	// References to Group in group to populate groupIds.
	// +kubebuilder:validation:Optional
	GroupIdsRefs []v1.Reference `json:"groupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Group in group to populate groupIds.
	// +kubebuilder:validation:Optional
	GroupIdsSelector *v1.Selector `json:"groupIdsSelector,omitempty" tf:"-"`

	// The realm this group exists in.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The ID of the user this resource should manage groups for.
	// +crossplane:generate:reference:type=User
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

type GroupsObservation struct {

	// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to true.
	Exhaustive *bool `json:"exhaustive,omitempty" tf:"exhaustive,omitempty"`

	// A list of group IDs that the user is member of.
	// +listType=set
	GroupIds []*string `json:"groupIds,omitempty" tf:"group_ids,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The realm this group exists in.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// The ID of the user this resource should manage groups for.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type GroupsParameters struct {

	// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to true.
	// +kubebuilder:validation:Optional
	Exhaustive *bool `json:"exhaustive,omitempty" tf:"exhaustive,omitempty"`

	// A list of group IDs that the user is member of.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/group/v1alpha1.Group
	// +kubebuilder:validation:Optional
	// +listType=set
	GroupIds []*string `json:"groupIds,omitempty" tf:"group_ids,omitempty"`

	// References to Group in group to populate groupIds.
	// +kubebuilder:validation:Optional
	GroupIdsRefs []v1.Reference `json:"groupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Group in group to populate groupIds.
	// +kubebuilder:validation:Optional
	GroupIdsSelector *v1.Selector `json:"groupIdsSelector,omitempty" tf:"-"`

	// The realm this group exists in.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The ID of the user this resource should manage groups for.
	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// GroupsSpec defines the desired state of Groups
type GroupsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupsParameters `json:"forProvider"`
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
	InitProvider GroupsInitParameters `json:"initProvider,omitempty"`
}

// GroupsStatus defines the observed state of Groups.
type GroupsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Groups is the Schema for the Groupss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type Groups struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupsSpec   `json:"spec"`
	Status            GroupsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupsList contains a list of Groupss
type GroupsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Groups `json:"items"`
}

// Repository type metadata.
var (
	Groups_Kind             = "Groups"
	Groups_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Groups_Kind}.String()
	Groups_KindAPIVersion   = Groups_Kind + "." + CRDGroupVersion.String()
	Groups_GroupVersionKind = CRDGroupVersion.WithKind(Groups_Kind)
)

func init() {
	SchemeBuilder.Register(&Groups{}, &GroupsList{})
}
