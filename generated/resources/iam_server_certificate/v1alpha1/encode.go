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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*IamServerCertificate)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IamServerCertificate.")
	}
	return EncodeIamServerCertificate(*r), nil
}

func EncodeIamServerCertificate(r IamServerCertificate) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamServerCertificate_PrivateKey(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_Arn(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_CertificateBody(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_CertificateChain(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeIamServerCertificate_Path(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
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

func EncodeIamServerCertificate_Name(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamServerCertificate_NamePrefix(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeIamServerCertificate_Path(p IamServerCertificateParameters, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}