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
	r, ok := mr.(*IamAccountPasswordPolicy)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IamAccountPasswordPolicy.")
	}
	return EncodeIamAccountPasswordPolicy(*r), nil
}

func EncodeIamAccountPasswordPolicy(r IamAccountPasswordPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamAccountPasswordPolicy_MinimumPasswordLength(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_RequireLowercaseCharacters(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_RequireNumbers(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_RequireSymbols(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_AllowUsersToChangePassword(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_HardExpiry(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_MaxPasswordAge(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_RequireUppercaseCharacters(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_PasswordReusePrevention(r.Spec.ForProvider, ctyVal)
	EncodeIamAccountPasswordPolicy_ExpirePasswords(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeIamAccountPasswordPolicy_MinimumPasswordLength(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["minimum_password_length"] = cty.NumberIntVal(p.MinimumPasswordLength)
}

func EncodeIamAccountPasswordPolicy_RequireLowercaseCharacters(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["require_lowercase_characters"] = cty.BoolVal(p.RequireLowercaseCharacters)
}

func EncodeIamAccountPasswordPolicy_RequireNumbers(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["require_numbers"] = cty.BoolVal(p.RequireNumbers)
}

func EncodeIamAccountPasswordPolicy_RequireSymbols(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["require_symbols"] = cty.BoolVal(p.RequireSymbols)
}

func EncodeIamAccountPasswordPolicy_AllowUsersToChangePassword(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["allow_users_to_change_password"] = cty.BoolVal(p.AllowUsersToChangePassword)
}

func EncodeIamAccountPasswordPolicy_HardExpiry(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["hard_expiry"] = cty.BoolVal(p.HardExpiry)
}

func EncodeIamAccountPasswordPolicy_MaxPasswordAge(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["max_password_age"] = cty.NumberIntVal(p.MaxPasswordAge)
}

func EncodeIamAccountPasswordPolicy_RequireUppercaseCharacters(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["require_uppercase_characters"] = cty.BoolVal(p.RequireUppercaseCharacters)
}

func EncodeIamAccountPasswordPolicy_Id(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamAccountPasswordPolicy_PasswordReusePrevention(p IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	vals["password_reuse_prevention"] = cty.NumberIntVal(p.PasswordReusePrevention)
}

func EncodeIamAccountPasswordPolicy_ExpirePasswords(p IamAccountPasswordPolicyObservation, vals map[string]cty.Value) {
	vals["expire_passwords"] = cty.BoolVal(p.ExpirePasswords)
}