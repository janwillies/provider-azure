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

type SentinelAlertRuleFusionInitParameters struct {

	// The GUID of the alert rule template which is used for this Sentinel Fusion Alert Rule. Changing this forces a new Sentinel Fusion Alert Rule to be created.
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// Should this Sentinel Fusion Alert Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name which should be used for this Sentinel Fusion Alert Rule. Changing this forces a new Sentinel Fusion Alert Rule to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more source blocks as defined below.
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type SentinelAlertRuleFusionObservation struct {

	// The GUID of the alert rule template which is used for this Sentinel Fusion Alert Rule. Changing this forces a new Sentinel Fusion Alert Rule to be created.
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// Should this Sentinel Fusion Alert Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Sentinel Fusion Alert Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Log Analytics Workspace this Sentinel Fusion Alert Rule belongs to. Changing this forces a new Sentinel Fusion Alert Rule to be created.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// The name which should be used for this Sentinel Fusion Alert Rule. Changing this forces a new Sentinel Fusion Alert Rule to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more source blocks as defined below.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`
}

type SentinelAlertRuleFusionParameters struct {

	// The GUID of the alert rule template which is used for this Sentinel Fusion Alert Rule. Changing this forces a new Sentinel Fusion Alert Rule to be created.
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// Should this Sentinel Fusion Alert Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Log Analytics Workspace this Sentinel Fusion Alert Rule belongs to. Changing this forces a new Sentinel Fusion Alert Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationsmanagement/v1beta1.LogAnalyticsSolution
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("workspace_resource_id",false)
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a LogAnalyticsSolution in operationsmanagement to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a LogAnalyticsSolution in operationsmanagement to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`

	// The name which should be used for this Sentinel Fusion Alert Rule. Changing this forces a new Sentinel Fusion Alert Rule to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more source blocks as defined below.
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type SourceInitParameters struct {

	// Whether this source signal is enabled or disabled in Fusion detection? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Fusion source signal. Refer to Fusion alert rule template for supported values.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more sub_type blocks as defined below.
	SubType []SubTypeInitParameters `json:"subType,omitempty" tf:"sub_type,omitempty"`
}

type SourceObservation struct {

	// Whether this source signal is enabled or disabled in Fusion detection? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Fusion source signal. Refer to Fusion alert rule template for supported values.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more sub_type blocks as defined below.
	SubType []SubTypeObservation `json:"subType,omitempty" tf:"sub_type,omitempty"`
}

type SourceParameters struct {

	// Whether this source signal is enabled or disabled in Fusion detection? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Fusion source signal. Refer to Fusion alert rule template for supported values.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more sub_type blocks as defined below.
	SubType []SubTypeParameters `json:"subType,omitempty" tf:"sub_type,omitempty"`
}

type SubTypeInitParameters struct {

	// Whether this source subtype under source signal is enabled or disabled in Fusion detection. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Name of the source subtype under a given source signal in Fusion detection. Refer to Fusion alert rule template for supported values.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of severities that are enabled for this source subtype consumed in Fusion detection. Possible values for each element are High, Medium, Low, Informational.
	SeveritiesAllowed []*string `json:"severitiesAllowed,omitempty" tf:"severities_allowed,omitempty"`
}

type SubTypeObservation struct {

	// Whether this source subtype under source signal is enabled or disabled in Fusion detection. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Name of the source subtype under a given source signal in Fusion detection. Refer to Fusion alert rule template for supported values.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of severities that are enabled for this source subtype consumed in Fusion detection. Possible values for each element are High, Medium, Low, Informational.
	SeveritiesAllowed []*string `json:"severitiesAllowed,omitempty" tf:"severities_allowed,omitempty"`
}

type SubTypeParameters struct {

	// Whether this source subtype under source signal is enabled or disabled in Fusion detection. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Name of the source subtype under a given source signal in Fusion detection. Refer to Fusion alert rule template for supported values.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of severities that are enabled for this source subtype consumed in Fusion detection. Possible values for each element are High, Medium, Low, Informational.
	SeveritiesAllowed []*string `json:"severitiesAllowed,omitempty" tf:"severities_allowed,omitempty"`
}

// SentinelAlertRuleFusionSpec defines the desired state of SentinelAlertRuleFusion
type SentinelAlertRuleFusionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelAlertRuleFusionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SentinelAlertRuleFusionInitParameters `json:"initProvider,omitempty"`
}

// SentinelAlertRuleFusionStatus defines the observed state of SentinelAlertRuleFusion.
type SentinelAlertRuleFusionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelAlertRuleFusionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleFusion is the Schema for the SentinelAlertRuleFusions API. Manages a Sentinel Fusion Alert Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelAlertRuleFusion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alertRuleTemplateGuid) || has(self.initProvider.alertRuleTemplateGuid)",message="alertRuleTemplateGuid is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   SentinelAlertRuleFusionSpec   `json:"spec"`
	Status SentinelAlertRuleFusionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleFusionList contains a list of SentinelAlertRuleFusions
type SentinelAlertRuleFusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelAlertRuleFusion `json:"items"`
}

// Repository type metadata.
var (
	SentinelAlertRuleFusion_Kind             = "SentinelAlertRuleFusion"
	SentinelAlertRuleFusion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelAlertRuleFusion_Kind}.String()
	SentinelAlertRuleFusion_KindAPIVersion   = SentinelAlertRuleFusion_Kind + "." + CRDGroupVersion.String()
	SentinelAlertRuleFusion_GroupVersionKind = CRDGroupVersion.WithKind(SentinelAlertRuleFusion_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelAlertRuleFusion{}, &SentinelAlertRuleFusionList{})
}
