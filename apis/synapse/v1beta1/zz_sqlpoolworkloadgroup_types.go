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

type SQLPoolWorkloadGroupInitParameters struct {

	// The workload group importance level. Defaults to normal.
	Importance *string `json:"importance,omitempty" tf:"importance,omitempty"`

	// The workload group cap percentage resource.
	MaxResourcePercent *float64 `json:"maxResourcePercent,omitempty" tf:"max_resource_percent,omitempty"`

	// The workload group request maximum grant percentage. Defaults to 3.
	MaxResourcePercentPerRequest *float64 `json:"maxResourcePercentPerRequest,omitempty" tf:"max_resource_percent_per_request,omitempty"`

	// The workload group minimum percentage resource.
	MinResourcePercent *float64 `json:"minResourcePercent,omitempty" tf:"min_resource_percent,omitempty"`

	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest *float64 `json:"minResourcePercentPerRequest,omitempty" tf:"min_resource_percent_per_request,omitempty"`

	// The workload group query execution timeout.
	QueryExecutionTimeoutInSeconds *float64 `json:"queryExecutionTimeoutInSeconds,omitempty" tf:"query_execution_timeout_in_seconds,omitempty"`
}

type SQLPoolWorkloadGroupObservation struct {

	// The ID of the Synapse SQL Pool Workload Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The workload group importance level. Defaults to normal.
	Importance *string `json:"importance,omitempty" tf:"importance,omitempty"`

	// The workload group cap percentage resource.
	MaxResourcePercent *float64 `json:"maxResourcePercent,omitempty" tf:"max_resource_percent,omitempty"`

	// The workload group request maximum grant percentage. Defaults to 3.
	MaxResourcePercentPerRequest *float64 `json:"maxResourcePercentPerRequest,omitempty" tf:"max_resource_percent_per_request,omitempty"`

	// The workload group minimum percentage resource.
	MinResourcePercent *float64 `json:"minResourcePercent,omitempty" tf:"min_resource_percent,omitempty"`

	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest *float64 `json:"minResourcePercentPerRequest,omitempty" tf:"min_resource_percent_per_request,omitempty"`

	// The workload group query execution timeout.
	QueryExecutionTimeoutInSeconds *float64 `json:"queryExecutionTimeoutInSeconds,omitempty" tf:"query_execution_timeout_in_seconds,omitempty"`

	// The ID of the Synapse SQL Pool. Changing this forces a new Synapse SQL Pool Workload Group to be created.
	SQLPoolID *string `json:"sqlPoolId,omitempty" tf:"sql_pool_id,omitempty"`
}

type SQLPoolWorkloadGroupParameters struct {

	// The workload group importance level. Defaults to normal.
	Importance *string `json:"importance,omitempty" tf:"importance,omitempty"`

	// The workload group cap percentage resource.
	MaxResourcePercent *float64 `json:"maxResourcePercent,omitempty" tf:"max_resource_percent,omitempty"`

	// The workload group request maximum grant percentage. Defaults to 3.
	MaxResourcePercentPerRequest *float64 `json:"maxResourcePercentPerRequest,omitempty" tf:"max_resource_percent_per_request,omitempty"`

	// The workload group minimum percentage resource.
	MinResourcePercent *float64 `json:"minResourcePercent,omitempty" tf:"min_resource_percent,omitempty"`

	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest *float64 `json:"minResourcePercentPerRequest,omitempty" tf:"min_resource_percent_per_request,omitempty"`

	// The workload group query execution timeout.
	QueryExecutionTimeoutInSeconds *float64 `json:"queryExecutionTimeoutInSeconds,omitempty" tf:"query_execution_timeout_in_seconds,omitempty"`

	// The ID of the Synapse SQL Pool. Changing this forces a new Synapse SQL Pool Workload Group to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.SQLPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SQLPoolID *string `json:"sqlPoolId,omitempty" tf:"sql_pool_id,omitempty"`

	// Reference to a SQLPool in synapse to populate sqlPoolId.
	// +kubebuilder:validation:Optional
	SQLPoolIDRef *v1.Reference `json:"sqlPoolIdRef,omitempty" tf:"-"`

	// Selector for a SQLPool in synapse to populate sqlPoolId.
	// +kubebuilder:validation:Optional
	SQLPoolIDSelector *v1.Selector `json:"sqlPoolIdSelector,omitempty" tf:"-"`
}

// SQLPoolWorkloadGroupSpec defines the desired state of SQLPoolWorkloadGroup
type SQLPoolWorkloadGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLPoolWorkloadGroupParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SQLPoolWorkloadGroupInitParameters `json:"initProvider,omitempty"`
}

// SQLPoolWorkloadGroupStatus defines the observed state of SQLPoolWorkloadGroup.
type SQLPoolWorkloadGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLPoolWorkloadGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolWorkloadGroup is the Schema for the SQLPoolWorkloadGroups API. Manages a Synapse SQL Pool Workload Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLPoolWorkloadGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxResourcePercent) || has(self.initProvider.maxResourcePercent)",message="maxResourcePercent is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.minResourcePercent) || has(self.initProvider.minResourcePercent)",message="minResourcePercent is a required parameter"
	Spec   SQLPoolWorkloadGroupSpec   `json:"spec"`
	Status SQLPoolWorkloadGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolWorkloadGroupList contains a list of SQLPoolWorkloadGroups
type SQLPoolWorkloadGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLPoolWorkloadGroup `json:"items"`
}

// Repository type metadata.
var (
	SQLPoolWorkloadGroup_Kind             = "SQLPoolWorkloadGroup"
	SQLPoolWorkloadGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLPoolWorkloadGroup_Kind}.String()
	SQLPoolWorkloadGroup_KindAPIVersion   = SQLPoolWorkloadGroup_Kind + "." + CRDGroupVersion.String()
	SQLPoolWorkloadGroup_GroupVersionKind = CRDGroupVersion.WithKind(SQLPoolWorkloadGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLPoolWorkloadGroup{}, &SQLPoolWorkloadGroupList{})
}
