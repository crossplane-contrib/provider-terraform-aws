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

func EncodeGameliftFleet(r GameliftFleet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftFleet_Ec2InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_Id(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_MetricGroups(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_BuildId(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_Description(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_NewGameSessionProtectionPolicy(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_FleetType(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_InstanceRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_Name(r.Spec.ForProvider, ctyVal)
	EncodeGameliftFleet_Ec2InboundPermission(r.Spec.ForProvider.Ec2InboundPermission, ctyVal)
	EncodeGameliftFleet_ResourceCreationLimitPolicy(r.Spec.ForProvider.ResourceCreationLimitPolicy, ctyVal)
	EncodeGameliftFleet_RuntimeConfiguration(r.Spec.ForProvider.RuntimeConfiguration, ctyVal)
	EncodeGameliftFleet_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeGameliftFleet_LogPaths(r.Status.AtProvider, ctyVal)
	EncodeGameliftFleet_Arn(r.Status.AtProvider, ctyVal)
	EncodeGameliftFleet_OperatingSystem(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGameliftFleet_Ec2InstanceType(p GameliftFleetParameters, vals map[string]cty.Value) {
	vals["ec2_instance_type"] = cty.StringVal(p.Ec2InstanceType)
}

func EncodeGameliftFleet_Id(p GameliftFleetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGameliftFleet_MetricGroups(p GameliftFleetParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.MetricGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["metric_groups"] = cty.ListVal(colVals)
}

func EncodeGameliftFleet_BuildId(p GameliftFleetParameters, vals map[string]cty.Value) {
	vals["build_id"] = cty.StringVal(p.BuildId)
}

func EncodeGameliftFleet_Description(p GameliftFleetParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGameliftFleet_NewGameSessionProtectionPolicy(p GameliftFleetParameters, vals map[string]cty.Value) {
	vals["new_game_session_protection_policy"] = cty.StringVal(p.NewGameSessionProtectionPolicy)
}

func EncodeGameliftFleet_Tags(p GameliftFleetParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGameliftFleet_FleetType(p GameliftFleetParameters, vals map[string]cty.Value) {
	vals["fleet_type"] = cty.StringVal(p.FleetType)
}

func EncodeGameliftFleet_InstanceRoleArn(p GameliftFleetParameters, vals map[string]cty.Value) {
	vals["instance_role_arn"] = cty.StringVal(p.InstanceRoleArn)
}

func EncodeGameliftFleet_Name(p GameliftFleetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGameliftFleet_Ec2InboundPermission(p []Ec2InboundPermission, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGameliftFleet_Ec2InboundPermission_FromPort(v, ctyVal)
		EncodeGameliftFleet_Ec2InboundPermission_IpRange(v, ctyVal)
		EncodeGameliftFleet_Ec2InboundPermission_Protocol(v, ctyVal)
		EncodeGameliftFleet_Ec2InboundPermission_ToPort(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ec2_inbound_permission"] = cty.ListVal(valsForCollection)
}

func EncodeGameliftFleet_Ec2InboundPermission_FromPort(p Ec2InboundPermission, vals map[string]cty.Value) {
	vals["from_port"] = cty.NumberIntVal(p.FromPort)
}

func EncodeGameliftFleet_Ec2InboundPermission_IpRange(p Ec2InboundPermission, vals map[string]cty.Value) {
	vals["ip_range"] = cty.StringVal(p.IpRange)
}

func EncodeGameliftFleet_Ec2InboundPermission_Protocol(p Ec2InboundPermission, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeGameliftFleet_Ec2InboundPermission_ToPort(p Ec2InboundPermission, vals map[string]cty.Value) {
	vals["to_port"] = cty.NumberIntVal(p.ToPort)
}

func EncodeGameliftFleet_ResourceCreationLimitPolicy(p ResourceCreationLimitPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftFleet_ResourceCreationLimitPolicy_PolicyPeriodInMinutes(p, ctyVal)
	EncodeGameliftFleet_ResourceCreationLimitPolicy_NewGameSessionsPerCreator(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["resource_creation_limit_policy"] = cty.ListVal(valsForCollection)
}

func EncodeGameliftFleet_ResourceCreationLimitPolicy_PolicyPeriodInMinutes(p ResourceCreationLimitPolicy, vals map[string]cty.Value) {
	vals["policy_period_in_minutes"] = cty.NumberIntVal(p.PolicyPeriodInMinutes)
}

func EncodeGameliftFleet_ResourceCreationLimitPolicy_NewGameSessionsPerCreator(p ResourceCreationLimitPolicy, vals map[string]cty.Value) {
	vals["new_game_sessions_per_creator"] = cty.NumberIntVal(p.NewGameSessionsPerCreator)
}

func EncodeGameliftFleet_RuntimeConfiguration(p RuntimeConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftFleet_RuntimeConfiguration_GameSessionActivationTimeoutSeconds(p, ctyVal)
	EncodeGameliftFleet_RuntimeConfiguration_MaxConcurrentGameSessionActivations(p, ctyVal)
	EncodeGameliftFleet_RuntimeConfiguration_ServerProcess(p.ServerProcess, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["runtime_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeGameliftFleet_RuntimeConfiguration_GameSessionActivationTimeoutSeconds(p RuntimeConfiguration, vals map[string]cty.Value) {
	vals["game_session_activation_timeout_seconds"] = cty.NumberIntVal(p.GameSessionActivationTimeoutSeconds)
}

func EncodeGameliftFleet_RuntimeConfiguration_MaxConcurrentGameSessionActivations(p RuntimeConfiguration, vals map[string]cty.Value) {
	vals["max_concurrent_game_session_activations"] = cty.NumberIntVal(p.MaxConcurrentGameSessionActivations)
}

func EncodeGameliftFleet_RuntimeConfiguration_ServerProcess(p []ServerProcess, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGameliftFleet_RuntimeConfiguration_ServerProcess_ConcurrentExecutions(v, ctyVal)
		EncodeGameliftFleet_RuntimeConfiguration_ServerProcess_LaunchPath(v, ctyVal)
		EncodeGameliftFleet_RuntimeConfiguration_ServerProcess_Parameters(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["server_process"] = cty.ListVal(valsForCollection)
}

func EncodeGameliftFleet_RuntimeConfiguration_ServerProcess_ConcurrentExecutions(p ServerProcess, vals map[string]cty.Value) {
	vals["concurrent_executions"] = cty.NumberIntVal(p.ConcurrentExecutions)
}

func EncodeGameliftFleet_RuntimeConfiguration_ServerProcess_LaunchPath(p ServerProcess, vals map[string]cty.Value) {
	vals["launch_path"] = cty.StringVal(p.LaunchPath)
}

func EncodeGameliftFleet_RuntimeConfiguration_ServerProcess_Parameters(p ServerProcess, vals map[string]cty.Value) {
	vals["parameters"] = cty.StringVal(p.Parameters)
}

func EncodeGameliftFleet_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftFleet_Timeouts_Delete(p, ctyVal)
	EncodeGameliftFleet_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeGameliftFleet_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeGameliftFleet_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeGameliftFleet_LogPaths(p GameliftFleetObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LogPaths {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["log_paths"] = cty.ListVal(colVals)
}

func EncodeGameliftFleet_Arn(p GameliftFleetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeGameliftFleet_OperatingSystem(p GameliftFleetObservation, vals map[string]cty.Value) {
	vals["operating_system"] = cty.StringVal(p.OperatingSystem)
}