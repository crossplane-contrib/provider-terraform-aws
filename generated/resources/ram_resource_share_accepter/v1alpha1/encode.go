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

func EncodeRamResourceShareAccepter(r RamResourceShareAccepter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRamResourceShareAccepter_Id(r.Spec.ForProvider, ctyVal)
	EncodeRamResourceShareAccepter_ShareArn(r.Spec.ForProvider, ctyVal)
	EncodeRamResourceShareAccepter_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRamResourceShareAccepter_Resources(r.Status.AtProvider, ctyVal)
	EncodeRamResourceShareAccepter_ShareId(r.Status.AtProvider, ctyVal)
	EncodeRamResourceShareAccepter_Status(r.Status.AtProvider, ctyVal)
	EncodeRamResourceShareAccepter_InvitationArn(r.Status.AtProvider, ctyVal)
	EncodeRamResourceShareAccepter_ReceiverAccountId(r.Status.AtProvider, ctyVal)
	EncodeRamResourceShareAccepter_SenderAccountId(r.Status.AtProvider, ctyVal)
	EncodeRamResourceShareAccepter_ShareName(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRamResourceShareAccepter_Id(p RamResourceShareAccepterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRamResourceShareAccepter_ShareArn(p RamResourceShareAccepterParameters, vals map[string]cty.Value) {
	vals["share_arn"] = cty.StringVal(p.ShareArn)
}

func EncodeRamResourceShareAccepter_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRamResourceShareAccepter_Timeouts_Delete(p, ctyVal)
	EncodeRamResourceShareAccepter_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRamResourceShareAccepter_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRamResourceShareAccepter_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRamResourceShareAccepter_Resources(p RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Resources {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["resources"] = cty.ListVal(colVals)
}

func EncodeRamResourceShareAccepter_ShareId(p RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	vals["share_id"] = cty.StringVal(p.ShareId)
}

func EncodeRamResourceShareAccepter_Status(p RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeRamResourceShareAccepter_InvitationArn(p RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	vals["invitation_arn"] = cty.StringVal(p.InvitationArn)
}

func EncodeRamResourceShareAccepter_ReceiverAccountId(p RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	vals["receiver_account_id"] = cty.StringVal(p.ReceiverAccountId)
}

func EncodeRamResourceShareAccepter_SenderAccountId(p RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	vals["sender_account_id"] = cty.StringVal(p.SenderAccountId)
}

func EncodeRamResourceShareAccepter_ShareName(p RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	vals["share_name"] = cty.StringVal(p.ShareName)
}