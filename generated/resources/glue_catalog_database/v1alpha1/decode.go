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
	r, ok := mr.(*GlueCatalogDatabase)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeGlueCatalogDatabase(r, ctyValue)
}

func DecodeGlueCatalogDatabase(prev *GlueCatalogDatabase, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeGlueCatalogDatabase_CatalogId(&new.Spec.ForProvider, valMap)
	DecodeGlueCatalogDatabase_Description(&new.Spec.ForProvider, valMap)
	DecodeGlueCatalogDatabase_Id(&new.Spec.ForProvider, valMap)
	DecodeGlueCatalogDatabase_LocationUri(&new.Spec.ForProvider, valMap)
	DecodeGlueCatalogDatabase_Name(&new.Spec.ForProvider, valMap)
	DecodeGlueCatalogDatabase_Parameters(&new.Spec.ForProvider, valMap)
	DecodeGlueCatalogDatabase_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeGlueCatalogDatabase_CatalogId(p *GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	p.CatalogId = ctwhy.ValueAsString(vals["catalog_id"])
}

//primitiveTypeDecodeTemplate
func DecodeGlueCatalogDatabase_Description(p *GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeGlueCatalogDatabase_Id(p *GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeGlueCatalogDatabase_LocationUri(p *GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	p.LocationUri = ctwhy.ValueAsString(vals["location_uri"])
}

//primitiveTypeDecodeTemplate
func DecodeGlueCatalogDatabase_Name(p *GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeGlueCatalogDatabase_Parameters(p *GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["parameters"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Parameters = vMap
}

//primitiveTypeDecodeTemplate
func DecodeGlueCatalogDatabase_Arn(p *GlueCatalogDatabaseObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}