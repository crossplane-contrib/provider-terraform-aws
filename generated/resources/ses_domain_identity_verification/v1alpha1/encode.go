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

package v1alpha1func EncodeSesDomainIdentityVerification(r SesDomainIdentityVerification) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSesDomainIdentityVerification_Domain(r.Spec.ForProvider, ctyVal)
	EncodeSesDomainIdentityVerification_Id(r.Spec.ForProvider, ctyVal)
	EncodeSesDomainIdentityVerification_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeSesDomainIdentityVerification_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSesDomainIdentityVerification_Domain(p *SesDomainIdentityVerificationParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeSesDomainIdentityVerification_Id(p *SesDomainIdentityVerificationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSesDomainIdentityVerification_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeSesDomainIdentityVerification_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeSesDomainIdentityVerification_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeSesDomainIdentityVerification_Arn(p *SesDomainIdentityVerificationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}