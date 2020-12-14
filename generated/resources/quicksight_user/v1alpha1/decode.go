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
	r, ok := mr.(*QuicksightUser)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeQuicksightUser(r, ctyValue)
}

func DecodeQuicksightUser(prev *QuicksightUser, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeQuicksightUser_UserRole(&new.Spec.ForProvider, valMap)
	DecodeQuicksightUser_AwsAccountId(&new.Spec.ForProvider, valMap)
	DecodeQuicksightUser_Namespace(&new.Spec.ForProvider, valMap)
	DecodeQuicksightUser_SessionName(&new.Spec.ForProvider, valMap)
	DecodeQuicksightUser_Id(&new.Spec.ForProvider, valMap)
	DecodeQuicksightUser_IdentityType(&new.Spec.ForProvider, valMap)
	DecodeQuicksightUser_UserName(&new.Spec.ForProvider, valMap)
	DecodeQuicksightUser_Email(&new.Spec.ForProvider, valMap)
	DecodeQuicksightUser_IamArn(&new.Spec.ForProvider, valMap)
	DecodeQuicksightUser_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeQuicksightUser_UserRole(p *QuicksightUserParameters, vals map[string]cty.Value) {
	p.UserRole = ctwhy.ValueAsString(vals["user_role"])
}

func DecodeQuicksightUser_AwsAccountId(p *QuicksightUserParameters, vals map[string]cty.Value) {
	p.AwsAccountId = ctwhy.ValueAsString(vals["aws_account_id"])
}

func DecodeQuicksightUser_Namespace(p *QuicksightUserParameters, vals map[string]cty.Value) {
	p.Namespace = ctwhy.ValueAsString(vals["namespace"])
}

func DecodeQuicksightUser_SessionName(p *QuicksightUserParameters, vals map[string]cty.Value) {
	p.SessionName = ctwhy.ValueAsString(vals["session_name"])
}

func DecodeQuicksightUser_Id(p *QuicksightUserParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeQuicksightUser_IdentityType(p *QuicksightUserParameters, vals map[string]cty.Value) {
	p.IdentityType = ctwhy.ValueAsString(vals["identity_type"])
}

func DecodeQuicksightUser_UserName(p *QuicksightUserParameters, vals map[string]cty.Value) {
	p.UserName = ctwhy.ValueAsString(vals["user_name"])
}

func DecodeQuicksightUser_Email(p *QuicksightUserParameters, vals map[string]cty.Value) {
	p.Email = ctwhy.ValueAsString(vals["email"])
}

func DecodeQuicksightUser_IamArn(p *QuicksightUserParameters, vals map[string]cty.Value) {
	p.IamArn = ctwhy.ValueAsString(vals["iam_arn"])
}

func DecodeQuicksightUser_Arn(p *QuicksightUserObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}