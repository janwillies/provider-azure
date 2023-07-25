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

type DeviceInitParameters struct {

	// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The sku_name is comprised of two segments separated by a hyphen (e.g. TEA_1Node_UPS_Heater-Standard). The first segment of the sku_name defines the name of the SKU, possible values are Gateway, EdgeMR_Mini, EdgeP_Base, EdgeP_High, EdgePR_Base, EdgePR_Base_UPS, GPU, RCA_Large, RCA_Small, RDC, TCA_Large, TCA_Small, TDC, TEA_1Node, TEA_1Node_UPS, TEA_1Node_Heater, TEA_1Node_UPS_Heater, TEA_4Node_Heater, TEA_4Node_UPS_Heater or TMA. The second segment defines the tier of the sku_name, possible values are Standard. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the Databox Edge Device.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DeviceObservation struct {

	// A device_properties block as defined below.
	DeviceProperties []DevicePropertiesObservation `json:"deviceProperties,omitempty" tf:"device_properties,omitempty"`

	// The ID of the Databox Edge Device.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The sku_name is comprised of two segments separated by a hyphen (e.g. TEA_1Node_UPS_Heater-Standard). The first segment of the sku_name defines the name of the SKU, possible values are Gateway, EdgeMR_Mini, EdgeP_Base, EdgeP_High, EdgePR_Base, EdgePR_Base_UPS, GPU, RCA_Large, RCA_Small, RDC, TCA_Large, TCA_Small, TDC, TEA_1Node, TEA_1Node_UPS, TEA_1Node_Heater, TEA_1Node_UPS_Heater, TEA_4Node_Heater, TEA_4Node_UPS_Heater or TMA. The second segment defines the tier of the sku_name, possible values are Standard. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the Databox Edge Device.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DeviceParameters struct {

	// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The sku_name is comprised of two segments separated by a hyphen (e.g. TEA_1Node_UPS_Heater-Standard). The first segment of the sku_name defines the name of the SKU, possible values are Gateway, EdgeMR_Mini, EdgeP_Base, EdgeP_High, EdgePR_Base, EdgePR_Base_UPS, GPU, RCA_Large, RCA_Small, RDC, TCA_Large, TCA_Small, TDC, TEA_1Node, TEA_1Node_UPS, TEA_1Node_Heater, TEA_1Node_UPS_Heater, TEA_4Node_Heater, TEA_4Node_UPS_Heater or TMA. The second segment defines the tier of the sku_name, possible values are Standard. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the Databox Edge Device.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DevicePropertiesInitParameters struct {
}

type DevicePropertiesObservation struct {

	// The Data Box Edge/Gateway device local capacity in MB.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Type of compute roles configured.
	ConfiguredRoleTypes []*string `json:"configuredRoleTypes,omitempty" tf:"configured_role_types,omitempty"`

	// The Data Box Edge/Gateway device culture.
	Culture *string `json:"culture,omitempty" tf:"culture,omitempty"`

	// The device software version number of the device (e.g. 1.2.18105.6).
	HcsVersion *string `json:"hcsVersion,omitempty" tf:"hcs_version,omitempty"`

	// The Data Box Edge/Gateway device model.
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// The number of nodes in the cluster.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The Serial Number of Data Box Edge/Gateway device.
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// The Data Box Edge/Gateway device software version.
	SoftwareVersion *string `json:"softwareVersion,omitempty" tf:"software_version,omitempty"`

	// The status of the Data Box Edge/Gateway device.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The Data Box Edge/Gateway device timezone.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// The type of the Data Box Edge/Gateway device.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DevicePropertiesParameters struct {
}

// DeviceSpec defines the desired state of Device
type DeviceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DeviceInitParameters `json:"initProvider,omitempty"`
}

// DeviceStatus defines the observed state of Device.
type DeviceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Device is the Schema for the Devices API. Manages a Databox Edge Device.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Device struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || has(self.initProvider.skuName)",message="skuName is a required parameter"
	Spec   DeviceSpec   `json:"spec"`
	Status DeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceList contains a list of Devices
type DeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Device `json:"items"`
}

// Repository type metadata.
var (
	Device_Kind             = "Device"
	Device_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Device_Kind}.String()
	Device_KindAPIVersion   = Device_Kind + "." + CRDGroupVersion.String()
	Device_GroupVersionKind = CRDGroupVersion.WithKind(Device_Kind)
)

func init() {
	SchemeBuilder.Register(&Device{}, &DeviceList{})
}
