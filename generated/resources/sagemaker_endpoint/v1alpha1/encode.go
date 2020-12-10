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
	r, ok := mr.(*SagemakerEndpoint)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SagemakerEndpoint.")
	}
	return EncodeSagemakerEndpoint(*r), nil
}

func EncodeSagemakerEndpoint(r SagemakerEndpoint) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSagemakerEndpoint_Name(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerEndpoint_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerEndpoint_EndpointConfigName(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerEndpoint_Id(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerEndpoint_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSagemakerEndpoint_Name(p SagemakerEndpointParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSagemakerEndpoint_Tags(p SagemakerEndpointParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSagemakerEndpoint_EndpointConfigName(p SagemakerEndpointParameters, vals map[string]cty.Value) {
	vals["endpoint_config_name"] = cty.StringVal(p.EndpointConfigName)
}

func EncodeSagemakerEndpoint_Id(p SagemakerEndpointParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSagemakerEndpoint_Arn(p SagemakerEndpointObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}