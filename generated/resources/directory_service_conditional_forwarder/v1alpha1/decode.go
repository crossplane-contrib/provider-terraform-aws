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
	r, ok := mr.(*DirectoryServiceConditionalForwarder)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDirectoryServiceConditionalForwarder(r, ctyValue)
}

func DecodeDirectoryServiceConditionalForwarder(prev *DirectoryServiceConditionalForwarder, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDirectoryServiceConditionalForwarder_RemoteDomainName(&new.Spec.ForProvider, valMap)
	DecodeDirectoryServiceConditionalForwarder_DirectoryId(&new.Spec.ForProvider, valMap)
	DecodeDirectoryServiceConditionalForwarder_DnsIps(&new.Spec.ForProvider, valMap)
	DecodeDirectoryServiceConditionalForwarder_Id(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDirectoryServiceConditionalForwarder_RemoteDomainName(p *DirectoryServiceConditionalForwarderParameters, vals map[string]cty.Value) {
	p.RemoteDomainName = ctwhy.ValueAsString(vals["remote_domain_name"])
}

func DecodeDirectoryServiceConditionalForwarder_DirectoryId(p *DirectoryServiceConditionalForwarderParameters, vals map[string]cty.Value) {
	p.DirectoryId = ctwhy.ValueAsString(vals["directory_id"])
}

func DecodeDirectoryServiceConditionalForwarder_DnsIps(p *DirectoryServiceConditionalForwarderParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["dns_ips"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.DnsIps = goVals
}

func DecodeDirectoryServiceConditionalForwarder_Id(p *DirectoryServiceConditionalForwarderParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}