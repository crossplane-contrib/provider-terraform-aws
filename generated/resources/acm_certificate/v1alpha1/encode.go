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

package v1alpha1func EncodeAcmCertificate(r AcmCertificate) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAcmCertificate_Id(r.Spec.ForProvider, ctyVal)
	EncodeAcmCertificate_PrivateKey(r.Spec.ForProvider, ctyVal)
	EncodeAcmCertificate_SubjectAlternativeNames(r.Spec.ForProvider, ctyVal)
	EncodeAcmCertificate_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAcmCertificate_CertificateChain(r.Spec.ForProvider, ctyVal)
	EncodeAcmCertificate_CertificateBody(r.Spec.ForProvider, ctyVal)
	EncodeAcmCertificate_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeAcmCertificate_ValidationMethod(r.Spec.ForProvider, ctyVal)
	EncodeAcmCertificate_CertificateAuthorityArn(r.Spec.ForProvider, ctyVal)
	EncodeAcmCertificate_Options(r.Spec.ForProvider.Options, ctyVal)
	EncodeAcmCertificate_Status(r.Status.AtProvider, ctyVal)
	EncodeAcmCertificate_ValidationEmails(r.Status.AtProvider, ctyVal)
	EncodeAcmCertificate_DomainValidationOptions(r.Status.AtProvider.DomainValidationOptions, ctyVal)
	EncodeAcmCertificate_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeAcmCertificate_Id(p *AcmCertificateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAcmCertificate_PrivateKey(p *AcmCertificateParameters, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodeAcmCertificate_SubjectAlternativeNames(p *AcmCertificateParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubjectAlternativeNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subject_alternative_names"] = cty.SetVal(colVals)
}

func EncodeAcmCertificate_Tags(p *AcmCertificateParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAcmCertificate_CertificateChain(p *AcmCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_chain"] = cty.StringVal(p.CertificateChain)
}

func EncodeAcmCertificate_CertificateBody(p *AcmCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_body"] = cty.StringVal(p.CertificateBody)
}

func EncodeAcmCertificate_DomainName(p *AcmCertificateParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeAcmCertificate_ValidationMethod(p *AcmCertificateParameters, vals map[string]cty.Value) {
	vals["validation_method"] = cty.StringVal(p.ValidationMethod)
}

func EncodeAcmCertificate_CertificateAuthorityArn(p *AcmCertificateParameters, vals map[string]cty.Value) {
	vals["certificate_authority_arn"] = cty.StringVal(p.CertificateAuthorityArn)
}

func EncodeAcmCertificate_Options(p *Options, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Options {
		ctyVal = make(map[string]cty.Value)
		EncodeAcmCertificate_Options_CertificateTransparencyLoggingPreference(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["options"] = cty.ListVal(valsForCollection)
}

func EncodeAcmCertificate_Options_CertificateTransparencyLoggingPreference(p *Options, vals map[string]cty.Value) {
	vals["certificate_transparency_logging_preference"] = cty.StringVal(p.CertificateTransparencyLoggingPreference)
}

func EncodeAcmCertificate_Status(p *AcmCertificateObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeAcmCertificate_ValidationEmails(p *AcmCertificateObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ValidationEmails {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["validation_emails"] = cty.ListVal(colVals)
}

func EncodeAcmCertificate_DomainValidationOptions(p *DomainValidationOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DomainValidationOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeAcmCertificate_DomainValidationOptions_DomainName(v, ctyVal)
		EncodeAcmCertificate_DomainValidationOptions_ResourceRecordName(v, ctyVal)
		EncodeAcmCertificate_DomainValidationOptions_ResourceRecordType(v, ctyVal)
		EncodeAcmCertificate_DomainValidationOptions_ResourceRecordValue(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["domain_validation_options"] = cty.SetVal(valsForCollection)
}

func EncodeAcmCertificate_DomainValidationOptions_DomainName(p *DomainValidationOptions, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeAcmCertificate_DomainValidationOptions_ResourceRecordName(p *DomainValidationOptions, vals map[string]cty.Value) {
	vals["resource_record_name"] = cty.StringVal(p.ResourceRecordName)
}

func EncodeAcmCertificate_DomainValidationOptions_ResourceRecordType(p *DomainValidationOptions, vals map[string]cty.Value) {
	vals["resource_record_type"] = cty.StringVal(p.ResourceRecordType)
}

func EncodeAcmCertificate_DomainValidationOptions_ResourceRecordValue(p *DomainValidationOptions, vals map[string]cty.Value) {
	vals["resource_record_value"] = cty.StringVal(p.ResourceRecordValue)
}

func EncodeAcmCertificate_Arn(p *AcmCertificateObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}