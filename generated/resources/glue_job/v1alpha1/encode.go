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

func EncodeGlueJob(r GlueJob) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueJob_SecurityConfiguration(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_Timeout(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_NonOverridableArguments(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_NumberOfWorkers(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_MaxCapacity(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_Description(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_MaxRetries(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_WorkerType(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_DefaultArguments(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_GlueVersion(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_Connections(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueJob_Command(r.Spec.ForProvider.Command, ctyVal)
	EncodeGlueJob_ExecutionProperty(r.Spec.ForProvider.ExecutionProperty, ctyVal)
	EncodeGlueJob_NotificationProperty(r.Spec.ForProvider.NotificationProperty, ctyVal)
	EncodeGlueJob_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlueJob_SecurityConfiguration(p GlueJobParameters, vals map[string]cty.Value) {
	vals["security_configuration"] = cty.StringVal(p.SecurityConfiguration)
}

func EncodeGlueJob_Timeout(p GlueJobParameters, vals map[string]cty.Value) {
	vals["timeout"] = cty.NumberIntVal(p.Timeout)
}

func EncodeGlueJob_NonOverridableArguments(p GlueJobParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.NonOverridableArguments {
		mVals[key] = cty.StringVal(value)
	}
	vals["non_overridable_arguments"] = cty.MapVal(mVals)
}

func EncodeGlueJob_NumberOfWorkers(p GlueJobParameters, vals map[string]cty.Value) {
	vals["number_of_workers"] = cty.NumberIntVal(p.NumberOfWorkers)
}

func EncodeGlueJob_MaxCapacity(p GlueJobParameters, vals map[string]cty.Value) {
	vals["max_capacity"] = cty.NumberIntVal(p.MaxCapacity)
}

func EncodeGlueJob_Description(p GlueJobParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGlueJob_MaxRetries(p GlueJobParameters, vals map[string]cty.Value) {
	vals["max_retries"] = cty.NumberIntVal(p.MaxRetries)
}

func EncodeGlueJob_Tags(p GlueJobParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGlueJob_WorkerType(p GlueJobParameters, vals map[string]cty.Value) {
	vals["worker_type"] = cty.StringVal(p.WorkerType)
}

func EncodeGlueJob_DefaultArguments(p GlueJobParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.DefaultArguments {
		mVals[key] = cty.StringVal(value)
	}
	vals["default_arguments"] = cty.MapVal(mVals)
}

func EncodeGlueJob_GlueVersion(p GlueJobParameters, vals map[string]cty.Value) {
	vals["glue_version"] = cty.StringVal(p.GlueVersion)
}

func EncodeGlueJob_Name(p GlueJobParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueJob_RoleArn(p GlueJobParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeGlueJob_Connections(p GlueJobParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Connections {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["connections"] = cty.ListVal(colVals)
}

func EncodeGlueJob_Id(p GlueJobParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueJob_Command(p Command, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueJob_Command_Name(p, ctyVal)
	EncodeGlueJob_Command_PythonVersion(p, ctyVal)
	EncodeGlueJob_Command_ScriptLocation(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["command"] = cty.ListVal(valsForCollection)
}

func EncodeGlueJob_Command_Name(p Command, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueJob_Command_PythonVersion(p Command, vals map[string]cty.Value) {
	vals["python_version"] = cty.StringVal(p.PythonVersion)
}

func EncodeGlueJob_Command_ScriptLocation(p Command, vals map[string]cty.Value) {
	vals["script_location"] = cty.StringVal(p.ScriptLocation)
}

func EncodeGlueJob_ExecutionProperty(p ExecutionProperty, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueJob_ExecutionProperty_MaxConcurrentRuns(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["execution_property"] = cty.ListVal(valsForCollection)
}

func EncodeGlueJob_ExecutionProperty_MaxConcurrentRuns(p ExecutionProperty, vals map[string]cty.Value) {
	vals["max_concurrent_runs"] = cty.NumberIntVal(p.MaxConcurrentRuns)
}

func EncodeGlueJob_NotificationProperty(p NotificationProperty, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueJob_NotificationProperty_NotifyDelayAfter(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["notification_property"] = cty.ListVal(valsForCollection)
}

func EncodeGlueJob_NotificationProperty_NotifyDelayAfter(p NotificationProperty, vals map[string]cty.Value) {
	vals["notify_delay_after"] = cty.NumberIntVal(p.NotifyDelayAfter)
}

func EncodeGlueJob_Arn(p GlueJobObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}