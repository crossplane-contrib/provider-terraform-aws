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

package v1alpha1func EncodeConfigConfigurationAggregator(r ConfigConfigurationAggregator) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeConfigConfigurationAggregator_Id(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigurationAggregator_Name(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigurationAggregator_Tags(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigurationAggregator_AccountAggregationSource(r.Spec.ForProvider.AccountAggregationSource, ctyVal)
	EncodeConfigConfigurationAggregator_OrganizationAggregationSource(r.Spec.ForProvider.OrganizationAggregationSource, ctyVal)
	EncodeConfigConfigurationAggregator_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeConfigConfigurationAggregator_Id(p *ConfigConfigurationAggregatorParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeConfigConfigurationAggregator_Name(p *ConfigConfigurationAggregatorParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeConfigConfigurationAggregator_Tags(p *ConfigConfigurationAggregatorParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeConfigConfigurationAggregator_AccountAggregationSource(p *AccountAggregationSource, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AccountAggregationSource {
		ctyVal = make(map[string]cty.Value)
		EncodeConfigConfigurationAggregator_AccountAggregationSource_AccountIds(v, ctyVal)
		EncodeConfigConfigurationAggregator_AccountAggregationSource_AllRegions(v, ctyVal)
		EncodeConfigConfigurationAggregator_AccountAggregationSource_Regions(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["account_aggregation_source"] = cty.ListVal(valsForCollection)
}

func EncodeConfigConfigurationAggregator_AccountAggregationSource_AccountIds(p *AccountAggregationSource, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AccountIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["account_ids"] = cty.ListVal(colVals)
}

func EncodeConfigConfigurationAggregator_AccountAggregationSource_AllRegions(p *AccountAggregationSource, vals map[string]cty.Value) {
	vals["all_regions"] = cty.BoolVal(p.AllRegions)
}

func EncodeConfigConfigurationAggregator_AccountAggregationSource_Regions(p *AccountAggregationSource, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Regions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["regions"] = cty.ListVal(colVals)
}

func EncodeConfigConfigurationAggregator_OrganizationAggregationSource(p *OrganizationAggregationSource, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OrganizationAggregationSource {
		ctyVal = make(map[string]cty.Value)
		EncodeConfigConfigurationAggregator_OrganizationAggregationSource_AllRegions(v, ctyVal)
		EncodeConfigConfigurationAggregator_OrganizationAggregationSource_Regions(v, ctyVal)
		EncodeConfigConfigurationAggregator_OrganizationAggregationSource_RoleArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["organization_aggregation_source"] = cty.ListVal(valsForCollection)
}

func EncodeConfigConfigurationAggregator_OrganizationAggregationSource_AllRegions(p *OrganizationAggregationSource, vals map[string]cty.Value) {
	vals["all_regions"] = cty.BoolVal(p.AllRegions)
}

func EncodeConfigConfigurationAggregator_OrganizationAggregationSource_Regions(p *OrganizationAggregationSource, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Regions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["regions"] = cty.ListVal(colVals)
}

func EncodeConfigConfigurationAggregator_OrganizationAggregationSource_RoleArn(p *OrganizationAggregationSource, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeConfigConfigurationAggregator_Arn(p *ConfigConfigurationAggregatorObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}