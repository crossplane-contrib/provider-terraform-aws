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
	r, ok := mr.(*CloudfrontPublicKey)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCloudfrontPublicKey(r, ctyValue)
}

func DecodeCloudfrontPublicKey(prev *CloudfrontPublicKey, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCloudfrontPublicKey_Id(&new.Spec.ForProvider, valMap)
	DecodeCloudfrontPublicKey_Name(&new.Spec.ForProvider, valMap)
	DecodeCloudfrontPublicKey_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeCloudfrontPublicKey_Comment(&new.Spec.ForProvider, valMap)
	DecodeCloudfrontPublicKey_EncodedKey(&new.Spec.ForProvider, valMap)
	DecodeCloudfrontPublicKey_CallerReference(&new.Status.AtProvider, valMap)
	DecodeCloudfrontPublicKey_Etag(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeCloudfrontPublicKey_Id(p *CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeCloudfrontPublicKey_Name(p *CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeCloudfrontPublicKey_NamePrefix(p *CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

func DecodeCloudfrontPublicKey_Comment(p *CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	p.Comment = ctwhy.ValueAsString(vals["comment"])
}

func DecodeCloudfrontPublicKey_EncodedKey(p *CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	p.EncodedKey = ctwhy.ValueAsString(vals["encoded_key"])
}

func DecodeCloudfrontPublicKey_CallerReference(p *CloudfrontPublicKeyObservation, vals map[string]cty.Value) {
	p.CallerReference = ctwhy.ValueAsString(vals["caller_reference"])
}

func DecodeCloudfrontPublicKey_Etag(p *CloudfrontPublicKeyObservation, vals map[string]cty.Value) {
	p.Etag = ctwhy.ValueAsString(vals["etag"])
}