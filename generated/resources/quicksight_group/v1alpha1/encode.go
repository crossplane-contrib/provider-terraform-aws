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
	r, ok := mr.(*QuicksightGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a QuicksightGroup.")
	}
	return EncodeQuicksightGroup(*r), nil
}

func EncodeQuicksightGroup(r QuicksightGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeQuicksightGroup_AwsAccountId(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightGroup_GroupName(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightGroup_Namespace(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightGroup_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeQuicksightGroup_AwsAccountId(p QuicksightGroupParameters, vals map[string]cty.Value) {
	vals["aws_account_id"] = cty.StringVal(p.AwsAccountId)
}

func EncodeQuicksightGroup_Description(p QuicksightGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeQuicksightGroup_GroupName(p QuicksightGroupParameters, vals map[string]cty.Value) {
	vals["group_name"] = cty.StringVal(p.GroupName)
}

func EncodeQuicksightGroup_Namespace(p QuicksightGroupParameters, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeQuicksightGroup_Arn(p QuicksightGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}