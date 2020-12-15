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
	r, ok := mr.(*IamAccountPasswordPolicy)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamAccountPasswordPolicy(r, ctyValue)
}

func DecodeIamAccountPasswordPolicy(prev *IamAccountPasswordPolicy, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamAccountPasswordPolicy_RequireUppercaseCharacters(&new.Spec.ForProvider, valMap)
	DecodeIamAccountPasswordPolicy_AllowUsersToChangePassword(&new.Spec.ForProvider, valMap)
	DecodeIamAccountPasswordPolicy_PasswordReusePrevention(&new.Spec.ForProvider, valMap)
	DecodeIamAccountPasswordPolicy_RequireNumbers(&new.Spec.ForProvider, valMap)
	DecodeIamAccountPasswordPolicy_HardExpiry(&new.Spec.ForProvider, valMap)
	DecodeIamAccountPasswordPolicy_MaxPasswordAge(&new.Spec.ForProvider, valMap)
	DecodeIamAccountPasswordPolicy_MinimumPasswordLength(&new.Spec.ForProvider, valMap)
	DecodeIamAccountPasswordPolicy_RequireLowercaseCharacters(&new.Spec.ForProvider, valMap)
	DecodeIamAccountPasswordPolicy_RequireSymbols(&new.Spec.ForProvider, valMap)
	DecodeIamAccountPasswordPolicy_ExpirePasswords(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_RequireUppercaseCharacters(p *IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	p.RequireUppercaseCharacters = ctwhy.ValueAsBool(vals["require_uppercase_characters"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_AllowUsersToChangePassword(p *IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	p.AllowUsersToChangePassword = ctwhy.ValueAsBool(vals["allow_users_to_change_password"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_PasswordReusePrevention(p *IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	p.PasswordReusePrevention = ctwhy.ValueAsInt64(vals["password_reuse_prevention"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_RequireNumbers(p *IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	p.RequireNumbers = ctwhy.ValueAsBool(vals["require_numbers"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_HardExpiry(p *IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	p.HardExpiry = ctwhy.ValueAsBool(vals["hard_expiry"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_MaxPasswordAge(p *IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	p.MaxPasswordAge = ctwhy.ValueAsInt64(vals["max_password_age"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_MinimumPasswordLength(p *IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	p.MinimumPasswordLength = ctwhy.ValueAsInt64(vals["minimum_password_length"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_RequireLowercaseCharacters(p *IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	p.RequireLowercaseCharacters = ctwhy.ValueAsBool(vals["require_lowercase_characters"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_RequireSymbols(p *IamAccountPasswordPolicyParameters, vals map[string]cty.Value) {
	p.RequireSymbols = ctwhy.ValueAsBool(vals["require_symbols"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccountPasswordPolicy_ExpirePasswords(p *IamAccountPasswordPolicyObservation, vals map[string]cty.Value) {
	p.ExpirePasswords = ctwhy.ValueAsBool(vals["expire_passwords"])
}