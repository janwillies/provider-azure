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

type ResourceGroupTemplateDeploymentInitParameters struct {

	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are none, requestContent, responseContent and requestContent, responseContent.
	DebugLevel *string `json:"debugLevel,omitempty" tf:"debug_level,omitempty"`

	// The Deployment Mode for this Resource Group Template Deployment. Possible values are Complete (where resources in the Resource Group not specified in the ARM Template will be destroyed) and Incremental (where resources are additive only).
	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`

	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent *string `json:"parametersContent,omitempty" tf:"parameters_content,omitempty"`

	// A mapping of tags which should be assigned to the Resource Group Template Deployment.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with template_spec_version_id.
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content,omitempty"`

	// The ID of the Template Spec Version to deploy. Cannot be specified with template_content.
	TemplateSpecVersionID *string `json:"templateSpecVersionId,omitempty" tf:"template_spec_version_id,omitempty"`
}

type ResourceGroupTemplateDeploymentObservation struct {

	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are none, requestContent, responseContent and requestContent, responseContent.
	DebugLevel *string `json:"debugLevel,omitempty" tf:"debug_level,omitempty"`

	// The Deployment Mode for this Resource Group Template Deployment. Possible values are Complete (where resources in the Resource Group not specified in the ARM Template will be destroyed) and Incremental (where resources are additive only).
	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`

	// The ID of the Resource Group Template Deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The JSON Content of the Outputs of the ARM Template Deployment.
	OutputContent *string `json:"outputContent,omitempty" tf:"output_content,omitempty"`

	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent *string `json:"parametersContent,omitempty" tf:"parameters_content,omitempty"`

	// The name of the Resource Group where the Resource Group Template Deployment should exist. Changing this forces a new Resource Group Template Deployment to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the Resource Group Template Deployment.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with template_spec_version_id.
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content,omitempty"`

	// The ID of the Template Spec Version to deploy. Cannot be specified with template_content.
	TemplateSpecVersionID *string `json:"templateSpecVersionId,omitempty" tf:"template_spec_version_id,omitempty"`
}

type ResourceGroupTemplateDeploymentParameters struct {

	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are none, requestContent, responseContent and requestContent, responseContent.
	DebugLevel *string `json:"debugLevel,omitempty" tf:"debug_level,omitempty"`

	// The Deployment Mode for this Resource Group Template Deployment. Possible values are Complete (where resources in the Resource Group not specified in the ARM Template will be destroyed) and Incremental (where resources are additive only).
	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`

	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent *string `json:"parametersContent,omitempty" tf:"parameters_content,omitempty"`

	// The name of the Resource Group where the Resource Group Template Deployment should exist. Changing this forces a new Resource Group Template Deployment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Resource Group Template Deployment.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with template_spec_version_id.
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content,omitempty"`

	// The ID of the Template Spec Version to deploy. Cannot be specified with template_content.
	TemplateSpecVersionID *string `json:"templateSpecVersionId,omitempty" tf:"template_spec_version_id,omitempty"`
}

// ResourceGroupTemplateDeploymentSpec defines the desired state of ResourceGroupTemplateDeployment
type ResourceGroupTemplateDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupTemplateDeploymentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ResourceGroupTemplateDeploymentInitParameters `json:"initProvider,omitempty"`
}

// ResourceGroupTemplateDeploymentStatus defines the observed state of ResourceGroupTemplateDeployment.
type ResourceGroupTemplateDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupTemplateDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupTemplateDeployment is the Schema for the ResourceGroupTemplateDeployments API. Manages a Resource Group Template Deployment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceGroupTemplateDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.deploymentMode) || has(self.initProvider.deploymentMode)",message="deploymentMode is a required parameter"
	Spec   ResourceGroupTemplateDeploymentSpec   `json:"spec"`
	Status ResourceGroupTemplateDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupTemplateDeploymentList contains a list of ResourceGroupTemplateDeployments
type ResourceGroupTemplateDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroupTemplateDeployment `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroupTemplateDeployment_Kind             = "ResourceGroupTemplateDeployment"
	ResourceGroupTemplateDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroupTemplateDeployment_Kind}.String()
	ResourceGroupTemplateDeployment_KindAPIVersion   = ResourceGroupTemplateDeployment_Kind + "." + CRDGroupVersion.String()
	ResourceGroupTemplateDeployment_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroupTemplateDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroupTemplateDeployment{}, &ResourceGroupTemplateDeploymentList{})
}
