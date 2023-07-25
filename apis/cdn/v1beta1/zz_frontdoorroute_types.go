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

type CacheInitParameters struct {

	// Is content compression enabled? Possible values are true or false. Defaults to false.
	CompressionEnabled *bool `json:"compressionEnabled,omitempty" tf:"compression_enabled,omitempty"`

	// A list of one or more Content types (formerly known as MIME types) to compress. Possible values include application/eot, application/font, application/font-sfnt, application/javascript, application/json, application/opentype, application/otf, application/pkcs7-mime, application/truetype, application/ttf, application/vnd.ms-fontobject, application/xhtml+xml, application/xml, application/xml+rss, application/x-font-opentype, application/x-font-truetype, application/x-font-ttf, application/x-httpd-cgi, application/x-mpegurl, application/x-opentype, application/x-otf, application/x-perl, application/x-ttf, application/x-javascript, font/eot, font/ttf, font/otf, font/opentype, image/svg+xml, text/css, text/csv, text/html, text/javascript, text/js, text/plain, text/richtext, text/tab-separated-values, text/xml, text/x-script, text/x-component or text/x-java-source.
	ContentTypesToCompress []*string `json:"contentTypesToCompress,omitempty" tf:"content_types_to_compress,omitempty"`

	// Defines how the Front Door Route will cache requests that include query strings. Possible values include IgnoreQueryString, IgnoreSpecifiedQueryStrings, IncludeSpecifiedQueryStrings or UseQueryString. Defaults it IgnoreQueryString.
	QueryStringCachingBehavior *string `json:"queryStringCachingBehavior,omitempty" tf:"query_string_caching_behavior,omitempty"`

	// Query strings to include or ignore.
	QueryStrings []*string `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type CacheObservation struct {

	// Is content compression enabled? Possible values are true or false. Defaults to false.
	CompressionEnabled *bool `json:"compressionEnabled,omitempty" tf:"compression_enabled,omitempty"`

	// A list of one or more Content types (formerly known as MIME types) to compress. Possible values include application/eot, application/font, application/font-sfnt, application/javascript, application/json, application/opentype, application/otf, application/pkcs7-mime, application/truetype, application/ttf, application/vnd.ms-fontobject, application/xhtml+xml, application/xml, application/xml+rss, application/x-font-opentype, application/x-font-truetype, application/x-font-ttf, application/x-httpd-cgi, application/x-mpegurl, application/x-opentype, application/x-otf, application/x-perl, application/x-ttf, application/x-javascript, font/eot, font/ttf, font/otf, font/opentype, image/svg+xml, text/css, text/csv, text/html, text/javascript, text/js, text/plain, text/richtext, text/tab-separated-values, text/xml, text/x-script, text/x-component or text/x-java-source.
	ContentTypesToCompress []*string `json:"contentTypesToCompress,omitempty" tf:"content_types_to_compress,omitempty"`

	// Defines how the Front Door Route will cache requests that include query strings. Possible values include IgnoreQueryString, IgnoreSpecifiedQueryStrings, IncludeSpecifiedQueryStrings or UseQueryString. Defaults it IgnoreQueryString.
	QueryStringCachingBehavior *string `json:"queryStringCachingBehavior,omitempty" tf:"query_string_caching_behavior,omitempty"`

	// Query strings to include or ignore.
	QueryStrings []*string `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type CacheParameters struct {

	// Is content compression enabled? Possible values are true or false. Defaults to false.
	CompressionEnabled *bool `json:"compressionEnabled,omitempty" tf:"compression_enabled,omitempty"`

	// A list of one or more Content types (formerly known as MIME types) to compress. Possible values include application/eot, application/font, application/font-sfnt, application/javascript, application/json, application/opentype, application/otf, application/pkcs7-mime, application/truetype, application/ttf, application/vnd.ms-fontobject, application/xhtml+xml, application/xml, application/xml+rss, application/x-font-opentype, application/x-font-truetype, application/x-font-ttf, application/x-httpd-cgi, application/x-mpegurl, application/x-opentype, application/x-otf, application/x-perl, application/x-ttf, application/x-javascript, font/eot, font/ttf, font/otf, font/opentype, image/svg+xml, text/css, text/csv, text/html, text/javascript, text/js, text/plain, text/richtext, text/tab-separated-values, text/xml, text/x-script, text/x-component or text/x-java-source.
	ContentTypesToCompress []*string `json:"contentTypesToCompress,omitempty" tf:"content_types_to_compress,omitempty"`

	// Defines how the Front Door Route will cache requests that include query strings. Possible values include IgnoreQueryString, IgnoreSpecifiedQueryStrings, IncludeSpecifiedQueryStrings or UseQueryString. Defaults it IgnoreQueryString.
	QueryStringCachingBehavior *string `json:"queryStringCachingBehavior,omitempty" tf:"query_string_caching_behavior,omitempty"`

	// Query strings to include or ignore.
	QueryStrings []*string `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type FrontdoorRouteInitParameters struct {

	// A cache block as defined below.
	Cache []CacheInitParameters `json:"cache,omitempty" tf:"cache,omitempty"`

	// A directory path on the Front Door Origin that can be used to retrieve content (e.g. contoso.cloudapp.net/originpath).
	CdnFrontdoorOriginPath *string `json:"cdnFrontdoorOriginPath,omitempty" tf:"cdn_frontdoor_origin_path,omitempty"`

	// Is this Front Door Route enabled? Possible values are true or false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Protocol that will be use when forwarding traffic to backends. Possible values are HttpOnly, HttpsOnly or MatchRequest.
	ForwardingProtocol *string `json:"forwardingProtocol,omitempty" tf:"forwarding_protocol,omitempty"`

	// Automatically redirect HTTP traffic to HTTPS traffic? Possible values are true or false. Defaults to true.
	HTTPSRedirectEnabled *bool `json:"httpsRedirectEnabled,omitempty" tf:"https_redirect_enabled,omitempty"`

	// Should this Front Door Route be linked to the default endpoint? Possible values include true or false. Defaults to true.
	LinkToDefaultDomain *bool `json:"linkToDefaultDomain,omitempty" tf:"link_to_default_domain,omitempty"`

	// The route patterns of the rule.
	PatternsToMatch []*string `json:"patternsToMatch,omitempty" tf:"patterns_to_match,omitempty"`

	// One or more Protocols supported by this Front Door Route. Possible values are Http or Https.
	SupportedProtocols []*string `json:"supportedProtocols,omitempty" tf:"supported_protocols,omitempty"`
}

type FrontdoorRouteObservation struct {

	// A cache block as defined below.
	Cache []CacheObservation `json:"cache,omitempty" tf:"cache,omitempty"`

	// The IDs of the Front Door Custom Domains which are associated with this Front Door Route.
	CdnFrontdoorCustomDomainIds []*string `json:"cdnFrontdoorCustomDomainIds,omitempty" tf:"cdn_frontdoor_custom_domain_ids,omitempty"`

	// The resource ID of the Front Door Endpoint where this Front Door Route should exist. Changing this forces a new Front Door Route to be created.
	CdnFrontdoorEndpointID *string `json:"cdnFrontdoorEndpointId,omitempty" tf:"cdn_frontdoor_endpoint_id,omitempty"`

	// The resource ID of the Front Door Origin Group where this Front Door Route should be created.
	CdnFrontdoorOriginGroupID *string `json:"cdnFrontdoorOriginGroupId,omitempty" tf:"cdn_frontdoor_origin_group_id,omitempty"`

	// One or more Front Door Origin resource IDs that this Front Door Route will link to.
	CdnFrontdoorOriginIds []*string `json:"cdnFrontdoorOriginIds,omitempty" tf:"cdn_frontdoor_origin_ids,omitempty"`

	// A directory path on the Front Door Origin that can be used to retrieve content (e.g. contoso.cloudapp.net/originpath).
	CdnFrontdoorOriginPath *string `json:"cdnFrontdoorOriginPath,omitempty" tf:"cdn_frontdoor_origin_path,omitempty"`

	// A list of the Front Door Rule Set IDs which should be assigned to this Front Door Route.
	CdnFrontdoorRuleSetIds []*string `json:"cdnFrontdoorRuleSetIds,omitempty" tf:"cdn_frontdoor_rule_set_ids,omitempty"`

	// Is this Front Door Route enabled? Possible values are true or false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Protocol that will be use when forwarding traffic to backends. Possible values are HttpOnly, HttpsOnly or MatchRequest.
	ForwardingProtocol *string `json:"forwardingProtocol,omitempty" tf:"forwarding_protocol,omitempty"`

	// Automatically redirect HTTP traffic to HTTPS traffic? Possible values are true or false. Defaults to true.
	HTTPSRedirectEnabled *bool `json:"httpsRedirectEnabled,omitempty" tf:"https_redirect_enabled,omitempty"`

	// The ID of the Front Door Route.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Should this Front Door Route be linked to the default endpoint? Possible values include true or false. Defaults to true.
	LinkToDefaultDomain *bool `json:"linkToDefaultDomain,omitempty" tf:"link_to_default_domain,omitempty"`

	// The route patterns of the rule.
	PatternsToMatch []*string `json:"patternsToMatch,omitempty" tf:"patterns_to_match,omitempty"`

	// One or more Protocols supported by this Front Door Route. Possible values are Http or Https.
	SupportedProtocols []*string `json:"supportedProtocols,omitempty" tf:"supported_protocols,omitempty"`
}

type FrontdoorRouteParameters struct {

	// A cache block as defined below.
	Cache []CacheParameters `json:"cache,omitempty" tf:"cache,omitempty"`

	// The IDs of the Front Door Custom Domains which are associated with this Front Door Route.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorCustomDomain
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorCustomDomainIds []*string `json:"cdnFrontdoorCustomDomainIds,omitempty" tf:"cdn_frontdoor_custom_domain_ids,omitempty"`

	// References to FrontdoorCustomDomain in cdn to populate cdnFrontdoorCustomDomainIds.
	// +kubebuilder:validation:Optional
	CdnFrontdoorCustomDomainIdsRefs []v1.Reference `json:"cdnFrontdoorCustomDomainIdsRefs,omitempty" tf:"-"`

	// Selector for a list of FrontdoorCustomDomain in cdn to populate cdnFrontdoorCustomDomainIds.
	// +kubebuilder:validation:Optional
	CdnFrontdoorCustomDomainIdsSelector *v1.Selector `json:"cdnFrontdoorCustomDomainIdsSelector,omitempty" tf:"-"`

	// The resource ID of the Front Door Endpoint where this Front Door Route should exist. Changing this forces a new Front Door Route to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorEndpoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorEndpointID *string `json:"cdnFrontdoorEndpointId,omitempty" tf:"cdn_frontdoor_endpoint_id,omitempty"`

	// Reference to a FrontdoorEndpoint in cdn to populate cdnFrontdoorEndpointId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorEndpointIDRef *v1.Reference `json:"cdnFrontdoorEndpointIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorEndpoint in cdn to populate cdnFrontdoorEndpointId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorEndpointIDSelector *v1.Selector `json:"cdnFrontdoorEndpointIdSelector,omitempty" tf:"-"`

	// The resource ID of the Front Door Origin Group where this Front Door Route should be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorOriginGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorOriginGroupID *string `json:"cdnFrontdoorOriginGroupId,omitempty" tf:"cdn_frontdoor_origin_group_id,omitempty"`

	// Reference to a FrontdoorOriginGroup in cdn to populate cdnFrontdoorOriginGroupId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorOriginGroupIDRef *v1.Reference `json:"cdnFrontdoorOriginGroupIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorOriginGroup in cdn to populate cdnFrontdoorOriginGroupId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorOriginGroupIDSelector *v1.Selector `json:"cdnFrontdoorOriginGroupIdSelector,omitempty" tf:"-"`

	// One or more Front Door Origin resource IDs that this Front Door Route will link to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorOrigin
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorOriginIds []*string `json:"cdnFrontdoorOriginIds,omitempty" tf:"cdn_frontdoor_origin_ids,omitempty"`

	// References to FrontdoorOrigin in cdn to populate cdnFrontdoorOriginIds.
	// +kubebuilder:validation:Optional
	CdnFrontdoorOriginIdsRefs []v1.Reference `json:"cdnFrontdoorOriginIdsRefs,omitempty" tf:"-"`

	// Selector for a list of FrontdoorOrigin in cdn to populate cdnFrontdoorOriginIds.
	// +kubebuilder:validation:Optional
	CdnFrontdoorOriginIdsSelector *v1.Selector `json:"cdnFrontdoorOriginIdsSelector,omitempty" tf:"-"`

	// A directory path on the Front Door Origin that can be used to retrieve content (e.g. contoso.cloudapp.net/originpath).
	CdnFrontdoorOriginPath *string `json:"cdnFrontdoorOriginPath,omitempty" tf:"cdn_frontdoor_origin_path,omitempty"`

	// A list of the Front Door Rule Set IDs which should be assigned to this Front Door Route.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorRuleSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorRuleSetIds []*string `json:"cdnFrontdoorRuleSetIds,omitempty" tf:"cdn_frontdoor_rule_set_ids,omitempty"`

	// References to FrontdoorRuleSet in cdn to populate cdnFrontdoorRuleSetIds.
	// +kubebuilder:validation:Optional
	CdnFrontdoorRuleSetIdsRefs []v1.Reference `json:"cdnFrontdoorRuleSetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of FrontdoorRuleSet in cdn to populate cdnFrontdoorRuleSetIds.
	// +kubebuilder:validation:Optional
	CdnFrontdoorRuleSetIdsSelector *v1.Selector `json:"cdnFrontdoorRuleSetIdsSelector,omitempty" tf:"-"`

	// Is this Front Door Route enabled? Possible values are true or false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Protocol that will be use when forwarding traffic to backends. Possible values are HttpOnly, HttpsOnly or MatchRequest.
	ForwardingProtocol *string `json:"forwardingProtocol,omitempty" tf:"forwarding_protocol,omitempty"`

	// Automatically redirect HTTP traffic to HTTPS traffic? Possible values are true or false. Defaults to true.
	HTTPSRedirectEnabled *bool `json:"httpsRedirectEnabled,omitempty" tf:"https_redirect_enabled,omitempty"`

	// Should this Front Door Route be linked to the default endpoint? Possible values include true or false. Defaults to true.
	LinkToDefaultDomain *bool `json:"linkToDefaultDomain,omitempty" tf:"link_to_default_domain,omitempty"`

	// The route patterns of the rule.
	PatternsToMatch []*string `json:"patternsToMatch,omitempty" tf:"patterns_to_match,omitempty"`

	// One or more Protocols supported by this Front Door Route. Possible values are Http or Https.
	SupportedProtocols []*string `json:"supportedProtocols,omitempty" tf:"supported_protocols,omitempty"`
}

// FrontdoorRouteSpec defines the desired state of FrontdoorRoute
type FrontdoorRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorRouteParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider FrontdoorRouteInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorRouteStatus defines the observed state of FrontdoorRoute.
type FrontdoorRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRoute is the Schema for the FrontdoorRoutes API. Manages a Front Door (standard/premium) Route.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.patternsToMatch) || has(self.initProvider.patternsToMatch)",message="patternsToMatch is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.supportedProtocols) || has(self.initProvider.supportedProtocols)",message="supportedProtocols is a required parameter"
	Spec   FrontdoorRouteSpec   `json:"spec"`
	Status FrontdoorRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRouteList contains a list of FrontdoorRoutes
type FrontdoorRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorRoute `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorRoute_Kind             = "FrontdoorRoute"
	FrontdoorRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorRoute_Kind}.String()
	FrontdoorRoute_KindAPIVersion   = FrontdoorRoute_Kind + "." + CRDGroupVersion.String()
	FrontdoorRoute_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorRoute{}, &FrontdoorRouteList{})
}
