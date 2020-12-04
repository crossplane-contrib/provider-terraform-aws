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

func EncodeServicecatalogPortfolio(r ServicecatalogPortfolio) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeServicecatalogPortfolio_Name(r.Spec.ForProvider, ctyVal)
	EncodeServicecatalogPortfolio_ProviderName(r.Spec.ForProvider, ctyVal)
	EncodeServicecatalogPortfolio_Tags(r.Spec.ForProvider, ctyVal)
	EncodeServicecatalogPortfolio_Description(r.Spec.ForProvider, ctyVal)
	EncodeServicecatalogPortfolio_Id(r.Spec.ForProvider, ctyVal)
	EncodeServicecatalogPortfolio_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeServicecatalogPortfolio_Arn(r.Status.AtProvider, ctyVal)
	EncodeServicecatalogPortfolio_CreatedTime(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeServicecatalogPortfolio_Name(p ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeServicecatalogPortfolio_ProviderName(p ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	vals["provider_name"] = cty.StringVal(p.ProviderName)
}

func EncodeServicecatalogPortfolio_Tags(p ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeServicecatalogPortfolio_Description(p ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeServicecatalogPortfolio_Id(p ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeServicecatalogPortfolio_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeServicecatalogPortfolio_Timeouts_Create(p, ctyVal)
	EncodeServicecatalogPortfolio_Timeouts_Delete(p, ctyVal)
	EncodeServicecatalogPortfolio_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeServicecatalogPortfolio_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeServicecatalogPortfolio_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeServicecatalogPortfolio_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeServicecatalogPortfolio_Arn(p ServicecatalogPortfolioObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeServicecatalogPortfolio_CreatedTime(p ServicecatalogPortfolioObservation, vals map[string]cty.Value) {
	vals["created_time"] = cty.StringVal(p.CreatedTime)
}