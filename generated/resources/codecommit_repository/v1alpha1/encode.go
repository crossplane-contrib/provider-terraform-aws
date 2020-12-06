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

func EncodeCodecommitRepository(r CodecommitRepository) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodecommitRepository_DefaultBranch(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitRepository_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitRepository_RepositoryName(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitRepository_Description(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitRepository_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitRepository_Arn(r.Status.AtProvider, ctyVal)
	EncodeCodecommitRepository_CloneUrlHttp(r.Status.AtProvider, ctyVal)
	EncodeCodecommitRepository_CloneUrlSsh(r.Status.AtProvider, ctyVal)
	EncodeCodecommitRepository_RepositoryId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCodecommitRepository_DefaultBranch(p CodecommitRepositoryParameters, vals map[string]cty.Value) {
	vals["default_branch"] = cty.StringVal(p.DefaultBranch)
}

func EncodeCodecommitRepository_Id(p CodecommitRepositoryParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodecommitRepository_RepositoryName(p CodecommitRepositoryParameters, vals map[string]cty.Value) {
	vals["repository_name"] = cty.StringVal(p.RepositoryName)
}

func EncodeCodecommitRepository_Description(p CodecommitRepositoryParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeCodecommitRepository_Tags(p CodecommitRepositoryParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCodecommitRepository_Arn(p CodecommitRepositoryObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCodecommitRepository_CloneUrlHttp(p CodecommitRepositoryObservation, vals map[string]cty.Value) {
	vals["clone_url_http"] = cty.StringVal(p.CloneUrlHttp)
}

func EncodeCodecommitRepository_CloneUrlSsh(p CodecommitRepositoryObservation, vals map[string]cty.Value) {
	vals["clone_url_ssh"] = cty.StringVal(p.CloneUrlSsh)
}

func EncodeCodecommitRepository_RepositoryId(p CodecommitRepositoryObservation, vals map[string]cty.Value) {
	vals["repository_id"] = cty.StringVal(p.RepositoryId)
}