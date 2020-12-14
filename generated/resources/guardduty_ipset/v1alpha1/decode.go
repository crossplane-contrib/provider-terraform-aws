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
	r, ok := mr.(*GuarddutyIpset)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeGuarddutyIpset(r, ctyValue)
}

func DecodeGuarddutyIpset(prev *GuarddutyIpset, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeGuarddutyIpset_Name(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyIpset_Tags(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyIpset_Activate(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyIpset_DetectorId(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyIpset_Format(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyIpset_Id(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyIpset_Location(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyIpset_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeGuarddutyIpset_Name(p *GuarddutyIpsetParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeGuarddutyIpset_Tags(p *GuarddutyIpsetParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeGuarddutyIpset_Activate(p *GuarddutyIpsetParameters, vals map[string]cty.Value) {
	p.Activate = ctwhy.ValueAsBool(vals["activate"])
}

func DecodeGuarddutyIpset_DetectorId(p *GuarddutyIpsetParameters, vals map[string]cty.Value) {
	p.DetectorId = ctwhy.ValueAsString(vals["detector_id"])
}

func DecodeGuarddutyIpset_Format(p *GuarddutyIpsetParameters, vals map[string]cty.Value) {
	p.Format = ctwhy.ValueAsString(vals["format"])
}

func DecodeGuarddutyIpset_Id(p *GuarddutyIpsetParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeGuarddutyIpset_Location(p *GuarddutyIpsetParameters, vals map[string]cty.Value) {
	p.Location = ctwhy.ValueAsString(vals["location"])
}

func DecodeGuarddutyIpset_Arn(p *GuarddutyIpsetObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}