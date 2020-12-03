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

package v1alpha1func EncodeLicensemanagerAssociation(r LicensemanagerAssociation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeLicensemanagerAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerAssociation_LicenseConfigurationArn(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerAssociation_ResourceArn(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeLicensemanagerAssociation_Id(p *LicensemanagerAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLicensemanagerAssociation_LicenseConfigurationArn(p *LicensemanagerAssociationParameters, vals map[string]cty.Value) {
	vals["license_configuration_arn"] = cty.StringVal(p.LicenseConfigurationArn)
}

func EncodeLicensemanagerAssociation_ResourceArn(p *LicensemanagerAssociationParameters, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}