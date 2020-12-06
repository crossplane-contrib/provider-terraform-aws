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

func EncodeEksNodeGroup(r EksNodeGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEksNodeGroup_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_AmiType(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_ClusterName(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_ForceUpdateVersion(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_Labels(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_NodeGroupName(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_InstanceTypes(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_NodeRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_ReleaseVersion(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_DiskSize(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_Version(r.Spec.ForProvider, ctyVal)
	EncodeEksNodeGroup_ScalingConfig(r.Spec.ForProvider.ScalingConfig, ctyVal)
	EncodeEksNodeGroup_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeEksNodeGroup_LaunchTemplate(r.Spec.ForProvider.LaunchTemplate, ctyVal)
	EncodeEksNodeGroup_RemoteAccess(r.Spec.ForProvider.RemoteAccess, ctyVal)
	EncodeEksNodeGroup_Arn(r.Status.AtProvider, ctyVal)
	EncodeEksNodeGroup_Resources(r.Status.AtProvider.Resources, ctyVal)
	EncodeEksNodeGroup_Status(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEksNodeGroup_SubnetIds(p EksNodeGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeEksNodeGroup_Tags(p EksNodeGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEksNodeGroup_AmiType(p EksNodeGroupParameters, vals map[string]cty.Value) {
	vals["ami_type"] = cty.StringVal(p.AmiType)
}

func EncodeEksNodeGroup_ClusterName(p EksNodeGroupParameters, vals map[string]cty.Value) {
	vals["cluster_name"] = cty.StringVal(p.ClusterName)
}

func EncodeEksNodeGroup_ForceUpdateVersion(p EksNodeGroupParameters, vals map[string]cty.Value) {
	vals["force_update_version"] = cty.BoolVal(p.ForceUpdateVersion)
}

func EncodeEksNodeGroup_Id(p EksNodeGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEksNodeGroup_Labels(p EksNodeGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Labels {
		mVals[key] = cty.StringVal(value)
	}
	vals["labels"] = cty.MapVal(mVals)
}

func EncodeEksNodeGroup_NodeGroupName(p EksNodeGroupParameters, vals map[string]cty.Value) {
	vals["node_group_name"] = cty.StringVal(p.NodeGroupName)
}

func EncodeEksNodeGroup_InstanceTypes(p EksNodeGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.InstanceTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["instance_types"] = cty.ListVal(colVals)
}

func EncodeEksNodeGroup_NodeRoleArn(p EksNodeGroupParameters, vals map[string]cty.Value) {
	vals["node_role_arn"] = cty.StringVal(p.NodeRoleArn)
}

func EncodeEksNodeGroup_ReleaseVersion(p EksNodeGroupParameters, vals map[string]cty.Value) {
	vals["release_version"] = cty.StringVal(p.ReleaseVersion)
}

func EncodeEksNodeGroup_DiskSize(p EksNodeGroupParameters, vals map[string]cty.Value) {
	vals["disk_size"] = cty.NumberIntVal(p.DiskSize)
}

func EncodeEksNodeGroup_Version(p EksNodeGroupParameters, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeEksNodeGroup_ScalingConfig(p ScalingConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEksNodeGroup_ScalingConfig_DesiredSize(p, ctyVal)
	EncodeEksNodeGroup_ScalingConfig_MaxSize(p, ctyVal)
	EncodeEksNodeGroup_ScalingConfig_MinSize(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["scaling_config"] = cty.ListVal(valsForCollection)
}

func EncodeEksNodeGroup_ScalingConfig_DesiredSize(p ScalingConfig, vals map[string]cty.Value) {
	vals["desired_size"] = cty.NumberIntVal(p.DesiredSize)
}

func EncodeEksNodeGroup_ScalingConfig_MaxSize(p ScalingConfig, vals map[string]cty.Value) {
	vals["max_size"] = cty.NumberIntVal(p.MaxSize)
}

func EncodeEksNodeGroup_ScalingConfig_MinSize(p ScalingConfig, vals map[string]cty.Value) {
	vals["min_size"] = cty.NumberIntVal(p.MinSize)
}

func EncodeEksNodeGroup_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeEksNodeGroup_Timeouts_Create(p, ctyVal)
	EncodeEksNodeGroup_Timeouts_Delete(p, ctyVal)
	EncodeEksNodeGroup_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeEksNodeGroup_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeEksNodeGroup_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeEksNodeGroup_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeEksNodeGroup_LaunchTemplate(p LaunchTemplate, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEksNodeGroup_LaunchTemplate_Version(p, ctyVal)
	EncodeEksNodeGroup_LaunchTemplate_Id(p, ctyVal)
	EncodeEksNodeGroup_LaunchTemplate_Name(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["launch_template"] = cty.ListVal(valsForCollection)
}

func EncodeEksNodeGroup_LaunchTemplate_Version(p LaunchTemplate, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeEksNodeGroup_LaunchTemplate_Id(p LaunchTemplate, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEksNodeGroup_LaunchTemplate_Name(p LaunchTemplate, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEksNodeGroup_RemoteAccess(p RemoteAccess, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEksNodeGroup_RemoteAccess_Ec2SshKey(p, ctyVal)
	EncodeEksNodeGroup_RemoteAccess_SourceSecurityGroupIds(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["remote_access"] = cty.ListVal(valsForCollection)
}

func EncodeEksNodeGroup_RemoteAccess_Ec2SshKey(p RemoteAccess, vals map[string]cty.Value) {
	vals["ec2_ssh_key"] = cty.StringVal(p.Ec2SshKey)
}

func EncodeEksNodeGroup_RemoteAccess_SourceSecurityGroupIds(p RemoteAccess, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SourceSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["source_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeEksNodeGroup_Arn(p EksNodeGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEksNodeGroup_Resources(p []Resources, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEksNodeGroup_Resources_RemoteAccessSecurityGroupId(v, ctyVal)
		EncodeEksNodeGroup_Resources_AutoscalingGroups(v.AutoscalingGroups, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["resources"] = cty.ListVal(valsForCollection)
}

func EncodeEksNodeGroup_Resources_RemoteAccessSecurityGroupId(p Resources, vals map[string]cty.Value) {
	vals["remote_access_security_group_id"] = cty.StringVal(p.RemoteAccessSecurityGroupId)
}

func EncodeEksNodeGroup_Resources_AutoscalingGroups(p []AutoscalingGroups, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEksNodeGroup_Resources_AutoscalingGroups_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["autoscaling_groups"] = cty.ListVal(valsForCollection)
}

func EncodeEksNodeGroup_Resources_AutoscalingGroups_Name(p AutoscalingGroups, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEksNodeGroup_Status(p EksNodeGroupObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}