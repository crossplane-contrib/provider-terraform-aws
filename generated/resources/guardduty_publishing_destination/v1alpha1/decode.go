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
	r, ok := mr.(*GuarddutyPublishingDestination)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeGuarddutyPublishingDestination(r, ctyValue)
}

func DecodeGuarddutyPublishingDestination(prev *GuarddutyPublishingDestination, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeGuarddutyPublishingDestination_DestinationArn(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyPublishingDestination_DestinationType(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyPublishingDestination_DetectorId(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyPublishingDestination_Id(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyPublishingDestination_KmsKeyArn(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyPublishingDestination_DestinationArn(p *GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	p.DestinationArn = ctwhy.ValueAsString(vals["destination_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyPublishingDestination_DestinationType(p *GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	p.DestinationType = ctwhy.ValueAsString(vals["destination_type"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyPublishingDestination_DetectorId(p *GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	p.DetectorId = ctwhy.ValueAsString(vals["detector_id"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyPublishingDestination_Id(p *GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyPublishingDestination_KmsKeyArn(p *GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	p.KmsKeyArn = ctwhy.ValueAsString(vals["kms_key_arn"])
}