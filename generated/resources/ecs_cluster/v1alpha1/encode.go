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

package v1alpha1func EncodeEcsCluster(r EcsCluster) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEcsCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEcsCluster_CapacityProviders(r.Spec.ForProvider, ctyVal)
	EncodeEcsCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeEcsCluster_Name(r.Spec.ForProvider, ctyVal)
	EncodeEcsCluster_DefaultCapacityProviderStrategy(r.Spec.ForProvider.DefaultCapacityProviderStrategy, ctyVal)
	EncodeEcsCluster_Setting(r.Spec.ForProvider.Setting, ctyVal)
	EncodeEcsCluster_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEcsCluster_Tags(p *EcsClusterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEcsCluster_CapacityProviders(p *EcsClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CapacityProviders {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["capacity_providers"] = cty.SetVal(colVals)
}

func EncodeEcsCluster_Id(p *EcsClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEcsCluster_Name(p *EcsClusterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEcsCluster_DefaultCapacityProviderStrategy(p *DefaultCapacityProviderStrategy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DefaultCapacityProviderStrategy {
		ctyVal = make(map[string]cty.Value)
		EncodeEcsCluster_DefaultCapacityProviderStrategy_Base(v, ctyVal)
		EncodeEcsCluster_DefaultCapacityProviderStrategy_CapacityProvider(v, ctyVal)
		EncodeEcsCluster_DefaultCapacityProviderStrategy_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["default_capacity_provider_strategy"] = cty.SetVal(valsForCollection)
}

func EncodeEcsCluster_DefaultCapacityProviderStrategy_Base(p *DefaultCapacityProviderStrategy, vals map[string]cty.Value) {
	vals["base"] = cty.IntVal(p.Base)
}

func EncodeEcsCluster_DefaultCapacityProviderStrategy_CapacityProvider(p *DefaultCapacityProviderStrategy, vals map[string]cty.Value) {
	vals["capacity_provider"] = cty.StringVal(p.CapacityProvider)
}

func EncodeEcsCluster_DefaultCapacityProviderStrategy_Weight(p *DefaultCapacityProviderStrategy, vals map[string]cty.Value) {
	vals["weight"] = cty.IntVal(p.Weight)
}

func EncodeEcsCluster_Setting(p *Setting, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Setting {
		ctyVal = make(map[string]cty.Value)
		EncodeEcsCluster_Setting_Name(v, ctyVal)
		EncodeEcsCluster_Setting_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["setting"] = cty.SetVal(valsForCollection)
}

func EncodeEcsCluster_Setting_Name(p *Setting, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEcsCluster_Setting_Value(p *Setting, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeEcsCluster_Arn(p *EcsClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}