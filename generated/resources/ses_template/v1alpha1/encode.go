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

func EncodeSesTemplate(r SesTemplate) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSesTemplate_Html(r.Spec.ForProvider, ctyVal)
	EncodeSesTemplate_Id(r.Spec.ForProvider, ctyVal)
	EncodeSesTemplate_Name(r.Spec.ForProvider, ctyVal)
	EncodeSesTemplate_Subject(r.Spec.ForProvider, ctyVal)
	EncodeSesTemplate_Text(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeSesTemplate_Html(p SesTemplateParameters, vals map[string]cty.Value) {
	vals["html"] = cty.StringVal(p.Html)
}

func EncodeSesTemplate_Id(p SesTemplateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSesTemplate_Name(p SesTemplateParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSesTemplate_Subject(p SesTemplateParameters, vals map[string]cty.Value) {
	vals["subject"] = cty.StringVal(p.Subject)
}

func EncodeSesTemplate_Text(p SesTemplateParameters, vals map[string]cty.Value) {
	vals["text"] = cty.StringVal(p.Text)
}