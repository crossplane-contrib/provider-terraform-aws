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

package v1alpha1func EncodeServiceDiscoveryService(r ServiceDiscoveryService) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeServiceDiscoveryService_Tags(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryService_Description(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryService_Id(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryService_Name(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryService_NamespaceId(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryService_DnsConfig(r.Spec.ForProvider.DnsConfig, ctyVal)
	EncodeServiceDiscoveryService_HealthCheckConfig(r.Spec.ForProvider.HealthCheckConfig, ctyVal)
	EncodeServiceDiscoveryService_HealthCheckCustomConfig(r.Spec.ForProvider.HealthCheckCustomConfig, ctyVal)
	EncodeServiceDiscoveryService_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeServiceDiscoveryService_Tags(p *ServiceDiscoveryServiceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeServiceDiscoveryService_Description(p *ServiceDiscoveryServiceParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeServiceDiscoveryService_Id(p *ServiceDiscoveryServiceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeServiceDiscoveryService_Name(p *ServiceDiscoveryServiceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeServiceDiscoveryService_NamespaceId(p *ServiceDiscoveryServiceParameters, vals map[string]cty.Value) {
	vals["namespace_id"] = cty.StringVal(p.NamespaceId)
}

func EncodeServiceDiscoveryService_DnsConfig(p *DnsConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DnsConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeServiceDiscoveryService_DnsConfig_NamespaceId(v, ctyVal)
		EncodeServiceDiscoveryService_DnsConfig_RoutingPolicy(v, ctyVal)
		EncodeServiceDiscoveryService_DnsConfig_DnsRecords(v.DnsRecords, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["dns_config"] = cty.ListVal(valsForCollection)
}

func EncodeServiceDiscoveryService_DnsConfig_NamespaceId(p *DnsConfig, vals map[string]cty.Value) {
	vals["namespace_id"] = cty.StringVal(p.NamespaceId)
}

func EncodeServiceDiscoveryService_DnsConfig_RoutingPolicy(p *DnsConfig, vals map[string]cty.Value) {
	vals["routing_policy"] = cty.StringVal(p.RoutingPolicy)
}

func EncodeServiceDiscoveryService_DnsConfig_DnsRecords(p *DnsRecords, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DnsRecords {
		ctyVal = make(map[string]cty.Value)
		EncodeServiceDiscoveryService_DnsConfig_DnsRecords_Ttl(v, ctyVal)
		EncodeServiceDiscoveryService_DnsConfig_DnsRecords_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["dns_records"] = cty.ListVal(valsForCollection)
}

func EncodeServiceDiscoveryService_DnsConfig_DnsRecords_Ttl(p *DnsRecords, vals map[string]cty.Value) {
	vals["ttl"] = cty.IntVal(p.Ttl)
}

func EncodeServiceDiscoveryService_DnsConfig_DnsRecords_Type(p *DnsRecords, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeServiceDiscoveryService_HealthCheckConfig(p *HealthCheckConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.HealthCheckConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeServiceDiscoveryService_HealthCheckConfig_FailureThreshold(v, ctyVal)
		EncodeServiceDiscoveryService_HealthCheckConfig_ResourcePath(v, ctyVal)
		EncodeServiceDiscoveryService_HealthCheckConfig_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["health_check_config"] = cty.ListVal(valsForCollection)
}

func EncodeServiceDiscoveryService_HealthCheckConfig_FailureThreshold(p *HealthCheckConfig, vals map[string]cty.Value) {
	vals["failure_threshold"] = cty.IntVal(p.FailureThreshold)
}

func EncodeServiceDiscoveryService_HealthCheckConfig_ResourcePath(p *HealthCheckConfig, vals map[string]cty.Value) {
	vals["resource_path"] = cty.StringVal(p.ResourcePath)
}

func EncodeServiceDiscoveryService_HealthCheckConfig_Type(p *HealthCheckConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeServiceDiscoveryService_HealthCheckCustomConfig(p *HealthCheckCustomConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.HealthCheckCustomConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeServiceDiscoveryService_HealthCheckCustomConfig_FailureThreshold(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["health_check_custom_config"] = cty.ListVal(valsForCollection)
}

func EncodeServiceDiscoveryService_HealthCheckCustomConfig_FailureThreshold(p *HealthCheckCustomConfig, vals map[string]cty.Value) {
	vals["failure_threshold"] = cty.IntVal(p.FailureThreshold)
}

func EncodeServiceDiscoveryService_Arn(p *ServiceDiscoveryServiceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}