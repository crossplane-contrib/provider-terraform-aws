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

package v1alpha1func EncodeSfnStateMachine(r SfnStateMachine) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSfnStateMachine_Definition(r.Spec.ForProvider, ctyVal)
	EncodeSfnStateMachine_Id(r.Spec.ForProvider, ctyVal)
	EncodeSfnStateMachine_Name(r.Spec.ForProvider, ctyVal)
	EncodeSfnStateMachine_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSfnStateMachine_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSfnStateMachine_CreationDate(r.Status.AtProvider, ctyVal)
	EncodeSfnStateMachine_Status(r.Status.AtProvider, ctyVal)
	EncodeSfnStateMachine_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSfnStateMachine_Definition(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	vals["definition"] = cty.StringVal(p.Definition)
}

func EncodeSfnStateMachine_Id(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSfnStateMachine_Name(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSfnStateMachine_RoleArn(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeSfnStateMachine_Tags(p *SfnStateMachineParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSfnStateMachine_CreationDate(p *SfnStateMachineObservation, vals map[string]cty.Value) {
	vals["creation_date"] = cty.StringVal(p.CreationDate)
}

func EncodeSfnStateMachine_Status(p *SfnStateMachineObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeSfnStateMachine_Arn(p *SfnStateMachineObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}