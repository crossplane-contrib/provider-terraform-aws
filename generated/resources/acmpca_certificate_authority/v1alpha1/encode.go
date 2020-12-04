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
	"github.com/zclconf/go-cty/cty"
)

func EncodeAcmpcaCertificateAuthority(r AcmpcaCertificateAuthority) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAcmpcaCertificateAuthority_Id(r.Spec.ForProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_Type(r.Spec.ForProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_PermanentDeletionTimeInDays(r.Spec.ForProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration(r.Spec.ForProvider.CertificateAuthorityConfiguration, ctyVal)
	EncodeAcmpcaCertificateAuthority_RevocationConfiguration(r.Spec.ForProvider.RevocationConfiguration, ctyVal)
	EncodeAcmpcaCertificateAuthority_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeAcmpcaCertificateAuthority_Certificate(r.Status.AtProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateSigningRequest(r.Status.AtProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_NotAfter(r.Status.AtProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_Serial(r.Status.AtProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_Arn(r.Status.AtProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateChain(r.Status.AtProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_NotBefore(r.Status.AtProvider, ctyVal)
	EncodeAcmpcaCertificateAuthority_Status(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAcmpcaCertificateAuthority_Id(p AcmpcaCertificateAuthorityParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAcmpcaCertificateAuthority_Type(p AcmpcaCertificateAuthorityParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeAcmpcaCertificateAuthority_Enabled(p AcmpcaCertificateAuthorityParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeAcmpcaCertificateAuthority_PermanentDeletionTimeInDays(p AcmpcaCertificateAuthorityParameters, vals map[string]cty.Value) {
	vals["permanent_deletion_time_in_days"] = cty.NumberIntVal(p.PermanentDeletionTimeInDays)
}

func EncodeAcmpcaCertificateAuthority_Tags(p AcmpcaCertificateAuthorityParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration(p CertificateAuthorityConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_KeyAlgorithm(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_SigningAlgorithm(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject(p.Subject, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["certificate_authority_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_KeyAlgorithm(p CertificateAuthorityConfiguration, vals map[string]cty.Value) {
	vals["key_algorithm"] = cty.StringVal(p.KeyAlgorithm)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_SigningAlgorithm(p CertificateAuthorityConfiguration, vals map[string]cty.Value) {
	vals["signing_algorithm"] = cty.StringVal(p.SigningAlgorithm)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject(p Subject, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Country(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_GivenName(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Title(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_OrganizationalUnit(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Pseudonym(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_CommonName(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_DistinguishedNameQualifier(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_GenerationQualifier(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Initials(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Locality(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Organization(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_State(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Surname(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["subject"] = cty.ListVal(valsForCollection)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Country(p Subject, vals map[string]cty.Value) {
	vals["country"] = cty.StringVal(p.Country)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_GivenName(p Subject, vals map[string]cty.Value) {
	vals["given_name"] = cty.StringVal(p.GivenName)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Title(p Subject, vals map[string]cty.Value) {
	vals["title"] = cty.StringVal(p.Title)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_OrganizationalUnit(p Subject, vals map[string]cty.Value) {
	vals["organizational_unit"] = cty.StringVal(p.OrganizationalUnit)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Pseudonym(p Subject, vals map[string]cty.Value) {
	vals["pseudonym"] = cty.StringVal(p.Pseudonym)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_CommonName(p Subject, vals map[string]cty.Value) {
	vals["common_name"] = cty.StringVal(p.CommonName)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_DistinguishedNameQualifier(p Subject, vals map[string]cty.Value) {
	vals["distinguished_name_qualifier"] = cty.StringVal(p.DistinguishedNameQualifier)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_GenerationQualifier(p Subject, vals map[string]cty.Value) {
	vals["generation_qualifier"] = cty.StringVal(p.GenerationQualifier)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Initials(p Subject, vals map[string]cty.Value) {
	vals["initials"] = cty.StringVal(p.Initials)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Locality(p Subject, vals map[string]cty.Value) {
	vals["locality"] = cty.StringVal(p.Locality)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Organization(p Subject, vals map[string]cty.Value) {
	vals["organization"] = cty.StringVal(p.Organization)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_State(p Subject, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeAcmpcaCertificateAuthority_CertificateAuthorityConfiguration_Subject_Surname(p Subject, vals map[string]cty.Value) {
	vals["surname"] = cty.StringVal(p.Surname)
}

func EncodeAcmpcaCertificateAuthority_RevocationConfiguration(p RevocationConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration(p.CrlConfiguration, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["revocation_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration(p CrlConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration_S3BucketName(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration_CustomCname(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration_Enabled(p, ctyVal)
	EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration_ExpirationInDays(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["crl_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration_S3BucketName(p CrlConfiguration, vals map[string]cty.Value) {
	vals["s3_bucket_name"] = cty.StringVal(p.S3BucketName)
}

func EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration_CustomCname(p CrlConfiguration, vals map[string]cty.Value) {
	vals["custom_cname"] = cty.StringVal(p.CustomCname)
}

func EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration_Enabled(p CrlConfiguration, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeAcmpcaCertificateAuthority_RevocationConfiguration_CrlConfiguration_ExpirationInDays(p CrlConfiguration, vals map[string]cty.Value) {
	vals["expiration_in_days"] = cty.NumberIntVal(p.ExpirationInDays)
}

func EncodeAcmpcaCertificateAuthority_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeAcmpcaCertificateAuthority_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeAcmpcaCertificateAuthority_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeAcmpcaCertificateAuthority_Certificate(p AcmpcaCertificateAuthorityObservation, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}

func EncodeAcmpcaCertificateAuthority_CertificateSigningRequest(p AcmpcaCertificateAuthorityObservation, vals map[string]cty.Value) {
	vals["certificate_signing_request"] = cty.StringVal(p.CertificateSigningRequest)
}

func EncodeAcmpcaCertificateAuthority_NotAfter(p AcmpcaCertificateAuthorityObservation, vals map[string]cty.Value) {
	vals["not_after"] = cty.StringVal(p.NotAfter)
}

func EncodeAcmpcaCertificateAuthority_Serial(p AcmpcaCertificateAuthorityObservation, vals map[string]cty.Value) {
	vals["serial"] = cty.StringVal(p.Serial)
}

func EncodeAcmpcaCertificateAuthority_Arn(p AcmpcaCertificateAuthorityObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAcmpcaCertificateAuthority_CertificateChain(p AcmpcaCertificateAuthorityObservation, vals map[string]cty.Value) {
	vals["certificate_chain"] = cty.StringVal(p.CertificateChain)
}

func EncodeAcmpcaCertificateAuthority_NotBefore(p AcmpcaCertificateAuthorityObservation, vals map[string]cty.Value) {
	vals["not_before"] = cty.StringVal(p.NotBefore)
}

func EncodeAcmpcaCertificateAuthority_Status(p AcmpcaCertificateAuthorityObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}