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

type ManagedHardwareSecurityModuleObservation struct {

	// The URI of the Key Vault Managed Hardware Security Module, used for performing operations on keys.
	HSMURI *string `json:"hsmUri,omitempty" tf:"hsm_uri,omitempty"`

	// The Key Vault Secret Managed Hardware Security Module ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagedHardwareSecurityModuleParameters struct {

	// Specifies a list of administrators object IDs for the key vault Managed Hardware Security Module. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	AdminObjectIds []*string `json:"adminObjectIds" tf:"admin_object_ids,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Is Purge Protection enabled for this Key Vault Managed Hardware Security Module? Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PurgeProtectionEnabled *bool `json:"purgeProtectionEnabled,omitempty" tf:"purge_protection_enabled,omitempty"`

	// The name of the resource group in which to create the Key Vault Managed Hardware Security Module. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Name of the SKU used for this Key Vault Managed Hardware Security Module. Possible value is Standard_B1. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// The number of days that items should be retained for once soft-deleted. This value can be between 7 and 90 days. Defaults to 90. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SoftDeleteRetentionDays *float64 `json:"softDeleteRetentionDays,omitempty" tf:"soft_delete_retention_days,omitempty"`

	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Azure Active Directory Tenant ID that should be used for authenticating requests to the key vault Managed Hardware Security Module. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

// ManagedHardwareSecurityModuleSpec defines the desired state of ManagedHardwareSecurityModule
type ManagedHardwareSecurityModuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedHardwareSecurityModuleParameters `json:"forProvider"`
}

// ManagedHardwareSecurityModuleStatus defines the observed state of ManagedHardwareSecurityModule.
type ManagedHardwareSecurityModuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedHardwareSecurityModuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedHardwareSecurityModule is the Schema for the ManagedHardwareSecurityModules API. Manages a Key Vault Managed Hardware Security Module.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagedHardwareSecurityModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedHardwareSecurityModuleSpec   `json:"spec"`
	Status            ManagedHardwareSecurityModuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedHardwareSecurityModuleList contains a list of ManagedHardwareSecurityModules
type ManagedHardwareSecurityModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedHardwareSecurityModule `json:"items"`
}

// Repository type metadata.
var (
	ManagedHardwareSecurityModule_Kind             = "ManagedHardwareSecurityModule"
	ManagedHardwareSecurityModule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedHardwareSecurityModule_Kind}.String()
	ManagedHardwareSecurityModule_KindAPIVersion   = ManagedHardwareSecurityModule_Kind + "." + CRDGroupVersion.String()
	ManagedHardwareSecurityModule_GroupVersionKind = CRDGroupVersion.WithKind(ManagedHardwareSecurityModule_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedHardwareSecurityModule{}, &ManagedHardwareSecurityModuleList{})
}
