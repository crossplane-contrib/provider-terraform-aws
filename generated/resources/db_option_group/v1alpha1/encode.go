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

func EncodeDbOptionGroup(r DbOptionGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbOptionGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeDbOptionGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeDbOptionGroup_OptionGroupDescription(r.Spec.ForProvider, ctyVal)
	EncodeDbOptionGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDbOptionGroup_EngineName(r.Spec.ForProvider, ctyVal)
	EncodeDbOptionGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbOptionGroup_MajorEngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeDbOptionGroup_Option(r.Spec.ForProvider.Option, ctyVal)
	EncodeDbOptionGroup_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDbOptionGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDbOptionGroup_Name(p DbOptionGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbOptionGroup_NamePrefix(p DbOptionGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeDbOptionGroup_OptionGroupDescription(p DbOptionGroupParameters, vals map[string]cty.Value) {
	vals["option_group_description"] = cty.StringVal(p.OptionGroupDescription)
}

func EncodeDbOptionGroup_Tags(p DbOptionGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDbOptionGroup_EngineName(p DbOptionGroupParameters, vals map[string]cty.Value) {
	vals["engine_name"] = cty.StringVal(p.EngineName)
}

func EncodeDbOptionGroup_Id(p DbOptionGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbOptionGroup_MajorEngineVersion(p DbOptionGroupParameters, vals map[string]cty.Value) {
	vals["major_engine_version"] = cty.StringVal(p.MajorEngineVersion)
}

func EncodeDbOptionGroup_Option(p Option, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDbOptionGroup_Option_Port(p, ctyVal)
	EncodeDbOptionGroup_Option_Version(p, ctyVal)
	EncodeDbOptionGroup_Option_VpcSecurityGroupMemberships(p, ctyVal)
	EncodeDbOptionGroup_Option_DbSecurityGroupMemberships(p, ctyVal)
	EncodeDbOptionGroup_Option_OptionName(p, ctyVal)
	EncodeDbOptionGroup_Option_OptionSettings(p.OptionSettings, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["option"] = cty.SetVal(valsForCollection)
}

func EncodeDbOptionGroup_Option_Port(p Option, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDbOptionGroup_Option_Version(p Option, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeDbOptionGroup_Option_VpcSecurityGroupMemberships(p Option, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupMemberships {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_memberships"] = cty.SetVal(colVals)
}

func EncodeDbOptionGroup_Option_DbSecurityGroupMemberships(p Option, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DbSecurityGroupMemberships {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["db_security_group_memberships"] = cty.SetVal(colVals)
}

func EncodeDbOptionGroup_Option_OptionName(p Option, vals map[string]cty.Value) {
	vals["option_name"] = cty.StringVal(p.OptionName)
}

func EncodeDbOptionGroup_Option_OptionSettings(p OptionSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDbOptionGroup_Option_OptionSettings_Name(p, ctyVal)
	EncodeDbOptionGroup_Option_OptionSettings_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["option_settings"] = cty.SetVal(valsForCollection)
}

func EncodeDbOptionGroup_Option_OptionSettings_Name(p OptionSettings, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbOptionGroup_Option_OptionSettings_Value(p OptionSettings, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeDbOptionGroup_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDbOptionGroup_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDbOptionGroup_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDbOptionGroup_Arn(p DbOptionGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}