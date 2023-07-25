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

type RedisFirewallRuleInitParameters struct {

	// The highest IP address included in the range.
	EndIP *string `json:"endIp,omitempty" tf:"end_ip,omitempty"`

	// The lowest IP address included in the range
	StartIP *string `json:"startIp,omitempty" tf:"start_ip,omitempty"`
}

type RedisFirewallRuleObservation struct {

	// The highest IP address included in the range.
	EndIP *string `json:"endIp,omitempty" tf:"end_ip,omitempty"`

	// The ID of the Redis Firewall Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Redis Cache. Changing this forces a new resource to be created.
	RedisCacheName *string `json:"redisCacheName,omitempty" tf:"redis_cache_name,omitempty"`

	// The name of the resource group in which this Redis Cache exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The lowest IP address included in the range
	StartIP *string `json:"startIp,omitempty" tf:"start_ip,omitempty"`
}

type RedisFirewallRuleParameters struct {

	// The highest IP address included in the range.
	EndIP *string `json:"endIp,omitempty" tf:"end_ip,omitempty"`

	// The name of the Redis Cache. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cache/v1beta1.RedisCache
	// +kubebuilder:validation:Optional
	RedisCacheName *string `json:"redisCacheName,omitempty" tf:"redis_cache_name,omitempty"`

	// Reference to a RedisCache in cache to populate redisCacheName.
	// +kubebuilder:validation:Optional
	RedisCacheNameRef *v1.Reference `json:"redisCacheNameRef,omitempty" tf:"-"`

	// Selector for a RedisCache in cache to populate redisCacheName.
	// +kubebuilder:validation:Optional
	RedisCacheNameSelector *v1.Selector `json:"redisCacheNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which this Redis Cache exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The lowest IP address included in the range
	StartIP *string `json:"startIp,omitempty" tf:"start_ip,omitempty"`
}

// RedisFirewallRuleSpec defines the desired state of RedisFirewallRule
type RedisFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisFirewallRuleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider RedisFirewallRuleInitParameters `json:"initProvider,omitempty"`
}

// RedisFirewallRuleStatus defines the observed state of RedisFirewallRule.
type RedisFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisFirewallRule is the Schema for the RedisFirewallRules API. Manages a Firewall Rule associated with a Redis Cache.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endIp) || has(self.initProvider.endIp)",message="endIp is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.startIp) || has(self.initProvider.startIp)",message="startIp is a required parameter"
	Spec   RedisFirewallRuleSpec   `json:"spec"`
	Status RedisFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisFirewallRuleList contains a list of RedisFirewallRules
type RedisFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	RedisFirewallRule_Kind             = "RedisFirewallRule"
	RedisFirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisFirewallRule_Kind}.String()
	RedisFirewallRule_KindAPIVersion   = RedisFirewallRule_Kind + "." + CRDGroupVersion.String()
	RedisFirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(RedisFirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisFirewallRule{}, &RedisFirewallRuleList{})
}
