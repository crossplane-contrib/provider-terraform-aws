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
	r, ok := mr.(*LightsailDomain)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LightsailDomain.")
	}
	return EncodeLightsailDomain(*r), nil
}

func EncodeLightsailDomain(r LightsailDomain) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLightsailDomain_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeLightsailDomain_Id(r.Spec.ForProvider, ctyVal)
	EncodeLightsailDomain_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLightsailDomain_DomainName(p LightsailDomainParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeLightsailDomain_Id(p LightsailDomainParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLightsailDomain_Arn(p LightsailDomainObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}