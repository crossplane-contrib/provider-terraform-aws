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

func EncodeCodeartifactRepository(r CodeartifactRepository) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodeartifactRepository_Repository(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactRepository_Description(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactRepository_Domain(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactRepository_DomainOwner(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactRepository_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactRepository_Upstream(r.Spec.ForProvider.Upstream, ctyVal)
	EncodeCodeartifactRepository_AdministratorAccount(r.Status.AtProvider, ctyVal)
	EncodeCodeartifactRepository_Arn(r.Status.AtProvider, ctyVal)
	EncodeCodeartifactRepository_ExternalConnections(r.Status.AtProvider.ExternalConnections, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCodeartifactRepository_Repository(p CodeartifactRepositoryParameters, vals map[string]cty.Value) {
	vals["repository"] = cty.StringVal(p.Repository)
}

func EncodeCodeartifactRepository_Description(p CodeartifactRepositoryParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeCodeartifactRepository_Domain(p CodeartifactRepositoryParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeCodeartifactRepository_DomainOwner(p CodeartifactRepositoryParameters, vals map[string]cty.Value) {
	vals["domain_owner"] = cty.StringVal(p.DomainOwner)
}

func EncodeCodeartifactRepository_Id(p CodeartifactRepositoryParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodeartifactRepository_Upstream(p Upstream, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodeartifactRepository_Upstream_RepositoryName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["upstream"] = cty.ListVal(valsForCollection)
}

func EncodeCodeartifactRepository_Upstream_RepositoryName(p Upstream, vals map[string]cty.Value) {
	vals["repository_name"] = cty.StringVal(p.RepositoryName)
}

func EncodeCodeartifactRepository_AdministratorAccount(p CodeartifactRepositoryObservation, vals map[string]cty.Value) {
	vals["administrator_account"] = cty.StringVal(p.AdministratorAccount)
}

func EncodeCodeartifactRepository_Arn(p CodeartifactRepositoryObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCodeartifactRepository_ExternalConnections(p []ExternalConnections, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeCodeartifactRepository_ExternalConnections_PackageFormat(v, ctyVal)
		EncodeCodeartifactRepository_ExternalConnections_Status(v, ctyVal)
		EncodeCodeartifactRepository_ExternalConnections_ExternalConnectionName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["external_connections"] = cty.ListVal(valsForCollection)
}

func EncodeCodeartifactRepository_ExternalConnections_PackageFormat(p ExternalConnections, vals map[string]cty.Value) {
	vals["package_format"] = cty.StringVal(p.PackageFormat)
}

func EncodeCodeartifactRepository_ExternalConnections_Status(p ExternalConnections, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeCodeartifactRepository_ExternalConnections_ExternalConnectionName(p ExternalConnections, vals map[string]cty.Value) {
	vals["external_connection_name"] = cty.StringVal(p.ExternalConnectionName)
}