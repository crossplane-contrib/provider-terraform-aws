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

// AppmeshVirtualNode is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppmeshVirtualNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppmeshVirtualNodeSpec   `json:"spec"`
	Status AppmeshVirtualNodeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualNode contains a list of AppmeshVirtualNodeList
type AppmeshVirtualNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshVirtualNode `json:"items"`
}

// A AppmeshVirtualNodeSpec defines the desired state of a AppmeshVirtualNode
type AppmeshVirtualNodeSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppmeshVirtualNodeParameters `json:",inline"`
}

// A AppmeshVirtualNodeParameters defines the desired state of a AppmeshVirtualNode
type AppmeshVirtualNodeParameters struct {
	Tags      map[string]string `json:"tags"`
	Id        string            `json:"id"`
	MeshName  string            `json:"mesh_name"`
	MeshOwner string            `json:"mesh_owner"`
	Name      string            `json:"name"`
	Spec      Spec              `json:"spec"`
}

type Spec struct {
	Backend          []Backend        `json:"backend"`
	BackendDefaults  BackendDefaults  `json:"backend_defaults"`
	Listener         Listener         `json:"listener"`
	Logging          Logging          `json:"logging"`
	ServiceDiscovery ServiceDiscovery `json:"service_discovery"`
}

type Backend struct {
	VirtualService VirtualService `json:"virtual_service"`
}

type VirtualService struct {
	VirtualServiceName string       `json:"virtual_service_name"`
	ClientPolicy       ClientPolicy `json:"client_policy"`
}

type ClientPolicy struct {
	Tls Tls `json:"tls"`
}

type Tls struct {
	Enforce    bool       `json:"enforce"`
	Ports      []int      `json:"ports"`
	Validation Validation `json:"validation"`
}

type Validation struct {
	Trust Trust `json:"trust"`
}

type Trust struct {
	Acm  Acm  `json:"acm"`
	File File `json:"file"`
}

type Acm struct {
	CertificateAuthorityArns []string `json:"certificate_authority_arns"`
}

type File struct {
	CertificateChain string `json:"certificate_chain"`
}

type BackendDefaults struct {
	ClientPolicy ClientPolicy `json:"client_policy"`
}

type ClientPolicy struct {
	Tls Tls `json:"tls"`
}

type Tls struct {
	Enforce    bool       `json:"enforce"`
	Ports      []int      `json:"ports"`
	Validation Validation `json:"validation"`
}

type Validation struct {
	Trust Trust `json:"trust"`
}

type Trust struct {
	File File `json:"file"`
	Acm  Acm  `json:"acm"`
}

type File struct {
	CertificateChain string `json:"certificate_chain"`
}

type Acm struct {
	CertificateAuthorityArns []string `json:"certificate_authority_arns"`
}

type Listener struct {
	HealthCheck HealthCheck `json:"health_check"`
	PortMapping PortMapping `json:"port_mapping"`
	Tls         Tls         `json:"tls"`
}

type HealthCheck struct {
	HealthyThreshold   int    `json:"healthy_threshold"`
	IntervalMillis     int    `json:"interval_millis"`
	Path               string `json:"path"`
	Port               int    `json:"port"`
	Protocol           string `json:"protocol"`
	TimeoutMillis      int    `json:"timeout_millis"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
}

type PortMapping struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type Tls struct {
	Mode        string      `json:"mode"`
	Certificate Certificate `json:"certificate"`
}

type Certificate struct {
	Acm  Acm  `json:"acm"`
	File File `json:"file"`
}

type Acm struct {
	CertificateArn string `json:"certificate_arn"`
}

type File struct {
	CertificateChain string `json:"certificate_chain"`
	PrivateKey       string `json:"private_key"`
}

type Logging struct {
	AccessLog AccessLog `json:"access_log"`
}

type AccessLog struct {
	File File `json:"file"`
}

type File struct {
	Path string `json:"path"`
}

type ServiceDiscovery struct {
	Dns         Dns         `json:"dns"`
	AwsCloudMap AwsCloudMap `json:"aws_cloud_map"`
}

type Dns struct {
	Hostname string `json:"hostname"`
}

type AwsCloudMap struct {
	Attributes    map[string]string `json:"attributes"`
	NamespaceName string            `json:"namespace_name"`
	ServiceName   string            `json:"service_name"`
}

// A AppmeshVirtualNodeStatus defines the observed state of a AppmeshVirtualNode
type AppmeshVirtualNodeStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppmeshVirtualNodeObservation `json:",inline"`
}

// A AppmeshVirtualNodeObservation records the observed state of a AppmeshVirtualNode
type AppmeshVirtualNodeObservation struct {
	ResourceOwner   string `json:"resource_owner"`
	Arn             string `json:"arn"`
	LastUpdatedDate string `json:"last_updated_date"`
	CreatedDate     string `json:"created_date"`
}