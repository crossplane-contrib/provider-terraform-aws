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

func EncodeMqBroker(r MqBroker) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeMqBroker_AutoMinorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_PubliclyAccessible(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_Tags(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_BrokerName(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_EngineType(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_Id(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_DeploymentMode(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_HostInstanceType(r.Spec.ForProvider, ctyVal)
	EncodeMqBroker_Configuration(r.Spec.ForProvider.Configuration, ctyVal)
	EncodeMqBroker_EncryptionOptions(r.Spec.ForProvider.EncryptionOptions, ctyVal)
	EncodeMqBroker_Logs(r.Spec.ForProvider.Logs, ctyVal)
	EncodeMqBroker_MaintenanceWindowStartTime(r.Spec.ForProvider.MaintenanceWindowStartTime, ctyVal)
	EncodeMqBroker_User(r.Spec.ForProvider.User, ctyVal)
	EncodeMqBroker_Arn(r.Status.AtProvider, ctyVal)
	EncodeMqBroker_Instances(r.Status.AtProvider.Instances, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeMqBroker_AutoMinorVersionUpgrade(p MqBrokerParameters, vals map[string]cty.Value) {
	vals["auto_minor_version_upgrade"] = cty.BoolVal(p.AutoMinorVersionUpgrade)
}

func EncodeMqBroker_PubliclyAccessible(p MqBrokerParameters, vals map[string]cty.Value) {
	vals["publicly_accessible"] = cty.BoolVal(p.PubliclyAccessible)
}

func EncodeMqBroker_SecurityGroups(p MqBrokerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeMqBroker_ApplyImmediately(p MqBrokerParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeMqBroker_EngineVersion(p MqBrokerParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeMqBroker_Tags(p MqBrokerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeMqBroker_BrokerName(p MqBrokerParameters, vals map[string]cty.Value) {
	vals["broker_name"] = cty.StringVal(p.BrokerName)
}

func EncodeMqBroker_EngineType(p MqBrokerParameters, vals map[string]cty.Value) {
	vals["engine_type"] = cty.StringVal(p.EngineType)
}

func EncodeMqBroker_Id(p MqBrokerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMqBroker_SubnetIds(p MqBrokerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeMqBroker_DeploymentMode(p MqBrokerParameters, vals map[string]cty.Value) {
	vals["deployment_mode"] = cty.StringVal(p.DeploymentMode)
}

func EncodeMqBroker_HostInstanceType(p MqBrokerParameters, vals map[string]cty.Value) {
	vals["host_instance_type"] = cty.StringVal(p.HostInstanceType)
}

func EncodeMqBroker_Configuration(p Configuration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMqBroker_Configuration_Revision(p, ctyVal)
	EncodeMqBroker_Configuration_Id(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["configuration"] = cty.ListVal(valsForCollection)
}

func EncodeMqBroker_Configuration_Revision(p Configuration, vals map[string]cty.Value) {
	vals["revision"] = cty.NumberIntVal(p.Revision)
}

func EncodeMqBroker_Configuration_Id(p Configuration, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMqBroker_EncryptionOptions(p EncryptionOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMqBroker_EncryptionOptions_KmsKeyId(p, ctyVal)
	EncodeMqBroker_EncryptionOptions_UseAwsOwnedKey(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption_options"] = cty.ListVal(valsForCollection)
}

func EncodeMqBroker_EncryptionOptions_KmsKeyId(p EncryptionOptions, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeMqBroker_EncryptionOptions_UseAwsOwnedKey(p EncryptionOptions, vals map[string]cty.Value) {
	vals["use_aws_owned_key"] = cty.BoolVal(p.UseAwsOwnedKey)
}

func EncodeMqBroker_Logs(p Logs, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMqBroker_Logs_Audit(p, ctyVal)
	EncodeMqBroker_Logs_General(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["logs"] = cty.ListVal(valsForCollection)
}

func EncodeMqBroker_Logs_Audit(p Logs, vals map[string]cty.Value) {
	vals["audit"] = cty.BoolVal(p.Audit)
}

func EncodeMqBroker_Logs_General(p Logs, vals map[string]cty.Value) {
	vals["general"] = cty.BoolVal(p.General)
}

func EncodeMqBroker_MaintenanceWindowStartTime(p MaintenanceWindowStartTime, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMqBroker_MaintenanceWindowStartTime_DayOfWeek(p, ctyVal)
	EncodeMqBroker_MaintenanceWindowStartTime_TimeOfDay(p, ctyVal)
	EncodeMqBroker_MaintenanceWindowStartTime_TimeZone(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["maintenance_window_start_time"] = cty.ListVal(valsForCollection)
}

func EncodeMqBroker_MaintenanceWindowStartTime_DayOfWeek(p MaintenanceWindowStartTime, vals map[string]cty.Value) {
	vals["day_of_week"] = cty.StringVal(p.DayOfWeek)
}

func EncodeMqBroker_MaintenanceWindowStartTime_TimeOfDay(p MaintenanceWindowStartTime, vals map[string]cty.Value) {
	vals["time_of_day"] = cty.StringVal(p.TimeOfDay)
}

func EncodeMqBroker_MaintenanceWindowStartTime_TimeZone(p MaintenanceWindowStartTime, vals map[string]cty.Value) {
	vals["time_zone"] = cty.StringVal(p.TimeZone)
}

func EncodeMqBroker_User(p []User, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeMqBroker_User_ConsoleAccess(v, ctyVal)
		EncodeMqBroker_User_Groups(v, ctyVal)
		EncodeMqBroker_User_Password(v, ctyVal)
		EncodeMqBroker_User_Username(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["user"] = cty.SetVal(valsForCollection)
}

func EncodeMqBroker_User_ConsoleAccess(p User, vals map[string]cty.Value) {
	vals["console_access"] = cty.BoolVal(p.ConsoleAccess)
}

func EncodeMqBroker_User_Groups(p User, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Groups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["groups"] = cty.SetVal(colVals)
}

func EncodeMqBroker_User_Password(p User, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeMqBroker_User_Username(p User, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeMqBroker_Arn(p MqBrokerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeMqBroker_Instances(p []Instances, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeMqBroker_Instances_ConsoleUrl(v, ctyVal)
		EncodeMqBroker_Instances_Endpoints(v, ctyVal)
		EncodeMqBroker_Instances_IpAddress(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["instances"] = cty.ListVal(valsForCollection)
}

func EncodeMqBroker_Instances_ConsoleUrl(p Instances, vals map[string]cty.Value) {
	vals["console_url"] = cty.StringVal(p.ConsoleUrl)
}

func EncodeMqBroker_Instances_Endpoints(p Instances, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Endpoints {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["endpoints"] = cty.ListVal(colVals)
}

func EncodeMqBroker_Instances_IpAddress(p Instances, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}