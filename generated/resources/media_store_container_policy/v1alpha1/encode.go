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

package v1alpha1func EncodeMediaStoreContainerPolicy(r MediaStoreContainerPolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeMediaStoreContainerPolicy_ContainerName(r.Spec.ForProvider, ctyVal)
	EncodeMediaStoreContainerPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeMediaStoreContainerPolicy_Policy(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeMediaStoreContainerPolicy_ContainerName(p *MediaStoreContainerPolicyParameters, vals map[string]cty.Value) {
	vals["container_name"] = cty.StringVal(p.ContainerName)
}

func EncodeMediaStoreContainerPolicy_Id(p *MediaStoreContainerPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMediaStoreContainerPolicy_Policy(p *MediaStoreContainerPolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}