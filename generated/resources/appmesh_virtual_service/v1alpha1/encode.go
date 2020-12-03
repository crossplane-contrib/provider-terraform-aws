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

package v1alpha1func EncodeAppmeshVirtualService(r AppmeshVirtualService) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAppmeshVirtualService_MeshOwner(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualService_MeshName(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualService_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualService_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualService_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualService_Spec(r.Spec.ForProvider.Spec, ctyVal)
	EncodeAppmeshVirtualService_ResourceOwner(r.Status.AtProvider, ctyVal)
	EncodeAppmeshVirtualService_LastUpdatedDate(r.Status.AtProvider, ctyVal)
	EncodeAppmeshVirtualService_Arn(r.Status.AtProvider, ctyVal)
	EncodeAppmeshVirtualService_CreatedDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeAppmeshVirtualService_MeshOwner(p *AppmeshVirtualServiceParameters, vals map[string]cty.Value) {
	vals["mesh_owner"] = cty.StringVal(p.MeshOwner)
}

func EncodeAppmeshVirtualService_MeshName(p *AppmeshVirtualServiceParameters, vals map[string]cty.Value) {
	vals["mesh_name"] = cty.StringVal(p.MeshName)
}

func EncodeAppmeshVirtualService_Name(p *AppmeshVirtualServiceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppmeshVirtualService_Tags(p *AppmeshVirtualServiceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAppmeshVirtualService_Id(p *AppmeshVirtualServiceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppmeshVirtualService_Spec(p *Spec, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Spec {
		ctyVal = make(map[string]cty.Value)
		EncodeAppmeshVirtualService_Spec_Provider(v.Provider, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["spec"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualService_Spec_Provider(p *Provider, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Provider {
		ctyVal = make(map[string]cty.Value)
		EncodeAppmeshVirtualService_Spec_Provider_VirtualNode(v.VirtualNode, ctyVal)
		EncodeAppmeshVirtualService_Spec_Provider_VirtualRouter(v.VirtualRouter, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["provider"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualService_Spec_Provider_VirtualNode(p *VirtualNode, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VirtualNode {
		ctyVal = make(map[string]cty.Value)
		EncodeAppmeshVirtualService_Spec_Provider_VirtualNode_VirtualNodeName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["virtual_node"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualService_Spec_Provider_VirtualNode_VirtualNodeName(p *VirtualNode, vals map[string]cty.Value) {
	vals["virtual_node_name"] = cty.StringVal(p.VirtualNodeName)
}

func EncodeAppmeshVirtualService_Spec_Provider_VirtualRouter(p *VirtualRouter, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VirtualRouter {
		ctyVal = make(map[string]cty.Value)
		EncodeAppmeshVirtualService_Spec_Provider_VirtualRouter_VirtualRouterName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["virtual_router"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualService_Spec_Provider_VirtualRouter_VirtualRouterName(p *VirtualRouter, vals map[string]cty.Value) {
	vals["virtual_router_name"] = cty.StringVal(p.VirtualRouterName)
}

func EncodeAppmeshVirtualService_ResourceOwner(p *AppmeshVirtualServiceObservation, vals map[string]cty.Value) {
	vals["resource_owner"] = cty.StringVal(p.ResourceOwner)
}

func EncodeAppmeshVirtualService_LastUpdatedDate(p *AppmeshVirtualServiceObservation, vals map[string]cty.Value) {
	vals["last_updated_date"] = cty.StringVal(p.LastUpdatedDate)
}

func EncodeAppmeshVirtualService_Arn(p *AppmeshVirtualServiceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAppmeshVirtualService_CreatedDate(p *AppmeshVirtualServiceObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}