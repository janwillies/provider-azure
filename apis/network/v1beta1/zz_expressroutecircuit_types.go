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

type ExpressRouteCircuitInitParameters struct {

	// Allow the circuit to interact with classic (RDFE) resources. Defaults to false.
	AllowClassicOperations *bool `json:"allowClassicOperations,omitempty" tf:"allow_classic_operations,omitempty"`

	// The bandwidth in Gbps of the circuit being created on the Express Route Port.
	BandwidthInGbps *float64 `json:"bandwidthInGbps,omitempty" tf:"bandwidth_in_gbps,omitempty"`

	// The bandwidth in Mbps of the circuit being created on the Service Provider.
	BandwidthInMbps *float64 `json:"bandwidthInMbps,omitempty" tf:"bandwidth_in_mbps,omitempty"`

	// The ID of the Express Route Port this Express Route Circuit is based on. Changing this forces a new resource to be created.
	ExpressRoutePortID *string `json:"expressRoutePortId,omitempty" tf:"express_route_port_id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the peering location and not the Azure resource location. Changing this forces a new resource to be created.
	PeeringLocation *string `json:"peeringLocation,omitempty" tf:"peering_location,omitempty"`

	// The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
	ServiceProviderName *string `json:"serviceProviderName,omitempty" tf:"service_provider_name,omitempty"`

	// A sku block for the ExpressRoute circuit as documented below.
	Sku []ExpressRouteCircuitSkuInitParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ExpressRouteCircuitObservation struct {

	// Allow the circuit to interact with classic (RDFE) resources. Defaults to false.
	AllowClassicOperations *bool `json:"allowClassicOperations,omitempty" tf:"allow_classic_operations,omitempty"`

	// The bandwidth in Gbps of the circuit being created on the Express Route Port.
	BandwidthInGbps *float64 `json:"bandwidthInGbps,omitempty" tf:"bandwidth_in_gbps,omitempty"`

	// The bandwidth in Mbps of the circuit being created on the Service Provider.
	BandwidthInMbps *float64 `json:"bandwidthInMbps,omitempty" tf:"bandwidth_in_mbps,omitempty"`

	// The ID of the Express Route Port this Express Route Circuit is based on. Changing this forces a new resource to be created.
	ExpressRoutePortID *string `json:"expressRoutePortId,omitempty" tf:"express_route_port_id,omitempty"`

	// The ID of the ExpressRoute circuit.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the peering location and not the Azure resource location. Changing this forces a new resource to be created.
	PeeringLocation *string `json:"peeringLocation,omitempty" tf:"peering_location,omitempty"`

	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
	ServiceProviderName *string `json:"serviceProviderName,omitempty" tf:"service_provider_name,omitempty"`

	// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are NotProvisioned, Provisioning, Provisioned, and Deprovisioning.
	ServiceProviderProvisioningState *string `json:"serviceProviderProvisioningState,omitempty" tf:"service_provider_provisioning_state,omitempty"`

	// A sku block for the ExpressRoute circuit as documented below.
	Sku []ExpressRouteCircuitSkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ExpressRouteCircuitParameters struct {

	// Allow the circuit to interact with classic (RDFE) resources. Defaults to false.
	AllowClassicOperations *bool `json:"allowClassicOperations,omitempty" tf:"allow_classic_operations,omitempty"`

	// The authorization key. This can be used to set up an ExpressRoute Circuit with an ExpressRoute Port from another subscription.
	AuthorizationKeySecretRef *v1.SecretKeySelector `json:"authorizationKeySecretRef,omitempty" tf:"-"`

	// The bandwidth in Gbps of the circuit being created on the Express Route Port.
	BandwidthInGbps *float64 `json:"bandwidthInGbps,omitempty" tf:"bandwidth_in_gbps,omitempty"`

	// The bandwidth in Mbps of the circuit being created on the Service Provider.
	BandwidthInMbps *float64 `json:"bandwidthInMbps,omitempty" tf:"bandwidth_in_mbps,omitempty"`

	// The ID of the Express Route Port this Express Route Circuit is based on. Changing this forces a new resource to be created.
	ExpressRoutePortID *string `json:"expressRoutePortId,omitempty" tf:"express_route_port_id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the peering location and not the Azure resource location. Changing this forces a new resource to be created.
	PeeringLocation *string `json:"peeringLocation,omitempty" tf:"peering_location,omitempty"`

	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
	ServiceProviderName *string `json:"serviceProviderName,omitempty" tf:"service_provider_name,omitempty"`

	// A sku block for the ExpressRoute circuit as documented below.
	Sku []ExpressRouteCircuitSkuParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ExpressRouteCircuitSkuInitParameters struct {

	// The billing mode for bandwidth. Possible values are MeteredData or UnlimitedData.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The service tier. Possible values are Basic, Local, Standard or Premium.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type ExpressRouteCircuitSkuObservation struct {

	// The billing mode for bandwidth. Possible values are MeteredData or UnlimitedData.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The service tier. Possible values are Basic, Local, Standard or Premium.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type ExpressRouteCircuitSkuParameters struct {

	// The billing mode for bandwidth. Possible values are MeteredData or UnlimitedData.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The service tier. Possible values are Basic, Local, Standard or Premium.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

// ExpressRouteCircuitSpec defines the desired state of ExpressRouteCircuit
type ExpressRouteCircuitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteCircuitParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ExpressRouteCircuitInitParameters `json:"initProvider,omitempty"`
}

// ExpressRouteCircuitStatus defines the observed state of ExpressRouteCircuit.
type ExpressRouteCircuitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteCircuitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuit is the Schema for the ExpressRouteCircuits API. Manages an ExpressRoute circuit.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteCircuit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || has(self.initProvider.sku)",message="sku is a required parameter"
	Spec   ExpressRouteCircuitSpec   `json:"spec"`
	Status ExpressRouteCircuitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitList contains a list of ExpressRouteCircuits
type ExpressRouteCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteCircuit `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteCircuit_Kind             = "ExpressRouteCircuit"
	ExpressRouteCircuit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRouteCircuit_Kind}.String()
	ExpressRouteCircuit_KindAPIVersion   = ExpressRouteCircuit_Kind + "." + CRDGroupVersion.String()
	ExpressRouteCircuit_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRouteCircuit_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteCircuit{}, &ExpressRouteCircuitList{})
}
