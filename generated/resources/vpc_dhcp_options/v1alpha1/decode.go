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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*VpcDhcpOptions)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVpcDhcpOptions(r, ctyValue)
}

func DecodeVpcDhcpOptions(prev *VpcDhcpOptions, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVpcDhcpOptions_DomainNameServers(&new.Spec.ForProvider, valMap)
	DecodeVpcDhcpOptions_Id(&new.Spec.ForProvider, valMap)
	DecodeVpcDhcpOptions_NetbiosNodeType(&new.Spec.ForProvider, valMap)
	DecodeVpcDhcpOptions_NtpServers(&new.Spec.ForProvider, valMap)
	DecodeVpcDhcpOptions_DomainName(&new.Spec.ForProvider, valMap)
	DecodeVpcDhcpOptions_NetbiosNameServers(&new.Spec.ForProvider, valMap)
	DecodeVpcDhcpOptions_Tags(&new.Spec.ForProvider, valMap)
	DecodeVpcDhcpOptions_OwnerId(&new.Status.AtProvider, valMap)
	DecodeVpcDhcpOptions_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeVpcDhcpOptions_DomainNameServers(p *VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["domain_name_servers"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.DomainNameServers = goVals
}

func DecodeVpcDhcpOptions_Id(p *VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeVpcDhcpOptions_NetbiosNodeType(p *VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	p.NetbiosNodeType = ctwhy.ValueAsString(vals["netbios_node_type"])
}

func DecodeVpcDhcpOptions_NtpServers(p *VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["ntp_servers"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.NtpServers = goVals
}

func DecodeVpcDhcpOptions_DomainName(p *VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	p.DomainName = ctwhy.ValueAsString(vals["domain_name"])
}

func DecodeVpcDhcpOptions_NetbiosNameServers(p *VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["netbios_name_servers"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.NetbiosNameServers = goVals
}

func DecodeVpcDhcpOptions_Tags(p *VpcDhcpOptionsParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeVpcDhcpOptions_OwnerId(p *VpcDhcpOptionsObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

func DecodeVpcDhcpOptions_Arn(p *VpcDhcpOptionsObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}