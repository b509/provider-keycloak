/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this Groups.
func (mg *Groups) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Groups.
func (mg *Groups) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Groups.
func (mg *Groups) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Groups.
func (mg *Groups) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this Groups.
func (mg *Groups) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Groups.
func (mg *Groups) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Groups.
func (mg *Groups) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Groups.
func (mg *Groups) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Groups.
func (mg *Groups) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Groups.
func (mg *Groups) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this Groups.
func (mg *Groups) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Groups.
func (mg *Groups) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Permissions.
func (mg *Permissions) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Permissions.
func (mg *Permissions) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Permissions.
func (mg *Permissions) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Permissions.
func (mg *Permissions) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this Permissions.
func (mg *Permissions) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Permissions.
func (mg *Permissions) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Permissions.
func (mg *Permissions) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Permissions.
func (mg *Permissions) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Permissions.
func (mg *Permissions) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Permissions.
func (mg *Permissions) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this Permissions.
func (mg *Permissions) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Permissions.
func (mg *Permissions) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Roles.
func (mg *Roles) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Roles.
func (mg *Roles) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Roles.
func (mg *Roles) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Roles.
func (mg *Roles) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this Roles.
func (mg *Roles) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Roles.
func (mg *Roles) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Roles.
func (mg *Roles) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Roles.
func (mg *Roles) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Roles.
func (mg *Roles) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Roles.
func (mg *Roles) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this Roles.
func (mg *Roles) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Roles.
func (mg *Roles) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this User.
func (mg *User) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this User.
func (mg *User) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this User.
func (mg *User) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this User.
func (mg *User) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this User.
func (mg *User) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this User.
func (mg *User) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this User.
func (mg *User) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this User.
func (mg *User) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this User.
func (mg *User) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this User.
func (mg *User) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this User.
func (mg *User) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this User.
func (mg *User) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
