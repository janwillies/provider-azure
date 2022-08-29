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

type CompositeIndexIndexObservation struct {
}

type CompositeIndexIndexParameters struct {

	// Order of the index. Possible values are Ascending or Descending.
	// +kubebuilder:validation:Required
	Order *string `json:"order" tf:"order,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type ExcludedPathObservation struct {
}

type ExcludedPathParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type IncludedPathObservation struct {
}

type IncludedPathParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type IndexingPolicyCompositeIndexObservation struct {
}

type IndexingPolicyCompositeIndexParameters struct {

	// One or more index blocks as defined below.
	// +kubebuilder:validation:Required
	Index []CompositeIndexIndexParameters `json:"index" tf:"index,omitempty"`
}

type IndexingPolicyObservation struct {

	// One or more spatial_index blocks as defined below.
	// +kubebuilder:validation:Optional
	SpatialIndex []IndexingPolicySpatialIndexObservation `json:"spatialIndex,omitempty" tf:"spatial_index,omitempty"`
}

type IndexingPolicyParameters struct {

	// One or more composite_index blocks as defined below.
	// +kubebuilder:validation:Optional
	CompositeIndex []IndexingPolicyCompositeIndexParameters `json:"compositeIndex,omitempty" tf:"composite_index,omitempty"`

	// One or more excluded_path blocks as defined below. Either included_path or excluded_path must contain the path /*
	// +kubebuilder:validation:Optional
	ExcludedPath []ExcludedPathParameters `json:"excludedPath,omitempty" tf:"excluded_path,omitempty"`

	// One or more included_path blocks as defined below. Either included_path or excluded_path must contain the path /*
	// +kubebuilder:validation:Optional
	IncludedPath []IncludedPathParameters `json:"includedPath,omitempty" tf:"included_path,omitempty"`

	// Indicates the indexing mode. Possible values include: Consistent and None. Defaults to Consistent.
	// +kubebuilder:validation:Optional
	IndexingMode *string `json:"indexingMode,omitempty" tf:"indexing_mode,omitempty"`

	// One or more spatial_index blocks as defined below.
	// +kubebuilder:validation:Optional
	SpatialIndex []IndexingPolicySpatialIndexParameters `json:"spatialIndex,omitempty" tf:"spatial_index,omitempty"`
}

type IndexingPolicySpatialIndexObservation struct {

	// A set of spatial types of the path.
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`
}

type IndexingPolicySpatialIndexParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type SQLContainerAutoscaleSettingsObservation struct {
}

type SQLContainerAutoscaleSettingsParameters struct {

	// The maximum throughput of the SQL container . Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type SQLContainerConflictResolutionPolicyObservation struct {
}

type SQLContainerConflictResolutionPolicyParameters struct {

	// The conflict resolution path in the case of LastWriterWins mode.
	// +kubebuilder:validation:Optional
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty" tf:"conflict_resolution_path,omitempty"`

	// The procedure to resolve conflicts in the case of Custom mode.
	// +kubebuilder:validation:Optional
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty" tf:"conflict_resolution_procedure,omitempty"`

	// Indicates the conflict resolution mode. Possible values include: LastWriterWins, Custom.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type SQLContainerObservation struct {

	// The ID of the CosmosDB SQL Container.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An indexing_policy block as defined below.
	// +kubebuilder:validation:Optional
	IndexingPolicy []IndexingPolicyObservation `json:"indexingPolicy,omitempty" tf:"indexing_policy,omitempty"`
}

type SQLContainerParameters struct {

	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/cosmosdb/v1beta1.Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// The default time to live of Analytical Storage for this SQL container. If present and the value is set to -1, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number n – items will expire n seconds after their last modified time. Changing this forces a new Cosmos DB SQL Container to be created when removing analytical_storage_ttl on an existing Cosmos DB SQL Container.
	// +kubebuilder:validation:Optional
	AnalyticalStorageTTL *float64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// An autoscale_settings block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual terraform destroy-apply. Requires partition_key_path to be set.
	// +kubebuilder:validation:Optional
	AutoscaleSettings []SQLContainerAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// A conflict_resolution_policy blocks as defined below.
	// +kubebuilder:validation:Optional
	ConflictResolutionPolicy []SQLContainerConflictResolutionPolicyParameters `json:"conflictResolutionPolicy,omitempty" tf:"conflict_resolution_policy,omitempty"`

	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=SQLDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a SQLDatabase to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a SQLDatabase to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to -1, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number n – items will expire n seconds after their last modified time.
	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// An indexing_policy block as defined below.
	// +kubebuilder:validation:Optional
	IndexingPolicy []IndexingPolicyParameters `json:"indexingPolicy,omitempty" tf:"indexing_policy,omitempty"`

	// Define a partition key. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	PartitionKeyPath *string `json:"partitionKeyPath" tf:"partition_key_path,omitempty"`

	// Define a partition key version. Changing this forces a new resource to be created. Possible values are 1 and 2. This should be set to 2 in order to use large partition keys.
	// +kubebuilder:validation:Optional
	PartitionKeyVersion *float64 `json:"partitionKeyVersion,omitempty" tf:"partition_key_version,omitempty"`

	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The throughput of SQL container . Must be set in increments of 100. The minimum value is 400. This must be set upon container creation otherwise it cannot be updated without a manual terraform destroy-apply.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// One or more unique_key blocks as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	UniqueKey []SQLContainerUniqueKeyParameters `json:"uniqueKey,omitempty" tf:"unique_key,omitempty"`
}

type SQLContainerUniqueKeyObservation struct {
}

type SQLContainerUniqueKeyParameters struct {

	// A list of paths to use for this unique key.
	// +kubebuilder:validation:Required
	Paths []*string `json:"paths" tf:"paths,omitempty"`
}

// SQLContainerSpec defines the desired state of SQLContainer
type SQLContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLContainerParameters `json:"forProvider"`
}

// SQLContainerStatus defines the observed state of SQLContainer.
type SQLContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLContainer is the Schema for the SQLContainers API. Manages a SQL Container within a Cosmos DB Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLContainerSpec   `json:"spec"`
	Status            SQLContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLContainerList contains a list of SQLContainers
type SQLContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLContainer `json:"items"`
}

// Repository type metadata.
var (
	SQLContainer_Kind             = "SQLContainer"
	SQLContainer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLContainer_Kind}.String()
	SQLContainer_KindAPIVersion   = SQLContainer_Kind + "." + CRDGroupVersion.String()
	SQLContainer_GroupVersionKind = CRDGroupVersion.WithKind(SQLContainer_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLContainer{}, &SQLContainerList{})
}
