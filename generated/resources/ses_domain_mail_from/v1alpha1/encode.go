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

package v1alpha1func EncodeSesDomainMailFrom(r SesDomainMailFrom) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSesDomainMailFrom_MailFromDomain(r.Spec.ForProvider, ctyVal)
	EncodeSesDomainMailFrom_BehaviorOnMxFailure(r.Spec.ForProvider, ctyVal)
	EncodeSesDomainMailFrom_Domain(r.Spec.ForProvider, ctyVal)
	EncodeSesDomainMailFrom_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeSesDomainMailFrom_MailFromDomain(p *SesDomainMailFromParameters, vals map[string]cty.Value) {
	vals["mail_from_domain"] = cty.StringVal(p.MailFromDomain)
}

func EncodeSesDomainMailFrom_BehaviorOnMxFailure(p *SesDomainMailFromParameters, vals map[string]cty.Value) {
	vals["behavior_on_mx_failure"] = cty.StringVal(p.BehaviorOnMxFailure)
}

func EncodeSesDomainMailFrom_Domain(p *SesDomainMailFromParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeSesDomainMailFrom_Id(p *SesDomainMailFromParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}