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

package v1alpha1func EncodeCodedeployDeploymentConfig(r CodedeployDeploymentConfig) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCodedeployDeploymentConfig_DeploymentConfigName(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployDeploymentConfig_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployDeploymentConfig_ComputePlatform(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployDeploymentConfig_MinimumHealthyHosts(r.Spec.ForProvider.MinimumHealthyHosts, ctyVal)
	EncodeCodedeployDeploymentConfig_TrafficRoutingConfig(r.Spec.ForProvider.TrafficRoutingConfig, ctyVal)
	EncodeCodedeployDeploymentConfig_DeploymentConfigId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeCodedeployDeploymentConfig_DeploymentConfigName(p *CodedeployDeploymentConfigParameters, vals map[string]cty.Value) {
	vals["deployment_config_name"] = cty.StringVal(p.DeploymentConfigName)
}

func EncodeCodedeployDeploymentConfig_Id(p *CodedeployDeploymentConfigParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodedeployDeploymentConfig_ComputePlatform(p *CodedeployDeploymentConfigParameters, vals map[string]cty.Value) {
	vals["compute_platform"] = cty.StringVal(p.ComputePlatform)
}

func EncodeCodedeployDeploymentConfig_MinimumHealthyHosts(p *MinimumHealthyHosts, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MinimumHealthyHosts {
		ctyVal = make(map[string]cty.Value)
		EncodeCodedeployDeploymentConfig_MinimumHealthyHosts_Type(v, ctyVal)
		EncodeCodedeployDeploymentConfig_MinimumHealthyHosts_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["minimum_healthy_hosts"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentConfig_MinimumHealthyHosts_Type(p *MinimumHealthyHosts, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodedeployDeploymentConfig_MinimumHealthyHosts_Value(p *MinimumHealthyHosts, vals map[string]cty.Value) {
	vals["value"] = cty.IntVal(p.Value)
}

func EncodeCodedeployDeploymentConfig_TrafficRoutingConfig(p *TrafficRoutingConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TrafficRoutingConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_Type(v, ctyVal)
		EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedCanary(v.TimeBasedCanary, ctyVal)
		EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedLinear(v.TimeBasedLinear, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["traffic_routing_config"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_Type(p *TrafficRoutingConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedCanary(p *TimeBasedCanary, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TimeBasedCanary {
		ctyVal = make(map[string]cty.Value)
		EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedCanary_Interval(v, ctyVal)
		EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedCanary_Percentage(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["time_based_canary"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedCanary_Interval(p *TimeBasedCanary, vals map[string]cty.Value) {
	vals["interval"] = cty.IntVal(p.Interval)
}

func EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedCanary_Percentage(p *TimeBasedCanary, vals map[string]cty.Value) {
	vals["percentage"] = cty.IntVal(p.Percentage)
}

func EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedLinear(p *TimeBasedLinear, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TimeBasedLinear {
		ctyVal = make(map[string]cty.Value)
		EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedLinear_Percentage(v, ctyVal)
		EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedLinear_Interval(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["time_based_linear"] = cty.ListVal(valsForCollection)
}

func EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedLinear_Percentage(p *TimeBasedLinear, vals map[string]cty.Value) {
	vals["percentage"] = cty.IntVal(p.Percentage)
}

func EncodeCodedeployDeploymentConfig_TrafficRoutingConfig_TimeBasedLinear_Interval(p *TimeBasedLinear, vals map[string]cty.Value) {
	vals["interval"] = cty.IntVal(p.Interval)
}

func EncodeCodedeployDeploymentConfig_DeploymentConfigId(p *CodedeployDeploymentConfigObservation, vals map[string]cty.Value) {
	vals["deployment_config_id"] = cty.StringVal(p.DeploymentConfigId)
}