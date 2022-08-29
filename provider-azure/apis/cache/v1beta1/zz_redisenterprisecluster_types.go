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

type RedisEnterpriseClusterObservation struct {

	// DNS name of the cluster endpoint.
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The ID of the Redis Enterprise Cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RedisEnterpriseClusterParameters struct {

	// The Azure Region where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The minimum TLS version.  Defaults to 1.2. Changing this forces a new Redis Enterprise Cluster to be created.
	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// The name of the Resource Group where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The sku_name is comprised of two segments separated by a hyphen . The first segment of the sku_name defines the name of the SKU, possible values are Enterprise_E10, Enterprise_E20", Enterprise_E50, Enterprise_E100, EnterpriseFlash_F300, EnterpriseFlash_F700 or EnterpriseFlash_F1500. The second segment defines the capacity of the sku_name, possible values for Enteprise SKUs are . Possible values for EnterpriseFlash SKUs are . Changing this forces a new Redis Enterprise Cluster to be created.
	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the Redis Enterprise Cluster.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of Availability Zones in which this Redis Enterprise Cluster should be located. Changing this forces a new Redis Enterprise Cluster to be created.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// RedisEnterpriseClusterSpec defines the desired state of RedisEnterpriseCluster
type RedisEnterpriseClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisEnterpriseClusterParameters `json:"forProvider"`
}

// RedisEnterpriseClusterStatus defines the observed state of RedisEnterpriseCluster.
type RedisEnterpriseClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisEnterpriseClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseCluster is the Schema for the RedisEnterpriseClusters API. Manages a Redis Enterprise Cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisEnterpriseCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterpriseClusterSpec   `json:"spec"`
	Status            RedisEnterpriseClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseClusterList contains a list of RedisEnterpriseClusters
type RedisEnterpriseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterpriseCluster `json:"items"`
}

// Repository type metadata.
var (
	RedisEnterpriseCluster_Kind             = "RedisEnterpriseCluster"
	RedisEnterpriseCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisEnterpriseCluster_Kind}.String()
	RedisEnterpriseCluster_KindAPIVersion   = RedisEnterpriseCluster_Kind + "." + CRDGroupVersion.String()
	RedisEnterpriseCluster_GroupVersionKind = CRDGroupVersion.WithKind(RedisEnterpriseCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisEnterpriseCluster{}, &RedisEnterpriseClusterList{})
}
