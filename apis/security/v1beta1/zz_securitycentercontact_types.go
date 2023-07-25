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

type SecurityCenterContactInitParameters struct {

	// Whether to send security alerts notifications to the security contact.
	AlertNotifications *bool `json:"alertNotifications,omitempty" tf:"alert_notifications,omitempty"`

	// Whether to send security alerts notifications to subscription admins.
	AlertsToAdmins *bool `json:"alertsToAdmins,omitempty" tf:"alerts_to_admins,omitempty"`

	// The email of the Security Center Contact.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The name of the Security Center Contact. Defaults to default1.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The phone number of the Security Center Contact.
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`
}

type SecurityCenterContactObservation struct {

	// Whether to send security alerts notifications to the security contact.
	AlertNotifications *bool `json:"alertNotifications,omitempty" tf:"alert_notifications,omitempty"`

	// Whether to send security alerts notifications to subscription admins.
	AlertsToAdmins *bool `json:"alertsToAdmins,omitempty" tf:"alerts_to_admins,omitempty"`

	// The email of the Security Center Contact.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The Security Center Contact ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Security Center Contact. Defaults to default1.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The phone number of the Security Center Contact.
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`
}

type SecurityCenterContactParameters struct {

	// Whether to send security alerts notifications to the security contact.
	AlertNotifications *bool `json:"alertNotifications,omitempty" tf:"alert_notifications,omitempty"`

	// Whether to send security alerts notifications to subscription admins.
	AlertsToAdmins *bool `json:"alertsToAdmins,omitempty" tf:"alerts_to_admins,omitempty"`

	// The email of the Security Center Contact.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The name of the Security Center Contact. Defaults to default1.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The phone number of the Security Center Contact.
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`
}

// SecurityCenterContactSpec defines the desired state of SecurityCenterContact
type SecurityCenterContactSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityCenterContactParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SecurityCenterContactInitParameters `json:"initProvider,omitempty"`
}

// SecurityCenterContactStatus defines the observed state of SecurityCenterContact.
type SecurityCenterContactStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityCenterContactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterContact is the Schema for the SecurityCenterContacts API. Manages the subscription's Security Center Contact.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterContact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alertNotifications) || has(self.initProvider.alertNotifications)",message="alertNotifications is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alertsToAdmins) || has(self.initProvider.alertsToAdmins)",message="alertsToAdmins is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || has(self.initProvider.email)",message="email is a required parameter"
	Spec   SecurityCenterContactSpec   `json:"spec"`
	Status SecurityCenterContactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterContactList contains a list of SecurityCenterContacts
type SecurityCenterContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterContact `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterContact_Kind             = "SecurityCenterContact"
	SecurityCenterContact_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityCenterContact_Kind}.String()
	SecurityCenterContact_KindAPIVersion   = SecurityCenterContact_Kind + "." + CRDGroupVersion.String()
	SecurityCenterContact_GroupVersionKind = CRDGroupVersion.WithKind(SecurityCenterContact_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterContact{}, &SecurityCenterContactList{})
}
