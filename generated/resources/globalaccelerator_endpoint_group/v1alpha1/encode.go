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

package v1alpha1func EncodeGlobalacceleratorEndpointGroup(r GlobalacceleratorEndpointGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeGlobalacceleratorEndpointGroup_EndpointGroupRegion(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorEndpointGroup_HealthCheckPort(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorEndpointGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorEndpointGroup_ThresholdCount(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorEndpointGroup_TrafficDialPercentage(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorEndpointGroup_HealthCheckIntervalSeconds(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorEndpointGroup_HealthCheckPath(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorEndpointGroup_HealthCheckProtocol(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorEndpointGroup_ListenerArn(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorEndpointGroup_EndpointConfiguration(r.Spec.ForProvider.EndpointConfiguration, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeGlobalacceleratorEndpointGroup_EndpointGroupRegion(p *GlobalacceleratorEndpointGroupParameters, vals map[string]cty.Value) {
	vals["endpoint_group_region"] = cty.StringVal(p.EndpointGroupRegion)
}

func EncodeGlobalacceleratorEndpointGroup_HealthCheckPort(p *GlobalacceleratorEndpointGroupParameters, vals map[string]cty.Value) {
	vals["health_check_port"] = cty.IntVal(p.HealthCheckPort)
}

func EncodeGlobalacceleratorEndpointGroup_Id(p *GlobalacceleratorEndpointGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlobalacceleratorEndpointGroup_ThresholdCount(p *GlobalacceleratorEndpointGroupParameters, vals map[string]cty.Value) {
	vals["threshold_count"] = cty.IntVal(p.ThresholdCount)
}

func EncodeGlobalacceleratorEndpointGroup_TrafficDialPercentage(p *GlobalacceleratorEndpointGroupParameters, vals map[string]cty.Value) {
	vals["traffic_dial_percentage"] = cty.IntVal(p.TrafficDialPercentage)
}

func EncodeGlobalacceleratorEndpointGroup_HealthCheckIntervalSeconds(p *GlobalacceleratorEndpointGroupParameters, vals map[string]cty.Value) {
	vals["health_check_interval_seconds"] = cty.IntVal(p.HealthCheckIntervalSeconds)
}

func EncodeGlobalacceleratorEndpointGroup_HealthCheckPath(p *GlobalacceleratorEndpointGroupParameters, vals map[string]cty.Value) {
	vals["health_check_path"] = cty.StringVal(p.HealthCheckPath)
}

func EncodeGlobalacceleratorEndpointGroup_HealthCheckProtocol(p *GlobalacceleratorEndpointGroupParameters, vals map[string]cty.Value) {
	vals["health_check_protocol"] = cty.StringVal(p.HealthCheckProtocol)
}

func EncodeGlobalacceleratorEndpointGroup_ListenerArn(p *GlobalacceleratorEndpointGroupParameters, vals map[string]cty.Value) {
	vals["listener_arn"] = cty.StringVal(p.ListenerArn)
}

func EncodeGlobalacceleratorEndpointGroup_EndpointConfiguration(p *EndpointConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EndpointConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeGlobalacceleratorEndpointGroup_EndpointConfiguration_ClientIpPreservationEnabled(v, ctyVal)
		EncodeGlobalacceleratorEndpointGroup_EndpointConfiguration_EndpointId(v, ctyVal)
		EncodeGlobalacceleratorEndpointGroup_EndpointConfiguration_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["endpoint_configuration"] = cty.SetVal(valsForCollection)
}

func EncodeGlobalacceleratorEndpointGroup_EndpointConfiguration_ClientIpPreservationEnabled(p *EndpointConfiguration, vals map[string]cty.Value) {
	vals["client_ip_preservation_enabled"] = cty.BoolVal(p.ClientIpPreservationEnabled)
}

func EncodeGlobalacceleratorEndpointGroup_EndpointConfiguration_EndpointId(p *EndpointConfiguration, vals map[string]cty.Value) {
	vals["endpoint_id"] = cty.StringVal(p.EndpointId)
}

func EncodeGlobalacceleratorEndpointGroup_EndpointConfiguration_Weight(p *EndpointConfiguration, vals map[string]cty.Value) {
	vals["weight"] = cty.IntVal(p.Weight)
}