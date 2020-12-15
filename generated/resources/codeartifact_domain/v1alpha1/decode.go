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
	r, ok := mr.(*CodeartifactDomain)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCodeartifactDomain(r, ctyValue)
}

func DecodeCodeartifactDomain(prev *CodeartifactDomain, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCodeartifactDomain_Domain(&new.Spec.ForProvider, valMap)
	DecodeCodeartifactDomain_EncryptionKey(&new.Spec.ForProvider, valMap)
	DecodeCodeartifactDomain_Id(&new.Spec.ForProvider, valMap)
	DecodeCodeartifactDomain_Owner(&new.Status.AtProvider, valMap)
	DecodeCodeartifactDomain_RepositoryCount(&new.Status.AtProvider, valMap)
	DecodeCodeartifactDomain_Arn(&new.Status.AtProvider, valMap)
	DecodeCodeartifactDomain_AssetSizeBytes(&new.Status.AtProvider, valMap)
	DecodeCodeartifactDomain_CreatedTime(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomain_Domain(p *CodeartifactDomainParameters, vals map[string]cty.Value) {
	p.Domain = ctwhy.ValueAsString(vals["domain"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomain_EncryptionKey(p *CodeartifactDomainParameters, vals map[string]cty.Value) {
	p.EncryptionKey = ctwhy.ValueAsString(vals["encryption_key"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomain_Id(p *CodeartifactDomainParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomain_Owner(p *CodeartifactDomainObservation, vals map[string]cty.Value) {
	p.Owner = ctwhy.ValueAsString(vals["owner"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomain_RepositoryCount(p *CodeartifactDomainObservation, vals map[string]cty.Value) {
	p.RepositoryCount = ctwhy.ValueAsInt64(vals["repository_count"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomain_Arn(p *CodeartifactDomainObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomain_AssetSizeBytes(p *CodeartifactDomainObservation, vals map[string]cty.Value) {
	p.AssetSizeBytes = ctwhy.ValueAsInt64(vals["asset_size_bytes"])
}

//primitiveTypeDecodeTemplate
func DecodeCodeartifactDomain_CreatedTime(p *CodeartifactDomainObservation, vals map[string]cty.Value) {
	p.CreatedTime = ctwhy.ValueAsString(vals["created_time"])
}