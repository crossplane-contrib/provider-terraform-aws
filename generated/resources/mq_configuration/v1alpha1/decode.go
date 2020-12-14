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
	r, ok := mr.(*MqConfiguration)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeMqConfiguration(r, ctyValue)
}

func DecodeMqConfiguration(prev *MqConfiguration, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeMqConfiguration_Data(&new.Spec.ForProvider, valMap)
	DecodeMqConfiguration_Id(&new.Spec.ForProvider, valMap)
	DecodeMqConfiguration_Name(&new.Spec.ForProvider, valMap)
	DecodeMqConfiguration_Tags(&new.Spec.ForProvider, valMap)
	DecodeMqConfiguration_EngineType(&new.Spec.ForProvider, valMap)
	DecodeMqConfiguration_EngineVersion(&new.Spec.ForProvider, valMap)
	DecodeMqConfiguration_Description(&new.Spec.ForProvider, valMap)
	DecodeMqConfiguration_LatestRevision(&new.Status.AtProvider, valMap)
	DecodeMqConfiguration_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeMqConfiguration_Data(p *MqConfigurationParameters, vals map[string]cty.Value) {
	p.Data = ctwhy.ValueAsString(vals["data"])
}

func DecodeMqConfiguration_Id(p *MqConfigurationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeMqConfiguration_Name(p *MqConfigurationParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeMqConfiguration_Tags(p *MqConfigurationParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeMqConfiguration_EngineType(p *MqConfigurationParameters, vals map[string]cty.Value) {
	p.EngineType = ctwhy.ValueAsString(vals["engine_type"])
}

func DecodeMqConfiguration_EngineVersion(p *MqConfigurationParameters, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

func DecodeMqConfiguration_Description(p *MqConfigurationParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeMqConfiguration_LatestRevision(p *MqConfigurationObservation, vals map[string]cty.Value) {
	p.LatestRevision = ctwhy.ValueAsInt64(vals["latest_revision"])
}

func DecodeMqConfiguration_Arn(p *MqConfigurationObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}