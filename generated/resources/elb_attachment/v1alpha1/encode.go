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
	r, ok := mr.(*ElbAttachment)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ElbAttachment.")
	}
	return EncodeElbAttachment(*r), nil
}

func EncodeElbAttachment(r ElbAttachment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeElbAttachment_Elb(r.Spec.ForProvider, ctyVal)
	EncodeElbAttachment_Instance(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeElbAttachment_Elb(p ElbAttachmentParameters, vals map[string]cty.Value) {
	vals["elb"] = cty.StringVal(p.Elb)
}

func EncodeElbAttachment_Instance(p ElbAttachmentParameters, vals map[string]cty.Value) {
	vals["instance"] = cty.StringVal(p.Instance)
}