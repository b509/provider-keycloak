/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1alpha1

import (
	"context"

	common "github.com/crossplane-contrib/provider-keycloak/config/common"
	apisresolver "github.com/crossplane-contrib/provider-keycloak/internal/apis"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Bindings.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *Bindings) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BrowserFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.ForProvider.BrowserFlowRef,
			Selector:     mg.Spec.ForProvider.BrowserFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BrowserFlow")
	}
	mg.Spec.ForProvider.BrowserFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BrowserFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientAuthenticationFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.ForProvider.ClientAuthenticationFlowRef,
			Selector:     mg.Spec.ForProvider.ClientAuthenticationFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientAuthenticationFlow")
	}
	mg.Spec.ForProvider.ClientAuthenticationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientAuthenticationFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DirectGrantFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.ForProvider.DirectGrantFlowRef,
			Selector:     mg.Spec.ForProvider.DirectGrantFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DirectGrantFlow")
	}
	mg.Spec.ForProvider.DirectGrantFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DirectGrantFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DockerAuthenticationFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.ForProvider.DockerAuthenticationFlowRef,
			Selector:     mg.Spec.ForProvider.DockerAuthenticationFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DockerAuthenticationFlow")
	}
	mg.Spec.ForProvider.DockerAuthenticationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DockerAuthenticationFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RegistrationFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.ForProvider.RegistrationFlowRef,
			Selector:     mg.Spec.ForProvider.RegistrationFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RegistrationFlow")
	}
	mg.Spec.ForProvider.RegistrationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RegistrationFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResetCredentialsFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.ForProvider.ResetCredentialsFlowRef,
			Selector:     mg.Spec.ForProvider.ResetCredentialsFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResetCredentialsFlow")
	}
	mg.Spec.ForProvider.ResetCredentialsFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResetCredentialsFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BrowserFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.InitProvider.BrowserFlowRef,
			Selector:     mg.Spec.InitProvider.BrowserFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BrowserFlow")
	}
	mg.Spec.InitProvider.BrowserFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BrowserFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientAuthenticationFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.InitProvider.ClientAuthenticationFlowRef,
			Selector:     mg.Spec.InitProvider.ClientAuthenticationFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientAuthenticationFlow")
	}
	mg.Spec.InitProvider.ClientAuthenticationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientAuthenticationFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DirectGrantFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.InitProvider.DirectGrantFlowRef,
			Selector:     mg.Spec.InitProvider.DirectGrantFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DirectGrantFlow")
	}
	mg.Spec.InitProvider.DirectGrantFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DirectGrantFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DockerAuthenticationFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.InitProvider.DockerAuthenticationFlowRef,
			Selector:     mg.Spec.InitProvider.DockerAuthenticationFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DockerAuthenticationFlow")
	}
	mg.Spec.InitProvider.DockerAuthenticationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DockerAuthenticationFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RegistrationFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.InitProvider.RegistrationFlowRef,
			Selector:     mg.Spec.InitProvider.RegistrationFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RegistrationFlow")
	}
	mg.Spec.InitProvider.RegistrationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RegistrationFlowRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResetCredentialsFlow),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.InitProvider.ResetCredentialsFlowRef,
			Selector:     mg.Spec.InitProvider.ResetCredentialsFlowSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResetCredentialsFlow")
	}
	mg.Spec.InitProvider.ResetCredentialsFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResetCredentialsFlowRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Execution.
func (mg *Execution) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ParentFlowAlias),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.ForProvider.ParentFlowAliasRef,
			Selector:     mg.Spec.ForProvider.ParentFlowAliasSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ParentFlowAlias")
	}
	mg.Spec.ForProvider.ParentFlowAlias = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentFlowAliasRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ParentFlowAlias),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.InitProvider.ParentFlowAliasRef,
			Selector:     mg.Spec.InitProvider.ParentFlowAliasSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ParentFlowAlias")
	}
	mg.Spec.InitProvider.ParentFlowAlias = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ParentFlowAliasRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ExecutionConfig.
func (mg *ExecutionConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Execution", "ExecutionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExecutionID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ExecutionIDRef,
			Selector:     mg.Spec.ForProvider.ExecutionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExecutionID")
	}
	mg.Spec.ForProvider.ExecutionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ExecutionIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Execution", "ExecutionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ExecutionID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ExecutionIDRef,
			Selector:     mg.Spec.InitProvider.ExecutionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ExecutionID")
	}
	mg.Spec.InitProvider.ExecutionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ExecutionIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Flow.
func (mg *Flow) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subflow.
func (mg *Subflow) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ParentFlowAlias),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.ForProvider.ParentFlowAliasRef,
			Selector:     mg.Spec.ForProvider.ParentFlowAliasSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ParentFlowAlias")
	}
	mg.Spec.ForProvider.ParentFlowAlias = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentFlowAliasRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("authenticationflow.keycloak.crossplane.io", "v1alpha1", "Flow", "FlowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ParentFlowAlias),
			Extract:      common.AuthenticationFlowAliasExtractor(),
			Reference:    mg.Spec.InitProvider.ParentFlowAliasRef,
			Selector:     mg.Spec.InitProvider.ParentFlowAliasSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ParentFlowAlias")
	}
	mg.Spec.InitProvider.ParentFlowAlias = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ParentFlowAliasRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}
