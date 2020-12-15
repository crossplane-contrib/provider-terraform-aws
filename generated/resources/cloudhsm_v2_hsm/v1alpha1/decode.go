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
	r, ok := mr.(*CloudhsmV2Hsm)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCloudhsmV2Hsm(r, ctyValue)
}

func DecodeCloudhsmV2Hsm(prev *CloudhsmV2Hsm, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCloudhsmV2Hsm_SubnetId(&new.Spec.ForProvider, valMap)
	DecodeCloudhsmV2Hsm_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeCloudhsmV2Hsm_ClusterId(&new.Spec.ForProvider, valMap)
	DecodeCloudhsmV2Hsm_IpAddress(&new.Spec.ForProvider, valMap)
	DecodeCloudhsmV2Hsm_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeCloudhsmV2Hsm_HsmEniId(&new.Status.AtProvider, valMap)
	DecodeCloudhsmV2Hsm_HsmId(&new.Status.AtProvider, valMap)
	DecodeCloudhsmV2Hsm_HsmState(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_SubnetId(p *CloudhsmV2HsmParameters, vals map[string]cty.Value) {
	p.SubnetId = ctwhy.ValueAsString(vals["subnet_id"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_AvailabilityZone(p *CloudhsmV2HsmParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_ClusterId(p *CloudhsmV2HsmParameters, vals map[string]cty.Value) {
	p.ClusterId = ctwhy.ValueAsString(vals["cluster_id"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_IpAddress(p *CloudhsmV2HsmParameters, vals map[string]cty.Value) {
	p.IpAddress = ctwhy.ValueAsString(vals["ip_address"])
}

//containerTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeCloudhsmV2Hsm_Timeouts_Delete(p, valMap)
	DecodeCloudhsmV2Hsm_Timeouts_Update(p, valMap)
	DecodeCloudhsmV2Hsm_Timeouts_Create(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_HsmEniId(p *CloudhsmV2HsmObservation, vals map[string]cty.Value) {
	p.HsmEniId = ctwhy.ValueAsString(vals["hsm_eni_id"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_HsmId(p *CloudhsmV2HsmObservation, vals map[string]cty.Value) {
	p.HsmId = ctwhy.ValueAsString(vals["hsm_id"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudhsmV2Hsm_HsmState(p *CloudhsmV2HsmObservation, vals map[string]cty.Value) {
	p.HsmState = ctwhy.ValueAsString(vals["hsm_state"])
}