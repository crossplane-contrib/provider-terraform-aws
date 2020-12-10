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
	r, ok := mr.(*LambdaEventSourceMapping)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LambdaEventSourceMapping.")
	}
	return EncodeLambdaEventSourceMapping(*r), nil
}

func EncodeLambdaEventSourceMapping(r LambdaEventSourceMapping) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaEventSourceMapping_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_EventSourceArn(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_FunctionName(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_Id(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_StartingPosition(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_StartingPositionTimestamp(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_BisectBatchOnFunctionError(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_MaximumBatchingWindowInSeconds(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_MaximumRetryAttempts(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_ParallelizationFactor(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_BatchSize(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_MaximumRecordAgeInSeconds(r.Spec.ForProvider, ctyVal)
	EncodeLambdaEventSourceMapping_DestinationConfig(r.Spec.ForProvider.DestinationConfig, ctyVal)
	EncodeLambdaEventSourceMapping_LastModified(r.Status.AtProvider, ctyVal)
	EncodeLambdaEventSourceMapping_LastProcessingResult(r.Status.AtProvider, ctyVal)
	EncodeLambdaEventSourceMapping_StateTransitionReason(r.Status.AtProvider, ctyVal)
	EncodeLambdaEventSourceMapping_FunctionArn(r.Status.AtProvider, ctyVal)
	EncodeLambdaEventSourceMapping_State(r.Status.AtProvider, ctyVal)
	EncodeLambdaEventSourceMapping_Uuid(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLambdaEventSourceMapping_Enabled(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeLambdaEventSourceMapping_EventSourceArn(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["event_source_arn"] = cty.StringVal(p.EventSourceArn)
}

func EncodeLambdaEventSourceMapping_FunctionName(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["function_name"] = cty.StringVal(p.FunctionName)
}

func EncodeLambdaEventSourceMapping_Id(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLambdaEventSourceMapping_StartingPosition(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["starting_position"] = cty.StringVal(p.StartingPosition)
}

func EncodeLambdaEventSourceMapping_StartingPositionTimestamp(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["starting_position_timestamp"] = cty.StringVal(p.StartingPositionTimestamp)
}

func EncodeLambdaEventSourceMapping_BisectBatchOnFunctionError(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["bisect_batch_on_function_error"] = cty.BoolVal(p.BisectBatchOnFunctionError)
}

func EncodeLambdaEventSourceMapping_MaximumBatchingWindowInSeconds(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["maximum_batching_window_in_seconds"] = cty.NumberIntVal(p.MaximumBatchingWindowInSeconds)
}

func EncodeLambdaEventSourceMapping_MaximumRetryAttempts(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["maximum_retry_attempts"] = cty.NumberIntVal(p.MaximumRetryAttempts)
}

func EncodeLambdaEventSourceMapping_ParallelizationFactor(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["parallelization_factor"] = cty.NumberIntVal(p.ParallelizationFactor)
}

func EncodeLambdaEventSourceMapping_BatchSize(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["batch_size"] = cty.NumberIntVal(p.BatchSize)
}

func EncodeLambdaEventSourceMapping_MaximumRecordAgeInSeconds(p LambdaEventSourceMappingParameters, vals map[string]cty.Value) {
	vals["maximum_record_age_in_seconds"] = cty.NumberIntVal(p.MaximumRecordAgeInSeconds)
}

func EncodeLambdaEventSourceMapping_DestinationConfig(p DestinationConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaEventSourceMapping_DestinationConfig_OnFailure(p.OnFailure, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["destination_config"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaEventSourceMapping_DestinationConfig_OnFailure(p OnFailure, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaEventSourceMapping_DestinationConfig_OnFailure_DestinationArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["on_failure"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaEventSourceMapping_DestinationConfig_OnFailure_DestinationArn(p OnFailure, vals map[string]cty.Value) {
	vals["destination_arn"] = cty.StringVal(p.DestinationArn)
}

func EncodeLambdaEventSourceMapping_LastModified(p LambdaEventSourceMappingObservation, vals map[string]cty.Value) {
	vals["last_modified"] = cty.StringVal(p.LastModified)
}

func EncodeLambdaEventSourceMapping_LastProcessingResult(p LambdaEventSourceMappingObservation, vals map[string]cty.Value) {
	vals["last_processing_result"] = cty.StringVal(p.LastProcessingResult)
}

func EncodeLambdaEventSourceMapping_StateTransitionReason(p LambdaEventSourceMappingObservation, vals map[string]cty.Value) {
	vals["state_transition_reason"] = cty.StringVal(p.StateTransitionReason)
}

func EncodeLambdaEventSourceMapping_FunctionArn(p LambdaEventSourceMappingObservation, vals map[string]cty.Value) {
	vals["function_arn"] = cty.StringVal(p.FunctionArn)
}

func EncodeLambdaEventSourceMapping_State(p LambdaEventSourceMappingObservation, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeLambdaEventSourceMapping_Uuid(p LambdaEventSourceMappingObservation, vals map[string]cty.Value) {
	vals["uuid"] = cty.StringVal(p.Uuid)
}