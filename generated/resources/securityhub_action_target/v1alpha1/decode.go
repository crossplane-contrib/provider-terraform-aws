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
	r, ok := mr.(*SecurityhubActionTarget)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSecurityhubActionTarget(r, ctyValue)
}

func DecodeSecurityhubActionTarget(prev *SecurityhubActionTarget, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSecurityhubActionTarget_Description(&new.Spec.ForProvider, valMap)
	DecodeSecurityhubActionTarget_Id(&new.Spec.ForProvider, valMap)
	DecodeSecurityhubActionTarget_Identifier(&new.Spec.ForProvider, valMap)
	DecodeSecurityhubActionTarget_Name(&new.Spec.ForProvider, valMap)
	DecodeSecurityhubActionTarget_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubActionTarget_Description(p *SecurityhubActionTargetParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubActionTarget_Id(p *SecurityhubActionTargetParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubActionTarget_Identifier(p *SecurityhubActionTargetParameters, vals map[string]cty.Value) {
	p.Identifier = ctwhy.ValueAsString(vals["identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubActionTarget_Name(p *SecurityhubActionTargetParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubActionTarget_Arn(p *SecurityhubActionTargetObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}