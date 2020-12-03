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

package v1alpha1func EncodeEmrInstanceGroup(r EmrInstanceGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEmrInstanceGroup_ClusterId(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_AutoscalingPolicy(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_BidPrice(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_ConfigurationsJson(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_EbsOptimized(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_InstanceCount(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_EbsConfig(r.Spec.ForProvider.EbsConfig, ctyVal)
	EncodeEmrInstanceGroup_RunningInstanceCount(r.Status.AtProvider, ctyVal)
	EncodeEmrInstanceGroup_Status(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEmrInstanceGroup_ClusterId(p *EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["cluster_id"] = cty.StringVal(p.ClusterId)
}

func EncodeEmrInstanceGroup_Id(p *EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrInstanceGroup_InstanceType(p *EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEmrInstanceGroup_AutoscalingPolicy(p *EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["autoscaling_policy"] = cty.StringVal(p.AutoscalingPolicy)
}

func EncodeEmrInstanceGroup_BidPrice(p *EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["bid_price"] = cty.StringVal(p.BidPrice)
}

func EncodeEmrInstanceGroup_ConfigurationsJson(p *EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["configurations_json"] = cty.StringVal(p.ConfigurationsJson)
}

func EncodeEmrInstanceGroup_EbsOptimized(p *EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["ebs_optimized"] = cty.BoolVal(p.EbsOptimized)
}

func EncodeEmrInstanceGroup_InstanceCount(p *EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["instance_count"] = cty.IntVal(p.InstanceCount)
}

func EncodeEmrInstanceGroup_Name(p *EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrInstanceGroup_EbsConfig(p *EbsConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrInstanceGroup_EbsConfig_Type(v, ctyVal)
		EncodeEmrInstanceGroup_EbsConfig_VolumesPerInstance(v, ctyVal)
		EncodeEmrInstanceGroup_EbsConfig_Iops(v, ctyVal)
		EncodeEmrInstanceGroup_EbsConfig_Size(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_config"] = cty.SetVal(valsForCollection)
}

func EncodeEmrInstanceGroup_EbsConfig_Type(p *EbsConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEmrInstanceGroup_EbsConfig_VolumesPerInstance(p *EbsConfig, vals map[string]cty.Value) {
	vals["volumes_per_instance"] = cty.IntVal(p.VolumesPerInstance)
}

func EncodeEmrInstanceGroup_EbsConfig_Iops(p *EbsConfig, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeEmrInstanceGroup_EbsConfig_Size(p *EbsConfig, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeEmrInstanceGroup_RunningInstanceCount(p *EmrInstanceGroupObservation, vals map[string]cty.Value) {
	vals["running_instance_count"] = cty.IntVal(p.RunningInstanceCount)
}

func EncodeEmrInstanceGroup_Status(p *EmrInstanceGroupObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}