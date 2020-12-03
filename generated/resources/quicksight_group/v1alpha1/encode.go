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

package v1alpha1func EncodeQuicksightGroup(r QuicksightGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeQuicksightGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightGroup_GroupName(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightGroup_Namespace(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightGroup_AwsAccountId(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeQuicksightGroup_Description(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeQuicksightGroup_GroupName(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	vals["group_name"] = cty.StringVal(p.GroupName)
}

func EncodeQuicksightGroup_Id(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeQuicksightGroup_Namespace(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeQuicksightGroup_AwsAccountId(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	vals["aws_account_id"] = cty.StringVal(p.AwsAccountId)
}

func EncodeQuicksightGroup_Arn(p *QuicksightGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}