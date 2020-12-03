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

package v1alpha1func EncodeSsmActivation(r SsmActivation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSsmActivation_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmActivation_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmActivation_RegistrationLimit(r.Spec.ForProvider, ctyVal)
	EncodeSsmActivation_Description(r.Spec.ForProvider, ctyVal)
	EncodeSsmActivation_ExpirationDate(r.Spec.ForProvider, ctyVal)
	EncodeSsmActivation_IamRole(r.Spec.ForProvider, ctyVal)
	EncodeSsmActivation_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSsmActivation_RegistrationCount(r.Status.AtProvider, ctyVal)
	EncodeSsmActivation_Expired(r.Status.AtProvider, ctyVal)
	EncodeSsmActivation_ActivationCode(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSsmActivation_Id(p *SsmActivationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmActivation_Name(p *SsmActivationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmActivation_RegistrationLimit(p *SsmActivationParameters, vals map[string]cty.Value) {
	vals["registration_limit"] = cty.IntVal(p.RegistrationLimit)
}

func EncodeSsmActivation_Description(p *SsmActivationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSsmActivation_ExpirationDate(p *SsmActivationParameters, vals map[string]cty.Value) {
	vals["expiration_date"] = cty.StringVal(p.ExpirationDate)
}

func EncodeSsmActivation_IamRole(p *SsmActivationParameters, vals map[string]cty.Value) {
	vals["iam_role"] = cty.StringVal(p.IamRole)
}

func EncodeSsmActivation_Tags(p *SsmActivationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSsmActivation_RegistrationCount(p *SsmActivationObservation, vals map[string]cty.Value) {
	vals["registration_count"] = cty.IntVal(p.RegistrationCount)
}

func EncodeSsmActivation_Expired(p *SsmActivationObservation, vals map[string]cty.Value) {
	vals["expired"] = cty.BoolVal(p.Expired)
}

func EncodeSsmActivation_ActivationCode(p *SsmActivationObservation, vals map[string]cty.Value) {
	vals["activation_code"] = cty.StringVal(p.ActivationCode)
}