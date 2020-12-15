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
	r, ok := mr.(*GuarddutyDetector)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeGuarddutyDetector(r, ctyValue)
}

func DecodeGuarddutyDetector(prev *GuarddutyDetector, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeGuarddutyDetector_Enable(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyDetector_FindingPublishingFrequency(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyDetector_Id(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyDetector_Tags(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyDetector_Arn(&new.Status.AtProvider, valMap)
	DecodeGuarddutyDetector_AccountId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyDetector_Enable(p *GuarddutyDetectorParameters, vals map[string]cty.Value) {
	p.Enable = ctwhy.ValueAsBool(vals["enable"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyDetector_FindingPublishingFrequency(p *GuarddutyDetectorParameters, vals map[string]cty.Value) {
	p.FindingPublishingFrequency = ctwhy.ValueAsString(vals["finding_publishing_frequency"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyDetector_Id(p *GuarddutyDetectorParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeGuarddutyDetector_Tags(p *GuarddutyDetectorParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyDetector_Arn(p *GuarddutyDetectorObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyDetector_AccountId(p *GuarddutyDetectorObservation, vals map[string]cty.Value) {
	p.AccountId = ctwhy.ValueAsString(vals["account_id"])
}