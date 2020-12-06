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

func EncodeGlueCatalogDatabase(r GlueCatalogDatabase) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCatalogDatabase_Description(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogDatabase_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogDatabase_LocationUri(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogDatabase_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogDatabase_Parameters(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogDatabase_CatalogId(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogDatabase_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlueCatalogDatabase_Description(p GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGlueCatalogDatabase_Id(p GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueCatalogDatabase_LocationUri(p GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	vals["location_uri"] = cty.StringVal(p.LocationUri)
}

func EncodeGlueCatalogDatabase_Name(p GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueCatalogDatabase_Parameters(p GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeGlueCatalogDatabase_CatalogId(p GlueCatalogDatabaseParameters, vals map[string]cty.Value) {
	vals["catalog_id"] = cty.StringVal(p.CatalogId)
}

func EncodeGlueCatalogDatabase_Arn(p GlueCatalogDatabaseObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}