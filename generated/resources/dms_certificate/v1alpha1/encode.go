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

func EncodeDmsCertificate(r DmsCertificate) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDmsCertificate_CertificateWallet(r.Spec.ForProvider, ctyVal)
	EncodeDmsCertificate_Id(r.Spec.ForProvider, ctyVal)
	EncodeDmsCertificate_CertificateId(r.Spec.ForProvider, ctyVal)
	EncodeDmsCertificate_CertificatePem(r.Spec.ForProvider, ctyVal)
	EncodeDmsCertificate_CertificateArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDmsCertificate_CertificateWallet(p DmsCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_wallet"] = cty.StringVal(p.CertificateWallet)
}

func EncodeDmsCertificate_Id(p DmsCertificateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDmsCertificate_CertificateId(p DmsCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_id"] = cty.StringVal(p.CertificateId)
}

func EncodeDmsCertificate_CertificatePem(p DmsCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_pem"] = cty.StringVal(p.CertificatePem)
}

func EncodeDmsCertificate_CertificateArn(p DmsCertificateObservation, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}