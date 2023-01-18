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

type CustomParametersObservation struct {
}

type CustomParametersParameters struct {

	// The ID of a Azure Machine Learning workspace to link with Databricks workspace. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId,omitempty" tf:"machine_learning_workspace_id,omitempty"`

	// Name of the NAT gateway for Secure Cluster Connectivity (No Public IP) workspace subnets. Defaults to nat-gateway. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	NATGatewayName *string `json:"natGatewayName,omitempty" tf:"nat_gateway_name,omitempty"`

	// Are public IP Addresses not allowed? Possible values are true or false. Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	NoPublicIP *bool `json:"noPublicIp,omitempty" tf:"no_public_ip,omitempty"`

	// The name of the Private Subnet within the Virtual Network. Required if virtual_network_id is set. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	PrivateSubnetName *string `json:"privateSubnetName,omitempty" tf:"private_subnet_name,omitempty"`

	// Reference to a Subnet in network to populate privateSubnetName.
	// +kubebuilder:validation:Optional
	PrivateSubnetNameRef *v1.Reference `json:"privateSubnetNameRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate privateSubnetName.
	// +kubebuilder:validation:Optional
	PrivateSubnetNameSelector *v1.Selector `json:"privateSubnetNameSelector,omitempty" tf:"-"`

	// The resource ID of the azurerm_subnet_network_security_group_association resource which is referred to by the private_subnet_name field. This is the same as the ID of the subnet referred to by the private_subnet_name field. Required if virtual_network_id is set.
	// +kubebuilder:validation:Optional
	PrivateSubnetNetworkSecurityGroupAssociationID *string `json:"privateSubnetNetworkSecurityGroupAssociationId,omitempty" tf:"private_subnet_network_security_group_association_id,omitempty"`

	// Name of the Public IP for No Public IP workspace with managed vNet. Defaults to nat-gw-public-ip. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PublicIPName *string `json:"publicIpName,omitempty" tf:"public_ip_name,omitempty"`

	// The name of the Public Subnet within the Virtual Network. Required if virtual_network_id is set. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	PublicSubnetName *string `json:"publicSubnetName,omitempty" tf:"public_subnet_name,omitempty"`

	// Reference to a Subnet in network to populate publicSubnetName.
	// +kubebuilder:validation:Optional
	PublicSubnetNameRef *v1.Reference `json:"publicSubnetNameRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate publicSubnetName.
	// +kubebuilder:validation:Optional
	PublicSubnetNameSelector *v1.Selector `json:"publicSubnetNameSelector,omitempty" tf:"-"`

	// The resource ID of the azurerm_subnet_network_security_group_association resource which is referred to by the public_subnet_name field. This is the same as the ID of the subnet referred to by the public_subnet_name field. Required if virtual_network_id is set.
	// +kubebuilder:validation:Optional
	PublicSubnetNetworkSecurityGroupAssociationID *string `json:"publicSubnetNetworkSecurityGroupAssociationId,omitempty" tf:"public_subnet_network_security_group_association_id,omitempty"`

	// Default Databricks File Storage account name. Defaults to a randomized name(e.g. dbstoragel6mfeghoe5kxu). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Storage account SKU name. Possible values include Standard_LRS, Standard_GRS, Standard_RAGRS, Standard_GZRS, Standard_RAGZRS, Standard_ZRS, Premium_LRS or Premium_ZRS. Defaults to Standard_GRS. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StorageAccountSkuName *string `json:"storageAccountSkuName,omitempty" tf:"storage_account_sku_name,omitempty"`

	// The ID of a Virtual Network where this Databricks Cluster should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`

	// Address prefix for Managed virtual network. Defaults to 10.139. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	VnetAddressPrefix *string `json:"vnetAddressPrefix,omitempty" tf:"vnet_address_prefix,omitempty"`
}

type StorageAccountIdentityObservation struct {

	// The principal UUID for the internal databricks storage account needed to provide access to the workspace for enabling Customer Managed Keys.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The UUID of the tenant where the internal databricks storage account was created.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The type of the internal databricks storage account.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StorageAccountIdentityParameters struct {
}

type WorkspaceObservation struct {

	// The ID of the Databricks Workspace in the Azure management plane.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Managed Resource Group created by the Databricks Workspace.
	ManagedResourceGroupID *string `json:"managedResourceGroupId,omitempty" tf:"managed_resource_group_id,omitempty"`

	// A storage_account_identity block as documented below.
	StorageAccountIdentity []StorageAccountIdentityObservation `json:"storageAccountIdentity,omitempty" tf:"storage_account_identity,omitempty"`

	// The unique identifier of the databricks workspace in Databricks control plane.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
	WorkspaceURL *string `json:"workspaceUrl,omitempty" tf:"workspace_url,omitempty"`
}

type WorkspaceParameters struct {

	// A custom_parameters block as documented below.
	// +kubebuilder:validation:Optional
	CustomParameters []CustomParametersParameters `json:"customParameters,omitempty" tf:"custom_parameters,omitempty"`

	// Is the workspace enabled for customer managed key encryption? If true this enables the Managed Identity for the managed storage account. Possible values are true or false. Defaults to false. This field is only valid if the Databricks Workspace sku is set to premium. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CustomerManagedKeyEnabled *bool `json:"customerManagedKeyEnabled,omitempty" tf:"customer_managed_key_enabled,omitempty"`

	// Is the Databricks File System root file system enabled with a secondary layer of encryption with platform managed keys? Possible values are true or false. Defaults to false. This field is only valid if the Databricks Workspace sku is set to premium. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	InfrastructureEncryptionEnabled *bool `json:"infrastructureEncryptionEnabled,omitempty" tf:"infrastructure_encryption_enabled,omitempty"`

	// Resource ID of the Outbound Load balancer Backend Address Pool for Secure Cluster Connectivity (No Public IP) workspace. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	LoadBalancerBackendAddressPoolID *string `json:"loadBalancerBackendAddressPoolId,omitempty" tf:"load_balancer_backend_address_pool_id,omitempty"`

	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate managedResourceGroupName.
	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameRef *v1.Reference `json:"managedResourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate managedResourceGroupName.
	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameSelector *v1.Selector `json:"managedResourceGroupNameSelector,omitempty" tf:"-"`

	// Customer managed encryption properties for the Databricks Workspace managed resources(e.g. Notebooks and Artifacts). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ManagedServicesCmkKeyVaultKeyID *string `json:"managedServicesCmkKeyVaultKeyId,omitempty" tf:"managed_services_cmk_key_vault_key_id,omitempty"`

	// Does the data plane (clusters) to control plane communication happen over private link endpoint only or publicly? Possible values AllRules, NoAzureDatabricksRules or NoAzureServiceRules. Required when public_network_access_enabled is set to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupRulesRequired *string `json:"networkSecurityGroupRulesRequired,omitempty" tf:"network_security_group_rules_required,omitempty"`

	// Allow public access for accessing workspace. Set value to false to access workspace only via private link endpoint. Possible values include true or false. Defaults to true. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The sku to use for the Databricks Workspace. Possible values are standard, premium, or trial. Changing this can force a new resource to be created in some circumstances.
	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workspace is the Schema for the Workspaces API. Manages a Databricks Workspace
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec"`
	Status            WorkspaceStatus `json:"status,omitempty"`
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
