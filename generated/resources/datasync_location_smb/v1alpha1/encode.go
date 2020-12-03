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

package v1alpha1func EncodeDatasyncLocationSmb(r DatasyncLocationSmb) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDatasyncLocationSmb_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_User(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_AgentArns(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_Domain(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_Password(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_ServerHostname(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_Subdirectory(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_MountOptions(r.Spec.ForProvider.MountOptions, ctyVal)
	EncodeDatasyncLocationSmb_Uri(r.Status.AtProvider, ctyVal)
	EncodeDatasyncLocationSmb_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeDatasyncLocationSmb_Tags(p *DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDatasyncLocationSmb_User(p *DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["user"] = cty.StringVal(p.User)
}

func EncodeDatasyncLocationSmb_AgentArns(p *DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AgentArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["agent_arns"] = cty.SetVal(colVals)
}

func EncodeDatasyncLocationSmb_Domain(p *DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeDatasyncLocationSmb_Password(p *DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeDatasyncLocationSmb_ServerHostname(p *DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["server_hostname"] = cty.StringVal(p.ServerHostname)
}

func EncodeDatasyncLocationSmb_Subdirectory(p *DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["subdirectory"] = cty.StringVal(p.Subdirectory)
}

func EncodeDatasyncLocationSmb_Id(p *DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncLocationSmb_MountOptions(p *MountOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MountOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeDatasyncLocationSmb_MountOptions_Version(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["mount_options"] = cty.ListVal(valsForCollection)
}

func EncodeDatasyncLocationSmb_MountOptions_Version(p *MountOptions, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeDatasyncLocationSmb_Uri(p *DatasyncLocationSmbObservation, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeDatasyncLocationSmb_Arn(p *DatasyncLocationSmbObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}