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

func EncodeVpcPeeringConnectionAccepter(r VpcPeeringConnectionAccepter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnectionAccepter_VpcPeeringConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnectionAccepter_AutoAccept(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnectionAccepter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnectionAccepter_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcPeeringConnectionAccepter_Accepter(r.Spec.ForProvider.Accepter, ctyVal)
	EncodeVpcPeeringConnectionAccepter_Requester(r.Spec.ForProvider.Requester, ctyVal)
	EncodeVpcPeeringConnectionAccepter_VpcId(r.Status.AtProvider, ctyVal)
	EncodeVpcPeeringConnectionAccepter_PeerVpcId(r.Status.AtProvider, ctyVal)
	EncodeVpcPeeringConnectionAccepter_PeerOwnerId(r.Status.AtProvider, ctyVal)
	EncodeVpcPeeringConnectionAccepter_PeerRegion(r.Status.AtProvider, ctyVal)
	EncodeVpcPeeringConnectionAccepter_AcceptStatus(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeVpcPeeringConnectionAccepter_VpcPeeringConnectionId(p VpcPeeringConnectionAccepterParameters, vals map[string]cty.Value) {
	vals["vpc_peering_connection_id"] = cty.StringVal(p.VpcPeeringConnectionId)
}

func EncodeVpcPeeringConnectionAccepter_AutoAccept(p VpcPeeringConnectionAccepterParameters, vals map[string]cty.Value) {
	vals["auto_accept"] = cty.BoolVal(p.AutoAccept)
}

func EncodeVpcPeeringConnectionAccepter_Tags(p VpcPeeringConnectionAccepterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeVpcPeeringConnectionAccepter_Id(p VpcPeeringConnectionAccepterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcPeeringConnectionAccepter_Accepter(p Accepter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnectionAccepter_Accepter_AllowVpcToRemoteClassicLink(p, ctyVal)
	EncodeVpcPeeringConnectionAccepter_Accepter_AllowClassicLinkToRemoteVpc(p, ctyVal)
	EncodeVpcPeeringConnectionAccepter_Accepter_AllowRemoteVpcDnsResolution(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["accepter"] = cty.ListVal(valsForCollection)
}

func EncodeVpcPeeringConnectionAccepter_Accepter_AllowVpcToRemoteClassicLink(p Accepter, vals map[string]cty.Value) {
	vals["allow_vpc_to_remote_classic_link"] = cty.BoolVal(p.AllowVpcToRemoteClassicLink)
}

func EncodeVpcPeeringConnectionAccepter_Accepter_AllowClassicLinkToRemoteVpc(p Accepter, vals map[string]cty.Value) {
	vals["allow_classic_link_to_remote_vpc"] = cty.BoolVal(p.AllowClassicLinkToRemoteVpc)
}

func EncodeVpcPeeringConnectionAccepter_Accepter_AllowRemoteVpcDnsResolution(p Accepter, vals map[string]cty.Value) {
	vals["allow_remote_vpc_dns_resolution"] = cty.BoolVal(p.AllowRemoteVpcDnsResolution)
}

func EncodeVpcPeeringConnectionAccepter_Requester(p Requester, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVpcPeeringConnectionAccepter_Requester_AllowClassicLinkToRemoteVpc(p, ctyVal)
	EncodeVpcPeeringConnectionAccepter_Requester_AllowRemoteVpcDnsResolution(p, ctyVal)
	EncodeVpcPeeringConnectionAccepter_Requester_AllowVpcToRemoteClassicLink(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["requester"] = cty.ListVal(valsForCollection)
}

func EncodeVpcPeeringConnectionAccepter_Requester_AllowClassicLinkToRemoteVpc(p Requester, vals map[string]cty.Value) {
	vals["allow_classic_link_to_remote_vpc"] = cty.BoolVal(p.AllowClassicLinkToRemoteVpc)
}

func EncodeVpcPeeringConnectionAccepter_Requester_AllowRemoteVpcDnsResolution(p Requester, vals map[string]cty.Value) {
	vals["allow_remote_vpc_dns_resolution"] = cty.BoolVal(p.AllowRemoteVpcDnsResolution)
}

func EncodeVpcPeeringConnectionAccepter_Requester_AllowVpcToRemoteClassicLink(p Requester, vals map[string]cty.Value) {
	vals["allow_vpc_to_remote_classic_link"] = cty.BoolVal(p.AllowVpcToRemoteClassicLink)
}

func EncodeVpcPeeringConnectionAccepter_VpcId(p VpcPeeringConnectionAccepterObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeVpcPeeringConnectionAccepter_PeerVpcId(p VpcPeeringConnectionAccepterObservation, vals map[string]cty.Value) {
	vals["peer_vpc_id"] = cty.StringVal(p.PeerVpcId)
}

func EncodeVpcPeeringConnectionAccepter_PeerOwnerId(p VpcPeeringConnectionAccepterObservation, vals map[string]cty.Value) {
	vals["peer_owner_id"] = cty.StringVal(p.PeerOwnerId)
}

func EncodeVpcPeeringConnectionAccepter_PeerRegion(p VpcPeeringConnectionAccepterObservation, vals map[string]cty.Value) {
	vals["peer_region"] = cty.StringVal(p.PeerRegion)
}

func EncodeVpcPeeringConnectionAccepter_AcceptStatus(p VpcPeeringConnectionAccepterObservation, vals map[string]cty.Value) {
	vals["accept_status"] = cty.StringVal(p.AcceptStatus)
}