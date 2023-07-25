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

type DiskPoolInitParameters struct {

	// The Azure Region where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The SKU of the Disk Pool. Possible values are Basic_B1, Standard_S1 and Premium_P1. Changing this forces a new Disk Pool to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the Disk Pool.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of Availability Zones in which this Disk Pool should be located. Changing this forces a new Disk Pool to be created.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type DiskPoolObservation struct {

	// The ID of the Disk Pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU of the Disk Pool. Possible values are Basic_B1, Standard_S1 and Premium_P1. Changing this forces a new Disk Pool to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The ID of the Subnet where the Disk Pool should be created. Changing this forces a new Disk Pool to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A mapping of tags which should be assigned to the Disk Pool.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of Availability Zones in which this Disk Pool should be located. Changing this forces a new Disk Pool to be created.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type DiskPoolParameters struct {

	// The Azure Region where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU of the Disk Pool. Possible values are Basic_B1, Standard_S1 and Premium_P1. Changing this forces a new Disk Pool to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The ID of the Subnet where the Disk Pool should be created. Changing this forces a new Disk Pool to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Disk Pool.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of Availability Zones in which this Disk Pool should be located. Changing this forces a new Disk Pool to be created.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// DiskPoolSpec defines the desired state of DiskPool
type DiskPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskPoolParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DiskPoolInitParameters `json:"initProvider,omitempty"`
}

// DiskPoolStatus defines the observed state of DiskPool.
type DiskPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskPool is the Schema for the DiskPools API. Manages a Disk Pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DiskPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || has(self.initProvider.skuName)",message="skuName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zones) || has(self.initProvider.zones)",message="zones is a required parameter"
	Spec   DiskPoolSpec   `json:"spec"`
	Status DiskPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskPoolList contains a list of DiskPools
type DiskPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskPool `json:"items"`
}

// Repository type metadata.
var (
	DiskPool_Kind             = "DiskPool"
	DiskPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskPool_Kind}.String()
	DiskPool_KindAPIVersion   = DiskPool_Kind + "." + CRDGroupVersion.String()
	DiskPool_GroupVersionKind = CRDGroupVersion.WithKind(DiskPool_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskPool{}, &DiskPoolList{})
}
