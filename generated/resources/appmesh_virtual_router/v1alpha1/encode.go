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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*AppmeshVirtualRouter)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a AppmeshVirtualRouter.")
	}
	return EncodeAppmeshVirtualRouter(*r), nil
}

func EncodeAppmeshVirtualRouter(r AppmeshVirtualRouter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualRouter_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualRouter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualRouter_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualRouter_MeshName(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualRouter_MeshOwner(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshVirtualRouter_Spec(r.Spec.ForProvider.Spec, ctyVal)
	EncodeAppmeshVirtualRouter_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeAppmeshVirtualRouter_ResourceOwner(r.Status.AtProvider, ctyVal)
	EncodeAppmeshVirtualRouter_Arn(r.Status.AtProvider, ctyVal)
	EncodeAppmeshVirtualRouter_LastUpdatedDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAppmeshVirtualRouter_Name(p AppmeshVirtualRouterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppmeshVirtualRouter_Tags(p AppmeshVirtualRouterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAppmeshVirtualRouter_Id(p AppmeshVirtualRouterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppmeshVirtualRouter_MeshName(p AppmeshVirtualRouterParameters, vals map[string]cty.Value) {
	vals["mesh_name"] = cty.StringVal(p.MeshName)
}

func EncodeAppmeshVirtualRouter_MeshOwner(p AppmeshVirtualRouterParameters, vals map[string]cty.Value) {
	vals["mesh_owner"] = cty.StringVal(p.MeshOwner)
}

func EncodeAppmeshVirtualRouter_Spec(p Spec, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualRouter_Spec_Listener(p.Listener, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["spec"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualRouter_Spec_Listener(p Listener, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualRouter_Spec_Listener_PortMapping(p.PortMapping, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["listener"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualRouter_Spec_Listener_PortMapping(p PortMapping, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshVirtualRouter_Spec_Listener_PortMapping_Port(p, ctyVal)
	EncodeAppmeshVirtualRouter_Spec_Listener_PortMapping_Protocol(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["port_mapping"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshVirtualRouter_Spec_Listener_PortMapping_Port(p PortMapping, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeAppmeshVirtualRouter_Spec_Listener_PortMapping_Protocol(p PortMapping, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeAppmeshVirtualRouter_CreatedDate(p AppmeshVirtualRouterObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeAppmeshVirtualRouter_ResourceOwner(p AppmeshVirtualRouterObservation, vals map[string]cty.Value) {
	vals["resource_owner"] = cty.StringVal(p.ResourceOwner)
}

func EncodeAppmeshVirtualRouter_Arn(p AppmeshVirtualRouterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAppmeshVirtualRouter_LastUpdatedDate(p AppmeshVirtualRouterObservation, vals map[string]cty.Value) {
	vals["last_updated_date"] = cty.StringVal(p.LastUpdatedDate)
}