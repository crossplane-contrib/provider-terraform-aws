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

package v1alpha1func EncodeSesDomainDkim(r SesDomainDkim) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSesDomainDkim_Domain(r.Spec.ForProvider, ctyVal)
	EncodeSesDomainDkim_Id(r.Spec.ForProvider, ctyVal)
	EncodeSesDomainDkim_DkimTokens(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSesDomainDkim_Domain(p *SesDomainDkimParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeSesDomainDkim_Id(p *SesDomainDkimParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSesDomainDkim_DkimTokens(p *SesDomainDkimObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DkimTokens {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["dkim_tokens"] = cty.ListVal(colVals)
}