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
	r, ok := mr.(*CodecommitRepository)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CodecommitRepository.")
	}
	return EncodeCodecommitRepository(*r), nil
}

func EncodeCodecommitRepository(r CodecommitRepository) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodecommitRepository_DefaultBranch(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitRepository_Description(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitRepository_RepositoryName(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitRepository_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitRepository_CloneUrlHttp(r.Status.AtProvider, ctyVal)
	EncodeCodecommitRepository_Arn(r.Status.AtProvider, ctyVal)
	EncodeCodecommitRepository_CloneUrlSsh(r.Status.AtProvider, ctyVal)
	EncodeCodecommitRepository_RepositoryId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeCodecommitRepository_DefaultBranch(p CodecommitRepositoryParameters, vals map[string]cty.Value) {
	vals["default_branch"] = cty.StringVal(p.DefaultBranch)
}

func EncodeCodecommitRepository_Description(p CodecommitRepositoryParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeCodecommitRepository_RepositoryName(p CodecommitRepositoryParameters, vals map[string]cty.Value) {
	vals["repository_name"] = cty.StringVal(p.RepositoryName)
}

func EncodeCodecommitRepository_Tags(p CodecommitRepositoryParameters, vals map[string]cty.Value) {
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

func EncodeCodecommitRepository_CloneUrlHttp(p CodecommitRepositoryObservation, vals map[string]cty.Value) {
	vals["clone_url_http"] = cty.StringVal(p.CloneUrlHttp)
}

func EncodeCodecommitRepository_Arn(p CodecommitRepositoryObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCodecommitRepository_CloneUrlSsh(p CodecommitRepositoryObservation, vals map[string]cty.Value) {
	vals["clone_url_ssh"] = cty.StringVal(p.CloneUrlSsh)
}

func EncodeCodecommitRepository_RepositoryId(p CodecommitRepositoryObservation, vals map[string]cty.Value) {
	vals["repository_id"] = cty.StringVal(p.RepositoryId)
}