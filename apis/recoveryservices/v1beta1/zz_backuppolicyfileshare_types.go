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

type BackupInitParameters struct {

	// Sets the backup frequency. Currently, only Daily is supported
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// The time of day to perform the backup in 24-hour format. Times must be either on the hour or half hour (e.g. 12:00, 12:30, 13:00, etc.)
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type BackupObservation struct {

	// Sets the backup frequency. Currently, only Daily is supported
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// The time of day to perform the backup in 24-hour format. Times must be either on the hour or half hour (e.g. 12:00, 12:30, 13:00, etc.)
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type BackupParameters struct {

	// Sets the backup frequency. Currently, only Daily is supported
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// The time of day to perform the backup in 24-hour format. Times must be either on the hour or half hour (e.g. 12:00, 12:30, 13:00, etc.)
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type BackupPolicyFileShareInitParameters struct {

	// Configures the Policy backup frequency and times as documented in the backup block below.
	Backup []BackupInitParameters `json:"backup,omitempty" tf:"backup,omitempty"`

	// Configures the policy daily retention as documented in the retention_daily block below.
	RetentionDaily []RetentionDailyInitParameters `json:"retentionDaily,omitempty" tf:"retention_daily,omitempty"`

	// Configures the policy monthly retention as documented in the retention_monthly block below.
	RetentionMonthly []RetentionMonthlyInitParameters `json:"retentionMonthly,omitempty" tf:"retention_monthly,omitempty"`

	// Configures the policy weekly retention as documented in the retention_weekly block below.
	RetentionWeekly []RetentionWeeklyInitParameters `json:"retentionWeekly,omitempty" tf:"retention_weekly,omitempty"`

	// Configures the policy yearly retention as documented in the retention_yearly block below.
	RetentionYearly []RetentionYearlyInitParameters `json:"retentionYearly,omitempty" tf:"retention_yearly,omitempty"`

	// Specifies the timezone. the possible values are defined here. Defaults to UTC
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type BackupPolicyFileShareObservation struct {

	// Configures the Policy backup frequency and times as documented in the backup block below.
	Backup []BackupObservation `json:"backup,omitempty" tf:"backup,omitempty"`

	// The ID of the Azure File Share Backup Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Configures the policy daily retention as documented in the retention_daily block below.
	RetentionDaily []RetentionDailyObservation `json:"retentionDaily,omitempty" tf:"retention_daily,omitempty"`

	// Configures the policy monthly retention as documented in the retention_monthly block below.
	RetentionMonthly []RetentionMonthlyObservation `json:"retentionMonthly,omitempty" tf:"retention_monthly,omitempty"`

	// Configures the policy weekly retention as documented in the retention_weekly block below.
	RetentionWeekly []RetentionWeeklyObservation `json:"retentionWeekly,omitempty" tf:"retention_weekly,omitempty"`

	// Configures the policy yearly retention as documented in the retention_yearly block below.
	RetentionYearly []RetentionYearlyObservation `json:"retentionYearly,omitempty" tf:"retention_yearly,omitempty"`

	// Specifies the timezone. the possible values are defined here. Defaults to UTC
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type BackupPolicyFileShareParameters struct {

	// Configures the Policy backup frequency and times as documented in the backup block below.
	Backup []BackupParameters `json:"backup,omitempty" tf:"backup,omitempty"`

	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Configures the policy daily retention as documented in the retention_daily block below.
	RetentionDaily []RetentionDailyParameters `json:"retentionDaily,omitempty" tf:"retention_daily,omitempty"`

	// Configures the policy monthly retention as documented in the retention_monthly block below.
	RetentionMonthly []RetentionMonthlyParameters `json:"retentionMonthly,omitempty" tf:"retention_monthly,omitempty"`

	// Configures the policy weekly retention as documented in the retention_weekly block below.
	RetentionWeekly []RetentionWeeklyParameters `json:"retentionWeekly,omitempty" tf:"retention_weekly,omitempty"`

	// Configures the policy yearly retention as documented in the retention_yearly block below.
	RetentionYearly []RetentionYearlyParameters `json:"retentionYearly,omitempty" tf:"retention_yearly,omitempty"`

	// Specifies the timezone. the possible values are defined here. Defaults to UTC
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type RetentionDailyInitParameters struct {

	// The number of daily backups to keep. Must be between 1 and 200 (inclusive)
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type RetentionDailyObservation struct {

	// The number of daily backups to keep. Must be between 1 and 200 (inclusive)
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type RetentionDailyParameters struct {

	// The number of daily backups to keep. Must be between 1 and 200 (inclusive)
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type RetentionMonthlyInitParameters struct {

	// The number of monthly backups to keep. Must be between 1 and 120
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type RetentionMonthlyObservation struct {

	// The number of monthly backups to keep. Must be between 1 and 120
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type RetentionMonthlyParameters struct {

	// The number of monthly backups to keep. Must be between 1 and 120
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type RetentionWeeklyInitParameters struct {

	// The number of daily backups to keep. Must be between 1 and 200 (inclusive)
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain. Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type RetentionWeeklyObservation struct {

	// The number of daily backups to keep. Must be between 1 and 200 (inclusive)
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain. Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type RetentionWeeklyParameters struct {

	// The number of daily backups to keep. Must be between 1 and 200 (inclusive)
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain. Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type RetentionYearlyInitParameters struct {

	// The number of yearly backups to keep. Must be between 1 and 10
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The months of the year to retain backups of. Must be one of January, February, March, April, May, June, July, Augest, September, October, November and December.
	Months []*string `json:"months,omitempty" tf:"months,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type RetentionYearlyObservation struct {

	// The number of yearly backups to keep. Must be between 1 and 10
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The months of the year to retain backups of. Must be one of January, February, March, April, May, June, July, Augest, September, October, November and December.
	Months []*string `json:"months,omitempty" tf:"months,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type RetentionYearlyParameters struct {

	// The number of yearly backups to keep. Must be between 1 and 10
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The months of the year to retain backups of. Must be one of January, February, March, April, May, June, July, Augest, September, October, November and December.
	Months []*string `json:"months,omitempty" tf:"months,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

// BackupPolicyFileShareSpec defines the desired state of BackupPolicyFileShare
type BackupPolicyFileShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyFileShareParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider BackupPolicyFileShareInitParameters `json:"initProvider,omitempty"`
}

// BackupPolicyFileShareStatus defines the observed state of BackupPolicyFileShare.
type BackupPolicyFileShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyFileShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyFileShare is the Schema for the BackupPolicyFileShares API. Manages an Azure File Share Backup Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupPolicyFileShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backup) || has(self.initProvider.backup)",message="backup is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.retentionDaily) || has(self.initProvider.retentionDaily)",message="retentionDaily is a required parameter"
	Spec   BackupPolicyFileShareSpec   `json:"spec"`
	Status BackupPolicyFileShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyFileShareList contains a list of BackupPolicyFileShares
type BackupPolicyFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicyFileShare `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicyFileShare_Kind             = "BackupPolicyFileShare"
	BackupPolicyFileShare_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicyFileShare_Kind}.String()
	BackupPolicyFileShare_KindAPIVersion   = BackupPolicyFileShare_Kind + "." + CRDGroupVersion.String()
	BackupPolicyFileShare_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicyFileShare_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicyFileShare{}, &BackupPolicyFileShareList{})
}
