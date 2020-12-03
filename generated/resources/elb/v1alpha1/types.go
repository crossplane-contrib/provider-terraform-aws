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

// Elb is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Elb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElbSpec   `json:"spec"`
	Status ElbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Elb contains a list of ElbList
type ElbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Elb `json:"items"`
}

// A ElbSpec defines the desired state of a Elb
type ElbSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElbParameters `json:",inline"`
}

// A ElbParameters defines the desired state of a Elb
type ElbParameters struct {
	ConnectionDraining        bool              `json:"connection_draining"`
	Internal                  bool              `json:"internal"`
	SourceSecurityGroup       string            `json:"source_security_group"`
	Subnets                   []string          `json:"subnets"`
	AvailabilityZones         []string          `json:"availability_zones"`
	NamePrefix                string            `json:"name_prefix"`
	CrossZoneLoadBalancing    bool              `json:"cross_zone_load_balancing"`
	Id                        string            `json:"id"`
	Name                      string            `json:"name"`
	Tags                      map[string]string `json:"tags"`
	ConnectionDrainingTimeout int               `json:"connection_draining_timeout"`
	IdleTimeout               int               `json:"idle_timeout"`
	Instances                 []string          `json:"instances"`
	SecurityGroups            []string          `json:"security_groups"`
	AccessLogs                AccessLogs        `json:"access_logs"`
	HealthCheck               HealthCheck       `json:"health_check"`
	Listener                  []Listener        `json:"listener"`
}

type AccessLogs struct {
	BucketPrefix string `json:"bucket_prefix"`
	Enabled      bool   `json:"enabled"`
	Interval     int    `json:"interval"`
	Bucket       string `json:"bucket"`
}

type HealthCheck struct {
	HealthyThreshold   int    `json:"healthy_threshold"`
	Interval           int    `json:"interval"`
	Target             string `json:"target"`
	Timeout            int    `json:"timeout"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
}

type Listener struct {
	InstancePort     int    `json:"instance_port"`
	InstanceProtocol string `json:"instance_protocol"`
	LbPort           int    `json:"lb_port"`
	LbProtocol       string `json:"lb_protocol"`
	SslCertificateId string `json:"ssl_certificate_id"`
}

// A ElbStatus defines the observed state of a Elb
type ElbStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElbObservation `json:",inline"`
}

// A ElbObservation records the observed state of a Elb
type ElbObservation struct {
	SourceSecurityGroupId string `json:"source_security_group_id"`
	ZoneId                string `json:"zone_id"`
	DnsName               string `json:"dns_name"`
	Arn                   string `json:"arn"`
}