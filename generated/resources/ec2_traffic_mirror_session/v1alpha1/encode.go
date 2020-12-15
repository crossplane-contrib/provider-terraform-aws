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

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*Ec2TrafficMirrorSession)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2TrafficMirrorSession.")
	}
	return EncodeEc2TrafficMirrorSession(*r), nil
}

func EncodeEc2TrafficMirrorSession(r Ec2TrafficMirrorSession) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TrafficMirrorSession_TrafficMirrorTargetId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorSession_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorSession_NetworkInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorSession_PacketLength(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorSession_SessionNumber(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorSession_TrafficMirrorFilterId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorSession_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorSession_VirtualNetworkId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorSession_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TrafficMirrorSession_TrafficMirrorTargetId(p Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	vals["traffic_mirror_target_id"] = cty.StringVal(p.TrafficMirrorTargetId)
}

func EncodeEc2TrafficMirrorSession_Description(p Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2TrafficMirrorSession_NetworkInterfaceId(p Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeEc2TrafficMirrorSession_PacketLength(p Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	vals["packet_length"] = cty.NumberIntVal(p.PacketLength)
}

func EncodeEc2TrafficMirrorSession_SessionNumber(p Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	vals["session_number"] = cty.NumberIntVal(p.SessionNumber)
}

func EncodeEc2TrafficMirrorSession_TrafficMirrorFilterId(p Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	vals["traffic_mirror_filter_id"] = cty.StringVal(p.TrafficMirrorFilterId)
}

func EncodeEc2TrafficMirrorSession_Tags(p Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2TrafficMirrorSession_VirtualNetworkId(p Ec2TrafficMirrorSessionParameters, vals map[string]cty.Value) {
	vals["virtual_network_id"] = cty.NumberIntVal(p.VirtualNetworkId)
}

func EncodeEc2TrafficMirrorSession_Arn(p Ec2TrafficMirrorSessionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}