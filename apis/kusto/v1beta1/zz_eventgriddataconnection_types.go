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

type EventGridDataConnectionInitParameters struct {

	// Specifies the blob storage event type that needs to be processed. Possible Values are Microsoft.Storage.BlobCreated and Microsoft.Storage.BlobRenamed. Defaults to Microsoft.Storage.BlobCreated.
	BlobStorageEventType *string `json:"blobStorageEventType,omitempty" tf:"blob_storage_event_type,omitempty"`

	// Specifies the data format of the EventHub messages. Allowed values: APACHEAVRO, AVRO, CSV, JSON, MULTIJSON, ORC, PARQUET, PSV, RAW, SCSV, SINGLEJSON, SOHSV, TSV, TSVE, TXT and W3CLOGFILE.
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// Indication for database routing information from the data connection, by default only database routing information is allowed. Allowed values: Single, Multi. Changing this forces a new resource to be created.
	DatabaseRoutingType *string `json:"databaseRoutingType,omitempty" tf:"database_routing_type,omitempty"`

	// The resource ID of the event grid that is subscribed to the storage account events.
	EventGridResourceID *string `json:"eventgridResourceId,omitempty" tf:"eventgrid_resource_id,omitempty"`

	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Empty for non-managed identity based data connection. For system assigned identity, provide cluster resource Id. For user assigned identity (UAI) provide the UAI resource Id.
	ManagedIdentityResourceID *string `json:"managedIdentityResourceId,omitempty" tf:"managed_identity_resource_id,omitempty"`

	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `json:"mappingRuleName,omitempty" tf:"mapping_rule_name,omitempty"`

	// is the first record of every file ignored? Defaults to false.
	SkipFirstRecord *bool `json:"skipFirstRecord,omitempty" tf:"skip_first_record,omitempty"`

	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type EventGridDataConnectionObservation struct {

	// Specifies the blob storage event type that needs to be processed. Possible Values are Microsoft.Storage.BlobCreated and Microsoft.Storage.BlobRenamed. Defaults to Microsoft.Storage.BlobCreated.
	BlobStorageEventType *string `json:"blobStorageEventType,omitempty" tf:"blob_storage_event_type,omitempty"`

	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Specifies the data format of the EventHub messages. Allowed values: APACHEAVRO, AVRO, CSV, JSON, MULTIJSON, ORC, PARQUET, PSV, RAW, SCSV, SINGLEJSON, SOHSV, TSV, TSVE, TXT and W3CLOGFILE.
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Indication for database routing information from the data connection, by default only database routing information is allowed. Allowed values: Single, Multi. Changing this forces a new resource to be created.
	DatabaseRoutingType *string `json:"databaseRoutingType,omitempty" tf:"database_routing_type,omitempty"`

	// The resource ID of the event grid that is subscribed to the storage account events.
	EventGridResourceID *string `json:"eventgridResourceId,omitempty" tf:"eventgrid_resource_id,omitempty"`

	// Specifies the Event Hub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventHubConsumerGroupName *string `json:"eventhubConsumerGroupName,omitempty" tf:"eventhub_consumer_group_name,omitempty"`

	// Specifies the resource id of the Event Hub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventHubID *string `json:"eventhubId,omitempty" tf:"eventhub_id,omitempty"`

	// The ID of the Kusto Event Grid Data Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Empty for non-managed identity based data connection. For system assigned identity, provide cluster resource Id. For user assigned identity (UAI) provide the UAI resource Id.
	ManagedIdentityResourceID *string `json:"managedIdentityResourceId,omitempty" tf:"managed_identity_resource_id,omitempty"`

	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `json:"mappingRuleName,omitempty" tf:"mapping_rule_name,omitempty"`

	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// is the first record of every file ignored? Defaults to false.
	SkipFirstRecord *bool `json:"skipFirstRecord,omitempty" tf:"skip_first_record,omitempty"`

	// Specifies the resource id of the Storage Account this data connection will use for ingestion. Changing this forces a new resource to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type EventGridDataConnectionParameters struct {

	// Specifies the blob storage event type that needs to be processed. Possible Values are Microsoft.Storage.BlobCreated and Microsoft.Storage.BlobRenamed. Defaults to Microsoft.Storage.BlobCreated.
	BlobStorageEventType *string `json:"blobStorageEventType,omitempty" tf:"blob_storage_event_type,omitempty"`

	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// Specifies the data format of the EventHub messages. Allowed values: APACHEAVRO, AVRO, CSV, JSON, MULTIJSON, ORC, PARQUET, PSV, RAW, SCSV, SINGLEJSON, SOHSV, TSV, TSVE, TXT and W3CLOGFILE.
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Database
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a Database in kusto to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a Database in kusto to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// Indication for database routing information from the data connection, by default only database routing information is allowed. Allowed values: Single, Multi. Changing this forces a new resource to be created.
	DatabaseRoutingType *string `json:"databaseRoutingType,omitempty" tf:"database_routing_type,omitempty"`

	// The resource ID of the event grid that is subscribed to the storage account events.
	EventGridResourceID *string `json:"eventgridResourceId,omitempty" tf:"eventgrid_resource_id,omitempty"`

	// Specifies the Event Hub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.ConsumerGroup
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupName *string `json:"eventhubConsumerGroupName,omitempty" tf:"eventhub_consumer_group_name,omitempty"`

	// Reference to a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameRef *v1.Reference `json:"eventhubConsumerGroupNameRef,omitempty" tf:"-"`

	// Selector for a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameSelector *v1.Selector `json:"eventhubConsumerGroupNameSelector,omitempty" tf:"-"`

	// Specifies the resource id of the Event Hub this data connection will use for ingestion. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EventHubID *string `json:"eventhubId,omitempty" tf:"eventhub_id,omitempty"`

	// Reference to a EventHub in eventhub to populate eventhubId.
	// +kubebuilder:validation:Optional
	EventHubIDRef *v1.Reference `json:"eventhubIdRef,omitempty" tf:"-"`

	// Selector for a EventHub in eventhub to populate eventhubId.
	// +kubebuilder:validation:Optional
	EventHubIDSelector *v1.Selector `json:"eventhubIdSelector,omitempty" tf:"-"`

	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Empty for non-managed identity based data connection. For system assigned identity, provide cluster resource Id. For user assigned identity (UAI) provide the UAI resource Id.
	ManagedIdentityResourceID *string `json:"managedIdentityResourceId,omitempty" tf:"managed_identity_resource_id,omitempty"`

	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `json:"mappingRuleName,omitempty" tf:"mapping_rule_name,omitempty"`

	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// is the first record of every file ignored? Defaults to false.
	SkipFirstRecord *bool `json:"skipFirstRecord,omitempty" tf:"skip_first_record,omitempty"`

	// Specifies the resource id of the Storage Account this data connection will use for ingestion. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`

	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

// EventGridDataConnectionSpec defines the desired state of EventGridDataConnection
type EventGridDataConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventGridDataConnectionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider EventGridDataConnectionInitParameters `json:"initProvider,omitempty"`
}

// EventGridDataConnectionStatus defines the observed state of EventGridDataConnection.
type EventGridDataConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventGridDataConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventGridDataConnection is the Schema for the EventGridDataConnections API. Manages Kusto / Data Explorer Event Grid Data Connection
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventGridDataConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   EventGridDataConnectionSpec   `json:"spec"`
	Status EventGridDataConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventGridDataConnectionList contains a list of EventGridDataConnections
type EventGridDataConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventGridDataConnection `json:"items"`
}

// Repository type metadata.
var (
	EventGridDataConnection_Kind             = "EventGridDataConnection"
	EventGridDataConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventGridDataConnection_Kind}.String()
	EventGridDataConnection_KindAPIVersion   = EventGridDataConnection_Kind + "." + CRDGroupVersion.String()
	EventGridDataConnection_GroupVersionKind = CRDGroupVersion.WithKind(EventGridDataConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&EventGridDataConnection{}, &EventGridDataConnectionList{})
}
