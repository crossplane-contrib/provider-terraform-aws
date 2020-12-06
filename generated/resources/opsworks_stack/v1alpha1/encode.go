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

func EncodeOpsworksStack(r OpsworksStack) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksStack_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_DefaultInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_DefaultOs(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_DefaultSubnetId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_UseOpsworksSecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_AgentVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_Color(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_DefaultAvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_DefaultRootDeviceType(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_UseCustomCookbooks(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_ConfigurationManagerName(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_ManageBerkshelf(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_Region(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_ServiceRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_BerkshelfVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_ConfigurationManagerVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_DefaultSshKeyName(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_HostnameTheme(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStack_CustomCookbooksSource(r.Spec.ForProvider.CustomCookbooksSource, ctyVal)
	EncodeOpsworksStack_StackEndpoint(r.Status.AtProvider, ctyVal)
	EncodeOpsworksStack_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksStack_Id(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksStack_Name(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksStack_CustomJson(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksStack_DefaultInstanceProfileArn(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["default_instance_profile_arn"] = cty.StringVal(p.DefaultInstanceProfileArn)
}

func EncodeOpsworksStack_DefaultOs(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["default_os"] = cty.StringVal(p.DefaultOs)
}

func EncodeOpsworksStack_DefaultSubnetId(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["default_subnet_id"] = cty.StringVal(p.DefaultSubnetId)
}

func EncodeOpsworksStack_UseOpsworksSecurityGroups(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["use_opsworks_security_groups"] = cty.BoolVal(p.UseOpsworksSecurityGroups)
}

func EncodeOpsworksStack_VpcId(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeOpsworksStack_AgentVersion(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["agent_version"] = cty.StringVal(p.AgentVersion)
}

func EncodeOpsworksStack_Color(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["color"] = cty.StringVal(p.Color)
}

func EncodeOpsworksStack_DefaultAvailabilityZone(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["default_availability_zone"] = cty.StringVal(p.DefaultAvailabilityZone)
}

func EncodeOpsworksStack_DefaultRootDeviceType(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["default_root_device_type"] = cty.StringVal(p.DefaultRootDeviceType)
}

func EncodeOpsworksStack_UseCustomCookbooks(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["use_custom_cookbooks"] = cty.BoolVal(p.UseCustomCookbooks)
}

func EncodeOpsworksStack_ConfigurationManagerName(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["configuration_manager_name"] = cty.StringVal(p.ConfigurationManagerName)
}

func EncodeOpsworksStack_ManageBerkshelf(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["manage_berkshelf"] = cty.BoolVal(p.ManageBerkshelf)
}

func EncodeOpsworksStack_Region(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeOpsworksStack_ServiceRoleArn(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["service_role_arn"] = cty.StringVal(p.ServiceRoleArn)
}

func EncodeOpsworksStack_Tags(p OpsworksStackParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksStack_BerkshelfVersion(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["berkshelf_version"] = cty.StringVal(p.BerkshelfVersion)
}

func EncodeOpsworksStack_ConfigurationManagerVersion(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["configuration_manager_version"] = cty.StringVal(p.ConfigurationManagerVersion)
}

func EncodeOpsworksStack_DefaultSshKeyName(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["default_ssh_key_name"] = cty.StringVal(p.DefaultSshKeyName)
}

func EncodeOpsworksStack_HostnameTheme(p OpsworksStackParameters, vals map[string]cty.Value) {
	vals["hostname_theme"] = cty.StringVal(p.HostnameTheme)
}

func EncodeOpsworksStack_CustomCookbooksSource(p CustomCookbooksSource, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksStack_CustomCookbooksSource_Url(p, ctyVal)
	EncodeOpsworksStack_CustomCookbooksSource_Username(p, ctyVal)
	EncodeOpsworksStack_CustomCookbooksSource_Password(p, ctyVal)
	EncodeOpsworksStack_CustomCookbooksSource_Revision(p, ctyVal)
	EncodeOpsworksStack_CustomCookbooksSource_SshKey(p, ctyVal)
	EncodeOpsworksStack_CustomCookbooksSource_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["custom_cookbooks_source"] = cty.ListVal(valsForCollection)
}

func EncodeOpsworksStack_CustomCookbooksSource_Url(p CustomCookbooksSource, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}

func EncodeOpsworksStack_CustomCookbooksSource_Username(p CustomCookbooksSource, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeOpsworksStack_CustomCookbooksSource_Password(p CustomCookbooksSource, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeOpsworksStack_CustomCookbooksSource_Revision(p CustomCookbooksSource, vals map[string]cty.Value) {
	vals["revision"] = cty.StringVal(p.Revision)
}

func EncodeOpsworksStack_CustomCookbooksSource_SshKey(p CustomCookbooksSource, vals map[string]cty.Value) {
	vals["ssh_key"] = cty.StringVal(p.SshKey)
}

func EncodeOpsworksStack_CustomCookbooksSource_Type(p CustomCookbooksSource, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksStack_StackEndpoint(p OpsworksStackObservation, vals map[string]cty.Value) {
	vals["stack_endpoint"] = cty.StringVal(p.StackEndpoint)
}

func EncodeOpsworksStack_Arn(p OpsworksStackObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}