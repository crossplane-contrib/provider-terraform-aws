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
	r, ok := mr.(*SesDomainIdentityVerification)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSesDomainIdentityVerification(r, ctyValue)
}

func DecodeSesDomainIdentityVerification(prev *SesDomainIdentityVerification, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSesDomainIdentityVerification_Domain(&new.Spec.ForProvider, valMap)
	DecodeSesDomainIdentityVerification_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeSesDomainIdentityVerification_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSesDomainIdentityVerification_Domain(p *SesDomainIdentityVerificationParameters, vals map[string]cty.Value) {
	p.Domain = ctwhy.ValueAsString(vals["domain"])
}

//containerTypeDecodeTemplate
func DecodeSesDomainIdentityVerification_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeSesDomainIdentityVerification_Timeouts_Create(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeSesDomainIdentityVerification_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeSesDomainIdentityVerification_Arn(p *SesDomainIdentityVerificationObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}