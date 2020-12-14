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
	r, ok := mr.(*GuarddutyInviteAccepter)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeGuarddutyInviteAccepter(r, ctyValue)
}

func DecodeGuarddutyInviteAccepter(prev *GuarddutyInviteAccepter, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeGuarddutyInviteAccepter_DetectorId(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyInviteAccepter_Id(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyInviteAccepter_MasterAccountId(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyInviteAccepter_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeGuarddutyInviteAccepter_DetectorId(p *GuarddutyInviteAccepterParameters, vals map[string]cty.Value) {
	p.DetectorId = ctwhy.ValueAsString(vals["detector_id"])
}

func DecodeGuarddutyInviteAccepter_Id(p *GuarddutyInviteAccepterParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeGuarddutyInviteAccepter_MasterAccountId(p *GuarddutyInviteAccepterParameters, vals map[string]cty.Value) {
	p.MasterAccountId = ctwhy.ValueAsString(vals["master_account_id"])
}

func DecodeGuarddutyInviteAccepter_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeGuarddutyInviteAccepter_Timeouts_Create(p, valMap)
}

func DecodeGuarddutyInviteAccepter_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}