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

package v1alpha1func EncodeLbListenerCertificate(r LbListenerCertificate) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeLbListenerCertificate_CertificateArn(r.Spec.ForProvider, ctyVal)
	EncodeLbListenerCertificate_Id(r.Spec.ForProvider, ctyVal)
	EncodeLbListenerCertificate_ListenerArn(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeLbListenerCertificate_CertificateArn(p *LbListenerCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}

func EncodeLbListenerCertificate_Id(p *LbListenerCertificateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLbListenerCertificate_ListenerArn(p *LbListenerCertificateParameters, vals map[string]cty.Value) {
	vals["listener_arn"] = cty.StringVal(p.ListenerArn)
}