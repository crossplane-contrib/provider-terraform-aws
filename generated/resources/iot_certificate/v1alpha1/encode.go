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
	r, ok := mr.(*IotCertificate)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IotCertificate.")
	}
	return EncodeIotCertificate(*r), nil
}

func EncodeIotCertificate(r IotCertificate) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIotCertificate_Active(r.Spec.ForProvider, ctyVal)
	EncodeIotCertificate_Csr(r.Spec.ForProvider, ctyVal)
	EncodeIotCertificate_PrivateKey(r.Status.AtProvider, ctyVal)
	EncodeIotCertificate_PublicKey(r.Status.AtProvider, ctyVal)
	EncodeIotCertificate_Arn(r.Status.AtProvider, ctyVal)
	EncodeIotCertificate_CertificatePem(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeIotCertificate_Active(p IotCertificateParameters, vals map[string]cty.Value) {
	vals["active"] = cty.BoolVal(p.Active)
}

func EncodeIotCertificate_Csr(p IotCertificateParameters, vals map[string]cty.Value) {
	vals["csr"] = cty.StringVal(p.Csr)
}

func EncodeIotCertificate_PrivateKey(p IotCertificateObservation, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodeIotCertificate_PublicKey(p IotCertificateObservation, vals map[string]cty.Value) {
	vals["public_key"] = cty.StringVal(p.PublicKey)
}

func EncodeIotCertificate_Arn(p IotCertificateObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeIotCertificate_CertificatePem(p IotCertificateObservation, vals map[string]cty.Value) {
	vals["certificate_pem"] = cty.StringVal(p.CertificatePem)
}