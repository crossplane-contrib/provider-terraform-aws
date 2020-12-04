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

func EncodeIamServerCertificate(r IamServerCertificate) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamServerCertificate_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_Path(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_PrivateKey(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_Arn(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_CertificateBody(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_CertificateChain(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_Name(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeIamServerCertificate_NamePrefix(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeIamServerCertificate_Path(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeIamServerCertificate_PrivateKey(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodeIamServerCertificate_Arn(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeIamServerCertificate_CertificateBody(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_body"] = cty.StringVal(p.CertificateBody)
}

func EncodeIamServerCertificate_CertificateChain(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_chain"] = cty.StringVal(p.CertificateChain)
}

func EncodeIamServerCertificate_Id(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamServerCertificate_Name(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}