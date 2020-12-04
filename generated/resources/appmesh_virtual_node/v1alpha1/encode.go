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
	"github.com/zclconf/go-cty/cty"
)

func EncodeAppmeshVirtualNode(r AppmeshVirtualNode) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualNode_MeshName(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualNode_MeshOwner(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualNode_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualNode_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualNode_Spec(r.Spec.ForProvider.Spec, ctyVal)
	EncodeAppmeshVirtualNode_ResourceOwner(r.Status.AtProvider, ctyVal)
	EncodeAppmeshVirtualNode_Arn(r.Status.AtProvider, ctyVal)
	EncodeAppmeshVirtualNode_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeAppmeshVirtualNode_LastUpdatedDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAppmeshVirtualNode_Id(p AppmeshVirtualNodeParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppmeshVirtualNode_MeshName(p AppmeshVirtualNodeParameters, vals map[string]cty.Value) {
	vals["mesh_name"] = cty.StringVal(p.MeshName)
}

func EncodeAppmeshVirtualNode_MeshOwner(p AppmeshVirtualNodeParameters, vals map[string]cty.Value) {
	vals["mesh_owner"] = cty.StringVal(p.MeshOwner)
}

func EncodeAppmeshVirtualNode_Tags(p AppmeshVirtualNodeParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAppmeshVirtualNode_Name(p AppmeshVirtualNodeParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppmeshVirtualNode_Spec(p Spec, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Backend(p.Backend, ctyVal)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults(p.BackendDefaults, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener(p.Listener, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Logging(p.Logging, ctyVal)
	EncodeAppmeshVirtualNode_Spec_ServiceDiscovery(p.ServiceDiscovery, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["spec"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Backend(p []Backend, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAppmeshVirtualNode_Spec_Backend_VirtualService(v.VirtualService, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["backend"] = cty.SetVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService(p VirtualService, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_VirtualServiceName(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy(p.ClientPolicy, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["virtual_service"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_VirtualServiceName(p VirtualService, vals map[string]cty.Value) {
	vals["virtual_service_name"] = cty.StringVal(p.VirtualServiceName)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy(p ClientPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls(p.Tls, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["client_policy"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls(p Tls, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Ports(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Enforce(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation(p.Validation, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["tls"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Ports(p Tls, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Ports {
		colVals = append(colVals, cty.NumberIntVal(value))
	}
	vals["ports"] = cty.SetVal(colVals)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Enforce(p Tls, vals map[string]cty.Value) {
	vals["enforce"] = cty.BoolVal(p.Enforce)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation(p Validation, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust(p.Trust, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["validation"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust(p Trust, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust_Acm(p.Acm, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust_File(p.File, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["trust"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust_Acm(p Acm, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust_Acm_CertificateAuthorityArns(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["acm"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust_Acm_CertificateAuthorityArns(p Acm, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CertificateAuthorityArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["certificate_authority_arns"] = cty.SetVal(colVals)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust_File(p File, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust_File_CertificateChain(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["file"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Backend_VirtualService_ClientPolicy_Tls_Validation_Trust_File_CertificateChain(p File, vals map[string]cty.Value) {
	vals["certificate_chain"] = cty.StringVal(p.CertificateChain)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults(p BackendDefaults, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy(p.ClientPolicy, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["backend_defaults"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy(p ClientPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls(p.Tls, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["client_policy"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls(p Tls, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Enforce(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Ports(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation(p.Validation, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["tls"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Enforce(p Tls, vals map[string]cty.Value) {
	vals["enforce"] = cty.BoolVal(p.Enforce)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Ports(p Tls, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Ports {
		colVals = append(colVals, cty.NumberIntVal(value))
	}
	vals["ports"] = cty.SetVal(colVals)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation(p Validation, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust(p.Trust, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["validation"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust(p Trust, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust_Acm(p.Acm, ctyVal)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust_File(p.File, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["trust"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust_Acm(p Acm, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust_Acm_CertificateAuthorityArns(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["acm"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust_Acm_CertificateAuthorityArns(p Acm, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CertificateAuthorityArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["certificate_authority_arns"] = cty.SetVal(colVals)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust_File(p File, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust_File_CertificateChain(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["file"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_BackendDefaults_ClientPolicy_Tls_Validation_Trust_File_CertificateChain(p File, vals map[string]cty.Value) {
	vals["certificate_chain"] = cty.StringVal(p.CertificateChain)
}

func EncodeAppmeshVirtualNode_Spec_Listener(p Listener, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck(p.HealthCheck, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_PortMapping(p.PortMapping, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_Tls(p.Tls, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["listener"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck(p HealthCheck, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_Port(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_Protocol(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_TimeoutMillis(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_UnhealthyThreshold(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_HealthyThreshold(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_IntervalMillis(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_Path(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["health_check"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_Port(p HealthCheck, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_Protocol(p HealthCheck, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_TimeoutMillis(p HealthCheck, vals map[string]cty.Value) {
	vals["timeout_millis"] = cty.NumberIntVal(p.TimeoutMillis)
}

func EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_UnhealthyThreshold(p HealthCheck, vals map[string]cty.Value) {
	vals["unhealthy_threshold"] = cty.NumberIntVal(p.UnhealthyThreshold)
}

func EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_HealthyThreshold(p HealthCheck, vals map[string]cty.Value) {
	vals["healthy_threshold"] = cty.NumberIntVal(p.HealthyThreshold)
}

func EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_IntervalMillis(p HealthCheck, vals map[string]cty.Value) {
	vals["interval_millis"] = cty.NumberIntVal(p.IntervalMillis)
}

func EncodeAppmeshVirtualNode_Spec_Listener_HealthCheck_Path(p HealthCheck, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeAppmeshVirtualNode_Spec_Listener_PortMapping(p PortMapping, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Listener_PortMapping_Protocol(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_PortMapping_Port(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["port_mapping"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Listener_PortMapping_Protocol(p PortMapping, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeAppmeshVirtualNode_Spec_Listener_PortMapping_Port(p PortMapping, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeAppmeshVirtualNode_Spec_Listener_Tls(p Tls, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Listener_Tls_Mode(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate(p.Certificate, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["tls"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Listener_Tls_Mode(p Tls, vals map[string]cty.Value) {
	vals["mode"] = cty.StringVal(p.Mode)
}

func EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate(p Certificate, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_Acm(p.Acm, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_File(p.File, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["certificate"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_Acm(p Acm, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_Acm_CertificateArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["acm"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_Acm_CertificateArn(p Acm, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}

func EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_File(p File, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_File_CertificateChain(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_File_PrivateKey(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["file"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_File_CertificateChain(p File, vals map[string]cty.Value) {
	vals["certificate_chain"] = cty.StringVal(p.CertificateChain)
}

func EncodeAppmeshVirtualNode_Spec_Listener_Tls_Certificate_File_PrivateKey(p File, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodeAppmeshVirtualNode_Spec_Logging(p Logging, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Logging_AccessLog(p.AccessLog, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["logging"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Logging_AccessLog(p AccessLog, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Logging_AccessLog_File(p.File, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["access_log"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Logging_AccessLog_File(p File, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_Logging_AccessLog_File_Path(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["file"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_Logging_AccessLog_File_Path(p File, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeAppmeshVirtualNode_Spec_ServiceDiscovery(p ServiceDiscovery, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_AwsCloudMap(p.AwsCloudMap, ctyVal)
	EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_Dns(p.Dns, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["service_discovery"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_AwsCloudMap(p AwsCloudMap, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_AwsCloudMap_Attributes(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_AwsCloudMap_NamespaceName(p, ctyVal)
	EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_AwsCloudMap_ServiceName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["aws_cloud_map"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_AwsCloudMap_Attributes(p AwsCloudMap, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Attributes {
		mVals[key] = cty.StringVal(value)
	}
	vals["attributes"] = cty.MapVal(mVals)
}

func EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_AwsCloudMap_NamespaceName(p AwsCloudMap, vals map[string]cty.Value) {
	vals["namespace_name"] = cty.StringVal(p.NamespaceName)
}

func EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_AwsCloudMap_ServiceName(p AwsCloudMap, vals map[string]cty.Value) {
	vals["service_name"] = cty.StringVal(p.ServiceName)
}

func EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_Dns(p Dns, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_Dns_Hostname(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dns"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualNode_Spec_ServiceDiscovery_Dns_Hostname(p Dns, vals map[string]cty.Value) {
	vals["hostname"] = cty.StringVal(p.Hostname)
}

func EncodeAppmeshVirtualNode_ResourceOwner(p AppmeshVirtualNodeObservation, vals map[string]cty.Value) {
	vals["resource_owner"] = cty.StringVal(p.ResourceOwner)
}

func EncodeAppmeshVirtualNode_Arn(p AppmeshVirtualNodeObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAppmeshVirtualNode_CreatedDate(p AppmeshVirtualNodeObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeAppmeshVirtualNode_LastUpdatedDate(p AppmeshVirtualNodeObservation, vals map[string]cty.Value) {
	vals["last_updated_date"] = cty.StringVal(p.LastUpdatedDate)
}