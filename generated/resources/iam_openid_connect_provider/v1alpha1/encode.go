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
	r, ok := mr.(*IamOpenidConnectProvider)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IamOpenidConnectProvider.")
	}
	return EncodeIamOpenidConnectProvider(*r), nil
}

func EncodeIamOpenidConnectProvider(r IamOpenidConnectProvider) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamOpenidConnectProvider_ClientIdList(r.Spec.ForProvider, ctyVal)
	EncodeIamOpenidConnectProvider_ThumbprintList(r.Spec.ForProvider, ctyVal)
	EncodeIamOpenidConnectProvider_Url(r.Spec.ForProvider, ctyVal)
	EncodeIamOpenidConnectProvider_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeIamOpenidConnectProvider_ClientIdList(p IamOpenidConnectProviderParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ClientIdList {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["client_id_list"] = cty.ListVal(colVals)
}

func EncodeIamOpenidConnectProvider_ThumbprintList(p IamOpenidConnectProviderParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ThumbprintList {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["thumbprint_list"] = cty.ListVal(colVals)
}

func EncodeIamOpenidConnectProvider_Url(p IamOpenidConnectProviderParameters, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}

func EncodeIamOpenidConnectProvider_Arn(p IamOpenidConnectProviderObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}