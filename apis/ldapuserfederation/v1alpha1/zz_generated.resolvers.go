/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this UserAttributeMapper.
func (mg *UserAttributeMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.ForProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LdapUserFederationID")
	}
	mg.Spec.ForProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LdapUserFederationIDRef = rsp.ResolvedReference

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

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LdapUserFederationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LdapUserFederationIDRef,
		Selector:     mg.Spec.InitProvider.LdapUserFederationIDSelector,
		To: reference.To{
			List:    &UserFederationList{},
			Managed: &UserFederation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LdapUserFederationID")
	}
	mg.Spec.InitProvider.LdapUserFederationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LdapUserFederationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserFederation.
func (mg *UserFederation) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}
