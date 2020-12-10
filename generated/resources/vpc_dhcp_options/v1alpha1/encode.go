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
	r, ok := mr.(*VpcDhcpOptions)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VpcDhcpOptions.")
	}
	return EncodeVpcDhcpOptions(*r), nil
}

func EncodeVpcDhcpOptions(r VpcDhcpOptions) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcDhcpOptions_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpcDhcpOptions_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeVpcDhcpOptions_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcDhcpOptions_NetbiosNodeType(r.Spec.ForProvider, ctyVal)
	EncodeVpcDhcpOptions_NtpServers(r.Spec.ForProvider, ctyVal)
	EncodeVpcDhcpOptions_DomainNameServers(r.Spec.ForProvider, ctyVal)
	EncodeVpcDhcpOptions_NetbiosNameServers(r.Spec.ForProvider, ctyVal)
	EncodeVpcDhcpOptions_Arn(r.Status.AtProvider, ctyVal)
	EncodeVpcDhcpOptions_OwnerId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeVpcDhcpOptions_Tags(p VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeVpcDhcpOptions_DomainName(p VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeVpcDhcpOptions_Id(p VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcDhcpOptions_NetbiosNodeType(p VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	vals["netbios_node_type"] = cty.StringVal(p.NetbiosNodeType)
}

func EncodeVpcDhcpOptions_NtpServers(p VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NtpServers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ntp_servers"] = cty.ListVal(colVals)
}

func EncodeVpcDhcpOptions_DomainNameServers(p VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DomainNameServers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["domain_name_servers"] = cty.ListVal(colVals)
}

func EncodeVpcDhcpOptions_NetbiosNameServers(p VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NetbiosNameServers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["netbios_name_servers"] = cty.ListVal(colVals)
}

func EncodeVpcDhcpOptions_Arn(p VpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeVpcDhcpOptions_OwnerId(p VpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}