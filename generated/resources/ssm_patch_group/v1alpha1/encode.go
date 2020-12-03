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

package v1alpha1func EncodeSsmPatchGroup(r SsmPatchGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSsmPatchGroup_BaselineId(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchGroup_PatchGroup(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeSsmPatchGroup_BaselineId(p *SsmPatchGroupParameters, vals map[string]cty.Value) {
	vals["baseline_id"] = cty.StringVal(p.BaselineId)
}

func EncodeSsmPatchGroup_Id(p *SsmPatchGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmPatchGroup_PatchGroup(p *SsmPatchGroupParameters, vals map[string]cty.Value) {
	vals["patch_group"] = cty.StringVal(p.PatchGroup)
}