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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*IamServerCertificate)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamServerCertificate(r, ctyValue)
}

func DecodeIamServerCertificate(prev *IamServerCertificate, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamServerCertificate_CertificateChain(&new.Spec.ForProvider, valMap)
	DecodeIamServerCertificate_Id(&new.Spec.ForProvider, valMap)
	DecodeIamServerCertificate_Name(&new.Spec.ForProvider, valMap)
	DecodeIamServerCertificate_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeIamServerCertificate_Path(&new.Spec.ForProvider, valMap)
	DecodeIamServerCertificate_PrivateKey(&new.Spec.ForProvider, valMap)
	DecodeIamServerCertificate_Arn(&new.Spec.ForProvider, valMap)
	DecodeIamServerCertificate_CertificateBody(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeIamServerCertificate_CertificateChain(p *IamServerCertificateParameters, vals map[string]cty.Value) {
	p.CertificateChain = ctwhy.ValueAsString(vals["certificate_chain"])
}

func DecodeIamServerCertificate_Id(p *IamServerCertificateParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeIamServerCertificate_Name(p *IamServerCertificateParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeIamServerCertificate_NamePrefix(p *IamServerCertificateParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

func DecodeIamServerCertificate_Path(p *IamServerCertificateParameters, vals map[string]cty.Value) {
	p.Path = ctwhy.ValueAsString(vals["path"])
}

func DecodeIamServerCertificate_PrivateKey(p *IamServerCertificateParameters, vals map[string]cty.Value) {
	p.PrivateKey = ctwhy.ValueAsString(vals["private_key"])
}

func DecodeIamServerCertificate_Arn(p *IamServerCertificateParameters, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeIamServerCertificate_CertificateBody(p *IamServerCertificateParameters, vals map[string]cty.Value) {
	p.CertificateBody = ctwhy.ValueAsString(vals["certificate_body"])
}