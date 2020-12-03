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

package v1alpha1func EncodeCloudwatchEventTarget(r CloudwatchEventTarget) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCloudwatchEventTarget_Arn(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventTarget_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventTarget_Input(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventTarget_InputPath(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventTarget_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventTarget_Rule(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventTarget_TargetId(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventTarget_BatchTarget(r.Spec.ForProvider.BatchTarget, ctyVal)
	EncodeCloudwatchEventTarget_EcsTarget(r.Spec.ForProvider.EcsTarget, ctyVal)
	EncodeCloudwatchEventTarget_InputTransformer(r.Spec.ForProvider.InputTransformer, ctyVal)
	EncodeCloudwatchEventTarget_KinesisTarget(r.Spec.ForProvider.KinesisTarget, ctyVal)
	EncodeCloudwatchEventTarget_RunCommandTargets(r.Spec.ForProvider.RunCommandTargets, ctyVal)
	EncodeCloudwatchEventTarget_SqsTarget(r.Spec.ForProvider.SqsTarget, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeCloudwatchEventTarget_Arn(p *CloudwatchEventTargetParameters, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCloudwatchEventTarget_Id(p *CloudwatchEventTargetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchEventTarget_Input(p *CloudwatchEventTargetParameters, vals map[string]cty.Value) {
	vals["input"] = cty.StringVal(p.Input)
}

func EncodeCloudwatchEventTarget_InputPath(p *CloudwatchEventTargetParameters, vals map[string]cty.Value) {
	vals["input_path"] = cty.StringVal(p.InputPath)
}

func EncodeCloudwatchEventTarget_RoleArn(p *CloudwatchEventTargetParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCloudwatchEventTarget_Rule(p *CloudwatchEventTargetParameters, vals map[string]cty.Value) {
	vals["rule"] = cty.StringVal(p.Rule)
}

func EncodeCloudwatchEventTarget_TargetId(p *CloudwatchEventTargetParameters, vals map[string]cty.Value) {
	vals["target_id"] = cty.StringVal(p.TargetId)
}

func EncodeCloudwatchEventTarget_BatchTarget(p *BatchTarget, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.BatchTarget {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudwatchEventTarget_BatchTarget_JobName(v, ctyVal)
		EncodeCloudwatchEventTarget_BatchTarget_ArraySize(v, ctyVal)
		EncodeCloudwatchEventTarget_BatchTarget_JobAttempts(v, ctyVal)
		EncodeCloudwatchEventTarget_BatchTarget_JobDefinition(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["batch_target"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchEventTarget_BatchTarget_JobName(p *BatchTarget, vals map[string]cty.Value) {
	vals["job_name"] = cty.StringVal(p.JobName)
}

func EncodeCloudwatchEventTarget_BatchTarget_ArraySize(p *BatchTarget, vals map[string]cty.Value) {
	vals["array_size"] = cty.IntVal(p.ArraySize)
}

func EncodeCloudwatchEventTarget_BatchTarget_JobAttempts(p *BatchTarget, vals map[string]cty.Value) {
	vals["job_attempts"] = cty.IntVal(p.JobAttempts)
}

func EncodeCloudwatchEventTarget_BatchTarget_JobDefinition(p *BatchTarget, vals map[string]cty.Value) {
	vals["job_definition"] = cty.StringVal(p.JobDefinition)
}

func EncodeCloudwatchEventTarget_EcsTarget(p *EcsTarget, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EcsTarget {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudwatchEventTarget_EcsTarget_Group(v, ctyVal)
		EncodeCloudwatchEventTarget_EcsTarget_LaunchType(v, ctyVal)
		EncodeCloudwatchEventTarget_EcsTarget_PlatformVersion(v, ctyVal)
		EncodeCloudwatchEventTarget_EcsTarget_TaskCount(v, ctyVal)
		EncodeCloudwatchEventTarget_EcsTarget_TaskDefinitionArn(v, ctyVal)
		EncodeCloudwatchEventTarget_EcsTarget_NetworkConfiguration(v.NetworkConfiguration, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ecs_target"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchEventTarget_EcsTarget_Group(p *EcsTarget, vals map[string]cty.Value) {
	vals["group"] = cty.StringVal(p.Group)
}

func EncodeCloudwatchEventTarget_EcsTarget_LaunchType(p *EcsTarget, vals map[string]cty.Value) {
	vals["launch_type"] = cty.StringVal(p.LaunchType)
}

func EncodeCloudwatchEventTarget_EcsTarget_PlatformVersion(p *EcsTarget, vals map[string]cty.Value) {
	vals["platform_version"] = cty.StringVal(p.PlatformVersion)
}

func EncodeCloudwatchEventTarget_EcsTarget_TaskCount(p *EcsTarget, vals map[string]cty.Value) {
	vals["task_count"] = cty.IntVal(p.TaskCount)
}

func EncodeCloudwatchEventTarget_EcsTarget_TaskDefinitionArn(p *EcsTarget, vals map[string]cty.Value) {
	vals["task_definition_arn"] = cty.StringVal(p.TaskDefinitionArn)
}

func EncodeCloudwatchEventTarget_EcsTarget_NetworkConfiguration(p *NetworkConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NetworkConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudwatchEventTarget_EcsTarget_NetworkConfiguration_SecurityGroups(v, ctyVal)
		EncodeCloudwatchEventTarget_EcsTarget_NetworkConfiguration_Subnets(v, ctyVal)
		EncodeCloudwatchEventTarget_EcsTarget_NetworkConfiguration_AssignPublicIp(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["network_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchEventTarget_EcsTarget_NetworkConfiguration_SecurityGroups(p *NetworkConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeCloudwatchEventTarget_EcsTarget_NetworkConfiguration_Subnets(p *NetworkConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Subnets {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnets"] = cty.SetVal(colVals)
}

func EncodeCloudwatchEventTarget_EcsTarget_NetworkConfiguration_AssignPublicIp(p *NetworkConfiguration, vals map[string]cty.Value) {
	vals["assign_public_ip"] = cty.BoolVal(p.AssignPublicIp)
}

func EncodeCloudwatchEventTarget_InputTransformer(p *InputTransformer, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.InputTransformer {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudwatchEventTarget_InputTransformer_InputPaths(v, ctyVal)
		EncodeCloudwatchEventTarget_InputTransformer_InputTemplate(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["input_transformer"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchEventTarget_InputTransformer_InputPaths(p *InputTransformer, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.InputPaths {
		mVals[key] = cty.StringVal(value)
	}
	vals["input_paths"] = cty.MapVal(mVals)
}

func EncodeCloudwatchEventTarget_InputTransformer_InputTemplate(p *InputTransformer, vals map[string]cty.Value) {
	vals["input_template"] = cty.StringVal(p.InputTemplate)
}

func EncodeCloudwatchEventTarget_KinesisTarget(p *KinesisTarget, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.KinesisTarget {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudwatchEventTarget_KinesisTarget_PartitionKeyPath(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["kinesis_target"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchEventTarget_KinesisTarget_PartitionKeyPath(p *KinesisTarget, vals map[string]cty.Value) {
	vals["partition_key_path"] = cty.StringVal(p.PartitionKeyPath)
}

func EncodeCloudwatchEventTarget_RunCommandTargets(p *RunCommandTargets, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RunCommandTargets {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudwatchEventTarget_RunCommandTargets_Key(v, ctyVal)
		EncodeCloudwatchEventTarget_RunCommandTargets_Values(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["run_command_targets"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchEventTarget_RunCommandTargets_Key(p *RunCommandTargets, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeCloudwatchEventTarget_RunCommandTargets_Values(p *RunCommandTargets, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}

func EncodeCloudwatchEventTarget_SqsTarget(p *SqsTarget, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqsTarget {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudwatchEventTarget_SqsTarget_MessageGroupId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqs_target"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchEventTarget_SqsTarget_MessageGroupId(p *SqsTarget, vals map[string]cty.Value) {
	vals["message_group_id"] = cty.StringVal(p.MessageGroupId)
}