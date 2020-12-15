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
	r, ok := mr.(*WafregionalWebAclAssociation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeWafregionalWebAclAssociation(r, ctyValue)
}

func DecodeWafregionalWebAclAssociation(prev *WafregionalWebAclAssociation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeWafregionalWebAclAssociation_ResourceArn(&new.Spec.ForProvider, valMap)
	DecodeWafregionalWebAclAssociation_WebAclId(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeWafregionalWebAclAssociation_ResourceArn(p *WafregionalWebAclAssociationParameters, vals map[string]cty.Value) {
	p.ResourceArn = ctwhy.ValueAsString(vals["resource_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeWafregionalWebAclAssociation_WebAclId(p *WafregionalWebAclAssociationParameters, vals map[string]cty.Value) {
	p.WebAclId = ctwhy.ValueAsString(vals["web_acl_id"])
}