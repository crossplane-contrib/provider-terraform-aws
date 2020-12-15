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
	r, ok := mr.(*PinpointApnsVoipChannel)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodePinpointApnsVoipChannel(r, ctyValue)
}

func DecodePinpointApnsVoipChannel(prev *PinpointApnsVoipChannel, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodePinpointApnsVoipChannel_TokenKeyId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipChannel_ApplicationId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipChannel_BundleId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipChannel_Certificate(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipChannel_DefaultAuthenticationMethod(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipChannel_Enabled(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipChannel_PrivateKey(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipChannel_TeamId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipChannel_Id(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipChannel_TokenKey(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_TokenKeyId(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.TokenKeyId = ctwhy.ValueAsString(vals["token_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_ApplicationId(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.ApplicationId = ctwhy.ValueAsString(vals["application_id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_BundleId(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.BundleId = ctwhy.ValueAsString(vals["bundle_id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_Certificate(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.Certificate = ctwhy.ValueAsString(vals["certificate"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_DefaultAuthenticationMethod(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.DefaultAuthenticationMethod = ctwhy.ValueAsString(vals["default_authentication_method"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_Enabled(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_PrivateKey(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.PrivateKey = ctwhy.ValueAsString(vals["private_key"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_TeamId(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.TeamId = ctwhy.ValueAsString(vals["team_id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_Id(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipChannel_TokenKey(p *PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	p.TokenKey = ctwhy.ValueAsString(vals["token_key"])
}