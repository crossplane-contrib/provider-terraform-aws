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

// S3Bucket is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type S3Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3BucketSpec   `json:"spec"`
	Status S3BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3Bucket contains a list of S3BucketList
type S3BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3Bucket `json:"items"`
}

// A S3BucketSpec defines the desired state of a S3Bucket
type S3BucketSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  S3BucketParameters `json:",inline"`
}

// A S3BucketParameters defines the desired state of a S3Bucket
type S3BucketParameters struct {
	ForceDestroy                      bool                              `json:"force_destroy"`
	Tags                              map[string]string                 `json:"tags"`
	BucketPrefix                      string                            `json:"bucket_prefix"`
	Policy                            string                            `json:"policy"`
	Acl                               string                            `json:"acl"`
	Grant                             []Grant                           `json:"grant"`
	LifecycleRule                     []LifecycleRule                   `json:"lifecycle_rule"`
	Logging                           []Logging                         `json:"logging"`
	ObjectLockConfiguration           ObjectLockConfiguration           `json:"object_lock_configuration"`
	ReplicationConfiguration          ReplicationConfiguration          `json:"replication_configuration"`
	ServerSideEncryptionConfiguration ServerSideEncryptionConfiguration `json:"server_side_encryption_configuration"`
	Website                           Website                           `json:"website"`
	CorsRule                          []CorsRule                        `json:"cors_rule"`
	Versioning                        Versioning                        `json:"versioning"`
}

type Grant struct {
	Id          string   `json:"id"`
	Permissions []string `json:"permissions"`
	Type        string   `json:"type"`
	Uri         string   `json:"uri"`
}

type LifecycleRule struct {
	Prefix                             string                        `json:"prefix"`
	Tags                               map[string]string             `json:"tags"`
	AbortIncompleteMultipartUploadDays int                           `json:"abort_incomplete_multipart_upload_days"`
	Enabled                            bool                          `json:"enabled"`
	Id                                 string                        `json:"id"`
	Transition                         []Transition                  `json:"transition"`
	Expiration                         Expiration                    `json:"expiration"`
	NoncurrentVersionExpiration        NoncurrentVersionExpiration   `json:"noncurrent_version_expiration"`
	NoncurrentVersionTransition        []NoncurrentVersionTransition `json:"noncurrent_version_transition"`
}

type Transition struct {
	Date         string `json:"date"`
	Days         int    `json:"days"`
	StorageClass string `json:"storage_class"`
}

type Expiration struct {
	Date                      string `json:"date"`
	Days                      int    `json:"days"`
	ExpiredObjectDeleteMarker bool   `json:"expired_object_delete_marker"`
}

type NoncurrentVersionExpiration struct {
	Days int `json:"days"`
}

type NoncurrentVersionTransition struct {
	Days         int    `json:"days"`
	StorageClass string `json:"storage_class"`
}

type Logging struct {
	TargetBucket string `json:"target_bucket"`
	TargetPrefix string `json:"target_prefix"`
}

type ObjectLockConfiguration struct {
	ObjectLockEnabled string `json:"object_lock_enabled"`
	Rule              Rule   `json:"rule"`
}

type Rule struct {
	DefaultRetention DefaultRetention `json:"default_retention"`
}

type DefaultRetention struct {
	Days  int    `json:"days"`
	Mode  string `json:"mode"`
	Years int    `json:"years"`
}

type ReplicationConfiguration struct {
	Role  string  `json:"role"`
	Rules []Rules `json:"rules"`
}

type Rules struct {
	Id                      string                  `json:"id"`
	Prefix                  string                  `json:"prefix"`
	Priority                int                     `json:"priority"`
	Status                  string                  `json:"status"`
	Filter                  Filter                  `json:"filter"`
	SourceSelectionCriteria SourceSelectionCriteria `json:"source_selection_criteria"`
	Destination             Destination             `json:"destination"`
}

type Filter struct {
	Prefix string            `json:"prefix"`
	Tags   map[string]string `json:"tags"`
}

type SourceSelectionCriteria struct {
	SseKmsEncryptedObjects SseKmsEncryptedObjects `json:"sse_kms_encrypted_objects"`
}

type SseKmsEncryptedObjects struct {
	Enabled bool `json:"enabled"`
}

type Destination struct {
	ReplicaKmsKeyId          string                   `json:"replica_kms_key_id"`
	StorageClass             string                   `json:"storage_class"`
	AccountId                string                   `json:"account_id"`
	Bucket                   string                   `json:"bucket"`
	AccessControlTranslation AccessControlTranslation `json:"access_control_translation"`
}

type AccessControlTranslation struct {
	Owner string `json:"owner"`
}

type ServerSideEncryptionConfiguration struct {
	Rule Rule `json:"rule"`
}

type Rule struct {
	ApplyServerSideEncryptionByDefault ApplyServerSideEncryptionByDefault `json:"apply_server_side_encryption_by_default"`
}

type ApplyServerSideEncryptionByDefault struct {
	KmsMasterKeyId string `json:"kms_master_key_id"`
	SseAlgorithm   string `json:"sse_algorithm"`
}

type Website struct {
	RoutingRules          string `json:"routing_rules"`
	ErrorDocument         string `json:"error_document"`
	IndexDocument         string `json:"index_document"`
	RedirectAllRequestsTo string `json:"redirect_all_requests_to"`
}

type CorsRule struct {
	AllowedHeaders []string `json:"allowed_headers"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedOrigins []string `json:"allowed_origins"`
	ExposeHeaders  []string `json:"expose_headers"`
	MaxAgeSeconds  int      `json:"max_age_seconds"`
}

type Versioning struct {
	Enabled   bool `json:"enabled"`
	MfaDelete bool `json:"mfa_delete"`
}

// A S3BucketStatus defines the observed state of a S3Bucket
type S3BucketStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3BucketObservation `json:",inline"`
}

// A S3BucketObservation records the observed state of a S3Bucket
type S3BucketObservation struct {
	RequestPayer             string `json:"request_payer"`
	Id                       string `json:"id"`
	Region                   string `json:"region"`
	WebsiteDomain            string `json:"website_domain"`
	WebsiteEndpoint          string `json:"website_endpoint"`
	AccelerationStatus       string `json:"acceleration_status"`
	BucketRegionalDomainName string `json:"bucket_regional_domain_name"`
	BucketDomainName         string `json:"bucket_domain_name"`
	Arn                      string `json:"arn"`
	Bucket                   string `json:"bucket"`
	HostedZoneId             string `json:"hosted_zone_id"`
}