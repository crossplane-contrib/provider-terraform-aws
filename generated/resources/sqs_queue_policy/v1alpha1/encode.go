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

package v1alpha1func EncodeSqsQueuePolicy(r SqsQueuePolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSqsQueuePolicy_QueueUrl(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueuePolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueuePolicy_Policy(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeSqsQueuePolicy_QueueUrl(p *SqsQueuePolicyParameters, vals map[string]cty.Value) {
	vals["queue_url"] = cty.StringVal(p.QueueUrl)
}

func EncodeSqsQueuePolicy_Id(p *SqsQueuePolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSqsQueuePolicy_Policy(p *SqsQueuePolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}