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

type PrivateDNSCNAMERecordObservation struct {

	// The FQDN of the DNS CNAME Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The Private DNS CNAME Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateDNSCNAMERecordParameters struct {

	// The target of the CNAME.
	// +kubebuilder:validation:Required
	Record *string `json:"record" tf:"record,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Time To Live  of the DNS record in seconds.
	// +kubebuilder:validation:Required
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=PrivateDNSZone
	// +kubebuilder:validation:Optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`

	// Reference to a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameRef *v1.Reference `json:"zoneNameRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameSelector *v1.Selector `json:"zoneNameSelector,omitempty" tf:"-"`
}

// PrivateDNSCNAMERecordSpec defines the desired state of PrivateDNSCNAMERecord
type PrivateDNSCNAMERecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSCNAMERecordParameters `json:"forProvider"`
}

// PrivateDNSCNAMERecordStatus defines the observed state of PrivateDNSCNAMERecord.
type PrivateDNSCNAMERecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSCNAMERecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSCNAMERecord is the Schema for the PrivateDNSCNAMERecords API. Manages a Private DNS CNAME Record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSCNAMERecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSCNAMERecordSpec   `json:"spec"`
	Status            PrivateDNSCNAMERecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSCNAMERecordList contains a list of PrivateDNSCNAMERecords
type PrivateDNSCNAMERecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSCNAMERecord `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSCNAMERecord_Kind             = "PrivateDNSCNAMERecord"
	PrivateDNSCNAMERecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSCNAMERecord_Kind}.String()
	PrivateDNSCNAMERecord_KindAPIVersion   = PrivateDNSCNAMERecord_Kind + "." + CRDGroupVersion.String()
	PrivateDNSCNAMERecord_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSCNAMERecord_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSCNAMERecord{}, &PrivateDNSCNAMERecordList{})
}
