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
	r, ok := mr.(*CloudformationStackSetInstance)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCloudformationStackSetInstance(r, ctyValue)
}

func DecodeCloudformationStackSetInstance(prev *CloudformationStackSetInstance, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCloudformationStackSetInstance_AccountId(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSetInstance_ParameterOverrides(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSetInstance_Region(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSetInstance_RetainStack(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSetInstance_StackSetName(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSetInstance_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeCloudformationStackSetInstance_StackId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_AccountId(p *CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	p.AccountId = ctwhy.ValueAsString(vals["account_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_ParameterOverrides(p *CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["parameter_overrides"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.ParameterOverrides = vMap
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_Region(p *CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	p.Region = ctwhy.ValueAsString(vals["region"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_RetainStack(p *CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	p.RetainStack = ctwhy.ValueAsBool(vals["retain_stack"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_StackSetName(p *CloudformationStackSetInstanceParameters, vals map[string]cty.Value) {
	p.StackSetName = ctwhy.ValueAsString(vals["stack_set_name"])
}

//containerTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeCloudformationStackSetInstance_Timeouts_Create(p, valMap)
	DecodeCloudformationStackSetInstance_Timeouts_Delete(p, valMap)
	DecodeCloudformationStackSetInstance_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSetInstance_StackId(p *CloudformationStackSetInstanceObservation, vals map[string]cty.Value) {
	p.StackId = ctwhy.ValueAsString(vals["stack_id"])
}