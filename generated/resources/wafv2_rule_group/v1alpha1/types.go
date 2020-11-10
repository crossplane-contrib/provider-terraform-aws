/*
	Copyright 2019 The Crossplane Authors.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	    http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true

// Wafv2RuleGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Wafv2RuleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Wafv2RuleGroupSpec   `json:"spec"`
	Status Wafv2RuleGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2RuleGroup contains a list of Wafv2RuleGroupList
type Wafv2RuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2RuleGroup `json:"items"`
}

// A Wafv2RuleGroupSpec defines the desired state of a Wafv2RuleGroup
type Wafv2RuleGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Wafv2RuleGroupParameters `json:",inline"`
}

// A Wafv2RuleGroupParameters defines the desired state of a Wafv2RuleGroup
type Wafv2RuleGroupParameters struct {
	Tags             map[string]string `json:"tags"`
	Capacity         int               `json:"capacity"`
	Description      string            `json:"description"`
	Name             string            `json:"name"`
	Scope            string            `json:"scope"`
	Rule             []Rule            `json:"rule"`
	VisibilityConfig VisibilityConfig  `json:"visibility_config"`
}

type Rule struct {
	Name             string           `json:"name"`
	Priority         int              `json:"priority"`
	Action           Action           `json:"action"`
	Statement        Statement        `json:"statement"`
	VisibilityConfig VisibilityConfig `json:"visibility_config"`
}

type Action struct {
	Allow Allow `json:"allow"`
	Block Block `json:"block"`
	Count Count `json:"count"`
}

type Allow struct{}

type Block struct{}

type Count struct{}

type Statement struct {
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	NotStatement                      NotStatement                      `json:"not_statement"`
	OrStatement                       OrStatement                       `json:"or_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	AndStatement                      AndStatement                      `json:"and_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
}

type NotStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	OrStatement                       OrStatement                       `json:"or_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	NotStatement                      NotStatement                      `json:"not_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
	AndStatement                      AndStatement                      `json:"and_statement"`
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
}

type OrStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
}

type NotStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type AndStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type OrStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	AndStatement                      AndStatement                      `json:"and_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	NotStatement                      NotStatement                      `json:"not_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	OrStatement                       OrStatement                       `json:"or_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type AndStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
	FallbackBehavior string `json:"fallback_behavior"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	TextTransformation []TextTransformation `json:"text_transformation"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type FieldToMatch struct {
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type NotStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SizeConstraintStatement struct {
	Size               int                  `json:"size"`
	ComparisonOperator string               `json:"comparison_operator"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type FieldToMatch struct {
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	TextTransformation []TextTransformation `json:"text_transformation"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type ByteMatchStatement struct {
	SearchString         string               `json:"search_string"`
	PositionalConstraint string               `json:"positional_constraint"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
}

type OrStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
	FallbackBehavior string `json:"fallback_behavior"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SqliMatchStatement struct {
	TextTransformation []TextTransformation `json:"text_transformation"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SqliMatchStatement struct {
	TextTransformation []TextTransformation `json:"text_transformation"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type AndStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	OrStatement                       OrStatement                       `json:"or_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
	AndStatement                      AndStatement                      `json:"and_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	NotStatement                      NotStatement                      `json:"not_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type OrStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	HeaderName       string `json:"header_name"`
	FallbackBehavior string `json:"fallback_behavior"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
	FallbackBehavior string `json:"fallback_behavior"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type XssMatchStatement struct {
	TextTransformation []TextTransformation `json:"text_transformation"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type FieldToMatch struct {
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type AndStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	HeaderName       string `json:"header_name"`
	FallbackBehavior string `json:"fallback_behavior"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	Position         string `json:"position"`
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
}

type NotStatement struct {
	Statement []Statement `json:"statement"`
}

type Statement struct {
	ByteMatchStatement                ByteMatchStatement                `json:"byte_match_statement"`
	GeoMatchStatement                 GeoMatchStatement                 `json:"geo_match_statement"`
	IpSetReferenceStatement           IpSetReferenceStatement           `json:"ip_set_reference_statement"`
	RegexPatternSetReferenceStatement RegexPatternSetReferenceStatement `json:"regex_pattern_set_reference_statement"`
	SizeConstraintStatement           SizeConstraintStatement           `json:"size_constraint_statement"`
	SqliMatchStatement                SqliMatchStatement                `json:"sqli_match_statement"`
	XssMatchStatement                 XssMatchStatement                 `json:"xss_match_statement"`
}

type ByteMatchStatement struct {
	PositionalConstraint string               `json:"positional_constraint"`
	SearchString         string               `json:"search_string"`
	TextTransformation   []TextTransformation `json:"text_transformation"`
	FieldToMatch         FieldToMatch         `json:"field_to_match"`
}

type TextTransformation struct {
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type FieldToMatch struct {
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type GeoMatchStatement struct {
	CountryCodes      []string          `json:"country_codes"`
	ForwardedIpConfig ForwardedIpConfig `json:"forwarded_ip_config"`
}

type ForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
}

type IpSetReferenceStatement struct {
	Arn                    string                 `json:"arn"`
	IpSetForwardedIpConfig IpSetForwardedIpConfig `json:"ip_set_forwarded_ip_config"`
}

type IpSetForwardedIpConfig struct {
	FallbackBehavior string `json:"fallback_behavior"`
	HeaderName       string `json:"header_name"`
	Position         string `json:"position"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SizeConstraintStatement struct {
	ComparisonOperator string               `json:"comparison_operator"`
	Size               int                  `json:"size"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type SqliMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type XssMatchStatement struct {
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
	TextTransformation []TextTransformation `json:"text_transformation"`
}

type FieldToMatch struct {
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type RegexPatternSetReferenceStatement struct {
	Arn                string               `json:"arn"`
	TextTransformation []TextTransformation `json:"text_transformation"`
	FieldToMatch       FieldToMatch         `json:"field_to_match"`
}

type TextTransformation struct {
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

type FieldToMatch struct {
	Method              Method              `json:"method"`
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
}

type Method struct{}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type VisibilityConfig struct {
	CloudwatchMetricsEnabled bool   `json:"cloudwatch_metrics_enabled"`
	MetricName               string `json:"metric_name"`
	SampledRequestsEnabled   bool   `json:"sampled_requests_enabled"`
}

type VisibilityConfig struct {
	MetricName               string `json:"metric_name"`
	SampledRequestsEnabled   bool   `json:"sampled_requests_enabled"`
	CloudwatchMetricsEnabled bool   `json:"cloudwatch_metrics_enabled"`
}

// A Wafv2RuleGroupStatus defines the observed state of a Wafv2RuleGroup
type Wafv2RuleGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Wafv2RuleGroupObservation `json:",inline"`
}

// A Wafv2RuleGroupObservation records the observed state of a Wafv2RuleGroup
type Wafv2RuleGroupObservation struct {
	Arn       string `json:"arn"`
	Id        string `json:"id"`
	LockToken string `json:"lock_token"`
}