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

// CloudfrontDistribution is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudfrontDistribution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudfrontDistributionSpec   `json:"spec"`
	Status CloudfrontDistributionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontDistribution contains a list of CloudfrontDistributionList
type CloudfrontDistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudfrontDistribution `json:"items"`
}

// A CloudfrontDistributionSpec defines the desired state of a CloudfrontDistribution
type CloudfrontDistributionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudfrontDistributionParameters `json:",inline"`
}

// A CloudfrontDistributionParameters defines the desired state of a CloudfrontDistribution
type CloudfrontDistributionParameters struct {
	Comment              string                 `json:"comment"`
	Enabled              bool                   `json:"enabled"`
	DefaultRootObject    string                 `json:"default_root_object"`
	HttpVersion          string                 `json:"http_version"`
	PriceClass           string                 `json:"price_class"`
	Tags                 map[string]string      `json:"tags"`
	Id                   string                 `json:"id"`
	WaitForDeployment    bool                   `json:"wait_for_deployment"`
	WebAclId             string                 `json:"web_acl_id"`
	Aliases              []string               `json:"aliases"`
	IsIpv6Enabled        bool                   `json:"is_ipv6_enabled"`
	RetainOnDelete       bool                   `json:"retain_on_delete"`
	OrderedCacheBehavior []OrderedCacheBehavior `json:"ordered_cache_behavior"`
	Origin               []Origin               `json:"origin"`
	OriginGroup          []OriginGroup          `json:"origin_group"`
	Restrictions         Restrictions           `json:"restrictions"`
	ViewerCertificate    ViewerCertificate      `json:"viewer_certificate"`
	CustomErrorResponse  []CustomErrorResponse  `json:"custom_error_response"`
	DefaultCacheBehavior DefaultCacheBehavior   `json:"default_cache_behavior"`
	LoggingConfig        LoggingConfig          `json:"logging_config"`
}

type OrderedCacheBehavior struct {
	DefaultTtl                int                         `json:"default_ttl"`
	FieldLevelEncryptionId    string                      `json:"field_level_encryption_id"`
	MaxTtl                    int                         `json:"max_ttl"`
	MinTtl                    int                         `json:"min_ttl"`
	PathPattern               string                      `json:"path_pattern"`
	SmoothStreaming           bool                        `json:"smooth_streaming"`
	AllowedMethods            []string                    `json:"allowed_methods"`
	CachedMethods             []string                    `json:"cached_methods"`
	TrustedSigners            []string                    `json:"trusted_signers"`
	ViewerProtocolPolicy      string                      `json:"viewer_protocol_policy"`
	Compress                  bool                        `json:"compress"`
	TargetOriginId            string                      `json:"target_origin_id"`
	ForwardedValues           ForwardedValues             `json:"forwarded_values"`
	LambdaFunctionAssociation []LambdaFunctionAssociation `json:"lambda_function_association"`
}

type ForwardedValues struct {
	Headers              []string `json:"headers"`
	QueryString          bool     `json:"query_string"`
	QueryStringCacheKeys []string `json:"query_string_cache_keys"`
	Cookies              Cookies  `json:"cookies"`
}

type Cookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type LambdaFunctionAssociation struct {
	EventType   string `json:"event_type"`
	IncludeBody bool   `json:"include_body"`
	LambdaArn   string `json:"lambda_arn"`
}

type Origin struct {
	DomainName         string             `json:"domain_name"`
	OriginId           string             `json:"origin_id"`
	OriginPath         string             `json:"origin_path"`
	S3OriginConfig     S3OriginConfig     `json:"s3_origin_config"`
	CustomHeader       []CustomHeader     `json:"custom_header"`
	CustomOriginConfig CustomOriginConfig `json:"custom_origin_config"`
}

type S3OriginConfig struct {
	OriginAccessIdentity string `json:"origin_access_identity"`
}

type CustomHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CustomOriginConfig struct {
	OriginKeepaliveTimeout int      `json:"origin_keepalive_timeout"`
	OriginProtocolPolicy   string   `json:"origin_protocol_policy"`
	OriginReadTimeout      int      `json:"origin_read_timeout"`
	OriginSslProtocols     []string `json:"origin_ssl_protocols"`
	HttpPort               int      `json:"http_port"`
	HttpsPort              int      `json:"https_port"`
}

type OriginGroup struct {
	OriginId         string           `json:"origin_id"`
	FailoverCriteria FailoverCriteria `json:"failover_criteria"`
	Member           []Member         `json:"member"`
}

type FailoverCriteria struct {
	StatusCodes []int `json:"status_codes"`
}

type Member struct {
	OriginId string `json:"origin_id"`
}

type Restrictions struct {
	GeoRestriction GeoRestriction `json:"geo_restriction"`
}

type GeoRestriction struct {
	Locations       []string `json:"locations"`
	RestrictionType string   `json:"restriction_type"`
}

type ViewerCertificate struct {
	AcmCertificateArn            string `json:"acm_certificate_arn"`
	CloudfrontDefaultCertificate bool   `json:"cloudfront_default_certificate"`
	IamCertificateId             string `json:"iam_certificate_id"`
	MinimumProtocolVersion       string `json:"minimum_protocol_version"`
	SslSupportMethod             string `json:"ssl_support_method"`
}

type CustomErrorResponse struct {
	ResponseCode       int    `json:"response_code"`
	ResponsePagePath   string `json:"response_page_path"`
	ErrorCachingMinTtl int    `json:"error_caching_min_ttl"`
	ErrorCode          int    `json:"error_code"`
}

type DefaultCacheBehavior struct {
	CachedMethods             []string                    `json:"cached_methods"`
	Compress                  bool                        `json:"compress"`
	MaxTtl                    int                         `json:"max_ttl"`
	TrustedSigners            []string                    `json:"trusted_signers"`
	AllowedMethods            []string                    `json:"allowed_methods"`
	DefaultTtl                int                         `json:"default_ttl"`
	FieldLevelEncryptionId    string                      `json:"field_level_encryption_id"`
	MinTtl                    int                         `json:"min_ttl"`
	SmoothStreaming           bool                        `json:"smooth_streaming"`
	TargetOriginId            string                      `json:"target_origin_id"`
	ViewerProtocolPolicy      string                      `json:"viewer_protocol_policy"`
	ForwardedValues           ForwardedValues             `json:"forwarded_values"`
	LambdaFunctionAssociation []LambdaFunctionAssociation `json:"lambda_function_association"`
}

type ForwardedValues struct {
	Headers              []string `json:"headers"`
	QueryString          bool     `json:"query_string"`
	QueryStringCacheKeys []string `json:"query_string_cache_keys"`
	Cookies              Cookies  `json:"cookies"`
}

type Cookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type LambdaFunctionAssociation struct {
	LambdaArn   string `json:"lambda_arn"`
	EventType   string `json:"event_type"`
	IncludeBody bool   `json:"include_body"`
}

type LoggingConfig struct {
	Bucket         string `json:"bucket"`
	IncludeCookies bool   `json:"include_cookies"`
	Prefix         string `json:"prefix"`
}

// A CloudfrontDistributionStatus defines the observed state of a CloudfrontDistribution
type CloudfrontDistributionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudfrontDistributionObservation `json:",inline"`
}

// A CloudfrontDistributionObservation records the observed state of a CloudfrontDistribution
type CloudfrontDistributionObservation struct {
	InProgressValidationBatches int              `json:"in_progress_validation_batches"`
	Status                      string           `json:"status"`
	DomainName                  string           `json:"domain_name"`
	HostedZoneId                string           `json:"hosted_zone_id"`
	LastModifiedTime            string           `json:"last_modified_time"`
	CallerReference             string           `json:"caller_reference"`
	TrustedSigners              []TrustedSigners `json:"trusted_signers"`
	Arn                         string           `json:"arn"`
	Etag                        string           `json:"etag"`
}

type TrustedSigners struct {
	Enabled bool    `json:"enabled"`
	Items   []Items `json:"items"`
}

type Items struct {
	AwsAccountNumber string   `json:"aws_account_number"`
	KeyPairIds       []string `json:"key_pair_ids"`
}