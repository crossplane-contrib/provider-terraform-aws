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
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*RamResourceShareAccepter)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRamResourceShareAccepter(r, ctyValue)
}

func DecodeRamResourceShareAccepter(prev *RamResourceShareAccepter, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRamResourceShareAccepter_ShareArn(&new.Spec.ForProvider, valMap)
	DecodeRamResourceShareAccepter_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeRamResourceShareAccepter_InvitationArn(&new.Status.AtProvider, valMap)
	DecodeRamResourceShareAccepter_ReceiverAccountId(&new.Status.AtProvider, valMap)
	DecodeRamResourceShareAccepter_SenderAccountId(&new.Status.AtProvider, valMap)
	DecodeRamResourceShareAccepter_ShareId(&new.Status.AtProvider, valMap)
	DecodeRamResourceShareAccepter_Status(&new.Status.AtProvider, valMap)
	DecodeRamResourceShareAccepter_Resources(&new.Status.AtProvider, valMap)
	DecodeRamResourceShareAccepter_ShareName(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeRamResourceShareAccepter_ShareArn(p *RamResourceShareAccepterParameters, vals map[string]cty.Value) {
	p.ShareArn = ctwhy.ValueAsString(vals["share_arn"])
}

//containerTypeDecodeTemplate
func DecodeRamResourceShareAccepter_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeRamResourceShareAccepter_Timeouts_Create(p, valMap)
	DecodeRamResourceShareAccepter_Timeouts_Delete(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeRamResourceShareAccepter_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeRamResourceShareAccepter_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeRamResourceShareAccepter_InvitationArn(p *RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	p.InvitationArn = ctwhy.ValueAsString(vals["invitation_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeRamResourceShareAccepter_ReceiverAccountId(p *RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	p.ReceiverAccountId = ctwhy.ValueAsString(vals["receiver_account_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRamResourceShareAccepter_SenderAccountId(p *RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	p.SenderAccountId = ctwhy.ValueAsString(vals["sender_account_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRamResourceShareAccepter_ShareId(p *RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	p.ShareId = ctwhy.ValueAsString(vals["share_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRamResourceShareAccepter_Status(p *RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeRamResourceShareAccepter_Resources(p *RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["resources"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Resources = goVals
}

//primitiveTypeDecodeTemplate
func DecodeRamResourceShareAccepter_ShareName(p *RamResourceShareAccepterObservation, vals map[string]cty.Value) {
	p.ShareName = ctwhy.ValueAsString(vals["share_name"])
}