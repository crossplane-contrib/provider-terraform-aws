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
	r, ok := mr.(*PinpointEmailChannel)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodePinpointEmailChannel(r, ctyValue)
}

func DecodePinpointEmailChannel(prev *PinpointEmailChannel, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodePinpointEmailChannel_FromAddress(&new.Spec.ForProvider, valMap)
	DecodePinpointEmailChannel_Id(&new.Spec.ForProvider, valMap)
	DecodePinpointEmailChannel_Identity(&new.Spec.ForProvider, valMap)
	DecodePinpointEmailChannel_RoleArn(&new.Spec.ForProvider, valMap)
	DecodePinpointEmailChannel_ApplicationId(&new.Spec.ForProvider, valMap)
	DecodePinpointEmailChannel_Enabled(&new.Spec.ForProvider, valMap)
	DecodePinpointEmailChannel_MessagesPerSecond(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodePinpointEmailChannel_FromAddress(p *PinpointEmailChannelParameters, vals map[string]cty.Value) {
	p.FromAddress = ctwhy.ValueAsString(vals["from_address"])
}

func DecodePinpointEmailChannel_Id(p *PinpointEmailChannelParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodePinpointEmailChannel_Identity(p *PinpointEmailChannelParameters, vals map[string]cty.Value) {
	p.Identity = ctwhy.ValueAsString(vals["identity"])
}

func DecodePinpointEmailChannel_RoleArn(p *PinpointEmailChannelParameters, vals map[string]cty.Value) {
	p.RoleArn = ctwhy.ValueAsString(vals["role_arn"])
}

func DecodePinpointEmailChannel_ApplicationId(p *PinpointEmailChannelParameters, vals map[string]cty.Value) {
	p.ApplicationId = ctwhy.ValueAsString(vals["application_id"])
}

func DecodePinpointEmailChannel_Enabled(p *PinpointEmailChannelParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

func DecodePinpointEmailChannel_MessagesPerSecond(p *PinpointEmailChannelObservation, vals map[string]cty.Value) {
	p.MessagesPerSecond = ctwhy.ValueAsInt64(vals["messages_per_second"])
}