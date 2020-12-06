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
	"github.com/zclconf/go-cty/cty"
)

func EncodeDefaultVpcDhcpOptions(r DefaultVpcDhcpOptions) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDefaultVpcDhcpOptions_Id(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_NetbiosNameServers(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_NetbiosNodeType(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_Arn(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_DomainNameServers(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_NtpServers(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_DomainName(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDefaultVpcDhcpOptions_Id(p DefaultVpcDhcpOptionsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDefaultVpcDhcpOptions_NetbiosNameServers(p DefaultVpcDhcpOptionsParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NetbiosNameServers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["netbios_name_servers"] = cty.ListVal(colVals)
}

func EncodeDefaultVpcDhcpOptions_NetbiosNodeType(p DefaultVpcDhcpOptionsParameters, vals map[string]cty.Value) {
	vals["netbios_node_type"] = cty.StringVal(p.NetbiosNodeType)
}

func EncodeDefaultVpcDhcpOptions_Tags(p DefaultVpcDhcpOptionsParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDefaultVpcDhcpOptions_OwnerId(p DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeDefaultVpcDhcpOptions_Arn(p DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDefaultVpcDhcpOptions_DomainNameServers(p DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["domain_name_servers"] = cty.StringVal(p.DomainNameServers)
}

func EncodeDefaultVpcDhcpOptions_NtpServers(p DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["ntp_servers"] = cty.StringVal(p.NtpServers)
}

func EncodeDefaultVpcDhcpOptions_DomainName(p DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}