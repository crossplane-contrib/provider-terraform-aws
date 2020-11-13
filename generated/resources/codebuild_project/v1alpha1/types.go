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

// CodebuildProject is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodebuildProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodebuildProjectSpec   `json:"spec"`
	Status CodebuildProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildProject contains a list of CodebuildProjectList
type CodebuildProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodebuildProject `json:"items"`
}

// A CodebuildProjectSpec defines the desired state of a CodebuildProject
type CodebuildProjectSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodebuildProjectParameters `json:",inline"`
}

// A CodebuildProjectParameters defines the desired state of a CodebuildProject
type CodebuildProjectParameters struct {
	Description        string               `json:"description"`
	EncryptionKey      string               `json:"encryption_key"`
	Name               string               `json:"name"`
	ServiceRole        string               `json:"service_role"`
	BadgeEnabled       bool                 `json:"badge_enabled"`
	QueuedTimeout      int                  `json:"queued_timeout"`
	SourceVersion      string               `json:"source_version"`
	Tags               map[string]string    `json:"tags"`
	BuildTimeout       int                  `json:"build_timeout"`
	Id                 string               `json:"id"`
	Cache              Cache                `json:"cache"`
	Environment        Environment          `json:"environment"`
	LogsConfig         LogsConfig           `json:"logs_config"`
	SecondaryArtifacts []SecondaryArtifacts `json:"secondary_artifacts"`
	SecondarySources   []SecondarySources   `json:"secondary_sources"`
	Source             Source               `json:"source"`
	VpcConfig          VpcConfig            `json:"vpc_config"`
	Artifacts          Artifacts            `json:"artifacts"`
}

type Cache struct {
	Modes    []string `json:"modes"`
	Type     string   `json:"type"`
	Location string   `json:"location"`
}

type Environment struct {
	Certificate              string                `json:"certificate"`
	ComputeType              string                `json:"compute_type"`
	Image                    string                `json:"image"`
	ImagePullCredentialsType string                `json:"image_pull_credentials_type"`
	PrivilegedMode           bool                  `json:"privileged_mode"`
	Type                     string                `json:"type"`
	EnvironmentVariable      []EnvironmentVariable `json:"environment_variable"`
	RegistryCredential       RegistryCredential    `json:"registry_credential"`
}

type EnvironmentVariable struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type RegistryCredential struct {
	Credential         string `json:"credential"`
	CredentialProvider string `json:"credential_provider"`
}

type LogsConfig struct {
	CloudwatchLogs CloudwatchLogs `json:"cloudwatch_logs"`
	S3Logs         S3Logs         `json:"s3_logs"`
}

type CloudwatchLogs struct {
	GroupName  string `json:"group_name"`
	Status     string `json:"status"`
	StreamName string `json:"stream_name"`
}

type S3Logs struct {
	Status             string `json:"status"`
	EncryptionDisabled bool   `json:"encryption_disabled"`
	Location           string `json:"location"`
}

type SecondaryArtifacts struct {
	ArtifactIdentifier   string `json:"artifact_identifier"`
	EncryptionDisabled   bool   `json:"encryption_disabled"`
	NamespaceType        string `json:"namespace_type"`
	OverrideArtifactName bool   `json:"override_artifact_name"`
	Packaging            string `json:"packaging"`
	Location             string `json:"location"`
	Name                 string `json:"name"`
	Path                 string `json:"path"`
	Type                 string `json:"type"`
}

type SecondarySources struct {
	Type                string              `json:"type"`
	Buildspec           string              `json:"buildspec"`
	GitCloneDepth       int                 `json:"git_clone_depth"`
	InsecureSsl         bool                `json:"insecure_ssl"`
	Location            string              `json:"location"`
	ReportBuildStatus   bool                `json:"report_build_status"`
	SourceIdentifier    string              `json:"source_identifier"`
	Auth                Auth                `json:"auth"`
	GitSubmodulesConfig GitSubmodulesConfig `json:"git_submodules_config"`
}

type Auth struct {
	Resource string `json:"resource"`
	Type     string `json:"type"`
}

type GitSubmodulesConfig struct {
	FetchSubmodules bool `json:"fetch_submodules"`
}

type Source struct {
	Buildspec           string              `json:"buildspec"`
	GitCloneDepth       int                 `json:"git_clone_depth"`
	InsecureSsl         bool                `json:"insecure_ssl"`
	Location            string              `json:"location"`
	ReportBuildStatus   bool                `json:"report_build_status"`
	Type                string              `json:"type"`
	Auth                Auth                `json:"auth"`
	GitSubmodulesConfig GitSubmodulesConfig `json:"git_submodules_config"`
}

type Auth struct {
	Type     string `json:"type"`
	Resource string `json:"resource"`
}

type GitSubmodulesConfig struct {
	FetchSubmodules bool `json:"fetch_submodules"`
}

type VpcConfig struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	Subnets          []string `json:"subnets"`
	VpcId            string   `json:"vpc_id"`
}

type Artifacts struct {
	OverrideArtifactName bool   `json:"override_artifact_name"`
	Type                 string `json:"type"`
	ArtifactIdentifier   string `json:"artifact_identifier"`
	EncryptionDisabled   bool   `json:"encryption_disabled"`
	Location             string `json:"location"`
	Path                 string `json:"path"`
	Name                 string `json:"name"`
	NamespaceType        string `json:"namespace_type"`
	Packaging            string `json:"packaging"`
}

// A CodebuildProjectStatus defines the observed state of a CodebuildProject
type CodebuildProjectStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodebuildProjectObservation `json:",inline"`
}

// A CodebuildProjectObservation records the observed state of a CodebuildProject
type CodebuildProjectObservation struct {
	BadgeUrl string `json:"badge_url"`
	Arn      string `json:"arn"`
}