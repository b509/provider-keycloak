/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Memberships.
func (mg *Memberships) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupIDRef,
		Selector:     mg.Spec.ForProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Roles.
func (mg *Roles) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupIDRef,
		Selector:     mg.Spec.ForProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.RoleIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.RoleIdsRefs,
		Selector:      mg.Spec.ForProvider.RoleIdsSelector,
		To: reference.To{
			List:    &v1alpha11.RoleList{},
			Managed: &v1alpha11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleIds")
	}
	mg.Spec.ForProvider.RoleIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.RoleIdsRefs = mrsp.ResolvedReferences

	return nil
}