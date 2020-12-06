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

func EncodeVpnGateway(r VpnGateway) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpnGateway_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeVpnGateway_AmazonSideAsn(r.Spec.ForProvider, ctyVal)
	EncodeVpnGateway_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeVpnGateway_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpnGateway_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpnGateway_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeVpnGateway_VpcId(p VpnGatewayParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeVpnGateway_AmazonSideAsn(p VpnGatewayParameters, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.StringVal(p.AmazonSideAsn)
}

func EncodeVpnGateway_AvailabilityZone(p VpnGatewayParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeVpnGateway_Id(p VpnGatewayParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpnGateway_Tags(p VpnGatewayParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeVpnGateway_Arn(p VpnGatewayObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}