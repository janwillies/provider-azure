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

type EncryptionInitParameters struct {
}

type EncryptionObservation struct {

	// The Key Vault URI to access the encryption key.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The ID of the keyVault where the customer owned encryption key is present.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// The Key Vault URI to access the encryption key.
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type EncryptionParameters struct {

	// The Key Vault URI to access the encryption key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a Key in keyvault to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`

	// The ID of the keyVault where the customer owned encryption key is present.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Vault
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Vault in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Vault in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// The Key Vault URI to access the encryption key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`

	// Reference to a UserAssignedIdentity in managedidentity to populate userAssignedIdentityId.
	// +kubebuilder:validation:Optional
	UserAssignedIdentityIDRef *v1.Reference `json:"userAssignedIdentityIdRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate userAssignedIdentityId.
	// +kubebuilder:validation:Optional
	UserAssignedIdentityIDSelector *v1.Selector `json:"userAssignedIdentityIdSelector,omitempty" tf:"-"`
}

type WorkspaceIdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Workspace.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Workspace. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WorkspaceIdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Workspace.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Workspace. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WorkspaceIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Workspace.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Workspace. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WorkspaceInitParameters struct {

	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryID *string `json:"containerRegistryId,omitempty" tf:"container_registry_id,omitempty"`

	// The description of this Machine Learning Workspace.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An encryption block as defined below. Changing this forces a new resource to be created.
	Encryption []EncryptionInitParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Display name for this Machine Learning Workspace.
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
	HighBusinessImpact *bool `json:"highBusinessImpact,omitempty" tf:"high_business_impact,omitempty"`

	// An identity block as defined below.
	Identity []WorkspaceIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The compute name for image build of the Machine Learning Workspace.
	ImageBuildComputeName *string `json:"imageBuildComputeName,omitempty" tf:"image_build_compute_name,omitempty"`

	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Enable public access when this Machine Learning Workspace is behind a VNet. Changing this forces a new resource to be created.
	PublicAccessBehindVirtualNetworkEnabled *bool `json:"publicAccessBehindVirtualNetworkEnabled,omitempty" tf:"public_access_behind_virtual_network_enabled,omitempty"`

	// Enable public access when this Machine Learning Workspace is behind VNet.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// SKU/edition of the Machine Learning Workspace, possible values are Basic. Defaults to Basic.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Enable V1 API features, enabling v1_legacy_mode may prevent you from using features provided by the v2 API. Defaults to false.
	V1LegacyModeEnabled *bool `json:"v1LegacyModeEnabled,omitempty" tf:"v1_legacy_mode_enabled,omitempty"`
}

type WorkspaceObservation struct {

	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryID *string `json:"containerRegistryId,omitempty" tf:"container_registry_id,omitempty"`

	// The description of this Machine Learning Workspace.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The url for the discovery service to identify regional endpoints for machine learning experimentation services.
	DiscoveryURL *string `json:"discoveryUrl,omitempty" tf:"discovery_url,omitempty"`

	// An encryption block as defined below. Changing this forces a new resource to be created.
	Encryption []EncryptionObservation `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Display name for this Machine Learning Workspace.
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
	HighBusinessImpact *bool `json:"highBusinessImpact,omitempty" tf:"high_business_impact,omitempty"`

	// The ID of the Machine Learning Workspace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []WorkspaceIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The compute name for image build of the Machine Learning Workspace.
	ImageBuildComputeName *string `json:"imageBuildComputeName,omitempty" tf:"image_build_compute_name,omitempty"`

	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The user assigned identity id that represents the workspace identity.
	PrimaryUserAssignedIdentity *string `json:"primaryUserAssignedIdentity,omitempty" tf:"primary_user_assigned_identity,omitempty"`

	// Enable public access when this Machine Learning Workspace is behind a VNet. Changing this forces a new resource to be created.
	PublicAccessBehindVirtualNetworkEnabled *bool `json:"publicAccessBehindVirtualNetworkEnabled,omitempty" tf:"public_access_behind_virtual_network_enabled,omitempty"`

	// Enable public access when this Machine Learning Workspace is behind VNet.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// SKU/edition of the Machine Learning Workspace, possible values are Basic. Defaults to Basic.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Enable V1 API features, enabling v1_legacy_mode may prevent you from using features provided by the v2 API. Defaults to false.
	V1LegacyModeEnabled *bool `json:"v1LegacyModeEnabled,omitempty" tf:"v1_legacy_mode_enabled,omitempty"`

	// The immutable id associated with this workspace.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type WorkspaceParameters struct {

	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDRef *v1.Reference `json:"applicationInsightsIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDSelector *v1.Selector `json:"applicationInsightsIdSelector,omitempty" tf:"-"`

	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryID *string `json:"containerRegistryId,omitempty" tf:"container_registry_id,omitempty"`

	// The description of this Machine Learning Workspace.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An encryption block as defined below. Changing this forces a new resource to be created.
	Encryption []EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Display name for this Machine Learning Workspace.
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
	HighBusinessImpact *bool `json:"highBusinessImpact,omitempty" tf:"high_business_impact,omitempty"`

	// An identity block as defined below.
	Identity []WorkspaceIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The compute name for image build of the Machine Learning Workspace.
	ImageBuildComputeName *string `json:"imageBuildComputeName,omitempty" tf:"image_build_compute_name,omitempty"`

	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Vault
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Vault in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Vault in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The user assigned identity id that represents the workspace identity.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrimaryUserAssignedIdentity *string `json:"primaryUserAssignedIdentity,omitempty" tf:"primary_user_assigned_identity,omitempty"`

	// Reference to a UserAssignedIdentity in managedidentity to populate primaryUserAssignedIdentity.
	// +kubebuilder:validation:Optional
	PrimaryUserAssignedIdentityRef *v1.Reference `json:"primaryUserAssignedIdentityRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate primaryUserAssignedIdentity.
	// +kubebuilder:validation:Optional
	PrimaryUserAssignedIdentitySelector *v1.Selector `json:"primaryUserAssignedIdentitySelector,omitempty" tf:"-"`

	// Enable public access when this Machine Learning Workspace is behind a VNet. Changing this forces a new resource to be created.
	PublicAccessBehindVirtualNetworkEnabled *bool `json:"publicAccessBehindVirtualNetworkEnabled,omitempty" tf:"public_access_behind_virtual_network_enabled,omitempty"`

	// Enable public access when this Machine Learning Workspace is behind VNet.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// SKU/edition of the Machine Learning Workspace, possible values are Basic. Defaults to Basic.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
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

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Enable V1 API features, enabling v1_legacy_mode may prevent you from using features provided by the v2 API. Defaults to false.
	V1LegacyModeEnabled *bool `json:"v1LegacyModeEnabled,omitempty" tf:"v1_legacy_mode_enabled,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider WorkspaceInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workspace is the Schema for the Workspaces API. Manages a Azure Machine Learning Workspace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identity) || has(self.initProvider.identity)",message="identity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   WorkspaceSpec   `json:"spec"`
	Status WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Repository type metadata.
var (
	Workspace_Kind             = "Workspace"
	Workspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workspace_Kind}.String()
	Workspace_KindAPIVersion   = Workspace_Kind + "." + CRDGroupVersion.String()
	Workspace_GroupVersionKind = CRDGroupVersion.WithKind(Workspace_Kind)
)

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
