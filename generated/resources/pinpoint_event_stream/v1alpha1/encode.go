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

package v1alpha1func EncodePinpointEventStream(r PinpointEventStream) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodePinpointEventStream_Id(r.Spec.ForProvider, ctyVal)
	EncodePinpointEventStream_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodePinpointEventStream_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointEventStream_DestinationStreamArn(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodePinpointEventStream_Id(p *PinpointEventStreamParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodePinpointEventStream_RoleArn(p *PinpointEventStreamParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodePinpointEventStream_ApplicationId(p *PinpointEventStreamParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointEventStream_DestinationStreamArn(p *PinpointEventStreamParameters, vals map[string]cty.Value) {
	vals["destination_stream_arn"] = cty.StringVal(p.DestinationStreamArn)
}