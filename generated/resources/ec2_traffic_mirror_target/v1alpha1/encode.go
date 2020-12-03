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

package v1alpha1func EncodeEc2TrafficMirrorTarget(r Ec2TrafficMirrorTarget) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEc2TrafficMirrorTarget_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorTarget_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorTarget_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorTarget_NetworkInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorTarget_NetworkLoadBalancerArn(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorTarget_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEc2TrafficMirrorTarget_Tags(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2TrafficMirrorTarget_Description(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2TrafficMirrorTarget_Id(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TrafficMirrorTarget_NetworkInterfaceId(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeEc2TrafficMirrorTarget_NetworkLoadBalancerArn(p *Ec2TrafficMirrorTargetParameters, vals map[string]cty.Value) {
	vals["network_load_balancer_arn"] = cty.StringVal(p.NetworkLoadBalancerArn)
}

func EncodeEc2TrafficMirrorTarget_Arn(p *Ec2TrafficMirrorTargetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}