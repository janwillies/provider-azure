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

type FirewallRuleObservation struct {

	// The ID of the PostgreSQL Firewall Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FirewallRuleParameters struct {

	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	EndIPAddress *string `json:"endIpAddress" tf:"end_ip_address,omitempty"`

	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Server
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a Server to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a Server to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	StartIPAddress *string `json:"startIpAddress" tf:"start_ip_address,omitempty"`
}

// FirewallRuleSpec defines the desired state of FirewallRule
type FirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallRuleParameters `json:"forProvider"`
}

// FirewallRuleStatus defines the observed state of FirewallRule.
type FirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRule is the Schema for the FirewallRules API. Manages a Firewall Rule for a PostgreSQL Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallRuleSpec   `json:"spec"`
	Status            FirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRuleList contains a list of FirewallRules
type FirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallRule `json:"items"`
}

// Repository type metadata.
var (
	FirewallRule_Kind             = "FirewallRule"
	FirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallRule_Kind}.String()
	FirewallRule_KindAPIVersion   = FirewallRule_Kind + "." + CRDGroupVersion.String()
	FirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(FirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallRule{}, &FirewallRuleList{})
}
