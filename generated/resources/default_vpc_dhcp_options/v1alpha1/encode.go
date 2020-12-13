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
	r, ok := mr.(*DefaultVpcDhcpOptions)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DefaultVpcDhcpOptions.")
	}
	return EncodeDefaultVpcDhcpOptions(*r), nil
}

func EncodeDefaultVpcDhcpOptions(r DefaultVpcDhcpOptions) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDefaultVpcDhcpOptions_Id(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_NetbiosNameServers(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_NetbiosNodeType(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_Arn(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_DomainName(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_DomainNameServers(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpcDhcpOptions_NtpServers(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
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
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
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

func EncodeDefaultVpcDhcpOptions_DomainName(p DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeDefaultVpcDhcpOptions_DomainNameServers(p DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["domain_name_servers"] = cty.StringVal(p.DomainNameServers)
}

func EncodeDefaultVpcDhcpOptions_NtpServers(p DefaultVpcDhcpOptionsObservation, vals map[string]cty.Value) {
	vals["ntp_servers"] = cty.StringVal(p.NtpServers)
}