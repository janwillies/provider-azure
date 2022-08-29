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

type RedisLinkedServerObservation struct {

	// The ID of the Redis.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the linked server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RedisLinkedServerParameters struct {

	// The ID of the linked Redis cache. Changing this forces a new Redis to be created.
	// +crossplane:generate:reference:type=RedisCache
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LinkedRedisCacheID *string `json:"linkedRedisCacheId,omitempty" tf:"linked_redis_cache_id,omitempty"`

	// Reference to a RedisCache to populate linkedRedisCacheId.
	// +kubebuilder:validation:Optional
	LinkedRedisCacheIDRef *v1.Reference `json:"linkedRedisCacheIdRef,omitempty" tf:"-"`

	// Selector for a RedisCache to populate linkedRedisCacheId.
	// +kubebuilder:validation:Optional
	LinkedRedisCacheIDSelector *v1.Selector `json:"linkedRedisCacheIdSelector,omitempty" tf:"-"`

	// The location of the linked Redis cache. Changing this forces a new Redis to be created.
	// +kubebuilder:validation:Required
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation" tf:"linked_redis_cache_location,omitempty"`

	// The name of the Resource Group where the Redis caches exists. Changing this forces a new Redis to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The role of the linked Redis cache . Changing this forces a new Redis to be created.
	// +kubebuilder:validation:Required
	ServerRole *string `json:"serverRole" tf:"server_role,omitempty"`

	// The name of Redis cache to link with. Changing this forces a new Redis to be created.
	// +crossplane:generate:reference:type=RedisCache
	// +kubebuilder:validation:Optional
	TargetRedisCacheName *string `json:"targetRedisCacheName,omitempty" tf:"target_redis_cache_name,omitempty"`

	// Reference to a RedisCache to populate targetRedisCacheName.
	// +kubebuilder:validation:Optional
	TargetRedisCacheNameRef *v1.Reference `json:"targetRedisCacheNameRef,omitempty" tf:"-"`

	// Selector for a RedisCache to populate targetRedisCacheName.
	// +kubebuilder:validation:Optional
	TargetRedisCacheNameSelector *v1.Selector `json:"targetRedisCacheNameSelector,omitempty" tf:"-"`
}

// RedisLinkedServerSpec defines the desired state of RedisLinkedServer
type RedisLinkedServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisLinkedServerParameters `json:"forProvider"`
}

// RedisLinkedServerStatus defines the observed state of RedisLinkedServer.
type RedisLinkedServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisLinkedServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisLinkedServer is the Schema for the RedisLinkedServers API. Manages a Redis Linked Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisLinkedServerSpec   `json:"spec"`
	Status            RedisLinkedServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisLinkedServerList contains a list of RedisLinkedServers
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

// Repository type metadata.
var (
	RedisLinkedServer_Kind             = "RedisLinkedServer"
	RedisLinkedServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisLinkedServer_Kind}.String()
	RedisLinkedServer_KindAPIVersion   = RedisLinkedServer_Kind + "." + CRDGroupVersion.String()
	RedisLinkedServer_GroupVersionKind = CRDGroupVersion.WithKind(RedisLinkedServer_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}
