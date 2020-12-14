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
	r, ok := mr.(*PinpointApnsChannel)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodePinpointApnsChannel(r, ctyValue)
}

func DecodePinpointApnsChannel(prev *PinpointApnsChannel, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodePinpointApnsChannel_ApplicationId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsChannel_BundleId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsChannel_DefaultAuthenticationMethod(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsChannel_Enabled(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsChannel_PrivateKey(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsChannel_Certificate(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsChannel_Id(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsChannel_TeamId(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsChannel_TokenKey(&new.Spec.ForProvider, valMap)
	DecodePinpointApnsChannel_TokenKeyId(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodePinpointApnsChannel_ApplicationId(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.ApplicationId = ctwhy.ValueAsString(vals["application_id"])
}

func DecodePinpointApnsChannel_BundleId(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.BundleId = ctwhy.ValueAsString(vals["bundle_id"])
}

func DecodePinpointApnsChannel_DefaultAuthenticationMethod(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.DefaultAuthenticationMethod = ctwhy.ValueAsString(vals["default_authentication_method"])
}

func DecodePinpointApnsChannel_Enabled(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

func DecodePinpointApnsChannel_PrivateKey(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.PrivateKey = ctwhy.ValueAsString(vals["private_key"])
}

func DecodePinpointApnsChannel_Certificate(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.Certificate = ctwhy.ValueAsString(vals["certificate"])
}

func DecodePinpointApnsChannel_Id(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodePinpointApnsChannel_TeamId(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.TeamId = ctwhy.ValueAsString(vals["team_id"])
}

func DecodePinpointApnsChannel_TokenKey(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.TokenKey = ctwhy.ValueAsString(vals["token_key"])
}

func DecodePinpointApnsChannel_TokenKeyId(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	p.TokenKeyId = ctwhy.ValueAsString(vals["token_key_id"])
}