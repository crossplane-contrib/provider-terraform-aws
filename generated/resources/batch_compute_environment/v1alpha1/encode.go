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

func EncodeBatchComputeEnvironment(r BatchComputeEnvironment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeBatchComputeEnvironment_ComputeEnvironmentNamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeBatchComputeEnvironment_ServiceRole(r.Spec.ForProvider, ctyVal)
	EncodeBatchComputeEnvironment_Type(r.Spec.ForProvider, ctyVal)
	EncodeBatchComputeEnvironment_ComputeEnvironmentName(r.Spec.ForProvider, ctyVal)
	EncodeBatchComputeEnvironment_Id(r.Spec.ForProvider, ctyVal)
	EncodeBatchComputeEnvironment_State(r.Spec.ForProvider, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources(r.Spec.ForProvider.ComputeResources, ctyVal)
	EncodeBatchComputeEnvironment_EcsClusterArn(r.Status.AtProvider, ctyVal)
	EncodeBatchComputeEnvironment_StatusReason(r.Status.AtProvider, ctyVal)
	EncodeBatchComputeEnvironment_Arn(r.Status.AtProvider, ctyVal)
	EncodeBatchComputeEnvironment_Status(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeBatchComputeEnvironment_ComputeEnvironmentNamePrefix(p BatchComputeEnvironmentParameters, vals map[string]cty.Value) {
	vals["compute_environment_name_prefix"] = cty.StringVal(p.ComputeEnvironmentNamePrefix)
}

func EncodeBatchComputeEnvironment_ServiceRole(p BatchComputeEnvironmentParameters, vals map[string]cty.Value) {
	vals["service_role"] = cty.StringVal(p.ServiceRole)
}

func EncodeBatchComputeEnvironment_Type(p BatchComputeEnvironmentParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeBatchComputeEnvironment_ComputeEnvironmentName(p BatchComputeEnvironmentParameters, vals map[string]cty.Value) {
	vals["compute_environment_name"] = cty.StringVal(p.ComputeEnvironmentName)
}

func EncodeBatchComputeEnvironment_Id(p BatchComputeEnvironmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeBatchComputeEnvironment_State(p BatchComputeEnvironmentParameters, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeBatchComputeEnvironment_ComputeResources(p ComputeResources, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeBatchComputeEnvironment_ComputeResources_DesiredVcpus(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_MinVcpus(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_AllocationStrategy(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_BidPercentage(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_Ec2KeyPair(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_Subnets(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_Tags(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_Type(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_ImageId(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_InstanceRole(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_InstanceType(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_MaxVcpus(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_SecurityGroupIds(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_SpotIamFleetRole(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_LaunchTemplate(p.LaunchTemplate, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["compute_resources"] = cty.ListVal(valsForCollection)
}

func EncodeBatchComputeEnvironment_ComputeResources_DesiredVcpus(p ComputeResources, vals map[string]cty.Value) {
	vals["desired_vcpus"] = cty.NumberIntVal(p.DesiredVcpus)
}

func EncodeBatchComputeEnvironment_ComputeResources_MinVcpus(p ComputeResources, vals map[string]cty.Value) {
	vals["min_vcpus"] = cty.NumberIntVal(p.MinVcpus)
}

func EncodeBatchComputeEnvironment_ComputeResources_AllocationStrategy(p ComputeResources, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeBatchComputeEnvironment_ComputeResources_BidPercentage(p ComputeResources, vals map[string]cty.Value) {
	vals["bid_percentage"] = cty.NumberIntVal(p.BidPercentage)
}

func EncodeBatchComputeEnvironment_ComputeResources_Ec2KeyPair(p ComputeResources, vals map[string]cty.Value) {
	vals["ec2_key_pair"] = cty.StringVal(p.Ec2KeyPair)
}

func EncodeBatchComputeEnvironment_ComputeResources_Subnets(p ComputeResources, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Subnets {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnets"] = cty.SetVal(colVals)
}

func EncodeBatchComputeEnvironment_ComputeResources_Tags(p ComputeResources, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeBatchComputeEnvironment_ComputeResources_Type(p ComputeResources, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeBatchComputeEnvironment_ComputeResources_ImageId(p ComputeResources, vals map[string]cty.Value) {
	vals["image_id"] = cty.StringVal(p.ImageId)
}

func EncodeBatchComputeEnvironment_ComputeResources_InstanceRole(p ComputeResources, vals map[string]cty.Value) {
	vals["instance_role"] = cty.StringVal(p.InstanceRole)
}

func EncodeBatchComputeEnvironment_ComputeResources_InstanceType(p ComputeResources, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.InstanceType {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["instance_type"] = cty.SetVal(colVals)
}

func EncodeBatchComputeEnvironment_ComputeResources_MaxVcpus(p ComputeResources, vals map[string]cty.Value) {
	vals["max_vcpus"] = cty.NumberIntVal(p.MaxVcpus)
}

func EncodeBatchComputeEnvironment_ComputeResources_SecurityGroupIds(p ComputeResources, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeBatchComputeEnvironment_ComputeResources_SpotIamFleetRole(p ComputeResources, vals map[string]cty.Value) {
	vals["spot_iam_fleet_role"] = cty.StringVal(p.SpotIamFleetRole)
}

func EncodeBatchComputeEnvironment_ComputeResources_LaunchTemplate(p LaunchTemplate, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeBatchComputeEnvironment_ComputeResources_LaunchTemplate_LaunchTemplateId(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_LaunchTemplate_LaunchTemplateName(p, ctyVal)
	EncodeBatchComputeEnvironment_ComputeResources_LaunchTemplate_Version(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["launch_template"] = cty.ListVal(valsForCollection)
}

func EncodeBatchComputeEnvironment_ComputeResources_LaunchTemplate_LaunchTemplateId(p LaunchTemplate, vals map[string]cty.Value) {
	vals["launch_template_id"] = cty.StringVal(p.LaunchTemplateId)
}

func EncodeBatchComputeEnvironment_ComputeResources_LaunchTemplate_LaunchTemplateName(p LaunchTemplate, vals map[string]cty.Value) {
	vals["launch_template_name"] = cty.StringVal(p.LaunchTemplateName)
}

func EncodeBatchComputeEnvironment_ComputeResources_LaunchTemplate_Version(p LaunchTemplate, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeBatchComputeEnvironment_EcsClusterArn(p BatchComputeEnvironmentObservation, vals map[string]cty.Value) {
	vals["ecs_cluster_arn"] = cty.StringVal(p.EcsClusterArn)
}

func EncodeBatchComputeEnvironment_StatusReason(p BatchComputeEnvironmentObservation, vals map[string]cty.Value) {
	vals["status_reason"] = cty.StringVal(p.StatusReason)
}

func EncodeBatchComputeEnvironment_Arn(p BatchComputeEnvironmentObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeBatchComputeEnvironment_Status(p BatchComputeEnvironmentObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}