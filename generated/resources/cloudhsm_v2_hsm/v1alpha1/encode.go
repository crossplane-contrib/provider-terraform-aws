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

func EncodeCloudhsmV2Hsm(r CloudhsmV2Hsm) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudhsmV2Hsm_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Hsm_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Hsm_ClusterId(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Hsm_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Hsm_IpAddress(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Hsm_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeCloudhsmV2Hsm_HsmEniId(r.Status.AtProvider, ctyVal)
	EncodeCloudhsmV2Hsm_HsmId(r.Status.AtProvider, ctyVal)
	EncodeCloudhsmV2Hsm_HsmState(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudhsmV2Hsm_SubnetId(p CloudhsmV2HsmParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeCloudhsmV2Hsm_AvailabilityZone(p CloudhsmV2HsmParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeCloudhsmV2Hsm_ClusterId(p CloudhsmV2HsmParameters, vals map[string]cty.Value) {
	vals["cluster_id"] = cty.StringVal(p.ClusterId)
}

func EncodeCloudhsmV2Hsm_Id(p CloudhsmV2HsmParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudhsmV2Hsm_IpAddress(p CloudhsmV2HsmParameters, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}

func EncodeCloudhsmV2Hsm_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudhsmV2Hsm_Timeouts_Create(p, ctyVal)
	EncodeCloudhsmV2Hsm_Timeouts_Delete(p, ctyVal)
	EncodeCloudhsmV2Hsm_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeCloudhsmV2Hsm_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeCloudhsmV2Hsm_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeCloudhsmV2Hsm_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeCloudhsmV2Hsm_HsmEniId(p CloudhsmV2HsmObservation, vals map[string]cty.Value) {
	vals["hsm_eni_id"] = cty.StringVal(p.HsmEniId)
}

func EncodeCloudhsmV2Hsm_HsmId(p CloudhsmV2HsmObservation, vals map[string]cty.Value) {
	vals["hsm_id"] = cty.StringVal(p.HsmId)
}

func EncodeCloudhsmV2Hsm_HsmState(p CloudhsmV2HsmObservation, vals map[string]cty.Value) {
	vals["hsm_state"] = cty.StringVal(p.HsmState)
}