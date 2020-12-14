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
	r, ok := mr.(*OpsworksUserProfile)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeOpsworksUserProfile(r, ctyValue)
}

func DecodeOpsworksUserProfile(prev *OpsworksUserProfile, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeOpsworksUserProfile_AllowSelfManagement(&new.Spec.ForProvider, valMap)
	DecodeOpsworksUserProfile_Id(&new.Spec.ForProvider, valMap)
	DecodeOpsworksUserProfile_SshPublicKey(&new.Spec.ForProvider, valMap)
	DecodeOpsworksUserProfile_SshUsername(&new.Spec.ForProvider, valMap)
	DecodeOpsworksUserProfile_UserArn(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeOpsworksUserProfile_AllowSelfManagement(p *OpsworksUserProfileParameters, vals map[string]cty.Value) {
	p.AllowSelfManagement = ctwhy.ValueAsBool(vals["allow_self_management"])
}

func DecodeOpsworksUserProfile_Id(p *OpsworksUserProfileParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeOpsworksUserProfile_SshPublicKey(p *OpsworksUserProfileParameters, vals map[string]cty.Value) {
	p.SshPublicKey = ctwhy.ValueAsString(vals["ssh_public_key"])
}

func DecodeOpsworksUserProfile_SshUsername(p *OpsworksUserProfileParameters, vals map[string]cty.Value) {
	p.SshUsername = ctwhy.ValueAsString(vals["ssh_username"])
}

func DecodeOpsworksUserProfile_UserArn(p *OpsworksUserProfileParameters, vals map[string]cty.Value) {
	p.UserArn = ctwhy.ValueAsString(vals["user_arn"])
}