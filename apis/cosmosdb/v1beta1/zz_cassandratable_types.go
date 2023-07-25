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

type CassandraTableAutoscaleSettingsInitParameters struct {

	// The maximum throughput of the Cassandra Table (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type CassandraTableAutoscaleSettingsObservation struct {

	// The maximum throughput of the Cassandra Table (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type CassandraTableAutoscaleSettingsParameters struct {

	// The maximum throughput of the Cassandra Table (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type CassandraTableInitParameters struct {

	// Time to live of the Analytical Storage. Possible values are between -1 and 2147483647 except 0. -1 means the Analytical Storage never expires. Changing this forces a new resource to be created.
	AnalyticalStorageTTL *float64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// An autoscale_settings block as defined below.
	AutoscaleSettings []CassandraTableAutoscaleSettingsInitParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// Time to live of the Cosmos DB Cassandra table. Possible values are at least -1. -1 means the Cassandra table never expires.
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// A schema block as defined below.
	Schema []SchemaInitParameters `json:"schema,omitempty" tf:"schema,omitempty"`

	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of 100. The minimum value is 400.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type CassandraTableObservation struct {

	// Time to live of the Analytical Storage. Possible values are between -1 and 2147483647 except 0. -1 means the Analytical Storage never expires. Changing this forces a new resource to be created.
	AnalyticalStorageTTL *float64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// An autoscale_settings block as defined below.
	AutoscaleSettings []CassandraTableAutoscaleSettingsObservation `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
	CassandraKeySpaceID *string `json:"cassandraKeyspaceId,omitempty" tf:"cassandra_keyspace_id,omitempty"`

	// Time to live of the Cosmos DB Cassandra table. Possible values are at least -1. -1 means the Cassandra table never expires.
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// the ID of the CosmosDB Cassandra Table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A schema block as defined below.
	Schema []SchemaObservation `json:"schema,omitempty" tf:"schema,omitempty"`

	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of 100. The minimum value is 400.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type CassandraTableParameters struct {

	// Time to live of the Analytical Storage. Possible values are between -1 and 2147483647 except 0. -1 means the Analytical Storage never expires. Changing this forces a new resource to be created.
	AnalyticalStorageTTL *float64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// An autoscale_settings block as defined below.
	AutoscaleSettings []CassandraTableAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=CassandraKeySpace
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CassandraKeySpaceID *string `json:"cassandraKeyspaceId,omitempty" tf:"cassandra_keyspace_id,omitempty"`

	// Reference to a CassandraKeySpace to populate cassandraKeyspaceId.
	// +kubebuilder:validation:Optional
	CassandraKeySpaceIDRef *v1.Reference `json:"cassandraKeyspaceIdRef,omitempty" tf:"-"`

	// Selector for a CassandraKeySpace to populate cassandraKeyspaceId.
	// +kubebuilder:validation:Optional
	CassandraKeySpaceIDSelector *v1.Selector `json:"cassandraKeyspaceIdSelector,omitempty" tf:"-"`

	// Time to live of the Cosmos DB Cassandra table. Possible values are at least -1. -1 means the Cassandra table never expires.
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// A schema block as defined below.
	Schema []SchemaParameters `json:"schema,omitempty" tf:"schema,omitempty"`

	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of 100. The minimum value is 400.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type ClusterKeyInitParameters struct {

	// Name of the column to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Order of the key. Currently supported values are Asc and Desc.
	OrderBy *string `json:"orderBy,omitempty" tf:"order_by,omitempty"`
}

type ClusterKeyObservation struct {

	// Name of the column to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Order of the key. Currently supported values are Asc and Desc.
	OrderBy *string `json:"orderBy,omitempty" tf:"order_by,omitempty"`
}

type ClusterKeyParameters struct {

	// Name of the column to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Order of the key. Currently supported values are Asc and Desc.
	OrderBy *string `json:"orderBy,omitempty" tf:"order_by,omitempty"`
}

type ColumnInitParameters struct {

	// Name of the column to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ColumnObservation struct {

	// Name of the column to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ColumnParameters struct {

	// Name of the column to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PartitionKeyInitParameters struct {

	// Name of the column to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PartitionKeyObservation struct {

	// Name of the column to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PartitionKeyParameters struct {

	// Name of the column to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SchemaInitParameters struct {

	// One or more cluster_key blocks as defined below.
	ClusterKey []ClusterKeyInitParameters `json:"clusterKey,omitempty" tf:"cluster_key,omitempty"`

	// One or more column blocks as defined below.
	Column []ColumnInitParameters `json:"column,omitempty" tf:"column,omitempty"`

	// One or more partition_key blocks as defined below.
	PartitionKey []PartitionKeyInitParameters `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`
}

type SchemaObservation struct {

	// One or more cluster_key blocks as defined below.
	ClusterKey []ClusterKeyObservation `json:"clusterKey,omitempty" tf:"cluster_key,omitempty"`

	// One or more column blocks as defined below.
	Column []ColumnObservation `json:"column,omitempty" tf:"column,omitempty"`

	// One or more partition_key blocks as defined below.
	PartitionKey []PartitionKeyObservation `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`
}

type SchemaParameters struct {

	// One or more cluster_key blocks as defined below.
	ClusterKey []ClusterKeyParameters `json:"clusterKey,omitempty" tf:"cluster_key,omitempty"`

	// One or more column blocks as defined below.
	Column []ColumnParameters `json:"column,omitempty" tf:"column,omitempty"`

	// One or more partition_key blocks as defined below.
	PartitionKey []PartitionKeyParameters `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`
}

// CassandraTableSpec defines the desired state of CassandraTable
type CassandraTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CassandraTableParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider CassandraTableInitParameters `json:"initProvider,omitempty"`
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
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schema) || has(self.initProvider.schema)",message="schema is a required parameter"
	Spec   CassandraTableSpec   `json:"spec"`
	Status CassandraTableStatus `json:"status,omitempty"`
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
