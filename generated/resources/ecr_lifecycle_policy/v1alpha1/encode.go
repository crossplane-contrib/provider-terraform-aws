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

package v1alpha1func EncodeEcrLifecyclePolicy(r EcrLifecyclePolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEcrLifecyclePolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeEcrLifecyclePolicy_Policy(r.Spec.ForProvider, ctyVal)
	EncodeEcrLifecyclePolicy_Repository(r.Spec.ForProvider, ctyVal)
	EncodeEcrLifecyclePolicy_RegistryId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEcrLifecyclePolicy_Id(p *EcrLifecyclePolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEcrLifecyclePolicy_Policy(p *EcrLifecyclePolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeEcrLifecyclePolicy_Repository(p *EcrLifecyclePolicyParameters, vals map[string]cty.Value) {
	vals["repository"] = cty.StringVal(p.Repository)
}

func EncodeEcrLifecyclePolicy_RegistryId(p *EcrLifecyclePolicyObservation, vals map[string]cty.Value) {
	vals["registry_id"] = cty.StringVal(p.RegistryId)
}