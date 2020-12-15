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
	r, ok := mr.(*OpsworksUserProfile)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a OpsworksUserProfile.")
	}
	return EncodeOpsworksUserProfile(*r), nil
}

func EncodeOpsworksUserProfile(r OpsworksUserProfile) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksUserProfile_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksUserProfile_SshPublicKey(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksUserProfile_SshUsername(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksUserProfile_UserArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksUserProfile_AllowSelfManagement(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksUserProfile_Id(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksUserProfile_SshPublicKey(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["ssh_public_key"] = cty.StringVal(p.SshPublicKey)
}

func EncodeOpsworksUserProfile_SshUsername(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["ssh_username"] = cty.StringVal(p.SshUsername)
}

func EncodeOpsworksUserProfile_UserArn(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["user_arn"] = cty.StringVal(p.UserArn)
}

func EncodeOpsworksUserProfile_AllowSelfManagement(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["allow_self_management"] = cty.BoolVal(p.AllowSelfManagement)
}