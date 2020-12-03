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

package v1alpha1func EncodeAlbListenerCertificate(r AlbListenerCertificate) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAlbListenerCertificate_CertificateArn(r.Spec.ForProvider, ctyVal)
	EncodeAlbListenerCertificate_Id(r.Spec.ForProvider, ctyVal)
	EncodeAlbListenerCertificate_ListenerArn(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeAlbListenerCertificate_CertificateArn(p *AlbListenerCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}

func EncodeAlbListenerCertificate_Id(p *AlbListenerCertificateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAlbListenerCertificate_ListenerArn(p *AlbListenerCertificateParameters, vals map[string]cty.Value) {
	vals["listener_arn"] = cty.StringVal(p.ListenerArn)
}