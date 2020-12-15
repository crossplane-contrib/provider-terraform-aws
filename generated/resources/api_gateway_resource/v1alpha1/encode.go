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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*ApiGatewayResource)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayResource.")
	}
	return EncodeApiGatewayResource(*r), nil
}

func EncodeApiGatewayResource(r ApiGatewayResource) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayResource_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayResource_ParentId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayResource_PathPart(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayResource_Path(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayResource_RestApiId(p ApiGatewayResourceParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayResource_ParentId(p ApiGatewayResourceParameters, vals map[string]cty.Value) {
	vals["parent_id"] = cty.StringVal(p.ParentId)
}

func EncodeApiGatewayResource_PathPart(p ApiGatewayResourceParameters, vals map[string]cty.Value) {
	vals["path_part"] = cty.StringVal(p.PathPart)
}

func EncodeApiGatewayResource_Path(p ApiGatewayResourceObservation, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}