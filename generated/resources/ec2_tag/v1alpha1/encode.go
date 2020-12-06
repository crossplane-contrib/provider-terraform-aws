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

func EncodeEc2Tag(r Ec2Tag) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2Tag_Value(r.Spec.ForProvider, ctyVal)
	EncodeEc2Tag_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2Tag_Key(r.Spec.ForProvider, ctyVal)
	EncodeEc2Tag_ResourceId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeEc2Tag_Value(p Ec2TagParameters, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeEc2Tag_Id(p Ec2TagParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2Tag_Key(p Ec2TagParameters, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeEc2Tag_ResourceId(p Ec2TagParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}