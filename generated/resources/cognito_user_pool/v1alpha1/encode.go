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

package v1alpha1func EncodeCognitoUserPool(r CognitoUserPool) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCognitoUserPool_Id(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_MfaConfiguration(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_EmailVerificationMessage(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_UsernameAttributes(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_AliasAttributes(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_SmsAuthenticationMessage(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_SmsVerificationMessage(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_AutoVerifiedAttributes(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_EmailVerificationSubject(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_Name(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPool_EmailConfiguration(r.Spec.ForProvider.EmailConfiguration, ctyVal)
	EncodeCognitoUserPool_SmsConfiguration(r.Spec.ForProvider.SmsConfiguration, ctyVal)
	EncodeCognitoUserPool_SoftwareTokenMfaConfiguration(r.Spec.ForProvider.SoftwareTokenMfaConfiguration, ctyVal)
	EncodeCognitoUserPool_AdminCreateUserConfig(r.Spec.ForProvider.AdminCreateUserConfig, ctyVal)
	EncodeCognitoUserPool_DeviceConfiguration(r.Spec.ForProvider.DeviceConfiguration, ctyVal)
	EncodeCognitoUserPool_Schema(r.Spec.ForProvider.Schema, ctyVal)
	EncodeCognitoUserPool_UserPoolAddOns(r.Spec.ForProvider.UserPoolAddOns, ctyVal)
	EncodeCognitoUserPool_UsernameConfiguration(r.Spec.ForProvider.UsernameConfiguration, ctyVal)
	EncodeCognitoUserPool_VerificationMessageTemplate(r.Spec.ForProvider.VerificationMessageTemplate, ctyVal)
	EncodeCognitoUserPool_LambdaConfig(r.Spec.ForProvider.LambdaConfig, ctyVal)
	EncodeCognitoUserPool_PasswordPolicy(r.Spec.ForProvider.PasswordPolicy, ctyVal)
	EncodeCognitoUserPool_LastModifiedDate(r.Status.AtProvider, ctyVal)
	EncodeCognitoUserPool_CreationDate(r.Status.AtProvider, ctyVal)
	EncodeCognitoUserPool_Endpoint(r.Status.AtProvider, ctyVal)
	EncodeCognitoUserPool_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeCognitoUserPool_Id(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCognitoUserPool_MfaConfiguration(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	vals["mfa_configuration"] = cty.StringVal(p.MfaConfiguration)
}

func EncodeCognitoUserPool_EmailVerificationMessage(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	vals["email_verification_message"] = cty.StringVal(p.EmailVerificationMessage)
}

func EncodeCognitoUserPool_UsernameAttributes(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.UsernameAttributes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["username_attributes"] = cty.ListVal(colVals)
}

func EncodeCognitoUserPool_AliasAttributes(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AliasAttributes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["alias_attributes"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPool_SmsAuthenticationMessage(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	vals["sms_authentication_message"] = cty.StringVal(p.SmsAuthenticationMessage)
}

func EncodeCognitoUserPool_SmsVerificationMessage(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	vals["sms_verification_message"] = cty.StringVal(p.SmsVerificationMessage)
}

func EncodeCognitoUserPool_Tags(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCognitoUserPool_AutoVerifiedAttributes(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AutoVerifiedAttributes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["auto_verified_attributes"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPool_EmailVerificationSubject(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	vals["email_verification_subject"] = cty.StringVal(p.EmailVerificationSubject)
}

func EncodeCognitoUserPool_Name(p *CognitoUserPoolParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCognitoUserPool_EmailConfiguration(p *EmailConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EmailConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_EmailConfiguration_FromEmailAddress(v, ctyVal)
		EncodeCognitoUserPool_EmailConfiguration_ReplyToEmailAddress(v, ctyVal)
		EncodeCognitoUserPool_EmailConfiguration_SourceArn(v, ctyVal)
		EncodeCognitoUserPool_EmailConfiguration_EmailSendingAccount(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["email_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_EmailConfiguration_FromEmailAddress(p *EmailConfiguration, vals map[string]cty.Value) {
	vals["from_email_address"] = cty.StringVal(p.FromEmailAddress)
}

func EncodeCognitoUserPool_EmailConfiguration_ReplyToEmailAddress(p *EmailConfiguration, vals map[string]cty.Value) {
	vals["reply_to_email_address"] = cty.StringVal(p.ReplyToEmailAddress)
}

func EncodeCognitoUserPool_EmailConfiguration_SourceArn(p *EmailConfiguration, vals map[string]cty.Value) {
	vals["source_arn"] = cty.StringVal(p.SourceArn)
}

func EncodeCognitoUserPool_EmailConfiguration_EmailSendingAccount(p *EmailConfiguration, vals map[string]cty.Value) {
	vals["email_sending_account"] = cty.StringVal(p.EmailSendingAccount)
}

func EncodeCognitoUserPool_SmsConfiguration(p *SmsConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SmsConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_SmsConfiguration_ExternalId(v, ctyVal)
		EncodeCognitoUserPool_SmsConfiguration_SnsCallerArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sms_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_SmsConfiguration_ExternalId(p *SmsConfiguration, vals map[string]cty.Value) {
	vals["external_id"] = cty.StringVal(p.ExternalId)
}

func EncodeCognitoUserPool_SmsConfiguration_SnsCallerArn(p *SmsConfiguration, vals map[string]cty.Value) {
	vals["sns_caller_arn"] = cty.StringVal(p.SnsCallerArn)
}

func EncodeCognitoUserPool_SoftwareTokenMfaConfiguration(p *SoftwareTokenMfaConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SoftwareTokenMfaConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_SoftwareTokenMfaConfiguration_Enabled(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["software_token_mfa_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_SoftwareTokenMfaConfiguration_Enabled(p *SoftwareTokenMfaConfiguration, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeCognitoUserPool_AdminCreateUserConfig(p *AdminCreateUserConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AdminCreateUserConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_AdminCreateUserConfig_AllowAdminCreateUserOnly(v, ctyVal)
		EncodeCognitoUserPool_AdminCreateUserConfig_InviteMessageTemplate(v.InviteMessageTemplate, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["admin_create_user_config"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_AdminCreateUserConfig_AllowAdminCreateUserOnly(p *AdminCreateUserConfig, vals map[string]cty.Value) {
	vals["allow_admin_create_user_only"] = cty.BoolVal(p.AllowAdminCreateUserOnly)
}

func EncodeCognitoUserPool_AdminCreateUserConfig_InviteMessageTemplate(p *InviteMessageTemplate, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.InviteMessageTemplate {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_AdminCreateUserConfig_InviteMessageTemplate_EmailMessage(v, ctyVal)
		EncodeCognitoUserPool_AdminCreateUserConfig_InviteMessageTemplate_EmailSubject(v, ctyVal)
		EncodeCognitoUserPool_AdminCreateUserConfig_InviteMessageTemplate_SmsMessage(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["invite_message_template"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_AdminCreateUserConfig_InviteMessageTemplate_EmailMessage(p *InviteMessageTemplate, vals map[string]cty.Value) {
	vals["email_message"] = cty.StringVal(p.EmailMessage)
}

func EncodeCognitoUserPool_AdminCreateUserConfig_InviteMessageTemplate_EmailSubject(p *InviteMessageTemplate, vals map[string]cty.Value) {
	vals["email_subject"] = cty.StringVal(p.EmailSubject)
}

func EncodeCognitoUserPool_AdminCreateUserConfig_InviteMessageTemplate_SmsMessage(p *InviteMessageTemplate, vals map[string]cty.Value) {
	vals["sms_message"] = cty.StringVal(p.SmsMessage)
}

func EncodeCognitoUserPool_DeviceConfiguration(p *DeviceConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DeviceConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_DeviceConfiguration_ChallengeRequiredOnNewDevice(v, ctyVal)
		EncodeCognitoUserPool_DeviceConfiguration_DeviceOnlyRememberedOnUserPrompt(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["device_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_DeviceConfiguration_ChallengeRequiredOnNewDevice(p *DeviceConfiguration, vals map[string]cty.Value) {
	vals["challenge_required_on_new_device"] = cty.BoolVal(p.ChallengeRequiredOnNewDevice)
}

func EncodeCognitoUserPool_DeviceConfiguration_DeviceOnlyRememberedOnUserPrompt(p *DeviceConfiguration, vals map[string]cty.Value) {
	vals["device_only_remembered_on_user_prompt"] = cty.BoolVal(p.DeviceOnlyRememberedOnUserPrompt)
}

func EncodeCognitoUserPool_Schema(p *Schema, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Schema {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_Schema_Mutable(v, ctyVal)
		EncodeCognitoUserPool_Schema_Name(v, ctyVal)
		EncodeCognitoUserPool_Schema_Required(v, ctyVal)
		EncodeCognitoUserPool_Schema_AttributeDataType(v, ctyVal)
		EncodeCognitoUserPool_Schema_DeveloperOnlyAttribute(v, ctyVal)
		EncodeCognitoUserPool_Schema_NumberAttributeConstraints(v.NumberAttributeConstraints, ctyVal)
		EncodeCognitoUserPool_Schema_StringAttributeConstraints(v.StringAttributeConstraints, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["schema"] = cty.SetVal(valsForCollection)
}

func EncodeCognitoUserPool_Schema_Mutable(p *Schema, vals map[string]cty.Value) {
	vals["mutable"] = cty.BoolVal(p.Mutable)
}

func EncodeCognitoUserPool_Schema_Name(p *Schema, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCognitoUserPool_Schema_Required(p *Schema, vals map[string]cty.Value) {
	vals["required"] = cty.BoolVal(p.Required)
}

func EncodeCognitoUserPool_Schema_AttributeDataType(p *Schema, vals map[string]cty.Value) {
	vals["attribute_data_type"] = cty.StringVal(p.AttributeDataType)
}

func EncodeCognitoUserPool_Schema_DeveloperOnlyAttribute(p *Schema, vals map[string]cty.Value) {
	vals["developer_only_attribute"] = cty.BoolVal(p.DeveloperOnlyAttribute)
}

func EncodeCognitoUserPool_Schema_NumberAttributeConstraints(p *NumberAttributeConstraints, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NumberAttributeConstraints {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_Schema_NumberAttributeConstraints_MaxValue(v, ctyVal)
		EncodeCognitoUserPool_Schema_NumberAttributeConstraints_MinValue(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["number_attribute_constraints"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_Schema_NumberAttributeConstraints_MaxValue(p *NumberAttributeConstraints, vals map[string]cty.Value) {
	vals["max_value"] = cty.StringVal(p.MaxValue)
}

func EncodeCognitoUserPool_Schema_NumberAttributeConstraints_MinValue(p *NumberAttributeConstraints, vals map[string]cty.Value) {
	vals["min_value"] = cty.StringVal(p.MinValue)
}

func EncodeCognitoUserPool_Schema_StringAttributeConstraints(p *StringAttributeConstraints, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.StringAttributeConstraints {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_Schema_StringAttributeConstraints_MaxLength(v, ctyVal)
		EncodeCognitoUserPool_Schema_StringAttributeConstraints_MinLength(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["string_attribute_constraints"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_Schema_StringAttributeConstraints_MaxLength(p *StringAttributeConstraints, vals map[string]cty.Value) {
	vals["max_length"] = cty.StringVal(p.MaxLength)
}

func EncodeCognitoUserPool_Schema_StringAttributeConstraints_MinLength(p *StringAttributeConstraints, vals map[string]cty.Value) {
	vals["min_length"] = cty.StringVal(p.MinLength)
}

func EncodeCognitoUserPool_UserPoolAddOns(p *UserPoolAddOns, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UserPoolAddOns {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_UserPoolAddOns_AdvancedSecurityMode(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["user_pool_add_ons"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_UserPoolAddOns_AdvancedSecurityMode(p *UserPoolAddOns, vals map[string]cty.Value) {
	vals["advanced_security_mode"] = cty.StringVal(p.AdvancedSecurityMode)
}

func EncodeCognitoUserPool_UsernameConfiguration(p *UsernameConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UsernameConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_UsernameConfiguration_CaseSensitive(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["username_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_UsernameConfiguration_CaseSensitive(p *UsernameConfiguration, vals map[string]cty.Value) {
	vals["case_sensitive"] = cty.BoolVal(p.CaseSensitive)
}

func EncodeCognitoUserPool_VerificationMessageTemplate(p *VerificationMessageTemplate, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VerificationMessageTemplate {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_VerificationMessageTemplate_SmsMessage(v, ctyVal)
		EncodeCognitoUserPool_VerificationMessageTemplate_DefaultEmailOption(v, ctyVal)
		EncodeCognitoUserPool_VerificationMessageTemplate_EmailMessage(v, ctyVal)
		EncodeCognitoUserPool_VerificationMessageTemplate_EmailMessageByLink(v, ctyVal)
		EncodeCognitoUserPool_VerificationMessageTemplate_EmailSubject(v, ctyVal)
		EncodeCognitoUserPool_VerificationMessageTemplate_EmailSubjectByLink(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["verification_message_template"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_VerificationMessageTemplate_SmsMessage(p *VerificationMessageTemplate, vals map[string]cty.Value) {
	vals["sms_message"] = cty.StringVal(p.SmsMessage)
}

func EncodeCognitoUserPool_VerificationMessageTemplate_DefaultEmailOption(p *VerificationMessageTemplate, vals map[string]cty.Value) {
	vals["default_email_option"] = cty.StringVal(p.DefaultEmailOption)
}

func EncodeCognitoUserPool_VerificationMessageTemplate_EmailMessage(p *VerificationMessageTemplate, vals map[string]cty.Value) {
	vals["email_message"] = cty.StringVal(p.EmailMessage)
}

func EncodeCognitoUserPool_VerificationMessageTemplate_EmailMessageByLink(p *VerificationMessageTemplate, vals map[string]cty.Value) {
	vals["email_message_by_link"] = cty.StringVal(p.EmailMessageByLink)
}

func EncodeCognitoUserPool_VerificationMessageTemplate_EmailSubject(p *VerificationMessageTemplate, vals map[string]cty.Value) {
	vals["email_subject"] = cty.StringVal(p.EmailSubject)
}

func EncodeCognitoUserPool_VerificationMessageTemplate_EmailSubjectByLink(p *VerificationMessageTemplate, vals map[string]cty.Value) {
	vals["email_subject_by_link"] = cty.StringVal(p.EmailSubjectByLink)
}

func EncodeCognitoUserPool_LambdaConfig(p *LambdaConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LambdaConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_LambdaConfig_CustomMessage(v, ctyVal)
		EncodeCognitoUserPool_LambdaConfig_PostConfirmation(v, ctyVal)
		EncodeCognitoUserPool_LambdaConfig_PreAuthentication(v, ctyVal)
		EncodeCognitoUserPool_LambdaConfig_PreSignUp(v, ctyVal)
		EncodeCognitoUserPool_LambdaConfig_PreTokenGeneration(v, ctyVal)
		EncodeCognitoUserPool_LambdaConfig_UserMigration(v, ctyVal)
		EncodeCognitoUserPool_LambdaConfig_VerifyAuthChallengeResponse(v, ctyVal)
		EncodeCognitoUserPool_LambdaConfig_CreateAuthChallenge(v, ctyVal)
		EncodeCognitoUserPool_LambdaConfig_DefineAuthChallenge(v, ctyVal)
		EncodeCognitoUserPool_LambdaConfig_PostAuthentication(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lambda_config"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_LambdaConfig_CustomMessage(p *LambdaConfig, vals map[string]cty.Value) {
	vals["custom_message"] = cty.StringVal(p.CustomMessage)
}

func EncodeCognitoUserPool_LambdaConfig_PostConfirmation(p *LambdaConfig, vals map[string]cty.Value) {
	vals["post_confirmation"] = cty.StringVal(p.PostConfirmation)
}

func EncodeCognitoUserPool_LambdaConfig_PreAuthentication(p *LambdaConfig, vals map[string]cty.Value) {
	vals["pre_authentication"] = cty.StringVal(p.PreAuthentication)
}

func EncodeCognitoUserPool_LambdaConfig_PreSignUp(p *LambdaConfig, vals map[string]cty.Value) {
	vals["pre_sign_up"] = cty.StringVal(p.PreSignUp)
}

func EncodeCognitoUserPool_LambdaConfig_PreTokenGeneration(p *LambdaConfig, vals map[string]cty.Value) {
	vals["pre_token_generation"] = cty.StringVal(p.PreTokenGeneration)
}

func EncodeCognitoUserPool_LambdaConfig_UserMigration(p *LambdaConfig, vals map[string]cty.Value) {
	vals["user_migration"] = cty.StringVal(p.UserMigration)
}

func EncodeCognitoUserPool_LambdaConfig_VerifyAuthChallengeResponse(p *LambdaConfig, vals map[string]cty.Value) {
	vals["verify_auth_challenge_response"] = cty.StringVal(p.VerifyAuthChallengeResponse)
}

func EncodeCognitoUserPool_LambdaConfig_CreateAuthChallenge(p *LambdaConfig, vals map[string]cty.Value) {
	vals["create_auth_challenge"] = cty.StringVal(p.CreateAuthChallenge)
}

func EncodeCognitoUserPool_LambdaConfig_DefineAuthChallenge(p *LambdaConfig, vals map[string]cty.Value) {
	vals["define_auth_challenge"] = cty.StringVal(p.DefineAuthChallenge)
}

func EncodeCognitoUserPool_LambdaConfig_PostAuthentication(p *LambdaConfig, vals map[string]cty.Value) {
	vals["post_authentication"] = cty.StringVal(p.PostAuthentication)
}

func EncodeCognitoUserPool_PasswordPolicy(p *PasswordPolicy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.PasswordPolicy {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoUserPool_PasswordPolicy_TemporaryPasswordValidityDays(v, ctyVal)
		EncodeCognitoUserPool_PasswordPolicy_MinimumLength(v, ctyVal)
		EncodeCognitoUserPool_PasswordPolicy_RequireLowercase(v, ctyVal)
		EncodeCognitoUserPool_PasswordPolicy_RequireNumbers(v, ctyVal)
		EncodeCognitoUserPool_PasswordPolicy_RequireSymbols(v, ctyVal)
		EncodeCognitoUserPool_PasswordPolicy_RequireUppercase(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["password_policy"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPool_PasswordPolicy_TemporaryPasswordValidityDays(p *PasswordPolicy, vals map[string]cty.Value) {
	vals["temporary_password_validity_days"] = cty.IntVal(p.TemporaryPasswordValidityDays)
}

func EncodeCognitoUserPool_PasswordPolicy_MinimumLength(p *PasswordPolicy, vals map[string]cty.Value) {
	vals["minimum_length"] = cty.IntVal(p.MinimumLength)
}

func EncodeCognitoUserPool_PasswordPolicy_RequireLowercase(p *PasswordPolicy, vals map[string]cty.Value) {
	vals["require_lowercase"] = cty.BoolVal(p.RequireLowercase)
}

func EncodeCognitoUserPool_PasswordPolicy_RequireNumbers(p *PasswordPolicy, vals map[string]cty.Value) {
	vals["require_numbers"] = cty.BoolVal(p.RequireNumbers)
}

func EncodeCognitoUserPool_PasswordPolicy_RequireSymbols(p *PasswordPolicy, vals map[string]cty.Value) {
	vals["require_symbols"] = cty.BoolVal(p.RequireSymbols)
}

func EncodeCognitoUserPool_PasswordPolicy_RequireUppercase(p *PasswordPolicy, vals map[string]cty.Value) {
	vals["require_uppercase"] = cty.BoolVal(p.RequireUppercase)
}

func EncodeCognitoUserPool_LastModifiedDate(p *CognitoUserPoolObservation, vals map[string]cty.Value) {
	vals["last_modified_date"] = cty.StringVal(p.LastModifiedDate)
}

func EncodeCognitoUserPool_CreationDate(p *CognitoUserPoolObservation, vals map[string]cty.Value) {
	vals["creation_date"] = cty.StringVal(p.CreationDate)
}

func EncodeCognitoUserPool_Endpoint(p *CognitoUserPoolObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeCognitoUserPool_Arn(p *CognitoUserPoolObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}