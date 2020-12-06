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

func EncodeEc2AvailabilityZoneGroup(r Ec2AvailabilityZoneGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2AvailabilityZoneGroup_GroupName(r.Spec.ForProvider, ctyVal)
	EncodeEc2AvailabilityZoneGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2AvailabilityZoneGroup_OptInStatus(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeEc2AvailabilityZoneGroup_GroupName(p Ec2AvailabilityZoneGroupParameters, vals map[string]cty.Value) {
	vals["group_name"] = cty.StringVal(p.GroupName)
}

func EncodeEc2AvailabilityZoneGroup_Id(p Ec2AvailabilityZoneGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2AvailabilityZoneGroup_OptInStatus(p Ec2AvailabilityZoneGroupParameters, vals map[string]cty.Value) {
	vals["opt_in_status"] = cty.StringVal(p.OptInStatus)
}