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
	r, ok := mr.(*ApiGatewayClientCertificate)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApiGatewayClientCertificate(r, ctyValue)
}

func DecodeApiGatewayClientCertificate(prev *ApiGatewayClientCertificate, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApiGatewayClientCertificate_Description(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayClientCertificate_Tags(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayClientCertificate_CreatedDate(&new.Status.AtProvider, valMap)
	DecodeApiGatewayClientCertificate_ExpirationDate(&new.Status.AtProvider, valMap)
	DecodeApiGatewayClientCertificate_PemEncodedCertificate(&new.Status.AtProvider, valMap)
	DecodeApiGatewayClientCertificate_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayClientCertificate_Description(p *ApiGatewayClientCertificateParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveMapTypeDecodeTemplate
func DecodeApiGatewayClientCertificate_Tags(p *ApiGatewayClientCertificateParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayClientCertificate_CreatedDate(p *ApiGatewayClientCertificateObservation, vals map[string]cty.Value) {
	p.CreatedDate = ctwhy.ValueAsString(vals["created_date"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayClientCertificate_ExpirationDate(p *ApiGatewayClientCertificateObservation, vals map[string]cty.Value) {
	p.ExpirationDate = ctwhy.ValueAsString(vals["expiration_date"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayClientCertificate_PemEncodedCertificate(p *ApiGatewayClientCertificateObservation, vals map[string]cty.Value) {
	p.PemEncodedCertificate = ctwhy.ValueAsString(vals["pem_encoded_certificate"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayClientCertificate_Arn(p *ApiGatewayClientCertificateObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}