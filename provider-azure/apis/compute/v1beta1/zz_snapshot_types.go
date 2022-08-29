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

type EncryptionSettingsDiskEncryptionKeyObservation struct {
}

type EncryptionSettingsDiskEncryptionKeyParameters struct {

	// +kubebuilder:validation:Required
	SecretURL *string `json:"secretUrl" tf:"secret_url,omitempty"`

	// +kubebuilder:validation:Required
	SourceVaultID *string `json:"sourceVaultId" tf:"source_vault_id,omitempty"`
}

type EncryptionSettingsKeyEncryptionKeyObservation struct {
}

type EncryptionSettingsKeyEncryptionKeyParameters struct {

	// +kubebuilder:validation:Required
	KeyURL *string `json:"keyUrl" tf:"key_url,omitempty"`

	// +kubebuilder:validation:Required
	SourceVaultID *string `json:"sourceVaultId" tf:"source_vault_id,omitempty"`
}

type SnapshotEncryptionSettingsObservation struct {
}

type SnapshotEncryptionSettingsParameters struct {

	// +kubebuilder:validation:Optional
	DiskEncryptionKey []EncryptionSettingsDiskEncryptionKeyParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KeyEncryptionKey []EncryptionSettingsKeyEncryptionKeyParameters `json:"keyEncryptionKey,omitempty" tf:"key_encryption_key,omitempty"`
}

type SnapshotObservation struct {

	// The Snapshot ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether Trusted Launch is enabled for the Snapshot.
	TrustedLaunchEnabled *bool `json:"trustedLaunchEnabled,omitempty" tf:"trusted_launch_enabled,omitempty"`
}

type SnapshotParameters struct {

	// Indicates how the snapshot is to be created. Possible values are Copy or Import. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	CreateOption *string `json:"createOption" tf:"create_option,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionSettings []SnapshotEncryptionSettingsParameters `json:"encryptionSettings,omitempty" tf:"encryption_settings,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies a reference to an existing snapshot, when create_option is Copy. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceResourceID *string `json:"sourceResourceId,omitempty" tf:"source_resource_id,omitempty"`

	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/compute/v1beta1.ManagedDisk
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// Reference to a ManagedDisk in compute to populate sourceUri.
	// +kubebuilder:validation:Optional
	SourceURIRef *v1.Reference `json:"sourceUriRef,omitempty" tf:"-"`

	// Selector for a ManagedDisk in compute to populate sourceUri.
	// +kubebuilder:validation:Optional
	SourceURISelector *v1.Selector `json:"sourceUriSelector,omitempty" tf:"-"`

	// Specifies the ID of an storage account. Used with source_uri to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Snapshot is the Schema for the Snapshots API. Manages a Disk Snapshot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
