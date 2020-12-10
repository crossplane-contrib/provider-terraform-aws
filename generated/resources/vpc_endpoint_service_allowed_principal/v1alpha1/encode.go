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
	r, ok := mr.(*VpcEndpointServiceAllowedPrincipal)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VpcEndpointServiceAllowedPrincipal.")
	}
	return EncodeVpcEndpointServiceAllowedPrincipal(*r), nil
}

func EncodeVpcEndpointServiceAllowedPrincipal(r VpcEndpointServiceAllowedPrincipal) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcEndpointServiceAllowedPrincipal_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointServiceAllowedPrincipal_PrincipalArn(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointServiceAllowedPrincipal_VpcEndpointServiceId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeVpcEndpointServiceAllowedPrincipal_Id(p VpcEndpointServiceAllowedPrincipalParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcEndpointServiceAllowedPrincipal_PrincipalArn(p VpcEndpointServiceAllowedPrincipalParameters, vals map[string]cty.Value) {
	vals["principal_arn"] = cty.StringVal(p.PrincipalArn)
}

func EncodeVpcEndpointServiceAllowedPrincipal_VpcEndpointServiceId(p VpcEndpointServiceAllowedPrincipalParameters, vals map[string]cty.Value) {
	vals["vpc_endpoint_service_id"] = cty.StringVal(p.VpcEndpointServiceId)
}