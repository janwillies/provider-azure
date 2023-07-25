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

type MSSQLVirtualNetworkRuleInitParameters struct {

	// Create the virtual network rule before the subnet has the virtual network service endpoint enabled. Defaults to false.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`
}

type MSSQLVirtualNetworkRuleObservation struct {

	// The ID of the SQL virtual network rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Create the virtual network rule before the subnet has the virtual network service endpoint enabled. Defaults to false.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The resource ID of the SQL Server to which this SQL virtual network rule will be applied. Changing this forces a new resource to be created.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// The ID of the subnet from which the SQL server will accept communications.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type MSSQLVirtualNetworkRuleParameters struct {

	// Create the virtual network rule before the subnet has the virtual network service endpoint enabled. Defaults to false.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The resource ID of the SQL Server to which this SQL virtual network rule will be applied. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLServer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a MSSQLServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// The ID of the subnet from which the SQL server will accept communications.
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
}

// MSSQLVirtualNetworkRuleSpec defines the desired state of MSSQLVirtualNetworkRule
type MSSQLVirtualNetworkRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLVirtualNetworkRuleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider MSSQLVirtualNetworkRuleInitParameters `json:"initProvider,omitempty"`
}

// MSSQLVirtualNetworkRuleStatus defines the observed state of MSSQLVirtualNetworkRule.
type MSSQLVirtualNetworkRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLVirtualNetworkRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLVirtualNetworkRule is the Schema for the MSSQLVirtualNetworkRules API. Manages an Azure SQL Virtual Network Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLVirtualNetworkRuleSpec   `json:"spec"`
	Status            MSSQLVirtualNetworkRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLVirtualNetworkRuleList contains a list of MSSQLVirtualNetworkRules
type MSSQLVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLVirtualNetworkRule `json:"items"`
}

// Repository type metadata.
var (
	MSSQLVirtualNetworkRule_Kind             = "MSSQLVirtualNetworkRule"
	MSSQLVirtualNetworkRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLVirtualNetworkRule_Kind}.String()
	MSSQLVirtualNetworkRule_KindAPIVersion   = MSSQLVirtualNetworkRule_Kind + "." + CRDGroupVersion.String()
	MSSQLVirtualNetworkRule_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLVirtualNetworkRule_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLVirtualNetworkRule{}, &MSSQLVirtualNetworkRuleList{})
}
