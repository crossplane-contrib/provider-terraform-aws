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

package v1alpha1func EncodeIotCertificate(r IotCertificate) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeIotCertificate_Active(r.Spec.ForProvider, ctyVal)
	EncodeIotCertificate_Csr(r.Spec.ForProvider, ctyVal)
	EncodeIotCertificate_Id(r.Spec.ForProvider, ctyVal)
	EncodeIotCertificate_Arn(r.Status.AtProvider, ctyVal)
	EncodeIotCertificate_CertificatePem(r.Status.AtProvider, ctyVal)
	EncodeIotCertificate_PrivateKey(r.Status.AtProvider, ctyVal)
	EncodeIotCertificate_PublicKey(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeIotCertificate_Active(p *IotCertificateParameters, vals map[string]cty.Value) {
	vals["active"] = cty.BoolVal(p.Active)
}

func EncodeIotCertificate_Csr(p *IotCertificateParameters, vals map[string]cty.Value) {
	vals["csr"] = cty.StringVal(p.Csr)
}

func EncodeIotCertificate_Id(p *IotCertificateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIotCertificate_Arn(p *IotCertificateObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeIotCertificate_CertificatePem(p *IotCertificateObservation, vals map[string]cty.Value) {
	vals["certificate_pem"] = cty.StringVal(p.CertificatePem)
}

func EncodeIotCertificate_PrivateKey(p *IotCertificateObservation, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodeIotCertificate_PublicKey(p *IotCertificateObservation, vals map[string]cty.Value) {
	vals["public_key"] = cty.StringVal(p.PublicKey)
}