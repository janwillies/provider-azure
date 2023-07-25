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

type AccessControlInitParameters struct {

	// A action block as defined below.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// A content block as defined below.
	Content []ContentInitParameters `json:"content,omitempty" tf:"content,omitempty"`

	// A trigger block as defined below.
	Trigger []TriggerInitParameters `json:"trigger,omitempty" tf:"trigger,omitempty"`

	// A workflow_management block as defined below.
	WorkflowManagement []WorkflowManagementInitParameters `json:"workflowManagement,omitempty" tf:"workflow_management,omitempty"`
}

type AccessControlObservation struct {

	// A action block as defined below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// A content block as defined below.
	Content []ContentObservation `json:"content,omitempty" tf:"content,omitempty"`

	// A trigger block as defined below.
	Trigger []TriggerObservation `json:"trigger,omitempty" tf:"trigger,omitempty"`

	// A workflow_management block as defined below.
	WorkflowManagement []WorkflowManagementObservation `json:"workflowManagement,omitempty" tf:"workflow_management,omitempty"`
}

type AccessControlParameters struct {

	// A action block as defined below.
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// A content block as defined below.
	Content []ContentParameters `json:"content,omitempty" tf:"content,omitempty"`

	// A trigger block as defined below.
	Trigger []TriggerParameters `json:"trigger,omitempty" tf:"trigger,omitempty"`

	// A workflow_management block as defined below.
	WorkflowManagement []WorkflowManagementParameters `json:"workflowManagement,omitempty" tf:"workflow_management,omitempty"`
}

type ActionInitParameters struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`
}

type ActionObservation struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`
}

type ActionParameters struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`
}

type AppWorkflowInitParameters struct {

	// A access_control block as defined below.
	AccessControl []AccessControlInitParameters `json:"accessControl,omitempty" tf:"access_control,omitempty"`

	// Is the Logic App Workflow enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The ID of the Integration Service Environment to which this Logic App Workflow belongs. Changing this forces a new Logic App Workflow to be created.
	IntegrationServiceEnvironmentID *string `json:"integrationServiceEnvironmentId,omitempty" tf:"integration_service_environment_id,omitempty"`

	// Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the integration account linked by this Logic App Workflow.
	LogicAppIntegrationAccountID *string `json:"logicAppIntegrationAccountId,omitempty" tf:"logic_app_integration_account_id,omitempty"`

	// A map of Key-Value pairs.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a map of Key-Value pairs of the Parameter Definitions to use for this Logic App Workflow. The key is the parameter name, and the value is a JSON encoded string of the parameter definition (see: https://docs.microsoft.com/azure/logic-apps/logic-apps-workflow-definition-language#parameters).
	WorkflowParameters map[string]*string `json:"workflowParameters,omitempty" tf:"workflow_parameters,omitempty"`

	// Specifies the Schema to use for this Logic App Workflow. Defaults to https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#. Changing this forces a new resource to be created.
	WorkflowSchema *string `json:"workflowSchema,omitempty" tf:"workflow_schema,omitempty"`

	// Specifies the version of the Schema used for this Logic App Workflow. Defaults to 1.0.0.0. Changing this forces a new resource to be created.
	WorkflowVersion *string `json:"workflowVersion,omitempty" tf:"workflow_version,omitempty"`
}

type AppWorkflowObservation struct {

	// A access_control block as defined below.
	AccessControl []AccessControlObservation `json:"accessControl,omitempty" tf:"access_control,omitempty"`

	// The Access Endpoint for the Logic App Workflow.
	AccessEndpoint *string `json:"accessEndpoint,omitempty" tf:"access_endpoint,omitempty"`

	// The list of access endpoint IP addresses of connector.
	ConnectorEndpointIPAddresses []*string `json:"connectorEndpointIpAddresses,omitempty" tf:"connector_endpoint_ip_addresses,omitempty"`

	// The list of outgoing IP addresses of connector.
	ConnectorOutboundIPAddresses []*string `json:"connectorOutboundIpAddresses,omitempty" tf:"connector_outbound_ip_addresses,omitempty"`

	// Is the Logic App Workflow enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Logic App Workflow ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The ID of the Integration Service Environment to which this Logic App Workflow belongs. Changing this forces a new Logic App Workflow to be created.
	IntegrationServiceEnvironmentID *string `json:"integrationServiceEnvironmentId,omitempty" tf:"integration_service_environment_id,omitempty"`

	// Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the integration account linked by this Logic App Workflow.
	LogicAppIntegrationAccountID *string `json:"logicAppIntegrationAccountId,omitempty" tf:"logic_app_integration_account_id,omitempty"`

	// A map of Key-Value pairs.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The list of access endpoint IP addresses of workflow.
	WorkflowEndpointIPAddresses []*string `json:"workflowEndpointIpAddresses,omitempty" tf:"workflow_endpoint_ip_addresses,omitempty"`

	// The list of outgoing IP addresses of workflow.
	WorkflowOutboundIPAddresses []*string `json:"workflowOutboundIpAddresses,omitempty" tf:"workflow_outbound_ip_addresses,omitempty"`

	// Specifies a map of Key-Value pairs of the Parameter Definitions to use for this Logic App Workflow. The key is the parameter name, and the value is a JSON encoded string of the parameter definition (see: https://docs.microsoft.com/azure/logic-apps/logic-apps-workflow-definition-language#parameters).
	WorkflowParameters map[string]*string `json:"workflowParameters,omitempty" tf:"workflow_parameters,omitempty"`

	// Specifies the Schema to use for this Logic App Workflow. Defaults to https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#. Changing this forces a new resource to be created.
	WorkflowSchema *string `json:"workflowSchema,omitempty" tf:"workflow_schema,omitempty"`

	// Specifies the version of the Schema used for this Logic App Workflow. Defaults to 1.0.0.0. Changing this forces a new resource to be created.
	WorkflowVersion *string `json:"workflowVersion,omitempty" tf:"workflow_version,omitempty"`
}

type AppWorkflowParameters struct {

	// A access_control block as defined below.
	AccessControl []AccessControlParameters `json:"accessControl,omitempty" tf:"access_control,omitempty"`

	// Is the Logic App Workflow enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// An identity block as defined below.
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The ID of the Integration Service Environment to which this Logic App Workflow belongs. Changing this forces a new Logic App Workflow to be created.
	IntegrationServiceEnvironmentID *string `json:"integrationServiceEnvironmentId,omitempty" tf:"integration_service_environment_id,omitempty"`

	// Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the integration account linked by this Logic App Workflow.
	LogicAppIntegrationAccountID *string `json:"logicAppIntegrationAccountId,omitempty" tf:"logic_app_integration_account_id,omitempty"`

	// A map of Key-Value pairs.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a map of Key-Value pairs of the Parameter Definitions to use for this Logic App Workflow. The key is the parameter name, and the value is a JSON encoded string of the parameter definition (see: https://docs.microsoft.com/azure/logic-apps/logic-apps-workflow-definition-language#parameters).
	WorkflowParameters map[string]*string `json:"workflowParameters,omitempty" tf:"workflow_parameters,omitempty"`

	// Specifies the Schema to use for this Logic App Workflow. Defaults to https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#. Changing this forces a new resource to be created.
	WorkflowSchema *string `json:"workflowSchema,omitempty" tf:"workflow_schema,omitempty"`

	// Specifies the version of the Schema used for this Logic App Workflow. Defaults to 1.0.0.0. Changing this forces a new resource to be created.
	WorkflowVersion *string `json:"workflowVersion,omitempty" tf:"workflow_version,omitempty"`
}

type ClaimInitParameters struct {

	// The OAuth policy name for the Logic App Workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the OAuth policy claim for the Logic App Workflow.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ClaimObservation struct {

	// The OAuth policy name for the Logic App Workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the OAuth policy claim for the Logic App Workflow.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ClaimParameters struct {

	// The OAuth policy name for the Logic App Workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the OAuth policy claim for the Logic App Workflow.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ContentInitParameters struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`
}

type ContentObservation struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`
}

type ContentParameters struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Logic App Workflow.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Logic App Workflow. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Logic App Workflow.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Managed Service Identity of this Logic App Workflow.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this Logic App Workflow.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Logic App Workflow. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Logic App Workflow.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Logic App Workflow. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OpenAuthenticationPolicyInitParameters struct {

	// A claim block as defined below.
	Claim []ClaimInitParameters `json:"claim,omitempty" tf:"claim,omitempty"`

	// The OAuth policy name for the Logic App Workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type OpenAuthenticationPolicyObservation struct {

	// A claim block as defined below.
	Claim []ClaimObservation `json:"claim,omitempty" tf:"claim,omitempty"`

	// The OAuth policy name for the Logic App Workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type OpenAuthenticationPolicyParameters struct {

	// A claim block as defined below.
	Claim []ClaimParameters `json:"claim,omitempty" tf:"claim,omitempty"`

	// The OAuth policy name for the Logic App Workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TriggerInitParameters struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`

	// A open_authentication_policy block as defined below.
	OpenAuthenticationPolicy []OpenAuthenticationPolicyInitParameters `json:"openAuthenticationPolicy,omitempty" tf:"open_authentication_policy,omitempty"`
}

type TriggerObservation struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`

	// A open_authentication_policy block as defined below.
	OpenAuthenticationPolicy []OpenAuthenticationPolicyObservation `json:"openAuthenticationPolicy,omitempty" tf:"open_authentication_policy,omitempty"`
}

type TriggerParameters struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`

	// A open_authentication_policy block as defined below.
	OpenAuthenticationPolicy []OpenAuthenticationPolicyParameters `json:"openAuthenticationPolicy,omitempty" tf:"open_authentication_policy,omitempty"`
}

type WorkflowManagementInitParameters struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`
}

type WorkflowManagementObservation struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`
}

type WorkflowManagementParameters struct {

	// A list of the allowed caller IP address ranges.
	AllowedCallerIPAddressRange []*string `json:"allowedCallerIpAddressRange,omitempty" tf:"allowed_caller_ip_address_range,omitempty"`
}

// AppWorkflowSpec defines the desired state of AppWorkflow
type AppWorkflowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppWorkflowParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AppWorkflowInitParameters `json:"initProvider,omitempty"`
}

// AppWorkflowStatus defines the observed state of AppWorkflow.
type AppWorkflowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppWorkflowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppWorkflow is the Schema for the AppWorkflows API. Manages a Logic App Workflow.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   AppWorkflowSpec   `json:"spec"`
	Status AppWorkflowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppWorkflowList contains a list of AppWorkflows
type AppWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppWorkflow `json:"items"`
}

// Repository type metadata.
var (
	AppWorkflow_Kind             = "AppWorkflow"
	AppWorkflow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppWorkflow_Kind}.String()
	AppWorkflow_KindAPIVersion   = AppWorkflow_Kind + "." + CRDGroupVersion.String()
	AppWorkflow_GroupVersionKind = CRDGroupVersion.WithKind(AppWorkflow_Kind)
)

func init() {
	SchemeBuilder.Register(&AppWorkflow{}, &AppWorkflowList{})
}
