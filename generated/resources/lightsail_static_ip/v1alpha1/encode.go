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
	r, ok := mr.(*LightsailStaticIp)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LightsailStaticIp.")
	}
	return EncodeLightsailStaticIp(*r), nil
}

func EncodeLightsailStaticIp(r LightsailStaticIp) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLightsailStaticIp_Id(r.Spec.ForProvider, ctyVal)
	EncodeLightsailStaticIp_Name(r.Spec.ForProvider, ctyVal)
	EncodeLightsailStaticIp_Arn(r.Status.AtProvider, ctyVal)
	EncodeLightsailStaticIp_IpAddress(r.Status.AtProvider, ctyVal)
	EncodeLightsailStaticIp_SupportCode(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLightsailStaticIp_Id(p LightsailStaticIpParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLightsailStaticIp_Name(p LightsailStaticIpParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLightsailStaticIp_Arn(p LightsailStaticIpObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLightsailStaticIp_IpAddress(p LightsailStaticIpObservation, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}

func EncodeLightsailStaticIp_SupportCode(p LightsailStaticIpObservation, vals map[string]cty.Value) {
	vals["support_code"] = cty.StringVal(p.SupportCode)
}