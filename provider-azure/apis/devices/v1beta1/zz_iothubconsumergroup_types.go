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

type IOTHubConsumerGroupObservation struct {

	// The ID of the IoTHub Consumer Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubConsumerGroupParameters struct {

	// The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	EventHubEndpointName *string `json:"eventhubEndpointName" tf:"eventhub_endpoint_name,omitempty"`

	// The name of the IoT Hub. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Reference to a IOTHub to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// Selector for a IOTHub to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IOTHubConsumerGroupSpec defines the desired state of IOTHubConsumerGroup
type IOTHubConsumerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubConsumerGroupParameters `json:"forProvider"`
}

// IOTHubConsumerGroupStatus defines the observed state of IOTHubConsumerGroup.
type IOTHubConsumerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubConsumerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubConsumerGroup is the Schema for the IOTHubConsumerGroups API. Manages a Consumer Group within an IotHub
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubConsumerGroupSpec   `json:"spec"`
	Status            IOTHubConsumerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubConsumerGroupList contains a list of IOTHubConsumerGroups
type IOTHubConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubConsumerGroup `json:"items"`
}

// Repository type metadata.
var (
	IOTHubConsumerGroup_Kind             = "IOTHubConsumerGroup"
	IOTHubConsumerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubConsumerGroup_Kind}.String()
	IOTHubConsumerGroup_KindAPIVersion   = IOTHubConsumerGroup_Kind + "." + CRDGroupVersion.String()
	IOTHubConsumerGroup_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubConsumerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubConsumerGroup{}, &IOTHubConsumerGroupList{})
}
