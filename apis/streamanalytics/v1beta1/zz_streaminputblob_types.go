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

type StreamInputBlobInitParameters struct {

	// The date format. Wherever {date} appears in path_pattern, the value of this property is used as the date format instead.
	DateFormat *string `json:"dateFormat,omitempty" tf:"date_format,omitempty"`

	// The name of the Stream Input Blob. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern *string `json:"pathPattern,omitempty" tf:"path_pattern,omitempty"`

	// A serialization block as defined below.
	Serialization []StreamInputBlobSerializationInitParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The time format. Wherever {time} appears in path_pattern, the value of this property is used as the time format instead.
	TimeFormat *string `json:"timeFormat,omitempty" tf:"time_format,omitempty"`
}

type StreamInputBlobObservation struct {

	// The date format. Wherever {date} appears in path_pattern, the value of this property is used as the date format instead.
	DateFormat *string `json:"dateFormat,omitempty" tf:"date_format,omitempty"`

	// The ID of the Stream Analytics Stream Input Blob.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Stream Input Blob. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern *string `json:"pathPattern,omitempty" tf:"path_pattern,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A serialization block as defined below.
	Serialization []StreamInputBlobSerializationObservation `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The name of the Storage Account.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// The name of the Container within the Storage Account.
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// The time format. Wherever {time} appears in path_pattern, the value of this property is used as the time format instead.
	TimeFormat *string `json:"timeFormat,omitempty" tf:"time_format,omitempty"`
}

type StreamInputBlobParameters struct {

	// The date format. Wherever {date} appears in path_pattern, the value of this property is used as the date format instead.
	DateFormat *string `json:"dateFormat,omitempty" tf:"date_format,omitempty"`

	// The name of the Stream Input Blob. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern *string `json:"pathPattern,omitempty" tf:"path_pattern,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A serialization block as defined below.
	Serialization []StreamInputBlobSerializationParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKeySecretRef v1.SecretKeySelector `json:"storageAccountKeySecretRef" tf:"-"`

	// The name of the Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`

	// The name of the Container within the Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`

	// Reference to a Container in storage to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameRef *v1.Reference `json:"storageContainerNameRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameSelector *v1.Selector `json:"storageContainerNameSelector,omitempty" tf:"-"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`

	// The time format. Wherever {time} appears in path_pattern, the value of this property is used as the time format instead.
	TimeFormat *string `json:"timeFormat,omitempty" tf:"time_format,omitempty"`
}

type StreamInputBlobSerializationInitParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The serialization format used for incoming data streams. Possible values are Avro, Csv and Json.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StreamInputBlobSerializationObservation struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The serialization format used for incoming data streams. Possible values are Avro, Csv and Json.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StreamInputBlobSerializationParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The serialization format used for incoming data streams. Possible values are Avro, Csv and Json.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// StreamInputBlobSpec defines the desired state of StreamInputBlob
type StreamInputBlobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamInputBlobParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider StreamInputBlobInitParameters `json:"initProvider,omitempty"`
}

// StreamInputBlobStatus defines the observed state of StreamInputBlob.
type StreamInputBlobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamInputBlobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StreamInputBlob is the Schema for the StreamInputBlobs API. Manages a Stream Analytics Stream Input Blob.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StreamInputBlob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dateFormat) || has(self.initProvider.dateFormat)",message="dateFormat is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pathPattern) || has(self.initProvider.pathPattern)",message="pathPattern is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serialization) || has(self.initProvider.serialization)",message="serialization is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageAccountKeySecretRef)",message="storageAccountKeySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timeFormat) || has(self.initProvider.timeFormat)",message="timeFormat is a required parameter"
	Spec   StreamInputBlobSpec   `json:"spec"`
	Status StreamInputBlobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamInputBlobList contains a list of StreamInputBlobs
type StreamInputBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamInputBlob `json:"items"`
}

// Repository type metadata.
var (
	StreamInputBlob_Kind             = "StreamInputBlob"
	StreamInputBlob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StreamInputBlob_Kind}.String()
	StreamInputBlob_KindAPIVersion   = StreamInputBlob_Kind + "." + CRDGroupVersion.String()
	StreamInputBlob_GroupVersionKind = CRDGroupVersion.WithKind(StreamInputBlob_Kind)
)

func init() {
	SchemeBuilder.Register(&StreamInputBlob{}, &StreamInputBlobList{})
}
