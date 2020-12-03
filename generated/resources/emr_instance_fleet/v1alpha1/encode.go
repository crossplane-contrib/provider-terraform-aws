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

package v1alpha1func EncodeEmrInstanceFleet(r EmrInstanceFleet) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEmrInstanceFleet_ClusterId(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_Id(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_Name(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_TargetOnDemandCapacity(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_TargetSpotCapacity(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs(r.Spec.ForProvider.InstanceTypeConfigs, ctyVal)
	EncodeEmrInstanceFleet_LaunchSpecifications(r.Spec.ForProvider.LaunchSpecifications, ctyVal)
	EncodeEmrInstanceFleet_ProvisionedOnDemandCapacity(r.Status.AtProvider, ctyVal)
	EncodeEmrInstanceFleet_ProvisionedSpotCapacity(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEmrInstanceFleet_ClusterId(p *EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["cluster_id"] = cty.StringVal(p.ClusterId)
}

func EncodeEmrInstanceFleet_Id(p *EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrInstanceFleet_Name(p *EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrInstanceFleet_TargetOnDemandCapacity(p *EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["target_on_demand_capacity"] = cty.IntVal(p.TargetOnDemandCapacity)
}

func EncodeEmrInstanceFleet_TargetSpotCapacity(p *EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["target_spot_capacity"] = cty.IntVal(p.TargetSpotCapacity)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.InstanceTypeConfigs {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_BidPrice(v, ctyVal)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_BidPriceAsPercentageOfOnDemandPrice(v, ctyVal)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_InstanceType(v, ctyVal)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_WeightedCapacity(v, ctyVal)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations(v.Configurations, ctyVal)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig(v.EbsConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["instance_type_configs"] = cty.SetVal(valsForCollection)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_BidPrice(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["bid_price"] = cty.StringVal(p.BidPrice)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_BidPriceAsPercentageOfOnDemandPrice(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["bid_price_as_percentage_of_on_demand_price"] = cty.IntVal(p.BidPriceAsPercentageOfOnDemandPrice)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_InstanceType(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_WeightedCapacity(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["weighted_capacity"] = cty.IntVal(p.WeightedCapacity)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations(p *Configurations, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Configurations {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations_Classification(v, ctyVal)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations_Properties(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["configurations"] = cty.SetVal(valsForCollection)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations_Classification(p *Configurations, vals map[string]cty.Value) {
	vals["classification"] = cty.StringVal(p.Classification)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations_Properties(p *Configurations, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Properties {
		mVals[key] = cty.StringVal(value)
	}
	vals["properties"] = cty.MapVal(mVals)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig(p *EbsConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Iops(v, ctyVal)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Size(v, ctyVal)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Type(v, ctyVal)
		EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_VolumesPerInstance(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_config"] = cty.SetVal(valsForCollection)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Iops(p *EbsConfig, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Size(p *EbsConfig, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Type(p *EbsConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_VolumesPerInstance(p *EbsConfig, vals map[string]cty.Value) {
	vals["volumes_per_instance"] = cty.IntVal(p.VolumesPerInstance)
}

func EncodeEmrInstanceFleet_LaunchSpecifications(p *LaunchSpecifications, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LaunchSpecifications {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrInstanceFleet_LaunchSpecifications_OnDemandSpecification(v.OnDemandSpecification, ctyVal)
		EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification(v.SpotSpecification, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["launch_specifications"] = cty.ListVal(valsForCollection)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_OnDemandSpecification(p *OnDemandSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OnDemandSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrInstanceFleet_LaunchSpecifications_OnDemandSpecification_AllocationStrategy(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["on_demand_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_OnDemandSpecification_AllocationStrategy(p *OnDemandSpecification, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification(p *SpotSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SpotSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_BlockDurationMinutes(v, ctyVal)
		EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutAction(v, ctyVal)
		EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutDurationMinutes(v, ctyVal)
		EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_AllocationStrategy(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["spot_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_BlockDurationMinutes(p *SpotSpecification, vals map[string]cty.Value) {
	vals["block_duration_minutes"] = cty.IntVal(p.BlockDurationMinutes)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutAction(p *SpotSpecification, vals map[string]cty.Value) {
	vals["timeout_action"] = cty.StringVal(p.TimeoutAction)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutDurationMinutes(p *SpotSpecification, vals map[string]cty.Value) {
	vals["timeout_duration_minutes"] = cty.IntVal(p.TimeoutDurationMinutes)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_AllocationStrategy(p *SpotSpecification, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEmrInstanceFleet_ProvisionedOnDemandCapacity(p *EmrInstanceFleetObservation, vals map[string]cty.Value) {
	vals["provisioned_on_demand_capacity"] = cty.IntVal(p.ProvisionedOnDemandCapacity)
}

func EncodeEmrInstanceFleet_ProvisionedSpotCapacity(p *EmrInstanceFleetObservation, vals map[string]cty.Value) {
	vals["provisioned_spot_capacity"] = cty.IntVal(p.ProvisionedSpotCapacity)
}