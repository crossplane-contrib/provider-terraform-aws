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
	r, ok := mr.(*IotThing)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIotThing(r, ctyValue)
}

func DecodeIotThing(prev *IotThing, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIotThing_Attributes(&new.Spec.ForProvider, valMap)
	DecodeIotThing_Id(&new.Spec.ForProvider, valMap)
	DecodeIotThing_Name(&new.Spec.ForProvider, valMap)
	DecodeIotThing_ThingTypeName(&new.Spec.ForProvider, valMap)
	DecodeIotThing_Version(&new.Status.AtProvider, valMap)
	DecodeIotThing_Arn(&new.Status.AtProvider, valMap)
	DecodeIotThing_DefaultClientId(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeIotThing_Attributes(p *IotThingParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["attributes"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Attributes = vMap
}

func DecodeIotThing_Id(p *IotThingParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeIotThing_Name(p *IotThingParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeIotThing_ThingTypeName(p *IotThingParameters, vals map[string]cty.Value) {
	p.ThingTypeName = ctwhy.ValueAsString(vals["thing_type_name"])
}

func DecodeIotThing_Version(p *IotThingObservation, vals map[string]cty.Value) {
	p.Version = ctwhy.ValueAsInt64(vals["version"])
}

func DecodeIotThing_Arn(p *IotThingObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeIotThing_DefaultClientId(p *IotThingObservation, vals map[string]cty.Value) {
	p.DefaultClientId = ctwhy.ValueAsString(vals["default_client_id"])
}