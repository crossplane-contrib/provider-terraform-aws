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
	r, ok := mr.(*EmrInstanceFleet)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EmrInstanceFleet.")
	}
	return EncodeEmrInstanceFleet(*r), nil
}

func EncodeEmrInstanceFleet(r EmrInstanceFleet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEmrInstanceFleet_ClusterId(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_Id(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_Name(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_TargetOnDemandCapacity(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_TargetSpotCapacity(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs(r.Spec.ForProvider.InstanceTypeConfigs, ctyVal)
	EncodeEmrInstanceFleet_LaunchSpecifications(r.Spec.ForProvider.LaunchSpecifications, ctyVal)
	EncodeEmrInstanceFleet_ProvisionedOnDemandCapacity(r.Status.AtProvider, ctyVal)
	EncodeEmrInstanceFleet_ProvisionedSpotCapacity(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeEmrInstanceFleet_ClusterId(p EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["cluster_id"] = cty.StringVal(p.ClusterId)
}

func EncodeEmrInstanceFleet_Id(p EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrInstanceFleet_Name(p EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrInstanceFleet_TargetOnDemandCapacity(p EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["target_on_demand_capacity"] = cty.NumberIntVal(p.TargetOnDemandCapacity)
}

func EncodeEmrInstanceFleet_TargetSpotCapacity(p EmrInstanceFleetParameters, vals map[string]cty.Value) {
	vals["target_spot_capacity"] = cty.NumberIntVal(p.TargetSpotCapacity)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs(p InstanceTypeConfigs, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_BidPrice(p, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_BidPriceAsPercentageOfOnDemandPrice(p, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_InstanceType(p, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_WeightedCapacity(p, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations(p.Configurations, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig(p.EbsConfig, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["instance_type_configs"] = cty.SetVal(valsForCollection)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_BidPrice(p InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["bid_price"] = cty.StringVal(p.BidPrice)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_BidPriceAsPercentageOfOnDemandPrice(p InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["bid_price_as_percentage_of_on_demand_price"] = cty.NumberIntVal(p.BidPriceAsPercentageOfOnDemandPrice)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_InstanceType(p InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_WeightedCapacity(p InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["weighted_capacity"] = cty.NumberIntVal(p.WeightedCapacity)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations(p Configurations, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations_Classification(p, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations_Properties(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["configurations"] = cty.SetVal(valsForCollection)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations_Classification(p Configurations, vals map[string]cty.Value) {
	vals["classification"] = cty.StringVal(p.Classification)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_Configurations_Properties(p Configurations, vals map[string]cty.Value) {
	if len(p.Properties) == 0 {
		vals["properties"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Properties {
		mVals[key] = cty.StringVal(value)
	}
	vals["properties"] = cty.MapVal(mVals)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig(p EbsConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Iops(p, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Size(p, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Type(p, ctyVal)
	EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_VolumesPerInstance(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_config"] = cty.SetVal(valsForCollection)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Iops(p EbsConfig, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Size(p EbsConfig, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_Type(p EbsConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEmrInstanceFleet_InstanceTypeConfigs_EbsConfig_VolumesPerInstance(p EbsConfig, vals map[string]cty.Value) {
	vals["volumes_per_instance"] = cty.NumberIntVal(p.VolumesPerInstance)
}

func EncodeEmrInstanceFleet_LaunchSpecifications(p LaunchSpecifications, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEmrInstanceFleet_LaunchSpecifications_OnDemandSpecification(p.OnDemandSpecification, ctyVal)
	EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification(p.SpotSpecification, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["launch_specifications"] = cty.ListVal(valsForCollection)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_OnDemandSpecification(p OnDemandSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEmrInstanceFleet_LaunchSpecifications_OnDemandSpecification_AllocationStrategy(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["on_demand_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_OnDemandSpecification_AllocationStrategy(p OnDemandSpecification, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification(p SpotSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_AllocationStrategy(p, ctyVal)
	EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_BlockDurationMinutes(p, ctyVal)
	EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutAction(p, ctyVal)
	EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutDurationMinutes(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["spot_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_AllocationStrategy(p SpotSpecification, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_BlockDurationMinutes(p SpotSpecification, vals map[string]cty.Value) {
	vals["block_duration_minutes"] = cty.NumberIntVal(p.BlockDurationMinutes)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutAction(p SpotSpecification, vals map[string]cty.Value) {
	vals["timeout_action"] = cty.StringVal(p.TimeoutAction)
}

func EncodeEmrInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutDurationMinutes(p SpotSpecification, vals map[string]cty.Value) {
	vals["timeout_duration_minutes"] = cty.NumberIntVal(p.TimeoutDurationMinutes)
}

func EncodeEmrInstanceFleet_ProvisionedOnDemandCapacity(p EmrInstanceFleetObservation, vals map[string]cty.Value) {
	vals["provisioned_on_demand_capacity"] = cty.NumberIntVal(p.ProvisionedOnDemandCapacity)
}

func EncodeEmrInstanceFleet_ProvisionedSpotCapacity(p EmrInstanceFleetObservation, vals map[string]cty.Value) {
	vals["provisioned_spot_capacity"] = cty.NumberIntVal(p.ProvisionedSpotCapacity)
}