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

type DataSetBlobStorageInitParameters struct {

	// The path of the file in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// The path of the folder in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// A storage_account block as defined below. Changing this forces a new resource to be created.
	StorageAccount []StorageAccountInitParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`
}

type DataSetBlobStorageObservation struct {

	// The name of the storage account container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// The ID of the Data Share in which this Data Share Blob Storage Dataset should be created. Changing this forces a new Data Share Blob Storage Dataset to be created.
	DataShareID *string `json:"dataShareId,omitempty" tf:"data_share_id,omitempty"`

	// The name of the Data Share Dataset.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The path of the file in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// The path of the folder in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// The ID of the Data Share Blob Storage Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A storage_account block as defined below. Changing this forces a new resource to be created.
	StorageAccount []StorageAccountObservation `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`
}

type DataSetBlobStorageParameters struct {

	// The name of the storage account container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Reference to a Container in storage to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// The ID of the Data Share in which this Data Share Blob Storage Dataset should be created. Changing this forces a new Data Share Blob Storage Dataset to be created.
	// +crossplane:generate:reference:type=DataShare
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataShareID *string `json:"dataShareId,omitempty" tf:"data_share_id,omitempty"`

	// Reference to a DataShare to populate dataShareId.
	// +kubebuilder:validation:Optional
	DataShareIDRef *v1.Reference `json:"dataShareIdRef,omitempty" tf:"-"`

	// Selector for a DataShare to populate dataShareId.
	// +kubebuilder:validation:Optional
	DataShareIDSelector *v1.Selector `json:"dataShareIdSelector,omitempty" tf:"-"`

	// The path of the file in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// The path of the folder in the storage container to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// A storage_account block as defined below. Changing this forces a new resource to be created.
	StorageAccount []StorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`
}

type StorageAccountInitParameters struct {

	// The subscription id of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type StorageAccountObservation struct {

	// The name of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The resource group name of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The subscription id of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type StorageAccountParameters struct {

	// The name of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Account in storage to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The resource group name of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The subscription id of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

// DataSetBlobStorageSpec defines the desired state of DataSetBlobStorage
type DataSetBlobStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetBlobStorageParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DataSetBlobStorageInitParameters `json:"initProvider,omitempty"`
}

// DataSetBlobStorageStatus defines the observed state of DataSetBlobStorage.
type DataSetBlobStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetBlobStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetBlobStorage is the Schema for the DataSetBlobStorages API. Manages a Data Share Blob Storage Dataset.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetBlobStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageAccount) || has(self.initProvider.storageAccount)",message="storageAccount is a required parameter"
	Spec   DataSetBlobStorageSpec   `json:"spec"`
	Status DataSetBlobStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetBlobStorageList contains a list of DataSetBlobStorages
type DataSetBlobStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetBlobStorage `json:"items"`
}

// Repository type metadata.
var (
	DataSetBlobStorage_Kind             = "DataSetBlobStorage"
	DataSetBlobStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetBlobStorage_Kind}.String()
	DataSetBlobStorage_KindAPIVersion   = DataSetBlobStorage_Kind + "." + CRDGroupVersion.String()
	DataSetBlobStorage_GroupVersionKind = CRDGroupVersion.WithKind(DataSetBlobStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetBlobStorage{}, &DataSetBlobStorageList{})
}
