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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*EcsCluster)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EcsCluster.")
	}
	return EncodeEcsCluster(*r), nil
}

func EncodeEcsCluster(r EcsCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEcsCluster_Name(r.Spec.ForProvider, ctyVal)
	EncodeEcsCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEcsCluster_CapacityProviders(r.Spec.ForProvider, ctyVal)
	EncodeEcsCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeEcsCluster_DefaultCapacityProviderStrategy(r.Spec.ForProvider.DefaultCapacityProviderStrategy, ctyVal)
	EncodeEcsCluster_Setting(r.Spec.ForProvider.Setting, ctyVal)
	EncodeEcsCluster_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeEcsCluster_Name(p EcsClusterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEcsCluster_Tags(p EcsClusterParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEcsCluster_CapacityProviders(p EcsClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CapacityProviders {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["capacity_providers"] = cty.SetVal(colVals)
}

func EncodeEcsCluster_Id(p EcsClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEcsCluster_DefaultCapacityProviderStrategy(p DefaultCapacityProviderStrategy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsCluster_DefaultCapacityProviderStrategy_Base(p, ctyVal)
	EncodeEcsCluster_DefaultCapacityProviderStrategy_CapacityProvider(p, ctyVal)
	EncodeEcsCluster_DefaultCapacityProviderStrategy_Weight(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["default_capacity_provider_strategy"] = cty.SetVal(valsForCollection)
}

func EncodeEcsCluster_DefaultCapacityProviderStrategy_Base(p DefaultCapacityProviderStrategy, vals map[string]cty.Value) {
	vals["base"] = cty.NumberIntVal(p.Base)
}

func EncodeEcsCluster_DefaultCapacityProviderStrategy_CapacityProvider(p DefaultCapacityProviderStrategy, vals map[string]cty.Value) {
	vals["capacity_provider"] = cty.StringVal(p.CapacityProvider)
}

func EncodeEcsCluster_DefaultCapacityProviderStrategy_Weight(p DefaultCapacityProviderStrategy, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeEcsCluster_Setting(p Setting, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsCluster_Setting_Name(p, ctyVal)
	EncodeEcsCluster_Setting_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["setting"] = cty.SetVal(valsForCollection)
}

func EncodeEcsCluster_Setting_Name(p Setting, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEcsCluster_Setting_Value(p Setting, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeEcsCluster_Arn(p EcsClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}