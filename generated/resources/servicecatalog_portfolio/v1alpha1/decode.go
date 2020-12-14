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
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*ServicecatalogPortfolio)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeServicecatalogPortfolio(r, ctyValue)
}

func DecodeServicecatalogPortfolio(prev *ServicecatalogPortfolio, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeServicecatalogPortfolio_Tags(&new.Spec.ForProvider, valMap)
	DecodeServicecatalogPortfolio_Description(&new.Spec.ForProvider, valMap)
	DecodeServicecatalogPortfolio_Id(&new.Spec.ForProvider, valMap)
	DecodeServicecatalogPortfolio_Name(&new.Spec.ForProvider, valMap)
	DecodeServicecatalogPortfolio_ProviderName(&new.Spec.ForProvider, valMap)
	DecodeServicecatalogPortfolio_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeServicecatalogPortfolio_Arn(&new.Status.AtProvider, valMap)
	DecodeServicecatalogPortfolio_CreatedTime(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeServicecatalogPortfolio_Tags(p *ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeServicecatalogPortfolio_Description(p *ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeServicecatalogPortfolio_Id(p *ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeServicecatalogPortfolio_Name(p *ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeServicecatalogPortfolio_ProviderName(p *ServicecatalogPortfolioParameters, vals map[string]cty.Value) {
	p.ProviderName = ctwhy.ValueAsString(vals["provider_name"])
}

func DecodeServicecatalogPortfolio_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeServicecatalogPortfolio_Timeouts_Create(p, valMap)
	DecodeServicecatalogPortfolio_Timeouts_Delete(p, valMap)
	DecodeServicecatalogPortfolio_Timeouts_Update(p, valMap)
}

func DecodeServicecatalogPortfolio_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeServicecatalogPortfolio_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeServicecatalogPortfolio_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeServicecatalogPortfolio_Arn(p *ServicecatalogPortfolioObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeServicecatalogPortfolio_CreatedTime(p *ServicecatalogPortfolioObservation, vals map[string]cty.Value) {
	p.CreatedTime = ctwhy.ValueAsString(vals["created_time"])
}