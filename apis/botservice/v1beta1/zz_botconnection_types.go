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

type BotConnectionInitParameters struct {

	// The Client ID that will be used to authenticate with the service provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A map of additional parameters to apply to the connection.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The Scopes at which the connection should be applied.
	Scopes *string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// The name of the service provider that will be associated with this connection. Changing this forces a new resource to be created.
	ServiceProviderName *string `json:"serviceProviderName,omitempty" tf:"service_provider_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BotConnectionObservation struct {

	// The name of the Bot Resource this connection will be associated with. Changing this forces a new resource to be created.
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// The Client ID that will be used to authenticate with the service provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The ID of the Bot Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A map of additional parameters to apply to the connection.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The name of the resource group in which to create the Bot Connection. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Scopes at which the connection should be applied.
	Scopes *string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// The name of the service provider that will be associated with this connection. Changing this forces a new resource to be created.
	ServiceProviderName *string `json:"serviceProviderName,omitempty" tf:"service_provider_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BotConnectionParameters struct {

	// The name of the Bot Resource this connection will be associated with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/botservice/v1beta1.BotChannelsRegistration
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// Reference to a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameRef *v1.Reference `json:"botNameRef,omitempty" tf:"-"`

	// Selector for a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameSelector *v1.Selector `json:"botNameSelector,omitempty" tf:"-"`

	// The Client ID that will be used to authenticate with the service provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The Client Secret that will be used to authenticate with the service provider.
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A map of additional parameters to apply to the connection.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The name of the resource group in which to create the Bot Connection. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Scopes at which the connection should be applied.
	Scopes *string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// The name of the service provider that will be associated with this connection. Changing this forces a new resource to be created.
	ServiceProviderName *string `json:"serviceProviderName,omitempty" tf:"service_provider_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// BotConnectionSpec defines the desired state of BotConnection
type BotConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotConnectionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider BotConnectionInitParameters `json:"initProvider,omitempty"`
}

// BotConnectionStatus defines the observed state of BotConnection.
type BotConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotConnection is the Schema for the BotConnections API. Manages a Bot Connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || has(self.initProvider.clientId)",message="clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretSecretRef)",message="clientSecretSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceProviderName) || has(self.initProvider.serviceProviderName)",message="serviceProviderName is a required parameter"
	Spec   BotConnectionSpec   `json:"spec"`
	Status BotConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotConnectionList contains a list of BotConnections
type BotConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotConnection `json:"items"`
}

// Repository type metadata.
var (
	BotConnection_Kind             = "BotConnection"
	BotConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotConnection_Kind}.String()
	BotConnection_KindAPIVersion   = BotConnection_Kind + "." + CRDGroupVersion.String()
	BotConnection_GroupVersionKind = CRDGroupVersion.WithKind(BotConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&BotConnection{}, &BotConnectionList{})
}
