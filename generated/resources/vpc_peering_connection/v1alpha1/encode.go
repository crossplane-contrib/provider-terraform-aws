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

func EncodeVpcPeeringConnection(r VpcPeeringConnection) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnection_PeerVpcId(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnection_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnection_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnection_AutoAccept(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnection_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnection_PeerOwnerId(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnection_PeerRegion(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnection_Accepter(r.Spec.ForProvider.Accepter, ctyVal)
	EncodeVpcPeeringConnection_Requester(r.Spec.ForProvider.Requester, ctyVal)
	EncodeVpcPeeringConnection_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeVpcPeeringConnection_AcceptStatus(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeVpcPeeringConnection_PeerVpcId(p VpcPeeringConnectionParameters, vals map[string]cty.Value) {
	vals["peer_vpc_id"] = cty.StringVal(p.PeerVpcId)
}

func EncodeVpcPeeringConnection_Tags(p VpcPeeringConnectionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeVpcPeeringConnection_VpcId(p VpcPeeringConnectionParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeVpcPeeringConnection_AutoAccept(p VpcPeeringConnectionParameters, vals map[string]cty.Value) {
	vals["auto_accept"] = cty.BoolVal(p.AutoAccept)
}

func EncodeVpcPeeringConnection_Id(p VpcPeeringConnectionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcPeeringConnection_PeerOwnerId(p VpcPeeringConnectionParameters, vals map[string]cty.Value) {
	vals["peer_owner_id"] = cty.StringVal(p.PeerOwnerId)
}

func EncodeVpcPeeringConnection_PeerRegion(p VpcPeeringConnectionParameters, vals map[string]cty.Value) {
	vals["peer_region"] = cty.StringVal(p.PeerRegion)
}

func EncodeVpcPeeringConnection_Accepter(p Accepter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnection_Accepter_AllowClassicLinkToRemoteVpc(p, ctyVal)
	EncodeVpcPeeringConnection_Accepter_AllowRemoteVpcDnsResolution(p, ctyVal)
	EncodeVpcPeeringConnection_Accepter_AllowVpcToRemoteClassicLink(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["accepter"] = cty.ListVal(valsForCollection)
}

func EncodeVpcPeeringConnection_Accepter_AllowClassicLinkToRemoteVpc(p Accepter, vals map[string]cty.Value) {
	vals["allow_classic_link_to_remote_vpc"] = cty.BoolVal(p.AllowClassicLinkToRemoteVpc)
}

func EncodeVpcPeeringConnection_Accepter_AllowRemoteVpcDnsResolution(p Accepter, vals map[string]cty.Value) {
	vals["allow_remote_vpc_dns_resolution"] = cty.BoolVal(p.AllowRemoteVpcDnsResolution)
}

func EncodeVpcPeeringConnection_Accepter_AllowVpcToRemoteClassicLink(p Accepter, vals map[string]cty.Value) {
	vals["allow_vpc_to_remote_classic_link"] = cty.BoolVal(p.AllowVpcToRemoteClassicLink)
}

func EncodeVpcPeeringConnection_Requester(p Requester, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnection_Requester_AllowClassicLinkToRemoteVpc(p, ctyVal)
	EncodeVpcPeeringConnection_Requester_AllowRemoteVpcDnsResolution(p, ctyVal)
	EncodeVpcPeeringConnection_Requester_AllowVpcToRemoteClassicLink(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["requester"] = cty.ListVal(valsForCollection)
}

func EncodeVpcPeeringConnection_Requester_AllowClassicLinkToRemoteVpc(p Requester, vals map[string]cty.Value) {
	vals["allow_classic_link_to_remote_vpc"] = cty.BoolVal(p.AllowClassicLinkToRemoteVpc)
}

func EncodeVpcPeeringConnection_Requester_AllowRemoteVpcDnsResolution(p Requester, vals map[string]cty.Value) {
	vals["allow_remote_vpc_dns_resolution"] = cty.BoolVal(p.AllowRemoteVpcDnsResolution)
}

func EncodeVpcPeeringConnection_Requester_AllowVpcToRemoteClassicLink(p Requester, vals map[string]cty.Value) {
	vals["allow_vpc_to_remote_classic_link"] = cty.BoolVal(p.AllowVpcToRemoteClassicLink)
}

func EncodeVpcPeeringConnection_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnection_Timeouts_Create(p, ctyVal)
	EncodeVpcPeeringConnection_Timeouts_Delete(p, ctyVal)
	EncodeVpcPeeringConnection_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeVpcPeeringConnection_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeVpcPeeringConnection_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeVpcPeeringConnection_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeVpcPeeringConnection_AcceptStatus(p VpcPeeringConnectionObservation, vals map[string]cty.Value) {
	vals["accept_status"] = cty.StringVal(p.AcceptStatus)
}