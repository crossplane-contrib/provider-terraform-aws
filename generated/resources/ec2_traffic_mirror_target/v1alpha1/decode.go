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
	r, ok := mr.(*Ec2TrafficMirrorTarget)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2TrafficMirrorTarget(r, ctyValue)
}

func DecodeEc2TrafficMirrorTarget(prev *Ec2TrafficMirrorTarget, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2TrafficMirrorTarget_Description(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorTarget_Id(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorTarget_NetworkInterfaceId(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorTarget_NetworkLoadBalancerArn(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorTarget_Tags(&new.Spec.ForProvider, valMap)
	DecodeEc2TrafficMirrorTarget_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeEc2TrafficMirrorTarget_Description(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeEc2TrafficMirrorTarget_Id(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeEc2TrafficMirrorTarget_NetworkInterfaceId(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	p.NetworkInterfaceId = ctwhy.ValueAsString(vals["network_interface_id"])
}

func DecodeEc2TrafficMirrorTarget_NetworkLoadBalancerArn(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	p.NetworkLoadBalancerArn = ctwhy.ValueAsString(vals["network_load_balancer_arn"])
}

func DecodeEc2TrafficMirrorTarget_Tags(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeEc2TrafficMirrorTarget_Arn(p *Ec2TrafficMirrorTargetObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}