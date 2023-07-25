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

type SQLPoolWorkloadClassifierInitParameters struct {

	// Specifies the session context value that a request can be classified against.
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// The workload classifier end time for classification. It's of the HH:MM format in UTC time zone.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The workload classifier importance. The allowed values are low, below_normal, normal, above_normal and high.
	Importance *string `json:"importance,omitempty" tf:"importance,omitempty"`

	// Specifies the label value that a request can be classified against.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The workload classifier member name used to classified against.
	MemberName *string `json:"memberName,omitempty" tf:"member_name,omitempty"`

	// The workload classifier start time for classification. It's of the HH:MM format in UTC time zone.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type SQLPoolWorkloadClassifierObservation struct {

	// Specifies the session context value that a request can be classified against.
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// The workload classifier end time for classification. It's of the HH:MM format in UTC time zone.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The ID of the Synapse SQL Pool Workload Classifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The workload classifier importance. The allowed values are low, below_normal, normal, above_normal and high.
	Importance *string `json:"importance,omitempty" tf:"importance,omitempty"`

	// Specifies the label value that a request can be classified against.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The workload classifier member name used to classified against.
	MemberName *string `json:"memberName,omitempty" tf:"member_name,omitempty"`

	// The workload classifier start time for classification. It's of the HH:MM format in UTC time zone.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The ID of the Synapse SQL Pool Workload Group. Changing this forces a new Synapse SQL Pool Workload Classifier to be created.
	WorkloadGroupID *string `json:"workloadGroupId,omitempty" tf:"workload_group_id,omitempty"`
}

type SQLPoolWorkloadClassifierParameters struct {

	// Specifies the session context value that a request can be classified against.
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// The workload classifier end time for classification. It's of the HH:MM format in UTC time zone.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The workload classifier importance. The allowed values are low, below_normal, normal, above_normal and high.
	Importance *string `json:"importance,omitempty" tf:"importance,omitempty"`

	// Specifies the label value that a request can be classified against.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The workload classifier member name used to classified against.
	MemberName *string `json:"memberName,omitempty" tf:"member_name,omitempty"`

	// The workload classifier start time for classification. It's of the HH:MM format in UTC time zone.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The ID of the Synapse SQL Pool Workload Group. Changing this forces a new Synapse SQL Pool Workload Classifier to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.SQLPoolWorkloadGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkloadGroupID *string `json:"workloadGroupId,omitempty" tf:"workload_group_id,omitempty"`

	// Reference to a SQLPoolWorkloadGroup in synapse to populate workloadGroupId.
	// +kubebuilder:validation:Optional
	WorkloadGroupIDRef *v1.Reference `json:"workloadGroupIdRef,omitempty" tf:"-"`

	// Selector for a SQLPoolWorkloadGroup in synapse to populate workloadGroupId.
	// +kubebuilder:validation:Optional
	WorkloadGroupIDSelector *v1.Selector `json:"workloadGroupIdSelector,omitempty" tf:"-"`
}

// SQLPoolWorkloadClassifierSpec defines the desired state of SQLPoolWorkloadClassifier
type SQLPoolWorkloadClassifierSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLPoolWorkloadClassifierParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SQLPoolWorkloadClassifierInitParameters `json:"initProvider,omitempty"`
}

// SQLPoolWorkloadClassifierStatus defines the observed state of SQLPoolWorkloadClassifier.
type SQLPoolWorkloadClassifierStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLPoolWorkloadClassifierObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolWorkloadClassifier is the Schema for the SQLPoolWorkloadClassifiers API. Manages a Synapse SQL Pool Workload Classifier.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLPoolWorkloadClassifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.memberName) || has(self.initProvider.memberName)",message="memberName is a required parameter"
	Spec   SQLPoolWorkloadClassifierSpec   `json:"spec"`
	Status SQLPoolWorkloadClassifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolWorkloadClassifierList contains a list of SQLPoolWorkloadClassifiers
type SQLPoolWorkloadClassifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLPoolWorkloadClassifier `json:"items"`
}

// Repository type metadata.
var (
	SQLPoolWorkloadClassifier_Kind             = "SQLPoolWorkloadClassifier"
	SQLPoolWorkloadClassifier_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLPoolWorkloadClassifier_Kind}.String()
	SQLPoolWorkloadClassifier_KindAPIVersion   = SQLPoolWorkloadClassifier_Kind + "." + CRDGroupVersion.String()
	SQLPoolWorkloadClassifier_GroupVersionKind = CRDGroupVersion.WithKind(SQLPoolWorkloadClassifier_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLPoolWorkloadClassifier{}, &SQLPoolWorkloadClassifierList{})
}
