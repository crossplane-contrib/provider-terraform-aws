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
	r, ok := mr.(*GameliftGameSessionQueue)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GameliftGameSessionQueue.")
	}
	return EncodeGameliftGameSessionQueue(*r), nil
}

func EncodeGameliftGameSessionQueue(r GameliftGameSessionQueue) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftGameSessionQueue_Destinations(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_Id(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_Name(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_TimeoutInSeconds(r.Spec.ForProvider, ctyVal)
	EncodeGameliftGameSessionQueue_PlayerLatencyPolicy(r.Spec.ForProvider.PlayerLatencyPolicy, ctyVal)
	EncodeGameliftGameSessionQueue_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGameliftGameSessionQueue_Destinations(p GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Destinations {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["destinations"] = cty.ListVal(colVals)
}

func EncodeGameliftGameSessionQueue_Id(p GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGameliftGameSessionQueue_Name(p GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGameliftGameSessionQueue_Tags(p GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGameliftGameSessionQueue_TimeoutInSeconds(p GameliftGameSessionQueueParameters, vals map[string]cty.Value) {
	vals["timeout_in_seconds"] = cty.NumberIntVal(p.TimeoutInSeconds)
}

func EncodeGameliftGameSessionQueue_PlayerLatencyPolicy(p PlayerLatencyPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftGameSessionQueue_PlayerLatencyPolicy_MaximumIndividualPlayerLatencyMilliseconds(p, ctyVal)
	EncodeGameliftGameSessionQueue_PlayerLatencyPolicy_PolicyDurationSeconds(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["player_latency_policy"] = cty.ListVal(valsForCollection)
}

func EncodeGameliftGameSessionQueue_PlayerLatencyPolicy_MaximumIndividualPlayerLatencyMilliseconds(p PlayerLatencyPolicy, vals map[string]cty.Value) {
	vals["maximum_individual_player_latency_milliseconds"] = cty.NumberIntVal(p.MaximumIndividualPlayerLatencyMilliseconds)
}

func EncodeGameliftGameSessionQueue_PlayerLatencyPolicy_PolicyDurationSeconds(p PlayerLatencyPolicy, vals map[string]cty.Value) {
	vals["policy_duration_seconds"] = cty.NumberIntVal(p.PolicyDurationSeconds)
}

func EncodeGameliftGameSessionQueue_Arn(p GameliftGameSessionQueueObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}