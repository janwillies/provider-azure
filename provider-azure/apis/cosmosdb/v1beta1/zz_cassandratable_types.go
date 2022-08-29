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

type CassandraTableAutoscaleSettingsObservation struct {
}

type CassandraTableAutoscaleSettingsParameters struct {

	// The maximum throughput of the Cassandra Table . Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type CassandraTableObservation struct {

	// the ID of the CosmosDB Cassandra Table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CassandraTableParameters struct {

	// Time to live of the Analytical Storage. Possible values are at least -1. -1 means the Analytical Storage never expires. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AnalyticalStorageTTL *float64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// An autoscale_settings block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual terraform destroy-apply.
	// +kubebuilder:validation:Optional
	AutoscaleSettings []CassandraTableAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=CassandraKeySpace
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CassandraKeySpaceID *string `json:"cassandraKeyspaceId,omitempty" tf:"cassandra_keyspace_id,omitempty"`

	// Reference to a CassandraKeySpace to populate cassandraKeyspaceId.
	// +kubebuilder:validation:Optional
	CassandraKeySpaceIDRef *v1.Reference `json:"cassandraKeyspaceIdRef,omitempty" tf:"-"`

	// Selector for a CassandraKeySpace to populate cassandraKeyspaceId.
	// +kubebuilder:validation:Optional
	CassandraKeySpaceIDSelector *v1.Selector `json:"cassandraKeyspaceIdSelector,omitempty" tf:"-"`

	// Time to live of the Cosmos DB Cassandra table. Possible values are at least -1. -1 means the Cassandra table never expires.
	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// A schema block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Schema []SchemaParameters `json:"schema" tf:"schema,omitempty"`

	// The throughput of Cassandra KeySpace . Must be set in increments of 100. The minimum value is 400. This must be set upon database creation otherwise it cannot be updated without a manual terraform destroy-apply.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type ClusterKeyObservation struct {
}

type ClusterKeyParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Order of the key. Currently supported values are Asc and Desc.
	// +kubebuilder:validation:Required
	OrderBy *string `json:"orderBy" tf:"order_by,omitempty"`
}

type ColumnObservation struct {
}

type ColumnParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Type of the column to be created.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PartitionKeyObservation struct {
}

type PartitionKeyParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type SchemaObservation struct {
}

type SchemaParameters struct {

	// One or more cluster_key blocks as defined below.
	// +kubebuilder:validation:Optional
	ClusterKey []ClusterKeyParameters `json:"clusterKey,omitempty" tf:"cluster_key,omitempty"`

	// One or more column blocks as defined below.
	// +kubebuilder:validation:Required
	Column []ColumnParameters `json:"column" tf:"column,omitempty"`

	// One or more partition_key blocks as defined below.
	// +kubebuilder:validation:Required
	PartitionKey []PartitionKeyParameters `json:"partitionKey" tf:"partition_key,omitempty"`
}

// CassandraTableSpec defines the desired state of CassandraTable
type CassandraTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CassandraTableParameters `json:"forProvider"`
}

// CassandraTableStatus defines the observed state of CassandraTable.
type CassandraTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CassandraTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraTable is the Schema for the CassandraTables API. Manages a Cassandra Table within a Cosmos DB Cassandra Keyspace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CassandraTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CassandraTableSpec   `json:"spec"`
	Status            CassandraTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraTableList contains a list of CassandraTables
type CassandraTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CassandraTable `json:"items"`
}

// Repository type metadata.
var (
	CassandraTable_Kind             = "CassandraTable"
	CassandraTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CassandraTable_Kind}.String()
	CassandraTable_KindAPIVersion   = CassandraTable_Kind + "." + CRDGroupVersion.String()
	CassandraTable_GroupVersionKind = CRDGroupVersion.WithKind(CassandraTable_Kind)
)

func init() {
	SchemeBuilder.Register(&CassandraTable{}, &CassandraTableList{})
}
