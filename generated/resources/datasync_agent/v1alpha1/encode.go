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

package v1alpha1func EncodeDatasyncAgent(r DatasyncAgent) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDatasyncAgent_ActivationKey(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_IpAddress(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_Name(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDatasyncAgent_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeDatasyncAgent_ActivationKey(p *DatasyncAgentParameters, vals map[string]cty.Value) {
	vals["activation_key"] = cty.StringVal(p.ActivationKey)
}

func EncodeDatasyncAgent_Id(p *DatasyncAgentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncAgent_IpAddress(p *DatasyncAgentParameters, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}

func EncodeDatasyncAgent_Name(p *DatasyncAgentParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDatasyncAgent_Tags(p *DatasyncAgentParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDatasyncAgent_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeDatasyncAgent_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDatasyncAgent_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDatasyncAgent_Arn(p *DatasyncAgentObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}