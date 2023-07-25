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

type NetworkInterfaceApplicationSecurityGroupAssociationInitParameters struct {
}

type NetworkInterfaceApplicationSecurityGroupAssociationObservation struct {

	// The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	ApplicationSecurityGroupID *string `json:"applicationSecurityGroupId,omitempty" tf:"application_security_group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`
}

type NetworkInterfaceApplicationSecurityGroupAssociationParameters struct {

	// The ID of the Application Security Group which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=ApplicationSecurityGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupID *string `json:"applicationSecurityGroupId,omitempty" tf:"application_security_group_id,omitempty"`

	// Reference to a ApplicationSecurityGroup to populate applicationSecurityGroupId.
	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupIDRef *v1.Reference `json:"applicationSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationSecurityGroup to populate applicationSecurityGroupId.
	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupIDSelector *v1.Selector `json:"applicationSecurityGroupIdSelector,omitempty" tf:"-"`

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`
}

// NetworkInterfaceApplicationSecurityGroupAssociationSpec defines the desired state of NetworkInterfaceApplicationSecurityGroupAssociation
type NetworkInterfaceApplicationSecurityGroupAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceApplicationSecurityGroupAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider NetworkInterfaceApplicationSecurityGroupAssociationInitParameters `json:"initProvider,omitempty"`
}

// NetworkInterfaceApplicationSecurityGroupAssociationStatus defines the observed state of NetworkInterfaceApplicationSecurityGroupAssociation.
type NetworkInterfaceApplicationSecurityGroupAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceApplicationSecurityGroupAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceApplicationSecurityGroupAssociation is the Schema for the NetworkInterfaceApplicationSecurityGroupAssociations API. Manages the association between a Network Interface and a Application Security Group
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkInterfaceApplicationSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceApplicationSecurityGroupAssociationSpec   `json:"spec"`
	Status            NetworkInterfaceApplicationSecurityGroupAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceApplicationSecurityGroupAssociationList contains a list of NetworkInterfaceApplicationSecurityGroupAssociations
type NetworkInterfaceApplicationSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceApplicationSecurityGroupAssociation `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceApplicationSecurityGroupAssociation_Kind             = "NetworkInterfaceApplicationSecurityGroupAssociation"
	NetworkInterfaceApplicationSecurityGroupAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterfaceApplicationSecurityGroupAssociation_Kind}.String()
	NetworkInterfaceApplicationSecurityGroupAssociation_KindAPIVersion   = NetworkInterfaceApplicationSecurityGroupAssociation_Kind + "." + CRDGroupVersion.String()
	NetworkInterfaceApplicationSecurityGroupAssociation_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterfaceApplicationSecurityGroupAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceApplicationSecurityGroupAssociation{}, &NetworkInterfaceApplicationSecurityGroupAssociationList{})
}
