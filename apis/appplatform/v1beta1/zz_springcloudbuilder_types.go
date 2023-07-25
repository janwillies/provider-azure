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

type BuildPackGroupInitParameters struct {

	// Specifies a list of the build pack's ID.
	BuildPackIds []*string `json:"buildPackIds,omitempty" tf:"build_pack_ids,omitempty"`

	// The name which should be used for this build pack group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BuildPackGroupObservation struct {

	// Specifies a list of the build pack's ID.
	BuildPackIds []*string `json:"buildPackIds,omitempty" tf:"build_pack_ids,omitempty"`

	// The name which should be used for this build pack group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BuildPackGroupParameters struct {

	// Specifies a list of the build pack's ID.
	BuildPackIds []*string `json:"buildPackIds,omitempty" tf:"build_pack_ids,omitempty"`

	// The name which should be used for this build pack group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SpringCloudBuilderInitParameters struct {

	// One or more build_pack_group blocks as defined below.
	BuildPackGroup []BuildPackGroupInitParameters `json:"buildPackGroup,omitempty" tf:"build_pack_group,omitempty"`

	// The name which should be used for this Spring Cloud Builder. Changing this forces a new Spring Cloud Builder to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A stack block as defined below.
	Stack []StackInitParameters `json:"stack,omitempty" tf:"stack,omitempty"`
}

type SpringCloudBuilderObservation struct {

	// One or more build_pack_group blocks as defined below.
	BuildPackGroup []BuildPackGroupObservation `json:"buildPackGroup,omitempty" tf:"build_pack_group,omitempty"`

	// The ID of the Spring Cloud Builder.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name which should be used for this Spring Cloud Builder. Changing this forces a new Spring Cloud Builder to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Builder to be created.
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// A stack block as defined below.
	Stack []StackObservation `json:"stack,omitempty" tf:"stack,omitempty"`
}

type SpringCloudBuilderParameters struct {

	// One or more build_pack_group blocks as defined below.
	BuildPackGroup []BuildPackGroupParameters `json:"buildPackGroup,omitempty" tf:"build_pack_group,omitempty"`

	// The name which should be used for this Spring Cloud Builder. Changing this forces a new Spring Cloud Builder to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Builder to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudService
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDRef *v1.Reference `json:"springCloudServiceIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDSelector *v1.Selector `json:"springCloudServiceIdSelector,omitempty" tf:"-"`

	// A stack block as defined below.
	Stack []StackParameters `json:"stack,omitempty" tf:"stack,omitempty"`
}

type StackInitParameters struct {

	// Specifies the ID of the ClusterStack.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the version of the ClusterStack
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type StackObservation struct {

	// Specifies the ID of the ClusterStack.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the version of the ClusterStack
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type StackParameters struct {

	// Specifies the ID of the ClusterStack.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the version of the ClusterStack
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// SpringCloudBuilderSpec defines the desired state of SpringCloudBuilder
type SpringCloudBuilderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudBuilderParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SpringCloudBuilderInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudBuilderStatus defines the observed state of SpringCloudBuilder.
type SpringCloudBuilderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudBuilderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudBuilder is the Schema for the SpringCloudBuilders API. Manages a Spring Cloud Builder.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.buildPackGroup) || has(self.initProvider.buildPackGroup)",message="buildPackGroup is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stack) || has(self.initProvider.stack)",message="stack is a required parameter"
	Spec   SpringCloudBuilderSpec   `json:"spec"`
	Status SpringCloudBuilderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudBuilderList contains a list of SpringCloudBuilders
type SpringCloudBuilderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudBuilder `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudBuilder_Kind             = "SpringCloudBuilder"
	SpringCloudBuilder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudBuilder_Kind}.String()
	SpringCloudBuilder_KindAPIVersion   = SpringCloudBuilder_Kind + "." + CRDGroupVersion.String()
	SpringCloudBuilder_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudBuilder_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudBuilder{}, &SpringCloudBuilderList{})
}
