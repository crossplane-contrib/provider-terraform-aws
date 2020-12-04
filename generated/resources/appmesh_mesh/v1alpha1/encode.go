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

func EncodeAppmeshMesh(r AppmeshMesh) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshMesh_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshMesh_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshMesh_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshMesh_Spec(r.Spec.ForProvider.Spec, ctyVal)
	EncodeAppmeshMesh_ResourceOwner(r.Status.AtProvider, ctyVal)
	EncodeAppmeshMesh_Arn(r.Status.AtProvider, ctyVal)
	EncodeAppmeshMesh_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeAppmeshMesh_LastUpdatedDate(r.Status.AtProvider, ctyVal)
	EncodeAppmeshMesh_MeshOwner(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAppmeshMesh_Name(p AppmeshMeshParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppmeshMesh_Tags(p AppmeshMeshParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAppmeshMesh_Id(p AppmeshMeshParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppmeshMesh_Spec(p Spec, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshMesh_Spec_EgressFilter(p.EgressFilter, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["spec"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshMesh_Spec_EgressFilter(p EgressFilter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshMesh_Spec_EgressFilter_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["egress_filter"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshMesh_Spec_EgressFilter_Type(p EgressFilter, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeAppmeshMesh_ResourceOwner(p AppmeshMeshObservation, vals map[string]cty.Value) {
	vals["resource_owner"] = cty.StringVal(p.ResourceOwner)
}

func EncodeAppmeshMesh_Arn(p AppmeshMeshObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAppmeshMesh_CreatedDate(p AppmeshMeshObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeAppmeshMesh_LastUpdatedDate(p AppmeshMeshObservation, vals map[string]cty.Value) {
	vals["last_updated_date"] = cty.StringVal(p.LastUpdatedDate)
}

func EncodeAppmeshMesh_MeshOwner(p AppmeshMeshObservation, vals map[string]cty.Value) {
	vals["mesh_owner"] = cty.StringVal(p.MeshOwner)
}