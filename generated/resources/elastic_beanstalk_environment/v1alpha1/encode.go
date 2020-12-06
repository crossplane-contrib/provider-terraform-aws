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

func EncodeElasticBeanstalkEnvironment(r ElasticBeanstalkEnvironment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeElasticBeanstalkEnvironment_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_PlatformArn(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_PollInterval(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_SolutionStackName(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Tags(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_TemplateName(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Tier(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_CnamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Name(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_VersionLabel(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Application(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Description(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_WaitForReadyTimeout(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Setting(r.Spec.ForProvider.Setting, ctyVal)
	EncodeElasticBeanstalkEnvironment_Instances(r.Status.AtProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_LaunchConfigurations(r.Status.AtProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_LoadBalancers(r.Status.AtProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Arn(r.Status.AtProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_EndpointUrl(r.Status.AtProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Triggers(r.Status.AtProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_AutoscalingGroups(r.Status.AtProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Cname(r.Status.AtProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_Queues(r.Status.AtProvider, ctyVal)
	EncodeElasticBeanstalkEnvironment_AllSettings(r.Status.AtProvider.AllSettings, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeElasticBeanstalkEnvironment_Id(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticBeanstalkEnvironment_PlatformArn(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["platform_arn"] = cty.StringVal(p.PlatformArn)
}

func EncodeElasticBeanstalkEnvironment_PollInterval(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["poll_interval"] = cty.StringVal(p.PollInterval)
}

func EncodeElasticBeanstalkEnvironment_SolutionStackName(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["solution_stack_name"] = cty.StringVal(p.SolutionStackName)
}

func EncodeElasticBeanstalkEnvironment_Tags(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeElasticBeanstalkEnvironment_TemplateName(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["template_name"] = cty.StringVal(p.TemplateName)
}

func EncodeElasticBeanstalkEnvironment_Tier(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["tier"] = cty.StringVal(p.Tier)
}

func EncodeElasticBeanstalkEnvironment_CnamePrefix(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["cname_prefix"] = cty.StringVal(p.CnamePrefix)
}

func EncodeElasticBeanstalkEnvironment_Name(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticBeanstalkEnvironment_VersionLabel(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["version_label"] = cty.StringVal(p.VersionLabel)
}

func EncodeElasticBeanstalkEnvironment_Application(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["application"] = cty.StringVal(p.Application)
}

func EncodeElasticBeanstalkEnvironment_Description(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeElasticBeanstalkEnvironment_WaitForReadyTimeout(p ElasticBeanstalkEnvironmentParameters, vals map[string]cty.Value) {
	vals["wait_for_ready_timeout"] = cty.StringVal(p.WaitForReadyTimeout)
}

func EncodeElasticBeanstalkEnvironment_Setting(p Setting, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeElasticBeanstalkEnvironment_Setting_Value(p, ctyVal)
	EncodeElasticBeanstalkEnvironment_Setting_Name(p, ctyVal)
	EncodeElasticBeanstalkEnvironment_Setting_Namespace(p, ctyVal)
	EncodeElasticBeanstalkEnvironment_Setting_Resource(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["setting"] = cty.SetVal(valsForCollection)
}

func EncodeElasticBeanstalkEnvironment_Setting_Value(p Setting, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeElasticBeanstalkEnvironment_Setting_Name(p Setting, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticBeanstalkEnvironment_Setting_Namespace(p Setting, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeElasticBeanstalkEnvironment_Setting_Resource(p Setting, vals map[string]cty.Value) {
	vals["resource"] = cty.StringVal(p.Resource)
}

func EncodeElasticBeanstalkEnvironment_Instances(p ElasticBeanstalkEnvironmentObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Instances {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["instances"] = cty.ListVal(colVals)
}

func EncodeElasticBeanstalkEnvironment_LaunchConfigurations(p ElasticBeanstalkEnvironmentObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LaunchConfigurations {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["launch_configurations"] = cty.ListVal(colVals)
}

func EncodeElasticBeanstalkEnvironment_LoadBalancers(p ElasticBeanstalkEnvironmentObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LoadBalancers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["load_balancers"] = cty.ListVal(colVals)
}

func EncodeElasticBeanstalkEnvironment_Arn(p ElasticBeanstalkEnvironmentObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeElasticBeanstalkEnvironment_EndpointUrl(p ElasticBeanstalkEnvironmentObservation, vals map[string]cty.Value) {
	vals["endpoint_url"] = cty.StringVal(p.EndpointUrl)
}

func EncodeElasticBeanstalkEnvironment_Triggers(p ElasticBeanstalkEnvironmentObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Triggers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["triggers"] = cty.ListVal(colVals)
}

func EncodeElasticBeanstalkEnvironment_AutoscalingGroups(p ElasticBeanstalkEnvironmentObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AutoscalingGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["autoscaling_groups"] = cty.ListVal(colVals)
}

func EncodeElasticBeanstalkEnvironment_Cname(p ElasticBeanstalkEnvironmentObservation, vals map[string]cty.Value) {
	vals["cname"] = cty.StringVal(p.Cname)
}

func EncodeElasticBeanstalkEnvironment_Queues(p ElasticBeanstalkEnvironmentObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Queues {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["queues"] = cty.ListVal(colVals)
}

func EncodeElasticBeanstalkEnvironment_AllSettings(p []AllSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeElasticBeanstalkEnvironment_AllSettings_Name(v, ctyVal)
		EncodeElasticBeanstalkEnvironment_AllSettings_Namespace(v, ctyVal)
		EncodeElasticBeanstalkEnvironment_AllSettings_Resource(v, ctyVal)
		EncodeElasticBeanstalkEnvironment_AllSettings_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_settings"] = cty.SetVal(valsForCollection)
}

func EncodeElasticBeanstalkEnvironment_AllSettings_Name(p AllSettings, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticBeanstalkEnvironment_AllSettings_Namespace(p AllSettings, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeElasticBeanstalkEnvironment_AllSettings_Resource(p AllSettings, vals map[string]cty.Value) {
	vals["resource"] = cty.StringVal(p.Resource)
}

func EncodeElasticBeanstalkEnvironment_AllSettings_Value(p AllSettings, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}