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
	r, ok := mr.(*PinpointApnsVoipSandboxChannel)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodePinpointApnsVoipSandboxChannel(r, ctyValue)
}

func DecodePinpointApnsVoipSandboxChannel(prev *PinpointApnsVoipSandboxChannel, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodePinpointApnsVoipSandboxChannel_BundleId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipSandboxChannel_Id(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipSandboxChannel_TeamId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipSandboxChannel_TokenKeyId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipSandboxChannel_ApplicationId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipSandboxChannel_Certificate(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipSandboxChannel_DefaultAuthenticationMethod(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipSandboxChannel_Enabled(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipSandboxChannel_PrivateKey(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsVoipSandboxChannel_TokenKey(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_BundleId(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.BundleId = ctwhy.ValueAsString(vals["bundle_id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_Id(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_TeamId(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.TeamId = ctwhy.ValueAsString(vals["team_id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_TokenKeyId(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.TokenKeyId = ctwhy.ValueAsString(vals["token_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_ApplicationId(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.ApplicationId = ctwhy.ValueAsString(vals["application_id"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_Certificate(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.Certificate = ctwhy.ValueAsString(vals["certificate"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_DefaultAuthenticationMethod(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.DefaultAuthenticationMethod = ctwhy.ValueAsString(vals["default_authentication_method"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_Enabled(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_PrivateKey(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.PrivateKey = ctwhy.ValueAsString(vals["private_key"])
}

//primitiveTypeDecodeTemplate
func DecodePinpointApnsVoipSandboxChannel_TokenKey(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	p.TokenKey = ctwhy.ValueAsString(vals["token_key"])
}