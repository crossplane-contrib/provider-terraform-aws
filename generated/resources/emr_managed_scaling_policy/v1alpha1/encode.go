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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*EmrManagedScalingPolicy)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EmrManagedScalingPolicy.")
	}
	return EncodeEmrManagedScalingPolicy(*r), nil
}

func EncodeEmrManagedScalingPolicy(r EmrManagedScalingPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEmrManagedScalingPolicy_ClusterId(r.Spec.ForProvider, ctyVal)
	EncodeEmrManagedScalingPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeEmrManagedScalingPolicy_ComputeLimits(r.Spec.ForProvider.ComputeLimits, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeEmrManagedScalingPolicy_ClusterId(p EmrManagedScalingPolicyParameters, vals map[string]cty.Value) {
	vals["cluster_id"] = cty.StringVal(p.ClusterId)
}

func EncodeEmrManagedScalingPolicy_Id(p EmrManagedScalingPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrManagedScalingPolicy_ComputeLimits(p []ComputeLimits, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEmrManagedScalingPolicy_ComputeLimits_MaximumOndemandCapacityUnits(v, ctyVal)
		EncodeEmrManagedScalingPolicy_ComputeLimits_MinimumCapacityUnits(v, ctyVal)
		EncodeEmrManagedScalingPolicy_ComputeLimits_UnitType(v, ctyVal)
		EncodeEmrManagedScalingPolicy_ComputeLimits_MaximumCapacityUnits(v, ctyVal)
		EncodeEmrManagedScalingPolicy_ComputeLimits_MaximumCoreCapacityUnits(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["compute_limits"] = cty.SetVal(valsForCollection)
}

func EncodeEmrManagedScalingPolicy_ComputeLimits_MaximumOndemandCapacityUnits(p ComputeLimits, vals map[string]cty.Value) {
	vals["maximum_ondemand_capacity_units"] = cty.NumberIntVal(p.MaximumOndemandCapacityUnits)
}

func EncodeEmrManagedScalingPolicy_ComputeLimits_MinimumCapacityUnits(p ComputeLimits, vals map[string]cty.Value) {
	vals["minimum_capacity_units"] = cty.NumberIntVal(p.MinimumCapacityUnits)
}

func EncodeEmrManagedScalingPolicy_ComputeLimits_UnitType(p ComputeLimits, vals map[string]cty.Value) {
	vals["unit_type"] = cty.StringVal(p.UnitType)
}

func EncodeEmrManagedScalingPolicy_ComputeLimits_MaximumCapacityUnits(p ComputeLimits, vals map[string]cty.Value) {
	vals["maximum_capacity_units"] = cty.NumberIntVal(p.MaximumCapacityUnits)
}

func EncodeEmrManagedScalingPolicy_ComputeLimits_MaximumCoreCapacityUnits(p ComputeLimits, vals map[string]cty.Value) {
	vals["maximum_core_capacity_units"] = cty.NumberIntVal(p.MaximumCoreCapacityUnits)
}