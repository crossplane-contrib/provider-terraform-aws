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

func EncodeVpcPeeringConnectionOptions(r VpcPeeringConnectionOptions) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnectionOptions_VpcPeeringConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnectionOptions_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnectionOptions_Requester(r.Spec.ForProvider.Requester, ctyVal)
	EncodeVpcPeeringConnectionOptions_Accepter(r.Spec.ForProvider.Accepter, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeVpcPeeringConnectionOptions_VpcPeeringConnectionId(p VpcPeeringConnectionOptionsParameters, vals map[string]cty.Value) {
	vals["vpc_peering_connection_id"] = cty.StringVal(p.VpcPeeringConnectionId)
}

func EncodeVpcPeeringConnectionOptions_Id(p VpcPeeringConnectionOptionsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcPeeringConnectionOptions_Requester(p Requester, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnectionOptions_Requester_AllowClassicLinkToRemoteVpc(p, ctyVal)
	EncodeVpcPeeringConnectionOptions_Requester_AllowRemoteVpcDnsResolution(p, ctyVal)
	EncodeVpcPeeringConnectionOptions_Requester_AllowVpcToRemoteClassicLink(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["requester"] = cty.ListVal(valsForCollection)
}

func EncodeVpcPeeringConnectionOptions_Requester_AllowClassicLinkToRemoteVpc(p Requester, vals map[string]cty.Value) {
	vals["allow_classic_link_to_remote_vpc"] = cty.BoolVal(p.AllowClassicLinkToRemoteVpc)
}

func EncodeVpcPeeringConnectionOptions_Requester_AllowRemoteVpcDnsResolution(p Requester, vals map[string]cty.Value) {
	vals["allow_remote_vpc_dns_resolution"] = cty.BoolVal(p.AllowRemoteVpcDnsResolution)
}

func EncodeVpcPeeringConnectionOptions_Requester_AllowVpcToRemoteClassicLink(p Requester, vals map[string]cty.Value) {
	vals["allow_vpc_to_remote_classic_link"] = cty.BoolVal(p.AllowVpcToRemoteClassicLink)
}

func EncodeVpcPeeringConnectionOptions_Accepter(p Accepter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnectionOptions_Accepter_AllowClassicLinkToRemoteVpc(p, ctyVal)
	EncodeVpcPeeringConnectionOptions_Accepter_AllowRemoteVpcDnsResolution(p, ctyVal)
	EncodeVpcPeeringConnectionOptions_Accepter_AllowVpcToRemoteClassicLink(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["accepter"] = cty.ListVal(valsForCollection)
}

func EncodeVpcPeeringConnectionOptions_Accepter_AllowClassicLinkToRemoteVpc(p Accepter, vals map[string]cty.Value) {
	vals["allow_classic_link_to_remote_vpc"] = cty.BoolVal(p.AllowClassicLinkToRemoteVpc)
}

func EncodeVpcPeeringConnectionOptions_Accepter_AllowRemoteVpcDnsResolution(p Accepter, vals map[string]cty.Value) {
	vals["allow_remote_vpc_dns_resolution"] = cty.BoolVal(p.AllowRemoteVpcDnsResolution)
}

func EncodeVpcPeeringConnectionOptions_Accepter_AllowVpcToRemoteClassicLink(p Accepter, vals map[string]cty.Value) {
	vals["allow_vpc_to_remote_classic_link"] = cty.BoolVal(p.AllowVpcToRemoteClassicLink)
}