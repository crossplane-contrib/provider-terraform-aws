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

package v1alpha1func EncodeSecurityhubActionTarget(r SecurityhubActionTarget) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSecurityhubActionTarget_Description(r.Spec.ForProvider, ctyVal)
	EncodeSecurityhubActionTarget_Id(r.Spec.ForProvider, ctyVal)
	EncodeSecurityhubActionTarget_Identifier(r.Spec.ForProvider, ctyVal)
	EncodeSecurityhubActionTarget_Name(r.Spec.ForProvider, ctyVal)
	EncodeSecurityhubActionTarget_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSecurityhubActionTarget_Description(p *SecurityhubActionTargetParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSecurityhubActionTarget_Id(p *SecurityhubActionTargetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSecurityhubActionTarget_Identifier(p *SecurityhubActionTargetParameters, vals map[string]cty.Value) {
	vals["identifier"] = cty.StringVal(p.Identifier)
}

func EncodeSecurityhubActionTarget_Name(p *SecurityhubActionTargetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSecurityhubActionTarget_Arn(p *SecurityhubActionTargetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}