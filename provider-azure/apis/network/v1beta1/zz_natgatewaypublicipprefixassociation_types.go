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

type NATGatewayPublicIPPrefixAssociationObservation struct {

	// The  ID of the Association between the NAT Gateway and the Public IP Prefix.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NATGatewayPublicIPPrefixAssociationParameters struct {

	// The ID of the NAT Gateway. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=NATGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Reference to a NATGateway to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDRef *v1.Reference `json:"natGatewayIdRef,omitempty" tf:"-"`

	// Selector for a NATGateway to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDSelector *v1.Selector `json:"natGatewayIdSelector,omitempty" tf:"-"`

	// The ID of the Public IP Prefix which this NAT Gateway which should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=PublicIPPrefix
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`

	// Reference to a PublicIPPrefix to populate publicIpPrefixId.
	// +kubebuilder:validation:Optional
	PublicIPPrefixIDRef *v1.Reference `json:"publicIpPrefixIdRef,omitempty" tf:"-"`

	// Selector for a PublicIPPrefix to populate publicIpPrefixId.
	// +kubebuilder:validation:Optional
	PublicIPPrefixIDSelector *v1.Selector `json:"publicIpPrefixIdSelector,omitempty" tf:"-"`
}

// NATGatewayPublicIPPrefixAssociationSpec defines the desired state of NATGatewayPublicIPPrefixAssociation
type NATGatewayPublicIPPrefixAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NATGatewayPublicIPPrefixAssociationParameters `json:"forProvider"`
}

// NATGatewayPublicIPPrefixAssociationStatus defines the observed state of NATGatewayPublicIPPrefixAssociation.
type NATGatewayPublicIPPrefixAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NATGatewayPublicIPPrefixAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NATGatewayPublicIPPrefixAssociation is the Schema for the NATGatewayPublicIPPrefixAssociations API. Manages the association between a NAT Gateway and a Public IP Prefix.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NATGatewayPublicIPPrefixAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NATGatewayPublicIPPrefixAssociationSpec   `json:"spec"`
	Status            NATGatewayPublicIPPrefixAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NATGatewayPublicIPPrefixAssociationList contains a list of NATGatewayPublicIPPrefixAssociations
type NATGatewayPublicIPPrefixAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NATGatewayPublicIPPrefixAssociation `json:"items"`
}

// Repository type metadata.
var (
	NATGatewayPublicIPPrefixAssociation_Kind             = "NATGatewayPublicIPPrefixAssociation"
	NATGatewayPublicIPPrefixAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NATGatewayPublicIPPrefixAssociation_Kind}.String()
	NATGatewayPublicIPPrefixAssociation_KindAPIVersion   = NATGatewayPublicIPPrefixAssociation_Kind + "." + CRDGroupVersion.String()
	NATGatewayPublicIPPrefixAssociation_GroupVersionKind = CRDGroupVersion.WithKind(NATGatewayPublicIPPrefixAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&NATGatewayPublicIPPrefixAssociation{}, &NATGatewayPublicIPPrefixAssociationList{})
}
