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

package v1alpha1func EncodeGlueSecurityConfiguration(r GlueSecurityConfiguration) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeGlueSecurityConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueSecurityConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueSecurityConfiguration_EncryptionConfiguration(r.Spec.ForProvider.EncryptionConfiguration, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeGlueSecurityConfiguration_Id(p *GlueSecurityConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueSecurityConfiguration_Name(p *GlueSecurityConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration(p *EncryptionConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EncryptionConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeGlueSecurityConfiguration_EncryptionConfiguration_CloudwatchEncryption(v.CloudwatchEncryption, ctyVal)
		EncodeGlueSecurityConfiguration_EncryptionConfiguration_JobBookmarksEncryption(v.JobBookmarksEncryption, ctyVal)
		EncodeGlueSecurityConfiguration_EncryptionConfiguration_S3Encryption(v.S3Encryption, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["encryption_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration_CloudwatchEncryption(p *CloudwatchEncryption, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchEncryption {
		ctyVal = make(map[string]cty.Value)
		EncodeGlueSecurityConfiguration_EncryptionConfiguration_CloudwatchEncryption_CloudwatchEncryptionMode(v, ctyVal)
		EncodeGlueSecurityConfiguration_EncryptionConfiguration_CloudwatchEncryption_KmsKeyArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_encryption"] = cty.ListVal(valsForCollection)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration_CloudwatchEncryption_CloudwatchEncryptionMode(p *CloudwatchEncryption, vals map[string]cty.Value) {
	vals["cloudwatch_encryption_mode"] = cty.StringVal(p.CloudwatchEncryptionMode)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration_CloudwatchEncryption_KmsKeyArn(p *CloudwatchEncryption, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration_JobBookmarksEncryption(p *JobBookmarksEncryption, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.JobBookmarksEncryption {
		ctyVal = make(map[string]cty.Value)
		EncodeGlueSecurityConfiguration_EncryptionConfiguration_JobBookmarksEncryption_JobBookmarksEncryptionMode(v, ctyVal)
		EncodeGlueSecurityConfiguration_EncryptionConfiguration_JobBookmarksEncryption_KmsKeyArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["job_bookmarks_encryption"] = cty.ListVal(valsForCollection)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration_JobBookmarksEncryption_JobBookmarksEncryptionMode(p *JobBookmarksEncryption, vals map[string]cty.Value) {
	vals["job_bookmarks_encryption_mode"] = cty.StringVal(p.JobBookmarksEncryptionMode)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration_JobBookmarksEncryption_KmsKeyArn(p *JobBookmarksEncryption, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration_S3Encryption(p *S3Encryption, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3Encryption {
		ctyVal = make(map[string]cty.Value)
		EncodeGlueSecurityConfiguration_EncryptionConfiguration_S3Encryption_KmsKeyArn(v, ctyVal)
		EncodeGlueSecurityConfiguration_EncryptionConfiguration_S3Encryption_S3EncryptionMode(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3_encryption"] = cty.ListVal(valsForCollection)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration_S3Encryption_KmsKeyArn(p *S3Encryption, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeGlueSecurityConfiguration_EncryptionConfiguration_S3Encryption_S3EncryptionMode(p *S3Encryption, vals map[string]cty.Value) {
	vals["s3_encryption_mode"] = cty.StringVal(p.S3EncryptionMode)
}