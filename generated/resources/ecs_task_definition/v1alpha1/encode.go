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

func EncodeEcsTaskDefinition(r EcsTaskDefinition) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEcsTaskDefinition_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_Cpu(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_Id(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_IpcMode(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_NetworkMode(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_ContainerDefinitions(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_RequiresCompatibilities(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_TaskRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_ExecutionRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_Family(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_Memory(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_PidMode(r.Spec.ForProvider, ctyVal)
	EncodeEcsTaskDefinition_InferenceAccelerator(r.Spec.ForProvider.InferenceAccelerator, ctyVal)
	EncodeEcsTaskDefinition_PlacementConstraints(r.Spec.ForProvider.PlacementConstraints, ctyVal)
	EncodeEcsTaskDefinition_ProxyConfiguration(r.Spec.ForProvider.ProxyConfiguration, ctyVal)
	EncodeEcsTaskDefinition_Volume(r.Spec.ForProvider.Volume, ctyVal)
	EncodeEcsTaskDefinition_Arn(r.Status.AtProvider, ctyVal)
	EncodeEcsTaskDefinition_Revision(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEcsTaskDefinition_Tags(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEcsTaskDefinition_Cpu(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["cpu"] = cty.StringVal(p.Cpu)
}

func EncodeEcsTaskDefinition_Id(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEcsTaskDefinition_IpcMode(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["ipc_mode"] = cty.StringVal(p.IpcMode)
}

func EncodeEcsTaskDefinition_NetworkMode(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["network_mode"] = cty.StringVal(p.NetworkMode)
}

func EncodeEcsTaskDefinition_ContainerDefinitions(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["container_definitions"] = cty.StringVal(p.ContainerDefinitions)
}

func EncodeEcsTaskDefinition_RequiresCompatibilities(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.RequiresCompatibilities {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["requires_compatibilities"] = cty.SetVal(colVals)
}

func EncodeEcsTaskDefinition_TaskRoleArn(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["task_role_arn"] = cty.StringVal(p.TaskRoleArn)
}

func EncodeEcsTaskDefinition_ExecutionRoleArn(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["execution_role_arn"] = cty.StringVal(p.ExecutionRoleArn)
}

func EncodeEcsTaskDefinition_Family(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["family"] = cty.StringVal(p.Family)
}

func EncodeEcsTaskDefinition_Memory(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["memory"] = cty.StringVal(p.Memory)
}

func EncodeEcsTaskDefinition_PidMode(p EcsTaskDefinitionParameters, vals map[string]cty.Value) {
	vals["pid_mode"] = cty.StringVal(p.PidMode)
}

func EncodeEcsTaskDefinition_InferenceAccelerator(p InferenceAccelerator, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsTaskDefinition_InferenceAccelerator_DeviceName(p, ctyVal)
	EncodeEcsTaskDefinition_InferenceAccelerator_DeviceType(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["inference_accelerator"] = cty.SetVal(valsForCollection)
}

func EncodeEcsTaskDefinition_InferenceAccelerator_DeviceName(p InferenceAccelerator, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeEcsTaskDefinition_InferenceAccelerator_DeviceType(p InferenceAccelerator, vals map[string]cty.Value) {
	vals["device_type"] = cty.StringVal(p.DeviceType)
}

func EncodeEcsTaskDefinition_PlacementConstraints(p []PlacementConstraints, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEcsTaskDefinition_PlacementConstraints_Expression(v, ctyVal)
		EncodeEcsTaskDefinition_PlacementConstraints_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["placement_constraints"] = cty.SetVal(valsForCollection)
}

func EncodeEcsTaskDefinition_PlacementConstraints_Expression(p PlacementConstraints, vals map[string]cty.Value) {
	vals["expression"] = cty.StringVal(p.Expression)
}

func EncodeEcsTaskDefinition_PlacementConstraints_Type(p PlacementConstraints, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEcsTaskDefinition_ProxyConfiguration(p ProxyConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsTaskDefinition_ProxyConfiguration_ContainerName(p, ctyVal)
	EncodeEcsTaskDefinition_ProxyConfiguration_Properties(p, ctyVal)
	EncodeEcsTaskDefinition_ProxyConfiguration_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["proxy_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeEcsTaskDefinition_ProxyConfiguration_ContainerName(p ProxyConfiguration, vals map[string]cty.Value) {
	vals["container_name"] = cty.StringVal(p.ContainerName)
}

func EncodeEcsTaskDefinition_ProxyConfiguration_Properties(p ProxyConfiguration, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Properties {
		mVals[key] = cty.StringVal(value)
	}
	vals["properties"] = cty.MapVal(mVals)
}

func EncodeEcsTaskDefinition_ProxyConfiguration_Type(p ProxyConfiguration, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEcsTaskDefinition_Volume(p Volume, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsTaskDefinition_Volume_HostPath(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_Name(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration(p.DockerVolumeConfiguration, ctyVal)
	EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration(p.EfsVolumeConfiguration, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["volume"] = cty.SetVal(valsForCollection)
}

func EncodeEcsTaskDefinition_Volume_HostPath(p Volume, vals map[string]cty.Value) {
	vals["host_path"] = cty.StringVal(p.HostPath)
}

func EncodeEcsTaskDefinition_Volume_Name(p Volume, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration(p DockerVolumeConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_DriverOpts(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_Labels(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_Scope(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_Autoprovision(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_Driver(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["docker_volume_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_DriverOpts(p DockerVolumeConfiguration, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.DriverOpts {
		mVals[key] = cty.StringVal(value)
	}
	vals["driver_opts"] = cty.MapVal(mVals)
}

func EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_Labels(p DockerVolumeConfiguration, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Labels {
		mVals[key] = cty.StringVal(value)
	}
	vals["labels"] = cty.MapVal(mVals)
}

func EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_Scope(p DockerVolumeConfiguration, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_Autoprovision(p DockerVolumeConfiguration, vals map[string]cty.Value) {
	vals["autoprovision"] = cty.BoolVal(p.Autoprovision)
}

func EncodeEcsTaskDefinition_Volume_DockerVolumeConfiguration_Driver(p DockerVolumeConfiguration, vals map[string]cty.Value) {
	vals["driver"] = cty.StringVal(p.Driver)
}

func EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration(p EfsVolumeConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_FileSystemId(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_RootDirectory(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_TransitEncryption(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_TransitEncryptionPort(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_AuthorizationConfig(p.AuthorizationConfig, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["efs_volume_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_FileSystemId(p EfsVolumeConfiguration, vals map[string]cty.Value) {
	vals["file_system_id"] = cty.StringVal(p.FileSystemId)
}

func EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_RootDirectory(p EfsVolumeConfiguration, vals map[string]cty.Value) {
	vals["root_directory"] = cty.StringVal(p.RootDirectory)
}

func EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_TransitEncryption(p EfsVolumeConfiguration, vals map[string]cty.Value) {
	vals["transit_encryption"] = cty.StringVal(p.TransitEncryption)
}

func EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_TransitEncryptionPort(p EfsVolumeConfiguration, vals map[string]cty.Value) {
	vals["transit_encryption_port"] = cty.NumberIntVal(p.TransitEncryptionPort)
}

func EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_AuthorizationConfig(p AuthorizationConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_AuthorizationConfig_Iam(p, ctyVal)
	EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_AuthorizationConfig_AccessPointId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["authorization_config"] = cty.ListVal(valsForCollection)
}

func EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_AuthorizationConfig_Iam(p AuthorizationConfig, vals map[string]cty.Value) {
	vals["iam"] = cty.StringVal(p.Iam)
}

func EncodeEcsTaskDefinition_Volume_EfsVolumeConfiguration_AuthorizationConfig_AccessPointId(p AuthorizationConfig, vals map[string]cty.Value) {
	vals["access_point_id"] = cty.StringVal(p.AccessPointId)
}

func EncodeEcsTaskDefinition_Arn(p EcsTaskDefinitionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEcsTaskDefinition_Revision(p EcsTaskDefinitionObservation, vals map[string]cty.Value) {
	vals["revision"] = cty.NumberIntVal(p.Revision)
}