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

func EncodeEc2CapacityReservation(r Ec2CapacityReservation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2CapacityReservation_InstancePlatform(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_Tenancy(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_EbsOptimized(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_EndDate(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_EndDateType(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_InstanceCount(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_InstanceMatchCriteria(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_EphemeralStorage(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2CapacityReservation_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2CapacityReservation_InstancePlatform(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["instance_platform"] = cty.StringVal(p.InstancePlatform)
}

func EncodeEc2CapacityReservation_InstanceType(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEc2CapacityReservation_Tags(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2CapacityReservation_Tenancy(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["tenancy"] = cty.StringVal(p.Tenancy)
}

func EncodeEc2CapacityReservation_EbsOptimized(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["ebs_optimized"] = cty.BoolVal(p.EbsOptimized)
}

func EncodeEc2CapacityReservation_EndDate(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["end_date"] = cty.StringVal(p.EndDate)
}

func EncodeEc2CapacityReservation_EndDateType(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["end_date_type"] = cty.StringVal(p.EndDateType)
}

func EncodeEc2CapacityReservation_InstanceCount(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["instance_count"] = cty.NumberIntVal(p.InstanceCount)
}

func EncodeEc2CapacityReservation_InstanceMatchCriteria(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["instance_match_criteria"] = cty.StringVal(p.InstanceMatchCriteria)
}

func EncodeEc2CapacityReservation_AvailabilityZone(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeEc2CapacityReservation_EphemeralStorage(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["ephemeral_storage"] = cty.BoolVal(p.EphemeralStorage)
}

func EncodeEc2CapacityReservation_Id(p Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2CapacityReservation_Arn(p Ec2CapacityReservationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}