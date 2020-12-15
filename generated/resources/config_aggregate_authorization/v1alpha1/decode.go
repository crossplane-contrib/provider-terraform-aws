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
	r, ok := mr.(*ConfigAggregateAuthorization)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeConfigAggregateAuthorization(r, ctyValue)
}

func DecodeConfigAggregateAuthorization(prev *ConfigAggregateAuthorization, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeConfigAggregateAuthorization_AccountId(&new.Spec.ForProvider, valMap)
	DecodeConfigAggregateAuthorization_Region(&new.Spec.ForProvider, valMap)
	DecodeConfigAggregateAuthorization_Tags(&new.Spec.ForProvider, valMap)
	DecodeConfigAggregateAuthorization_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeConfigAggregateAuthorization_AccountId(p *ConfigAggregateAuthorizationParameters, vals map[string]cty.Value) {
	p.AccountId = ctwhy.ValueAsString(vals["account_id"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigAggregateAuthorization_Region(p *ConfigAggregateAuthorizationParameters, vals map[string]cty.Value) {
	p.Region = ctwhy.ValueAsString(vals["region"])
}

//primitiveMapTypeDecodeTemplate
func DecodeConfigAggregateAuthorization_Tags(p *ConfigAggregateAuthorizationParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeConfigAggregateAuthorization_Arn(p *ConfigAggregateAuthorizationObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}