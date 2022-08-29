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

type DiskEncryptionKeyObservation struct {
}

type DiskEncryptionKeyParameters struct {

	// The URL to the Key Vault Secret used as the Disk Encryption Key. This can be found as id on the azurerm_key_vault_secret resource.
	// +kubebuilder:validation:Required
	SecretURL *string `json:"secretUrl" tf:"secret_url,omitempty"`

	// +kubebuilder:validation:Required
	SourceVaultID *string `json:"sourceVaultId" tf:"source_vault_id,omitempty"`
}

type EncryptionSettingsObservation struct {
}

type EncryptionSettingsParameters struct {

	// A disk_encryption_key block as defined above.
	// +kubebuilder:validation:Optional
	DiskEncryptionKey []DiskEncryptionKeyParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// Is Encryption enabled on this Managed Disk? Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// A key_encryption_key block as defined below.
	// +kubebuilder:validation:Optional
	KeyEncryptionKey []KeyEncryptionKeyParameters `json:"keyEncryptionKey,omitempty" tf:"key_encryption_key,omitempty"`
}

type KeyEncryptionKeyObservation struct {
}

type KeyEncryptionKeyParameters struct {

	// The URL to the Key Vault Key used as the Key Encryption Key. This can be found as id on the azurerm_key_vault_key resource.
	// +kubebuilder:validation:Required
	KeyURL *string `json:"keyUrl" tf:"key_url,omitempty"`

	// +kubebuilder:validation:Required
	SourceVaultID *string `json:"sourceVaultId" tf:"source_vault_id,omitempty"`
}

type ManagedDiskObservation struct {

	// The ID of the Managed Disk.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagedDiskParameters struct {

	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
	// +kubebuilder:validation:Required
	CreateOption *string `json:"createOption" tf:"create_option,omitempty"`

	// The ID of the disk access resource for using private endpoints on disks.
	// +kubebuilder:validation:Optional
	DiskAccessID *string `json:"diskAccessId,omitempty" tf:"disk_access_id,omitempty"`

	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk. Conflicts with secure_vm_disk_encryption_set_id.
	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// The number of IOPS allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. One operation can transfer between 4k and 256k bytes.
	// +kubebuilder:validation:Optional
	DiskIopsReadOnly *float64 `json:"diskIopsReadOnly,omitempty" tf:"disk_iops_read_only,omitempty"`

	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	// +kubebuilder:validation:Optional
	DiskIopsReadWrite *float64 `json:"diskIopsReadWrite,omitempty" tf:"disk_iops_read_write,omitempty"`

	// The bandwidth allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. MBps means millions of bytes per second.
	// +kubebuilder:validation:Optional
	DiskMbpsReadOnly *float64 `json:"diskMbpsReadOnly,omitempty" tf:"disk_mbps_read_only,omitempty"`

	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	// +kubebuilder:validation:Optional
	DiskMbpsReadWrite *float64 `json:"diskMbpsReadWrite,omitempty" tf:"disk_mbps_read_write,omitempty"`

	// Specifies the size of the managed disk to create in gigabytes. If create_option is Copy or FromImage, then the value must be equal to or greater than the source's size. The size can only be increased.
	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Managed Disk should exist. Changing this forces a new Managed Disk to be created.
	// +kubebuilder:validation:Optional
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// A encryption_settings block as defined below.
	// +kubebuilder:validation:Optional
	EncryptionSettings []EncryptionSettingsParameters `json:"encryptionSettings,omitempty" tf:"encryption_settings,omitempty"`

	// ID of a Gallery Image Version to copy when create_option is FromImage. This field cannot be specified if image_reference_id is specified.
	// +kubebuilder:validation:Optional
	GalleryImageReferenceID *string `json:"galleryImageReferenceId,omitempty" tf:"gallery_image_reference_id,omitempty"`

	// The HyperV Generation of the Disk when the source of an Import or Copy operation targets a source that contains an operating system. Possible values are V1 and V2. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	HyperVGeneration *string `json:"hyperVGeneration,omitempty" tf:"hyper_v_generation,omitempty"`

	// ID of an existing platform/marketplace disk image to copy when create_option is FromImage. This field cannot be specified if gallery_image_reference_id is specified.
	// +kubebuilder:validation:Optional
	ImageReferenceID *string `json:"imageReferenceId,omitempty" tf:"image_reference_id,omitempty"`

	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Logical Sector Size. Possible values are: 512 and 4096. Defaults to 4096. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	LogicalSectorSize *float64 `json:"logicalSectorSize,omitempty" tf:"logical_sector_size,omitempty"`

	// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
	// +kubebuilder:validation:Optional
	MaxShares *float64 `json:"maxShares,omitempty" tf:"max_shares,omitempty"`

	// Policy for accessing the disk via network. Allowed values are AllowAll, AllowPrivate, and DenyAll.
	// +kubebuilder:validation:Optional
	NetworkAccessPolicy *string `json:"networkAccessPolicy,omitempty" tf:"network_access_policy,omitempty"`

	// emand Bursting is enabled for the Managed Disk. Defaults to false.
	// +kubebuilder:validation:Optional
	OnDemandBurstingEnabled *bool `json:"onDemandBurstingEnabled,omitempty" tf:"on_demand_bursting_enabled,omitempty"`

	// Specify a value when the source of an Import or Copy operation targets a source that contains an operating system. Valid values are Linux or Windows.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Whether it is allowed to access the disk via public network. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the Resource Group where the Managed Disk should exist.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The ID of the Disk Encryption Set which should be used to Encrypt this OS Disk when the Virtual Machine is a Confidential VM. Conflicts with disk_encryption_set_id. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SecureVMDiskEncryptionSetID *string `json:"secureVmDiskEncryptionSetId,omitempty" tf:"secure_vm_disk_encryption_set_id,omitempty"`

	// Security Type of the Managed Disk when it is used for a Confidential VM. Possible values are ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey, ConfidentialVM_DiskEncryptedWithPlatformKey and ConfidentialVM_DiskEncryptedWithCustomerKey. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SecurityType *string `json:"securityType,omitempty" tf:"security_type,omitempty"`

	// The ID of an existing Managed Disk to copy create_option is Copy or the recovery point to restore when create_option is Restore
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/compute/v1beta1.ManagedDisk
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceResourceID *string `json:"sourceResourceId,omitempty" tf:"source_resource_id,omitempty"`

	// Reference to a ManagedDisk in compute to populate sourceResourceId.
	// +kubebuilder:validation:Optional
	SourceResourceIDRef *v1.Reference `json:"sourceResourceIdRef,omitempty" tf:"-"`

	// Selector for a ManagedDisk in compute to populate sourceResourceId.
	// +kubebuilder:validation:Optional
	SourceResourceIDSelector *v1.Selector `json:"sourceResourceIdSelector,omitempty" tf:"-"`

	// URI to a valid VHD file to be used when create_option is Import.
	// +kubebuilder:validation:Optional
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// The ID of the Storage Account where the source_uri is located. Required when create_option is set to Import.  Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// The type of storage to use for the managed disk. Possible values are Standard_LRS, StandardSSD_ZRS, Premium_LRS, Premium_ZRS, StandardSSD_LRS or UltraSSD_LRS.
	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The disk performance tier to use. Possible values are documented here. This feature is currently supported only for premium SSDs.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// Specifies if Trusted Launch is enabled for the Managed Disk. Defaults to false.
	// +kubebuilder:validation:Optional
	TrustedLaunchEnabled *bool `json:"trustedLaunchEnabled,omitempty" tf:"trusted_launch_enabled,omitempty"`

	// Specifies the Availability Zone in which this Managed Disk should be located. Changing this property forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// ManagedDiskSpec defines the desired state of ManagedDisk
type ManagedDiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedDiskParameters `json:"forProvider"`
}

// ManagedDiskStatus defines the observed state of ManagedDisk.
type ManagedDiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedDiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedDisk is the Schema for the ManagedDisks API. Manages a Managed Disk.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagedDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedDiskSpec   `json:"spec"`
	Status            ManagedDiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedDiskList contains a list of ManagedDisks
type ManagedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedDisk `json:"items"`
}

// Repository type metadata.
var (
	ManagedDisk_Kind             = "ManagedDisk"
	ManagedDisk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedDisk_Kind}.String()
	ManagedDisk_KindAPIVersion   = ManagedDisk_Kind + "." + CRDGroupVersion.String()
	ManagedDisk_GroupVersionKind = CRDGroupVersion.WithKind(ManagedDisk_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedDisk{}, &ManagedDiskList{})
}
