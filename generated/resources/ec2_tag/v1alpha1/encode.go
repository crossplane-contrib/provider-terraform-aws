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
	r, ok := mr.(*Ec2Tag)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2Tag.")
	}
	return EncodeEc2Tag(*r), nil
}

func EncodeEc2Tag(r Ec2Tag) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2Tag_Key(r.Spec.ForProvider, ctyVal)
	EncodeEc2Tag_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeEc2Tag_Value(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2Tag_Key(p Ec2TagParameters, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeEc2Tag_ResourceId(p Ec2TagParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeEc2Tag_Value(p Ec2TagParameters, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}