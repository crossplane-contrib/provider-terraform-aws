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
	r, ok := mr.(*LbListenerCertificate)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLbListenerCertificate(r, ctyValue)
}

func DecodeLbListenerCertificate(prev *LbListenerCertificate, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLbListenerCertificate_Id(&new.Spec.ForProvider, valMap)
	DecodeLbListenerCertificate_ListenerArn(&new.Spec.ForProvider, valMap)
	DecodeLbListenerCertificate_CertificateArn(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeLbListenerCertificate_Id(p *LbListenerCertificateParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeLbListenerCertificate_ListenerArn(p *LbListenerCertificateParameters, vals map[string]cty.Value) {
	p.ListenerArn = ctwhy.ValueAsString(vals["listener_arn"])
}

func DecodeLbListenerCertificate_CertificateArn(p *LbListenerCertificateParameters, vals map[string]cty.Value) {
	p.CertificateArn = ctwhy.ValueAsString(vals["certificate_arn"])
}