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

type ManagementGroupSubscriptionAssociationInitParameters struct {
}

type ManagementGroupSubscriptionAssociationObservation struct {

	// The ID of the Management Group Subscription Association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Management Group to associate the Subscription with. Changing this forces a new Management to be created.
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// The ID of the Subscription to be associated with the Management Group. Changing this forces a new Management to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type ManagementGroupSubscriptionAssociationParameters struct {

	// The ID of the Management Group to associate the Subscription with. Changing this forces a new Management to be created.
	// +crossplane:generate:reference:type=ManagementGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// Reference to a ManagementGroup to populate managementGroupId.
	// +kubebuilder:validation:Optional
	ManagementGroupIDRef *v1.Reference `json:"managementGroupIdRef,omitempty" tf:"-"`

	// Selector for a ManagementGroup to populate managementGroupId.
	// +kubebuilder:validation:Optional
	ManagementGroupIDSelector *v1.Selector `json:"managementGroupIdSelector,omitempty" tf:"-"`

	// The ID of the Subscription to be associated with the Management Group. Changing this forces a new Management to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.Subscription
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// Reference to a Subscription in azure to populate subscriptionId.
	// +kubebuilder:validation:Optional
	SubscriptionIDRef *v1.Reference `json:"subscriptionIdRef,omitempty" tf:"-"`

	// Selector for a Subscription in azure to populate subscriptionId.
	// +kubebuilder:validation:Optional
	SubscriptionIDSelector *v1.Selector `json:"subscriptionIdSelector,omitempty" tf:"-"`
}

// ManagementGroupSubscriptionAssociationSpec defines the desired state of ManagementGroupSubscriptionAssociation
type ManagementGroupSubscriptionAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementGroupSubscriptionAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ManagementGroupSubscriptionAssociationInitParameters `json:"initProvider,omitempty"`
}

// ManagementGroupSubscriptionAssociationStatus defines the observed state of ManagementGroupSubscriptionAssociation.
type ManagementGroupSubscriptionAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementGroupSubscriptionAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementGroupSubscriptionAssociation is the Schema for the ManagementGroupSubscriptionAssociations API. Manages a Management Group Subscription Association.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagementGroupSubscriptionAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementGroupSubscriptionAssociationSpec   `json:"spec"`
	Status            ManagementGroupSubscriptionAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementGroupSubscriptionAssociationList contains a list of ManagementGroupSubscriptionAssociations
type ManagementGroupSubscriptionAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementGroupSubscriptionAssociation `json:"items"`
}

// Repository type metadata.
var (
	ManagementGroupSubscriptionAssociation_Kind             = "ManagementGroupSubscriptionAssociation"
	ManagementGroupSubscriptionAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementGroupSubscriptionAssociation_Kind}.String()
	ManagementGroupSubscriptionAssociation_KindAPIVersion   = ManagementGroupSubscriptionAssociation_Kind + "." + CRDGroupVersion.String()
	ManagementGroupSubscriptionAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ManagementGroupSubscriptionAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementGroupSubscriptionAssociation{}, &ManagementGroupSubscriptionAssociationList{})
}
