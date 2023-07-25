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

type ResourcePolicyRemediationInitParameters struct {

	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
	FailurePercentage *float64 `json:"failurePercentage,omitempty" tf:"failure_percentage,omitempty"`

	// A list of the resource locations that will be remediated.
	LocationFilters []*string `json:"locationFilters,omitempty" tf:"location_filters,omitempty"`

	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	ParallelDeployments *float64 `json:"parallelDeployments,omitempty" tf:"parallel_deployments,omitempty"`

	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// The unique ID for the policy definition reference within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`

	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	ResourceCount *float64 `json:"resourceCount,omitempty" tf:"resource_count,omitempty"`

	// The way that resources to remediate are discovered. Possible values are ExistingNonCompliant, ReEvaluateCompliance. Defaults to ExistingNonCompliant.
	ResourceDiscoveryMode *string `json:"resourceDiscoveryMode,omitempty" tf:"resource_discovery_mode,omitempty"`
}

type ResourcePolicyRemediationObservation struct {

	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
	FailurePercentage *float64 `json:"failurePercentage,omitempty" tf:"failure_percentage,omitempty"`

	// The ID of the Policy Remediation.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of the resource locations that will be remediated.
	LocationFilters []*string `json:"locationFilters,omitempty" tf:"location_filters,omitempty"`

	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	ParallelDeployments *float64 `json:"parallelDeployments,omitempty" tf:"parallel_deployments,omitempty"`

	// The ID of the Policy Assignment that should be remediated.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// The unique ID for the policy definition reference within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`

	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	ResourceCount *float64 `json:"resourceCount,omitempty" tf:"resource_count,omitempty"`

	// The way that resources to remediate are discovered. Possible values are ExistingNonCompliant, ReEvaluateCompliance. Defaults to ExistingNonCompliant.
	ResourceDiscoveryMode *string `json:"resourceDiscoveryMode,omitempty" tf:"resource_discovery_mode,omitempty"`

	// The Resource ID at which the Policy Remediation should be applied. Changing this forces a new resource to be created.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type ResourcePolicyRemediationParameters struct {

	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
	FailurePercentage *float64 `json:"failurePercentage,omitempty" tf:"failure_percentage,omitempty"`

	// A list of the resource locations that will be remediated.
	LocationFilters []*string `json:"locationFilters,omitempty" tf:"location_filters,omitempty"`

	// The name of the Policy Remediation. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	ParallelDeployments *float64 `json:"parallelDeployments,omitempty" tf:"parallel_deployments,omitempty"`

	// The ID of the Policy Assignment that should be remediated.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.ResourceGroupPolicyAssignment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// Reference to a ResourceGroupPolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDRef *v1.Reference `json:"policyAssignmentIdRef,omitempty" tf:"-"`

	// Selector for a ResourceGroupPolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDSelector *v1.Selector `json:"policyAssignmentIdSelector,omitempty" tf:"-"`

	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// The unique ID for the policy definition reference within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`

	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	ResourceCount *float64 `json:"resourceCount,omitempty" tf:"resource_count,omitempty"`

	// The way that resources to remediate are discovered. Possible values are ExistingNonCompliant, ReEvaluateCompliance. Defaults to ExistingNonCompliant.
	ResourceDiscoveryMode *string `json:"resourceDiscoveryMode,omitempty" tf:"resource_discovery_mode,omitempty"`

	// The Resource ID at which the Policy Remediation should be applied. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a VirtualNetwork in network to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork in network to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

// ResourcePolicyRemediationSpec defines the desired state of ResourcePolicyRemediation
type ResourcePolicyRemediationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourcePolicyRemediationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ResourcePolicyRemediationInitParameters `json:"initProvider,omitempty"`
}

// ResourcePolicyRemediationStatus defines the observed state of ResourcePolicyRemediation.
type ResourcePolicyRemediationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourcePolicyRemediationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicyRemediation is the Schema for the ResourcePolicyRemediations API. Manages an Azure Resource Policy Remediation.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourcePolicyRemediation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ResourcePolicyRemediationSpec   `json:"spec"`
	Status ResourcePolicyRemediationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicyRemediationList contains a list of ResourcePolicyRemediations
type ResourcePolicyRemediationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePolicyRemediation `json:"items"`
}

// Repository type metadata.
var (
	ResourcePolicyRemediation_Kind             = "ResourcePolicyRemediation"
	ResourcePolicyRemediation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourcePolicyRemediation_Kind}.String()
	ResourcePolicyRemediation_KindAPIVersion   = ResourcePolicyRemediation_Kind + "." + CRDGroupVersion.String()
	ResourcePolicyRemediation_GroupVersionKind = CRDGroupVersion.WithKind(ResourcePolicyRemediation_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourcePolicyRemediation{}, &ResourcePolicyRemediationList{})
}
