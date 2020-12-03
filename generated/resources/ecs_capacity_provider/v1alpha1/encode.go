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

package v1alpha1func EncodeEcsCapacityProvider(r EcsCapacityProvider) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEcsCapacityProvider_Id(r.Spec.ForProvider, ctyVal)
	EncodeEcsCapacityProvider_Name(r.Spec.ForProvider, ctyVal)
	EncodeEcsCapacityProvider_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEcsCapacityProvider_AutoScalingGroupProvider(r.Spec.ForProvider.AutoScalingGroupProvider, ctyVal)
	EncodeEcsCapacityProvider_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEcsCapacityProvider_Id(p *EcsCapacityProviderParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEcsCapacityProvider_Name(p *EcsCapacityProviderParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEcsCapacityProvider_Tags(p *EcsCapacityProviderParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEcsCapacityProvider_AutoScalingGroupProvider(p *AutoScalingGroupProvider, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AutoScalingGroupProvider {
		ctyVal = make(map[string]cty.Value)
		EncodeEcsCapacityProvider_AutoScalingGroupProvider_AutoScalingGroupArn(v, ctyVal)
		EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedTerminationProtection(v, ctyVal)
		EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling(v.ManagedScaling, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["auto_scaling_group_provider"] = cty.ListVal(valsForCollection)
}

func EncodeEcsCapacityProvider_AutoScalingGroupProvider_AutoScalingGroupArn(p *AutoScalingGroupProvider, vals map[string]cty.Value) {
	vals["auto_scaling_group_arn"] = cty.StringVal(p.AutoScalingGroupArn)
}

func EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedTerminationProtection(p *AutoScalingGroupProvider, vals map[string]cty.Value) {
	vals["managed_termination_protection"] = cty.StringVal(p.ManagedTerminationProtection)
}

func EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling(p *ManagedScaling, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ManagedScaling {
		ctyVal = make(map[string]cty.Value)
		EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling_MaximumScalingStepSize(v, ctyVal)
		EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling_MinimumScalingStepSize(v, ctyVal)
		EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling_Status(v, ctyVal)
		EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling_TargetCapacity(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["managed_scaling"] = cty.ListVal(valsForCollection)
}

func EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling_MaximumScalingStepSize(p *ManagedScaling, vals map[string]cty.Value) {
	vals["maximum_scaling_step_size"] = cty.IntVal(p.MaximumScalingStepSize)
}

func EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling_MinimumScalingStepSize(p *ManagedScaling, vals map[string]cty.Value) {
	vals["minimum_scaling_step_size"] = cty.IntVal(p.MinimumScalingStepSize)
}

func EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling_Status(p *ManagedScaling, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeEcsCapacityProvider_AutoScalingGroupProvider_ManagedScaling_TargetCapacity(p *ManagedScaling, vals map[string]cty.Value) {
	vals["target_capacity"] = cty.IntVal(p.TargetCapacity)
}

func EncodeEcsCapacityProvider_Arn(p *EcsCapacityProviderObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}