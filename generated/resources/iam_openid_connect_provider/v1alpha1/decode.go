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
	r, ok := mr.(*IamOpenidConnectProvider)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamOpenidConnectProvider(r, ctyValue)
}

func DecodeIamOpenidConnectProvider(prev *IamOpenidConnectProvider, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamOpenidConnectProvider_ClientIdList(&new.Spec.ForProvider, valMap)
	DecodeIamOpenidConnectProvider_Id(&new.Spec.ForProvider, valMap)
	DecodeIamOpenidConnectProvider_ThumbprintList(&new.Spec.ForProvider, valMap)
	DecodeIamOpenidConnectProvider_Url(&new.Spec.ForProvider, valMap)
	DecodeIamOpenidConnectProvider_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeIamOpenidConnectProvider_ClientIdList(p *IamOpenidConnectProviderParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["client_id_list"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ClientIdList = goVals
}

//primitiveTypeDecodeTemplate
func DecodeIamOpenidConnectProvider_Id(p *IamOpenidConnectProviderParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeIamOpenidConnectProvider_ThumbprintList(p *IamOpenidConnectProviderParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["thumbprint_list"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ThumbprintList = goVals
}

//primitiveTypeDecodeTemplate
func DecodeIamOpenidConnectProvider_Url(p *IamOpenidConnectProviderParameters, vals map[string]cty.Value) {
	p.Url = ctwhy.ValueAsString(vals["url"])
}

//primitiveTypeDecodeTemplate
func DecodeIamOpenidConnectProvider_Arn(p *IamOpenidConnectProviderObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}