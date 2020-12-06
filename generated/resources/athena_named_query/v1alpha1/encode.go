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

func EncodeAthenaNamedQuery(r AthenaNamedQuery) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAthenaNamedQuery_Id(r.Spec.ForProvider, ctyVal)
	EncodeAthenaNamedQuery_Name(r.Spec.ForProvider, ctyVal)
	EncodeAthenaNamedQuery_Query(r.Spec.ForProvider, ctyVal)
	EncodeAthenaNamedQuery_Workgroup(r.Spec.ForProvider, ctyVal)
	EncodeAthenaNamedQuery_Database(r.Spec.ForProvider, ctyVal)
	EncodeAthenaNamedQuery_Description(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeAthenaNamedQuery_Id(p AthenaNamedQueryParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAthenaNamedQuery_Name(p AthenaNamedQueryParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAthenaNamedQuery_Query(p AthenaNamedQueryParameters, vals map[string]cty.Value) {
	vals["query"] = cty.StringVal(p.Query)
}

func EncodeAthenaNamedQuery_Workgroup(p AthenaNamedQueryParameters, vals map[string]cty.Value) {
	vals["workgroup"] = cty.StringVal(p.Workgroup)
}

func EncodeAthenaNamedQuery_Database(p AthenaNamedQueryParameters, vals map[string]cty.Value) {
	vals["database"] = cty.StringVal(p.Database)
}

func EncodeAthenaNamedQuery_Description(p AthenaNamedQueryParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}