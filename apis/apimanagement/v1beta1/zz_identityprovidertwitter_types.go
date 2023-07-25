/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IdentityProviderTwitterInitParameters struct {
}

type IdentityProviderTwitterObservation struct {

	// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The ID of the API Management Twitter Identity Provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type IdentityProviderTwitterParameters struct {

	// App Consumer API key for Twitter.
	APIKeySecretRef v1.SecretKeySelector `json:"apiKeySecretRef" tf:"-"`

	// The Name of the API Management Service where this Twitter Identity Provider should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// App Consumer API secret key for Twitter.
	APISecretKeySecretRef v1.SecretKeySelector `json:"apiSecretKeySecretRef" tf:"-"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IdentityProviderTwitterSpec defines the desired state of IdentityProviderTwitter
type IdentityProviderTwitterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderTwitterParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider IdentityProviderTwitterInitParameters `json:"initProvider,omitempty"`
}

// IdentityProviderTwitterStatus defines the observed state of IdentityProviderTwitter.
type IdentityProviderTwitterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderTwitterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderTwitter is the Schema for the IdentityProviderTwitters API. Manages an API Management Twitter Identity Provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IdentityProviderTwitter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.apiKeySecretRef)",message="apiKeySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.apiSecretKeySecretRef)",message="apiSecretKeySecretRef is a required parameter"
	Spec   IdentityProviderTwitterSpec   `json:"spec"`
	Status IdentityProviderTwitterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderTwitterList contains a list of IdentityProviderTwitters
type IdentityProviderTwitterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderTwitter `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderTwitter_Kind             = "IdentityProviderTwitter"
	IdentityProviderTwitter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderTwitter_Kind}.String()
	IdentityProviderTwitter_KindAPIVersion   = IdentityProviderTwitter_Kind + "." + CRDGroupVersion.String()
	IdentityProviderTwitter_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProviderTwitter_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProviderTwitter{}, &IdentityProviderTwitterList{})
}
