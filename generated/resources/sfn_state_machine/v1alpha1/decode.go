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
	r, ok := mr.(*SfnStateMachine)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSfnStateMachine(r, ctyValue)
}

func DecodeSfnStateMachine(prev *SfnStateMachine, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSfnStateMachine_RoleArn(&new.Spec.ForProvider, valMap)
	DecodeSfnStateMachine_Tags(&new.Spec.ForProvider, valMap)
	DecodeSfnStateMachine_Definition(&new.Spec.ForProvider, valMap)
	DecodeSfnStateMachine_Id(&new.Spec.ForProvider, valMap)
	DecodeSfnStateMachine_Name(&new.Spec.ForProvider, valMap)
	DecodeSfnStateMachine_Status(&new.Status.AtProvider, valMap)
	DecodeSfnStateMachine_Arn(&new.Status.AtProvider, valMap)
	DecodeSfnStateMachine_CreationDate(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeSfnStateMachine_RoleArn(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	p.RoleArn = ctwhy.ValueAsString(vals["role_arn"])
}

func DecodeSfnStateMachine_Tags(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeSfnStateMachine_Definition(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	p.Definition = ctwhy.ValueAsString(vals["definition"])
}

func DecodeSfnStateMachine_Id(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeSfnStateMachine_Name(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeSfnStateMachine_Status(p *SfnStateMachineObservation, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

func DecodeSfnStateMachine_Arn(p *SfnStateMachineObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeSfnStateMachine_CreationDate(p *SfnStateMachineObservation, vals map[string]cty.Value) {
	p.CreationDate = ctwhy.ValueAsString(vals["creation_date"])
}