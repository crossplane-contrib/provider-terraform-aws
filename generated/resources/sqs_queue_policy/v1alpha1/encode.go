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
	
	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*SqsQueuePolicy)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SqsQueuePolicy.")
	}
	return EncodeSqsQueuePolicy(*r), nil
}

func EncodeSqsQueuePolicy(r SqsQueuePolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSqsQueuePolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueuePolicy_Policy(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueuePolicy_QueueUrl(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeSqsQueuePolicy_Id(p SqsQueuePolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSqsQueuePolicy_Policy(p SqsQueuePolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeSqsQueuePolicy_QueueUrl(p SqsQueuePolicyParameters, vals map[string]cty.Value) {
	vals["queue_url"] = cty.StringVal(p.QueueUrl)
}