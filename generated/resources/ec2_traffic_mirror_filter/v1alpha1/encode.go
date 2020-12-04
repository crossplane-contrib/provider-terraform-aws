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
	"github.com/zclconf/go-cty/cty"
)

func EncodeEc2TrafficMirrorFilter(r Ec2TrafficMirrorFilter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TrafficMirrorFilter_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilter_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilter_NetworkServices(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilter_Tags(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TrafficMirrorFilter_Description(p Ec2TrafficMirrorFilterParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2TrafficMirrorFilter_Id(p Ec2TrafficMirrorFilterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TrafficMirrorFilter_NetworkServices(p Ec2TrafficMirrorFilterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NetworkServices {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["network_services"] = cty.SetVal(colVals)
}

func EncodeEc2TrafficMirrorFilter_Tags(p Ec2TrafficMirrorFilterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}