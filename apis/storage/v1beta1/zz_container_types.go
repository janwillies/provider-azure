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

type ContainerInitParameters struct {

	// The Access Level configured for this Container. Possible values are blob, container or private. Defaults to private.
	ContainerAccessType *string `json:"containerAccessType,omitempty" tf:"container_access_type,omitempty"`

	// A mapping of MetaData for this Container. All metadata keys should be lowercase.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

type ContainerObservation struct {

	// The Access Level configured for this Container. Possible values are blob, container or private. Defaults to private.
	ContainerAccessType *string `json:"containerAccessType,omitempty" tf:"container_access_type,omitempty"`

	// Is there an Immutability Policy configured on this Storage Container?
	HasImmutabilityPolicy *bool `json:"hasImmutabilityPolicy,omitempty" tf:"has_immutability_policy,omitempty"`

	// Is there a Legal Hold configured on this Storage Container?
	HasLegalHold *bool `json:"hasLegalHold,omitempty" tf:"has_legal_hold,omitempty"`

	// The ID of the Storage Container.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A mapping of MetaData for this Container. All metadata keys should be lowercase.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The Resource Manager ID of this Storage Container.
	ResourceManagerID *string `json:"resourceManagerId,omitempty" tf:"resource_manager_id,omitempty"`

	// The name of the Storage Account where the Container should be created. Changing this forces a new resource to be created.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`
}

type ContainerParameters struct {

	// The Access Level configured for this Container. Possible values are blob, container or private. Defaults to private.
	ContainerAccessType *string `json:"containerAccessType,omitempty" tf:"container_access_type,omitempty"`

	// A mapping of MetaData for this Container. All metadata keys should be lowercase.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the Storage Account where the Container should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

// ContainerSpec defines the desired state of Container
type ContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ContainerInitParameters `json:"initProvider,omitempty"`
}

// ContainerStatus defines the observed state of Container.
type ContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Container is the Schema for the Containers API. Manages a Container within an Azure Storage Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Container struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerSpec   `json:"spec"`
	Status            ContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerList contains a list of Containers
type ContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Container `json:"items"`
}

// Repository type metadata.
var (
	Container_Kind             = "Container"
	Container_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Container_Kind}.String()
	Container_KindAPIVersion   = Container_Kind + "." + CRDGroupVersion.String()
	Container_GroupVersionKind = CRDGroupVersion.WithKind(Container_Kind)
)

func init() {
	SchemeBuilder.Register(&Container{}, &ContainerList{})
}
