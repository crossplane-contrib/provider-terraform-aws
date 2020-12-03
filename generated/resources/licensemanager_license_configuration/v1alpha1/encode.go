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

package v1alpha1func EncodeLicensemanagerLicenseConfiguration(r LicensemanagerLicenseConfiguration) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeLicensemanagerLicenseConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_Tags(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_Description(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_LicenseCount(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_LicenseCountHardLimit(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_LicenseCountingType(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_LicenseRules(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeLicensemanagerLicenseConfiguration_Name(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLicensemanagerLicenseConfiguration_Tags(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeLicensemanagerLicenseConfiguration_Description(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLicensemanagerLicenseConfiguration_Id(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLicensemanagerLicenseConfiguration_LicenseCount(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["license_count"] = cty.IntVal(p.LicenseCount)
}

func EncodeLicensemanagerLicenseConfiguration_LicenseCountHardLimit(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["license_count_hard_limit"] = cty.BoolVal(p.LicenseCountHardLimit)
}

func EncodeLicensemanagerLicenseConfiguration_LicenseCountingType(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["license_counting_type"] = cty.StringVal(p.LicenseCountingType)
}

func EncodeLicensemanagerLicenseConfiguration_LicenseRules(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LicenseRules {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["license_rules"] = cty.ListVal(colVals)
}