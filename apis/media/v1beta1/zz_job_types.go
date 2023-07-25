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

type InputAssetInitParameters struct {

	// A label that is assigned to a JobInputClip, that is used to satisfy a reference used in the Transform. For example, a Transform can be authored so as to take an image file with the label 'xyz' and apply it as an overlay onto the input video before it is encoded. When submitting a Job, exactly one of the JobInputs should be the image file, and it should have the label 'xyz'. Changing this forces a new resource to be created.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`
}

type InputAssetObservation struct {

	// A label that is assigned to a JobInputClip, that is used to satisfy a reference used in the Transform. For example, a Transform can be authored so as to take an image file with the label 'xyz' and apply it as an overlay onto the input video before it is encoded. When submitting a Job, exactly one of the JobInputs should be the image file, and it should have the label 'xyz'. Changing this forces a new resource to be created.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The name of the input Asset. Changing this forces a new Media Job to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type InputAssetParameters struct {

	// A label that is assigned to a JobInputClip, that is used to satisfy a reference used in the Transform. For example, a Transform can be authored so as to take an image file with the label 'xyz' and apply it as an overlay onto the input video before it is encoded. When submitting a Job, exactly one of the JobInputs should be the image file, and it should have the label 'xyz'. Changing this forces a new resource to be created.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The name of the input Asset. Changing this forces a new Media Job to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/media/v1beta1.Asset
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Asset in media to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Asset in media to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

type JobInitParameters struct {

	// Optional customer supplied description of the Job.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A input_asset block as defined below. Changing this forces a new Media Job to be created.
	InputAsset []InputAssetInitParameters `json:"inputAsset,omitempty" tf:"input_asset,omitempty"`

	// One or more output_asset blocks as defined below. Changing this forces a new Media Job to be created.
	OutputAsset []OutputAssetInitParameters `json:"outputAsset,omitempty" tf:"output_asset,omitempty"`

	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal. Changing this forces a new Media Job to be created. Possible values are High, Normal and Low.
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`
}

type JobObservation struct {

	// Optional customer supplied description of the Job.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Media Job.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A input_asset block as defined below. Changing this forces a new Media Job to be created.
	InputAsset []InputAssetObservation `json:"inputAsset,omitempty" tf:"input_asset,omitempty"`

	// The Media Services account name. Changing this forces a new Transform to be created.
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// One or more output_asset blocks as defined below. Changing this forces a new Media Job to be created.
	OutputAsset []OutputAssetObservation `json:"outputAsset,omitempty" tf:"output_asset,omitempty"`

	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal. Changing this forces a new Media Job to be created. Possible values are High, Normal and Low.
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// The name of the Resource Group where the Media Job should exist. Changing this forces a new Media Job to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Transform name. Changing this forces a new Media Job to be created.
	TransformName *string `json:"transformName,omitempty" tf:"transform_name,omitempty"`
}

type JobParameters struct {

	// Optional customer supplied description of the Job.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A input_asset block as defined below. Changing this forces a new Media Job to be created.
	InputAsset []InputAssetParameters `json:"inputAsset,omitempty" tf:"input_asset,omitempty"`

	// The Media Services account name. Changing this forces a new Transform to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/media/v1beta1.ServicesAccount
	// +kubebuilder:validation:Optional
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// Reference to a ServicesAccount in media to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameRef *v1.Reference `json:"mediaServicesAccountNameRef,omitempty" tf:"-"`

	// Selector for a ServicesAccount in media to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameSelector *v1.Selector `json:"mediaServicesAccountNameSelector,omitempty" tf:"-"`

	// One or more output_asset blocks as defined below. Changing this forces a new Media Job to be created.
	OutputAsset []OutputAssetParameters `json:"outputAsset,omitempty" tf:"output_asset,omitempty"`

	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal. Changing this forces a new Media Job to be created. Possible values are High, Normal and Low.
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// The name of the Resource Group where the Media Job should exist. Changing this forces a new Media Job to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Transform name. Changing this forces a new Media Job to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/media/v1beta1.Transform
	// +kubebuilder:validation:Optional
	TransformName *string `json:"transformName,omitempty" tf:"transform_name,omitempty"`

	// Reference to a Transform in media to populate transformName.
	// +kubebuilder:validation:Optional
	TransformNameRef *v1.Reference `json:"transformNameRef,omitempty" tf:"-"`

	// Selector for a Transform in media to populate transformName.
	// +kubebuilder:validation:Optional
	TransformNameSelector *v1.Selector `json:"transformNameSelector,omitempty" tf:"-"`
}

type OutputAssetInitParameters struct {

	// A label that is assigned to a JobOutput in order to help uniquely identify it. This is useful when your Transform has more than one TransformOutput, whereby your Job has more than one JobOutput. In such cases, when you submit the Job, you will add two or more JobOutputs, in the same order as TransformOutputs in the Transform. Subsequently, when you retrieve the Job, either through events or on a GET request, you can use the label to easily identify the JobOutput. If a label is not provided, a default value of '{presetName}_{outputIndex}' will be used, where the preset name is the name of the preset in the corresponding TransformOutput and the output index is the relative index of the this JobOutput within the Job. Note that this index is the same as the relative index of the corresponding TransformOutput within its Transform. Changing this forces a new resource to be created.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`
}

type OutputAssetObservation struct {

	// A label that is assigned to a JobOutput in order to help uniquely identify it. This is useful when your Transform has more than one TransformOutput, whereby your Job has more than one JobOutput. In such cases, when you submit the Job, you will add two or more JobOutputs, in the same order as TransformOutputs in the Transform. Subsequently, when you retrieve the Job, either through events or on a GET request, you can use the label to easily identify the JobOutput. If a label is not provided, a default value of '{presetName}_{outputIndex}' will be used, where the preset name is the name of the preset in the corresponding TransformOutput and the output index is the relative index of the this JobOutput within the Job. Note that this index is the same as the relative index of the corresponding TransformOutput within its Transform. Changing this forces a new resource to be created.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The name of the output Asset. Changing this forces a new Media Job to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type OutputAssetParameters struct {

	// A label that is assigned to a JobOutput in order to help uniquely identify it. This is useful when your Transform has more than one TransformOutput, whereby your Job has more than one JobOutput. In such cases, when you submit the Job, you will add two or more JobOutputs, in the same order as TransformOutputs in the Transform. Subsequently, when you retrieve the Job, either through events or on a GET request, you can use the label to easily identify the JobOutput. If a label is not provided, a default value of '{presetName}_{outputIndex}' will be used, where the preset name is the name of the preset in the corresponding TransformOutput and the output index is the relative index of the this JobOutput within the Job. Note that this index is the same as the relative index of the corresponding TransformOutput within its Transform. Changing this forces a new resource to be created.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The name of the output Asset. Changing this forces a new Media Job to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/media/v1beta1.Asset
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Asset in media to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Asset in media to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

// JobSpec defines the desired state of Job
type JobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider JobInitParameters `json:"initProvider,omitempty"`
}

// JobStatus defines the observed state of Job.
type JobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Job is the Schema for the Jobs API. Manages a Media Job.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.inputAsset) || has(self.initProvider.inputAsset)",message="inputAsset is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.outputAsset) || has(self.initProvider.outputAsset)",message="outputAsset is a required parameter"
	Spec   JobSpec   `json:"spec"`
	Status JobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobList contains a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

// Repository type metadata.
var (
	Job_Kind             = "Job"
	Job_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Job_Kind}.String()
	Job_KindAPIVersion   = Job_Kind + "." + CRDGroupVersion.String()
	Job_GroupVersionKind = CRDGroupVersion.WithKind(Job_Kind)
)

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}
