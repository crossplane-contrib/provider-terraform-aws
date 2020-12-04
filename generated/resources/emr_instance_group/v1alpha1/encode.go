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
	"github.com/zclconf/go-cty/cty"
)

func EncodeEmrInstanceGroup(r EmrInstanceGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEmrInstanceGroup_AutoscalingPolicy(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_InstanceCount(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_BidPrice(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_ClusterId(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_ConfigurationsJson(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_EbsOptimized(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeEmrInstanceGroup_EbsConfig(r.Spec.ForProvider.EbsConfig, ctyVal)
	EncodeEmrInstanceGroup_Status(r.Status.AtProvider, ctyVal)
	EncodeEmrInstanceGroup_RunningInstanceCount(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEmrInstanceGroup_AutoscalingPolicy(p EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["autoscaling_policy"] = cty.StringVal(p.AutoscalingPolicy)
}

func EncodeEmrInstanceGroup_InstanceCount(p EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["instance_count"] = cty.NumberIntVal(p.InstanceCount)
}

func EncodeEmrInstanceGroup_InstanceType(p EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEmrInstanceGroup_Name(p EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrInstanceGroup_BidPrice(p EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["bid_price"] = cty.StringVal(p.BidPrice)
}

func EncodeEmrInstanceGroup_ClusterId(p EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["cluster_id"] = cty.StringVal(p.ClusterId)
}

func EncodeEmrInstanceGroup_ConfigurationsJson(p EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["configurations_json"] = cty.StringVal(p.ConfigurationsJson)
}

func EncodeEmrInstanceGroup_EbsOptimized(p EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["ebs_optimized"] = cty.BoolVal(p.EbsOptimized)
}

func EncodeEmrInstanceGroup_Id(p EmrInstanceGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrInstanceGroup_EbsConfig(p EbsConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEmrInstanceGroup_EbsConfig_Iops(p, ctyVal)
	EncodeEmrInstanceGroup_EbsConfig_Size(p, ctyVal)
	EncodeEmrInstanceGroup_EbsConfig_Type(p, ctyVal)
	EncodeEmrInstanceGroup_EbsConfig_VolumesPerInstance(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_config"] = cty.SetVal(valsForCollection)
}

func EncodeEmrInstanceGroup_EbsConfig_Iops(p EbsConfig, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeEmrInstanceGroup_EbsConfig_Size(p EbsConfig, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeEmrInstanceGroup_EbsConfig_Type(p EbsConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEmrInstanceGroup_EbsConfig_VolumesPerInstance(p EbsConfig, vals map[string]cty.Value) {
	vals["volumes_per_instance"] = cty.NumberIntVal(p.VolumesPerInstance)
}

func EncodeEmrInstanceGroup_Status(p EmrInstanceGroupObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeEmrInstanceGroup_RunningInstanceCount(p EmrInstanceGroupObservation, vals map[string]cty.Value) {
	vals["running_instance_count"] = cty.NumberIntVal(p.RunningInstanceCount)
}