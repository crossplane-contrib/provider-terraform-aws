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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*QuicksightUser)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a QuicksightUser.")
	}
	return EncodeQuicksightUser(*r), nil
}

func EncodeQuicksightUser(r QuicksightUser) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeQuicksightUser_UserName(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightUser_UserRole(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightUser_AwsAccountId(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightUser_Email(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightUser_IdentityType(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightUser_IamArn(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightUser_Id(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightUser_Namespace(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightUser_SessionName(r.Spec.ForProvider, ctyVal)
	EncodeQuicksightUser_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeQuicksightUser_UserName(p QuicksightUserParameters, vals map[string]cty.Value) {
	vals["user_name"] = cty.StringVal(p.UserName)
}

func EncodeQuicksightUser_UserRole(p QuicksightUserParameters, vals map[string]cty.Value) {
	vals["user_role"] = cty.StringVal(p.UserRole)
}

func EncodeQuicksightUser_AwsAccountId(p QuicksightUserParameters, vals map[string]cty.Value) {
	vals["aws_account_id"] = cty.StringVal(p.AwsAccountId)
}

func EncodeQuicksightUser_Email(p QuicksightUserParameters, vals map[string]cty.Value) {
	vals["email"] = cty.StringVal(p.Email)
}

func EncodeQuicksightUser_IdentityType(p QuicksightUserParameters, vals map[string]cty.Value) {
	vals["identity_type"] = cty.StringVal(p.IdentityType)
}

func EncodeQuicksightUser_IamArn(p QuicksightUserParameters, vals map[string]cty.Value) {
	vals["iam_arn"] = cty.StringVal(p.IamArn)
}

func EncodeQuicksightUser_Id(p QuicksightUserParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeQuicksightUser_Namespace(p QuicksightUserParameters, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeQuicksightUser_SessionName(p QuicksightUserParameters, vals map[string]cty.Value) {
	vals["session_name"] = cty.StringVal(p.SessionName)
}

func EncodeQuicksightUser_Arn(p QuicksightUserObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}