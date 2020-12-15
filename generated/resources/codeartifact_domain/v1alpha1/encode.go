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
	r, ok := mr.(*CodeartifactDomain)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CodeartifactDomain.")
	}
	return EncodeCodeartifactDomain(*r), nil
}

func EncodeCodeartifactDomain(r CodeartifactDomain) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodeartifactDomain_Domain(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactDomain_EncryptionKey(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactDomain_RepositoryCount(r.Status.AtProvider, ctyVal)
	EncodeCodeartifactDomain_Arn(r.Status.AtProvider, ctyVal)
	EncodeCodeartifactDomain_AssetSizeBytes(r.Status.AtProvider, ctyVal)
	EncodeCodeartifactDomain_CreatedTime(r.Status.AtProvider, ctyVal)
	EncodeCodeartifactDomain_Owner(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeCodeartifactDomain_Domain(p CodeartifactDomainParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeCodeartifactDomain_EncryptionKey(p CodeartifactDomainParameters, vals map[string]cty.Value) {
	vals["encryption_key"] = cty.StringVal(p.EncryptionKey)
}

func EncodeCodeartifactDomain_RepositoryCount(p CodeartifactDomainObservation, vals map[string]cty.Value) {
	vals["repository_count"] = cty.NumberIntVal(p.RepositoryCount)
}

func EncodeCodeartifactDomain_Arn(p CodeartifactDomainObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCodeartifactDomain_AssetSizeBytes(p CodeartifactDomainObservation, vals map[string]cty.Value) {
	vals["asset_size_bytes"] = cty.NumberIntVal(p.AssetSizeBytes)
}

func EncodeCodeartifactDomain_CreatedTime(p CodeartifactDomainObservation, vals map[string]cty.Value) {
	vals["created_time"] = cty.StringVal(p.CreatedTime)
}

func EncodeCodeartifactDomain_Owner(p CodeartifactDomainObservation, vals map[string]cty.Value) {
	vals["owner"] = cty.StringVal(p.Owner)
}