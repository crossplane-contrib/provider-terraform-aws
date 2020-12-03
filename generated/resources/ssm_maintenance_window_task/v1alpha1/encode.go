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

package v1alpha1func EncodeSsmMaintenanceWindowTask(r SsmMaintenanceWindowTask) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSsmMaintenanceWindowTask_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_MaxConcurrency(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_Priority(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_ServiceRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_TaskArn(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_TaskType(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_WindowId(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_Description(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_MaxErrors(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTask_Targets(r.Spec.ForProvider.Targets, ctyVal)
	EncodeSsmMaintenanceWindowTask_TaskInvocationParameters(r.Spec.ForProvider.TaskInvocationParameters, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeSsmMaintenanceWindowTask_Name(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmMaintenanceWindowTask_Id(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmMaintenanceWindowTask_MaxConcurrency(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["max_concurrency"] = cty.StringVal(p.MaxConcurrency)
}

func EncodeSsmMaintenanceWindowTask_Priority(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeSsmMaintenanceWindowTask_ServiceRoleArn(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["service_role_arn"] = cty.StringVal(p.ServiceRoleArn)
}

func EncodeSsmMaintenanceWindowTask_TaskArn(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["task_arn"] = cty.StringVal(p.TaskArn)
}

func EncodeSsmMaintenanceWindowTask_TaskType(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["task_type"] = cty.StringVal(p.TaskType)
}

func EncodeSsmMaintenanceWindowTask_WindowId(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["window_id"] = cty.StringVal(p.WindowId)
}

func EncodeSsmMaintenanceWindowTask_Description(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSsmMaintenanceWindowTask_MaxErrors(p *SsmMaintenanceWindowTaskParameters, vals map[string]cty.Value) {
	vals["max_errors"] = cty.StringVal(p.MaxErrors)
}

func EncodeSsmMaintenanceWindowTask_Targets(p *Targets, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Targets {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTask_Targets_Key(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_Targets_Values(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["targets"] = cty.ListVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTask_Targets_Key(p *Targets, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeSsmMaintenanceWindowTask_Targets_Values(p *Targets, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters(p *TaskInvocationParameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TaskInvocationParameters {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_LambdaParameters(v.LambdaParameters, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters(v.RunCommandParameters, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_StepFunctionsParameters(v.StepFunctionsParameters, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters(v.AutomationParameters, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["task_invocation_parameters"] = cty.ListVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_LambdaParameters(p *LambdaParameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LambdaParameters {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_LambdaParameters_Qualifier(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_LambdaParameters_ClientContext(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_LambdaParameters_Payload(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lambda_parameters"] = cty.ListVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_LambdaParameters_Qualifier(p *LambdaParameters, vals map[string]cty.Value) {
	vals["qualifier"] = cty.StringVal(p.Qualifier)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_LambdaParameters_ClientContext(p *LambdaParameters, vals map[string]cty.Value) {
	vals["client_context"] = cty.StringVal(p.ClientContext)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_LambdaParameters_Payload(p *LambdaParameters, vals map[string]cty.Value) {
	vals["payload"] = cty.StringVal(p.Payload)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters(p *RunCommandParameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RunCommandParameters {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_TimeoutSeconds(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_Comment(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_DocumentHash(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_DocumentHashType(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_OutputS3Bucket(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_OutputS3KeyPrefix(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_ServiceRoleArn(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_NotificationConfig(v.NotificationConfig, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_Parameter(v.Parameter, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["run_command_parameters"] = cty.ListVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_TimeoutSeconds(p *RunCommandParameters, vals map[string]cty.Value) {
	vals["timeout_seconds"] = cty.IntVal(p.TimeoutSeconds)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_Comment(p *RunCommandParameters, vals map[string]cty.Value) {
	vals["comment"] = cty.StringVal(p.Comment)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_DocumentHash(p *RunCommandParameters, vals map[string]cty.Value) {
	vals["document_hash"] = cty.StringVal(p.DocumentHash)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_DocumentHashType(p *RunCommandParameters, vals map[string]cty.Value) {
	vals["document_hash_type"] = cty.StringVal(p.DocumentHashType)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_OutputS3Bucket(p *RunCommandParameters, vals map[string]cty.Value) {
	vals["output_s3_bucket"] = cty.StringVal(p.OutputS3Bucket)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_OutputS3KeyPrefix(p *RunCommandParameters, vals map[string]cty.Value) {
	vals["output_s3_key_prefix"] = cty.StringVal(p.OutputS3KeyPrefix)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_ServiceRoleArn(p *RunCommandParameters, vals map[string]cty.Value) {
	vals["service_role_arn"] = cty.StringVal(p.ServiceRoleArn)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_NotificationConfig(p *NotificationConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NotificationConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_NotificationConfig_NotificationEvents(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_NotificationConfig_NotificationType(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_NotificationConfig_NotificationArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["notification_config"] = cty.ListVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_NotificationConfig_NotificationEvents(p *NotificationConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NotificationEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["notification_events"] = cty.ListVal(colVals)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_NotificationConfig_NotificationType(p *NotificationConfig, vals map[string]cty.Value) {
	vals["notification_type"] = cty.StringVal(p.NotificationType)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_NotificationConfig_NotificationArn(p *NotificationConfig, vals map[string]cty.Value) {
	vals["notification_arn"] = cty.StringVal(p.NotificationArn)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_Parameter(p *Parameter, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameter {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_Parameter_Name(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_Parameter_Values(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_Parameter_Name(p *Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_RunCommandParameters_Parameter_Values(p *Parameter, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_StepFunctionsParameters(p *StepFunctionsParameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.StepFunctionsParameters {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_StepFunctionsParameters_Input(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_StepFunctionsParameters_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["step_functions_parameters"] = cty.ListVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_StepFunctionsParameters_Input(p *StepFunctionsParameters, vals map[string]cty.Value) {
	vals["input"] = cty.StringVal(p.Input)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_StepFunctionsParameters_Name(p *StepFunctionsParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters(p *AutomationParameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AutomationParameters {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters_DocumentVersion(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters_Parameter(v.Parameter, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["automation_parameters"] = cty.ListVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters_DocumentVersion(p *AutomationParameters, vals map[string]cty.Value) {
	vals["document_version"] = cty.StringVal(p.DocumentVersion)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters_Parameter(p *Parameter, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameter {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters_Parameter_Name(v, ctyVal)
		EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters_Parameter_Values(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters_Parameter_Name(p *Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmMaintenanceWindowTask_TaskInvocationParameters_AutomationParameters_Parameter_Values(p *Parameter, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}