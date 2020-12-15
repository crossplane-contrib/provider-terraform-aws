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
	r, ok := mr.(*Ec2TrafficMirrorSession)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2TrafficMirrorSession(r, ctyValue)
}

func DecodeEc2TrafficMirrorSession(prev *Ec2TrafficMirrorSession, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2TrafficMirrorSession_NetworkInterfaceId(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorSession_SessionNumber(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorSession_TrafficMirrorFilterId(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorSession_Id(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorSession_Description(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorSession_PacketLength(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorSession_Tags(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorSession_TrafficMirrorTargetId(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorSession_VirtualNetworkId(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorSession_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_NetworkInterfaceId(p *Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	p.NetworkInterfaceId = ctwhy.ValueAsString(vals["network_interface_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_SessionNumber(p *Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	p.SessionNumber = ctwhy.ValueAsInt64(vals["session_number"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_TrafficMirrorFilterId(p *Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	p.TrafficMirrorFilterId = ctwhy.ValueAsString(vals["traffic_mirror_filter_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_Id(p *Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_Description(p *Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_PacketLength(p *Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	p.PacketLength = ctwhy.ValueAsInt64(vals["packet_length"])
}

//primitiveMapTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_Tags(p *Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_TrafficMirrorTargetId(p *Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	p.TrafficMirrorTargetId = ctwhy.ValueAsString(vals["traffic_mirror_target_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_VirtualNetworkId(p *Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	p.VirtualNetworkId = ctwhy.ValueAsInt64(vals["virtual_network_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TrafficMirrorSession_Arn(p *Ec2TrafficMirrorSessionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}