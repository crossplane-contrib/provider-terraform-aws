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
	r, ok := mr.(*DxConnection)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxConnection(r, ctyValue)
}

func DecodeDxConnection(prev *DxConnection, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxConnection_Name(&new.Spec.ForProvider, valMap)
	DecodeDxConnection_Location(&new.Spec.ForProvider, valMap)
	DecodeDxConnection_Bandwidth(&new.Spec.ForProvider, valMap)
	DecodeDxConnection_Tags(&new.Spec.ForProvider, valMap)
	DecodeDxConnection_HasLogicalRedundancy(&new.Status.AtProvider, valMap)
	DecodeDxConnection_JumboFrameCapable(&new.Status.AtProvider, valMap)
	DecodeDxConnection_Arn(&new.Status.AtProvider, valMap)
	DecodeDxConnection_AwsDevice(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDxConnection_Name(p *DxConnectionParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeDxConnection_Location(p *DxConnectionParameters, vals map[string]cty.Value) {
	p.Location = ctwhy.ValueAsString(vals["location"])
}

//primitiveTypeDecodeTemplate
func DecodeDxConnection_Bandwidth(p *DxConnectionParameters, vals map[string]cty.Value) {
	p.Bandwidth = ctwhy.ValueAsString(vals["bandwidth"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDxConnection_Tags(p *DxConnectionParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeDxConnection_HasLogicalRedundancy(p *DxConnectionObservation, vals map[string]cty.Value) {
	p.HasLogicalRedundancy = ctwhy.ValueAsString(vals["has_logical_redundancy"])
}

//primitiveTypeDecodeTemplate
func DecodeDxConnection_JumboFrameCapable(p *DxConnectionObservation, vals map[string]cty.Value) {
	p.JumboFrameCapable = ctwhy.ValueAsBool(vals["jumbo_frame_capable"])
}

//primitiveTypeDecodeTemplate
func DecodeDxConnection_Arn(p *DxConnectionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxConnection_AwsDevice(p *DxConnectionObservation, vals map[string]cty.Value) {
	p.AwsDevice = ctwhy.ValueAsString(vals["aws_device"])
}