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

type ScopeMapInitParameters struct {

	// A list of actions to attach to the scope map (e.g. repo/content/read, repo2/content/delete).
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// The description of the Container Registry.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type ScopeMapObservation struct {

	// A list of actions to attach to the scope map (e.g. repo/content/read, repo2/content/delete).
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// The name of the Container Registry. Changing this forces a new resource to be created.
	ContainerRegistryName *string `json:"containerRegistryName,omitempty" tf:"container_registry_name,omitempty"`

	// The description of the Container Registry.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Container Registry scope map.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which to create the Container Registry token. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type ScopeMapParameters struct {

	// A list of actions to attach to the scope map (e.g. repo/content/read, repo2/content/delete).
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// The name of the Container Registry. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerregistry/v1beta1.Registry
	// +kubebuilder:validation:Optional
	ContainerRegistryName *string `json:"containerRegistryName,omitempty" tf:"container_registry_name,omitempty"`

	// Reference to a Registry in containerregistry to populate containerRegistryName.
	// +kubebuilder:validation:Optional
	ContainerRegistryNameRef *v1.Reference `json:"containerRegistryNameRef,omitempty" tf:"-"`

	// Selector for a Registry in containerregistry to populate containerRegistryName.
	// +kubebuilder:validation:Optional
	ContainerRegistryNameSelector *v1.Selector `json:"containerRegistryNameSelector,omitempty" tf:"-"`

	// The description of the Container Registry.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the resource group in which to create the Container Registry token. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// ScopeMapSpec defines the desired state of ScopeMap
type ScopeMapSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScopeMapParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ScopeMapInitParameters `json:"initProvider,omitempty"`
}

// ScopeMapStatus defines the observed state of ScopeMap.
type ScopeMapStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScopeMapObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScopeMap is the Schema for the ScopeMaps API. Manages an Azure Container Registry scope map.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ScopeMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actions) || has(self.initProvider.actions)",message="actions is a required parameter"
	Spec   ScopeMapSpec   `json:"spec"`
	Status ScopeMapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScopeMapList contains a list of ScopeMaps
type ScopeMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScopeMap `json:"items"`
}

// Repository type metadata.
var (
	ScopeMap_Kind             = "ScopeMap"
	ScopeMap_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScopeMap_Kind}.String()
	ScopeMap_KindAPIVersion   = ScopeMap_Kind + "." + CRDGroupVersion.String()
	ScopeMap_GroupVersionKind = CRDGroupVersion.WithKind(ScopeMap_Kind)
)

func init() {
	SchemeBuilder.Register(&ScopeMap{}, &ScopeMapList{})
}
