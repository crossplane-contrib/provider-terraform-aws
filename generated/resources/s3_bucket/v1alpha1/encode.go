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

package v1alpha1func EncodeS3Bucket(r S3Bucket) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeS3Bucket_Arn(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_Policy(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_WebsiteEndpoint(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_AccelerationStatus(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_BucketPrefix(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_Acl(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_HostedZoneId(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_Tags(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_WebsiteDomain(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_RequestPayer(r.Spec.ForProvider, ctyVal)
	EncodeS3Bucket_CorsRule(r.Spec.ForProvider.CorsRule, ctyVal)
	EncodeS3Bucket_Website(r.Spec.ForProvider.Website, ctyVal)
	EncodeS3Bucket_Versioning(r.Spec.ForProvider.Versioning, ctyVal)
	EncodeS3Bucket_Grant(r.Spec.ForProvider.Grant, ctyVal)
	EncodeS3Bucket_LifecycleRule(r.Spec.ForProvider.LifecycleRule, ctyVal)
	EncodeS3Bucket_Logging(r.Spec.ForProvider.Logging, ctyVal)
	EncodeS3Bucket_ObjectLockConfiguration(r.Spec.ForProvider.ObjectLockConfiguration, ctyVal)
	EncodeS3Bucket_ReplicationConfiguration(r.Spec.ForProvider.ReplicationConfiguration, ctyVal)
	EncodeS3Bucket_ServerSideEncryptionConfiguration(r.Spec.ForProvider.ServerSideEncryptionConfiguration, ctyVal)
	EncodeS3Bucket_BucketDomainName(r.Status.AtProvider, ctyVal)
	EncodeS3Bucket_BucketRegionalDomainName(r.Status.AtProvider, ctyVal)
	EncodeS3Bucket_Region(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeS3Bucket_Arn(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeS3Bucket_Bucket(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3Bucket_Id(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3Bucket_Policy(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeS3Bucket_WebsiteEndpoint(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["website_endpoint"] = cty.StringVal(p.WebsiteEndpoint)
}

func EncodeS3Bucket_AccelerationStatus(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["acceleration_status"] = cty.StringVal(p.AccelerationStatus)
}

func EncodeS3Bucket_BucketPrefix(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["bucket_prefix"] = cty.StringVal(p.BucketPrefix)
}

func EncodeS3Bucket_Acl(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["acl"] = cty.StringVal(p.Acl)
}

func EncodeS3Bucket_HostedZoneId(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}

func EncodeS3Bucket_Tags(p *S3BucketParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeS3Bucket_WebsiteDomain(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["website_domain"] = cty.StringVal(p.WebsiteDomain)
}

func EncodeS3Bucket_ForceDestroy(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeS3Bucket_RequestPayer(p *S3BucketParameters, vals map[string]cty.Value) {
	vals["request_payer"] = cty.StringVal(p.RequestPayer)
}

func EncodeS3Bucket_CorsRule(p *CorsRule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CorsRule {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_CorsRule_ExposeHeaders(v, ctyVal)
		EncodeS3Bucket_CorsRule_MaxAgeSeconds(v, ctyVal)
		EncodeS3Bucket_CorsRule_AllowedHeaders(v, ctyVal)
		EncodeS3Bucket_CorsRule_AllowedMethods(v, ctyVal)
		EncodeS3Bucket_CorsRule_AllowedOrigins(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cors_rule"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_CorsRule_ExposeHeaders(p *CorsRule, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ExposeHeaders {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["expose_headers"] = cty.ListVal(colVals)
}

func EncodeS3Bucket_CorsRule_MaxAgeSeconds(p *CorsRule, vals map[string]cty.Value) {
	vals["max_age_seconds"] = cty.IntVal(p.MaxAgeSeconds)
}

func EncodeS3Bucket_CorsRule_AllowedHeaders(p *CorsRule, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedHeaders {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_headers"] = cty.ListVal(colVals)
}

func EncodeS3Bucket_CorsRule_AllowedMethods(p *CorsRule, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedMethods {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_methods"] = cty.ListVal(colVals)
}

func EncodeS3Bucket_CorsRule_AllowedOrigins(p *CorsRule, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedOrigins {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_origins"] = cty.ListVal(colVals)
}

func EncodeS3Bucket_Website(p *Website, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Website {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_Website_ErrorDocument(v, ctyVal)
		EncodeS3Bucket_Website_IndexDocument(v, ctyVal)
		EncodeS3Bucket_Website_RedirectAllRequestsTo(v, ctyVal)
		EncodeS3Bucket_Website_RoutingRules(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["website"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_Website_ErrorDocument(p *Website, vals map[string]cty.Value) {
	vals["error_document"] = cty.StringVal(p.ErrorDocument)
}

func EncodeS3Bucket_Website_IndexDocument(p *Website, vals map[string]cty.Value) {
	vals["index_document"] = cty.StringVal(p.IndexDocument)
}

func EncodeS3Bucket_Website_RedirectAllRequestsTo(p *Website, vals map[string]cty.Value) {
	vals["redirect_all_requests_to"] = cty.StringVal(p.RedirectAllRequestsTo)
}

func EncodeS3Bucket_Website_RoutingRules(p *Website, vals map[string]cty.Value) {
	vals["routing_rules"] = cty.StringVal(p.RoutingRules)
}

func EncodeS3Bucket_Versioning(p *Versioning, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Versioning {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_Versioning_MfaDelete(v, ctyVal)
		EncodeS3Bucket_Versioning_Enabled(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["versioning"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_Versioning_MfaDelete(p *Versioning, vals map[string]cty.Value) {
	vals["mfa_delete"] = cty.BoolVal(p.MfaDelete)
}

func EncodeS3Bucket_Versioning_Enabled(p *Versioning, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeS3Bucket_Grant(p *Grant, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Grant {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_Grant_Id(v, ctyVal)
		EncodeS3Bucket_Grant_Permissions(v, ctyVal)
		EncodeS3Bucket_Grant_Type(v, ctyVal)
		EncodeS3Bucket_Grant_Uri(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["grant"] = cty.SetVal(valsForCollection)
}

func EncodeS3Bucket_Grant_Id(p *Grant, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3Bucket_Grant_Permissions(p *Grant, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Permissions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["permissions"] = cty.SetVal(colVals)
}

func EncodeS3Bucket_Grant_Type(p *Grant, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeS3Bucket_Grant_Uri(p *Grant, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeS3Bucket_LifecycleRule(p *LifecycleRule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LifecycleRule {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_LifecycleRule_AbortIncompleteMultipartUploadDays(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_Enabled(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_Id(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_Prefix(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_Tags(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_Transition(v.Transition, ctyVal)
		EncodeS3Bucket_LifecycleRule_Expiration(v.Expiration, ctyVal)
		EncodeS3Bucket_LifecycleRule_NoncurrentVersionExpiration(v.NoncurrentVersionExpiration, ctyVal)
		EncodeS3Bucket_LifecycleRule_NoncurrentVersionTransition(v.NoncurrentVersionTransition, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lifecycle_rule"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_LifecycleRule_AbortIncompleteMultipartUploadDays(p *LifecycleRule, vals map[string]cty.Value) {
	vals["abort_incomplete_multipart_upload_days"] = cty.IntVal(p.AbortIncompleteMultipartUploadDays)
}

func EncodeS3Bucket_LifecycleRule_Enabled(p *LifecycleRule, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeS3Bucket_LifecycleRule_Id(p *LifecycleRule, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3Bucket_LifecycleRule_Prefix(p *LifecycleRule, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeS3Bucket_LifecycleRule_Tags(p *LifecycleRule, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeS3Bucket_LifecycleRule_Transition(p *Transition, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Transition {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_LifecycleRule_Transition_Date(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_Transition_Days(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_Transition_StorageClass(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["transition"] = cty.SetVal(valsForCollection)
}

func EncodeS3Bucket_LifecycleRule_Transition_Date(p *Transition, vals map[string]cty.Value) {
	vals["date"] = cty.StringVal(p.Date)
}

func EncodeS3Bucket_LifecycleRule_Transition_Days(p *Transition, vals map[string]cty.Value) {
	vals["days"] = cty.IntVal(p.Days)
}

func EncodeS3Bucket_LifecycleRule_Transition_StorageClass(p *Transition, vals map[string]cty.Value) {
	vals["storage_class"] = cty.StringVal(p.StorageClass)
}

func EncodeS3Bucket_LifecycleRule_Expiration(p *Expiration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Expiration {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_LifecycleRule_Expiration_Date(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_Expiration_Days(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_Expiration_ExpiredObjectDeleteMarker(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["expiration"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_LifecycleRule_Expiration_Date(p *Expiration, vals map[string]cty.Value) {
	vals["date"] = cty.StringVal(p.Date)
}

func EncodeS3Bucket_LifecycleRule_Expiration_Days(p *Expiration, vals map[string]cty.Value) {
	vals["days"] = cty.IntVal(p.Days)
}

func EncodeS3Bucket_LifecycleRule_Expiration_ExpiredObjectDeleteMarker(p *Expiration, vals map[string]cty.Value) {
	vals["expired_object_delete_marker"] = cty.BoolVal(p.ExpiredObjectDeleteMarker)
}

func EncodeS3Bucket_LifecycleRule_NoncurrentVersionExpiration(p *NoncurrentVersionExpiration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NoncurrentVersionExpiration {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_LifecycleRule_NoncurrentVersionExpiration_Days(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["noncurrent_version_expiration"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_LifecycleRule_NoncurrentVersionExpiration_Days(p *NoncurrentVersionExpiration, vals map[string]cty.Value) {
	vals["days"] = cty.IntVal(p.Days)
}

func EncodeS3Bucket_LifecycleRule_NoncurrentVersionTransition(p *NoncurrentVersionTransition, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NoncurrentVersionTransition {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_LifecycleRule_NoncurrentVersionTransition_Days(v, ctyVal)
		EncodeS3Bucket_LifecycleRule_NoncurrentVersionTransition_StorageClass(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["noncurrent_version_transition"] = cty.SetVal(valsForCollection)
}

func EncodeS3Bucket_LifecycleRule_NoncurrentVersionTransition_Days(p *NoncurrentVersionTransition, vals map[string]cty.Value) {
	vals["days"] = cty.IntVal(p.Days)
}

func EncodeS3Bucket_LifecycleRule_NoncurrentVersionTransition_StorageClass(p *NoncurrentVersionTransition, vals map[string]cty.Value) {
	vals["storage_class"] = cty.StringVal(p.StorageClass)
}

func EncodeS3Bucket_Logging(p *Logging, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Logging {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_Logging_TargetBucket(v, ctyVal)
		EncodeS3Bucket_Logging_TargetPrefix(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["logging"] = cty.SetVal(valsForCollection)
}

func EncodeS3Bucket_Logging_TargetBucket(p *Logging, vals map[string]cty.Value) {
	vals["target_bucket"] = cty.StringVal(p.TargetBucket)
}

func EncodeS3Bucket_Logging_TargetPrefix(p *Logging, vals map[string]cty.Value) {
	vals["target_prefix"] = cty.StringVal(p.TargetPrefix)
}

func EncodeS3Bucket_ObjectLockConfiguration(p *ObjectLockConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ObjectLockConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ObjectLockConfiguration_ObjectLockEnabled(v, ctyVal)
		EncodeS3Bucket_ObjectLockConfiguration_Rule(v.Rule, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["object_lock_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ObjectLockConfiguration_ObjectLockEnabled(p *ObjectLockConfiguration, vals map[string]cty.Value) {
	vals["object_lock_enabled"] = cty.StringVal(p.ObjectLockEnabled)
}

func EncodeS3Bucket_ObjectLockConfiguration_Rule(p *Rule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Rule {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ObjectLockConfiguration_Rule_DefaultRetention(v.DefaultRetention, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["rule"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ObjectLockConfiguration_Rule_DefaultRetention(p *DefaultRetention, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DefaultRetention {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ObjectLockConfiguration_Rule_DefaultRetention_Days(v, ctyVal)
		EncodeS3Bucket_ObjectLockConfiguration_Rule_DefaultRetention_Mode(v, ctyVal)
		EncodeS3Bucket_ObjectLockConfiguration_Rule_DefaultRetention_Years(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["default_retention"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ObjectLockConfiguration_Rule_DefaultRetention_Days(p *DefaultRetention, vals map[string]cty.Value) {
	vals["days"] = cty.IntVal(p.Days)
}

func EncodeS3Bucket_ObjectLockConfiguration_Rule_DefaultRetention_Mode(p *DefaultRetention, vals map[string]cty.Value) {
	vals["mode"] = cty.StringVal(p.Mode)
}

func EncodeS3Bucket_ObjectLockConfiguration_Rule_DefaultRetention_Years(p *DefaultRetention, vals map[string]cty.Value) {
	vals["years"] = cty.IntVal(p.Years)
}

func EncodeS3Bucket_ReplicationConfiguration(p *ReplicationConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ReplicationConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ReplicationConfiguration_Role(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules(v.Rules, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["replication_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ReplicationConfiguration_Role(p *ReplicationConfiguration, vals map[string]cty.Value) {
	vals["role"] = cty.StringVal(p.Role)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules(p *Rules, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Rules {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Id(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Prefix(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Priority(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Status(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Destination(v.Destination, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Filter(v.Filter, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_SourceSelectionCriteria(v.SourceSelectionCriteria, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["rules"] = cty.SetVal(valsForCollection)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Id(p *Rules, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Prefix(p *Rules, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Priority(p *Rules, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Status(p *Rules, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Destination(p *Destination, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Destination {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_AccountId(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_Bucket(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_ReplicaKmsKeyId(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_StorageClass(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_AccessControlTranslation(v.AccessControlTranslation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["destination"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_AccountId(p *Destination, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_Bucket(p *Destination, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_ReplicaKmsKeyId(p *Destination, vals map[string]cty.Value) {
	vals["replica_kms_key_id"] = cty.StringVal(p.ReplicaKmsKeyId)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_StorageClass(p *Destination, vals map[string]cty.Value) {
	vals["storage_class"] = cty.StringVal(p.StorageClass)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_AccessControlTranslation(p *AccessControlTranslation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AccessControlTranslation {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_AccessControlTranslation_Owner(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["access_control_translation"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Destination_AccessControlTranslation_Owner(p *AccessControlTranslation, vals map[string]cty.Value) {
	vals["owner"] = cty.StringVal(p.Owner)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Filter(p *Filter, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Filter {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Filter_Prefix(v, ctyVal)
		EncodeS3Bucket_ReplicationConfiguration_Rules_Filter_Tags(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["filter"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Filter_Prefix(p *Filter, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_Filter_Tags(p *Filter, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_SourceSelectionCriteria(p *SourceSelectionCriteria, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SourceSelectionCriteria {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ReplicationConfiguration_Rules_SourceSelectionCriteria_SseKmsEncryptedObjects(v.SseKmsEncryptedObjects, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["source_selection_criteria"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_SourceSelectionCriteria_SseKmsEncryptedObjects(p *SseKmsEncryptedObjects, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SseKmsEncryptedObjects {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ReplicationConfiguration_Rules_SourceSelectionCriteria_SseKmsEncryptedObjects_Enabled(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sse_kms_encrypted_objects"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ReplicationConfiguration_Rules_SourceSelectionCriteria_SseKmsEncryptedObjects_Enabled(p *SseKmsEncryptedObjects, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeS3Bucket_ServerSideEncryptionConfiguration(p *ServerSideEncryptionConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ServerSideEncryptionConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ServerSideEncryptionConfiguration_Rule(v.Rule, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["server_side_encryption_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ServerSideEncryptionConfiguration_Rule(p *Rule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Rule {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ServerSideEncryptionConfiguration_Rule_ApplyServerSideEncryptionByDefault(v.ApplyServerSideEncryptionByDefault, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["rule"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ServerSideEncryptionConfiguration_Rule_ApplyServerSideEncryptionByDefault(p *ApplyServerSideEncryptionByDefault, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ApplyServerSideEncryptionByDefault {
		ctyVal = make(map[string]cty.Value)
		EncodeS3Bucket_ServerSideEncryptionConfiguration_Rule_ApplyServerSideEncryptionByDefault_KmsMasterKeyId(v, ctyVal)
		EncodeS3Bucket_ServerSideEncryptionConfiguration_Rule_ApplyServerSideEncryptionByDefault_SseAlgorithm(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["apply_server_side_encryption_by_default"] = cty.ListVal(valsForCollection)
}

func EncodeS3Bucket_ServerSideEncryptionConfiguration_Rule_ApplyServerSideEncryptionByDefault_KmsMasterKeyId(p *ApplyServerSideEncryptionByDefault, vals map[string]cty.Value) {
	vals["kms_master_key_id"] = cty.StringVal(p.KmsMasterKeyId)
}

func EncodeS3Bucket_ServerSideEncryptionConfiguration_Rule_ApplyServerSideEncryptionByDefault_SseAlgorithm(p *ApplyServerSideEncryptionByDefault, vals map[string]cty.Value) {
	vals["sse_algorithm"] = cty.StringVal(p.SseAlgorithm)
}

func EncodeS3Bucket_BucketDomainName(p *S3BucketObservation, vals map[string]cty.Value) {
	vals["bucket_domain_name"] = cty.StringVal(p.BucketDomainName)
}

func EncodeS3Bucket_BucketRegionalDomainName(p *S3BucketObservation, vals map[string]cty.Value) {
	vals["bucket_regional_domain_name"] = cty.StringVal(p.BucketRegionalDomainName)
}

func EncodeS3Bucket_Region(p *S3BucketObservation, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}