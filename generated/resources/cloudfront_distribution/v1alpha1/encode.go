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

package v1alpha1func EncodeCloudfrontDistribution(r CloudfrontDistribution) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCloudfrontDistribution_Comment(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_DefaultRootObject(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_HttpVersion(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_PriceClass(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_WaitForDeployment(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_WebAclId(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_Aliases(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_IsIpv6Enabled(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_RetainOnDelete(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontDistribution_OrderedCacheBehavior(r.Spec.ForProvider.OrderedCacheBehavior, ctyVal)
	EncodeCloudfrontDistribution_Origin(r.Spec.ForProvider.Origin, ctyVal)
	EncodeCloudfrontDistribution_OriginGroup(r.Spec.ForProvider.OriginGroup, ctyVal)
	EncodeCloudfrontDistribution_Restrictions(r.Spec.ForProvider.Restrictions, ctyVal)
	EncodeCloudfrontDistribution_ViewerCertificate(r.Spec.ForProvider.ViewerCertificate, ctyVal)
	EncodeCloudfrontDistribution_CustomErrorResponse(r.Spec.ForProvider.CustomErrorResponse, ctyVal)
	EncodeCloudfrontDistribution_DefaultCacheBehavior(r.Spec.ForProvider.DefaultCacheBehavior, ctyVal)
	EncodeCloudfrontDistribution_LoggingConfig(r.Spec.ForProvider.LoggingConfig, ctyVal)
	EncodeCloudfrontDistribution_InProgressValidationBatches(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontDistribution_Status(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontDistribution_DomainName(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontDistribution_HostedZoneId(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontDistribution_LastModifiedTime(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontDistribution_CallerReference(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontDistribution_TrustedSigners(r.Status.AtProvider.TrustedSigners, ctyVal)
	EncodeCloudfrontDistribution_Arn(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontDistribution_Etag(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeCloudfrontDistribution_Comment(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["comment"] = cty.StringVal(p.Comment)
}

func EncodeCloudfrontDistribution_Enabled(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeCloudfrontDistribution_DefaultRootObject(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["default_root_object"] = cty.StringVal(p.DefaultRootObject)
}

func EncodeCloudfrontDistribution_HttpVersion(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["http_version"] = cty.StringVal(p.HttpVersion)
}

func EncodeCloudfrontDistribution_PriceClass(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["price_class"] = cty.StringVal(p.PriceClass)
}

func EncodeCloudfrontDistribution_Tags(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudfrontDistribution_Id(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudfrontDistribution_WaitForDeployment(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["wait_for_deployment"] = cty.BoolVal(p.WaitForDeployment)
}

func EncodeCloudfrontDistribution_WebAclId(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["web_acl_id"] = cty.StringVal(p.WebAclId)
}

func EncodeCloudfrontDistribution_Aliases(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Aliases {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["aliases"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_IsIpv6Enabled(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["is_ipv6_enabled"] = cty.BoolVal(p.IsIpv6Enabled)
}

func EncodeCloudfrontDistribution_RetainOnDelete(p *CloudfrontDistributionParameters, vals map[string]cty.Value) {
	vals["retain_on_delete"] = cty.BoolVal(p.RetainOnDelete)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OrderedCacheBehavior {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_DefaultTtl(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_FieldLevelEncryptionId(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_MaxTtl(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_MinTtl(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_PathPattern(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_SmoothStreaming(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_AllowedMethods(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_CachedMethods(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_TrustedSigners(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_ViewerProtocolPolicy(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_Compress(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_TargetOriginId(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues(v.ForwardedValues, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_LambdaFunctionAssociation(v.LambdaFunctionAssociation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ordered_cache_behavior"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_DefaultTtl(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	vals["default_ttl"] = cty.IntVal(p.DefaultTtl)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_FieldLevelEncryptionId(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	vals["field_level_encryption_id"] = cty.StringVal(p.FieldLevelEncryptionId)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_MaxTtl(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	vals["max_ttl"] = cty.IntVal(p.MaxTtl)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_MinTtl(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	vals["min_ttl"] = cty.IntVal(p.MinTtl)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_PathPattern(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	vals["path_pattern"] = cty.StringVal(p.PathPattern)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_SmoothStreaming(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	vals["smooth_streaming"] = cty.BoolVal(p.SmoothStreaming)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_AllowedMethods(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedMethods {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_methods"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_CachedMethods(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CachedMethods {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cached_methods"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_TrustedSigners(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TrustedSigners {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["trusted_signers"] = cty.ListVal(colVals)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_ViewerProtocolPolicy(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	vals["viewer_protocol_policy"] = cty.StringVal(p.ViewerProtocolPolicy)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_Compress(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	vals["compress"] = cty.BoolVal(p.Compress)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_TargetOriginId(p *OrderedCacheBehavior, vals map[string]cty.Value) {
	vals["target_origin_id"] = cty.StringVal(p.TargetOriginId)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues(p *ForwardedValues, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedValues {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_Headers(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_QueryString(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_QueryStringCacheKeys(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_Cookies(v.Cookies, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_values"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_Headers(p *ForwardedValues, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Headers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["headers"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_QueryString(p *ForwardedValues, vals map[string]cty.Value) {
	vals["query_string"] = cty.BoolVal(p.QueryString)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_QueryStringCacheKeys(p *ForwardedValues, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.QueryStringCacheKeys {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["query_string_cache_keys"] = cty.ListVal(colVals)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_Cookies(p *Cookies, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Cookies {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_Cookies_Forward(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_Cookies_WhitelistedNames(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cookies"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_Cookies_Forward(p *Cookies, vals map[string]cty.Value) {
	vals["forward"] = cty.StringVal(p.Forward)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_ForwardedValues_Cookies_WhitelistedNames(p *Cookies, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.WhitelistedNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["whitelisted_names"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_LambdaFunctionAssociation(p *LambdaFunctionAssociation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LambdaFunctionAssociation {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_LambdaFunctionAssociation_EventType(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_LambdaFunctionAssociation_IncludeBody(v, ctyVal)
		EncodeCloudfrontDistribution_OrderedCacheBehavior_LambdaFunctionAssociation_LambdaArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lambda_function_association"] = cty.SetVal(valsForCollection)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_LambdaFunctionAssociation_EventType(p *LambdaFunctionAssociation, vals map[string]cty.Value) {
	vals["event_type"] = cty.StringVal(p.EventType)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_LambdaFunctionAssociation_IncludeBody(p *LambdaFunctionAssociation, vals map[string]cty.Value) {
	vals["include_body"] = cty.BoolVal(p.IncludeBody)
}

func EncodeCloudfrontDistribution_OrderedCacheBehavior_LambdaFunctionAssociation_LambdaArn(p *LambdaFunctionAssociation, vals map[string]cty.Value) {
	vals["lambda_arn"] = cty.StringVal(p.LambdaArn)
}

func EncodeCloudfrontDistribution_Origin(p *Origin, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Origin {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_Origin_DomainName(v, ctyVal)
		EncodeCloudfrontDistribution_Origin_OriginId(v, ctyVal)
		EncodeCloudfrontDistribution_Origin_OriginPath(v, ctyVal)
		EncodeCloudfrontDistribution_Origin_S3OriginConfig(v.S3OriginConfig, ctyVal)
		EncodeCloudfrontDistribution_Origin_CustomHeader(v.CustomHeader, ctyVal)
		EncodeCloudfrontDistribution_Origin_CustomOriginConfig(v.CustomOriginConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["origin"] = cty.SetVal(valsForCollection)
}

func EncodeCloudfrontDistribution_Origin_DomainName(p *Origin, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeCloudfrontDistribution_Origin_OriginId(p *Origin, vals map[string]cty.Value) {
	vals["origin_id"] = cty.StringVal(p.OriginId)
}

func EncodeCloudfrontDistribution_Origin_OriginPath(p *Origin, vals map[string]cty.Value) {
	vals["origin_path"] = cty.StringVal(p.OriginPath)
}

func EncodeCloudfrontDistribution_Origin_S3OriginConfig(p *S3OriginConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3OriginConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_Origin_S3OriginConfig_OriginAccessIdentity(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3_origin_config"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_Origin_S3OriginConfig_OriginAccessIdentity(p *S3OriginConfig, vals map[string]cty.Value) {
	vals["origin_access_identity"] = cty.StringVal(p.OriginAccessIdentity)
}

func EncodeCloudfrontDistribution_Origin_CustomHeader(p *CustomHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CustomHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_Origin_CustomHeader_Name(v, ctyVal)
		EncodeCloudfrontDistribution_Origin_CustomHeader_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["custom_header"] = cty.SetVal(valsForCollection)
}

func EncodeCloudfrontDistribution_Origin_CustomHeader_Name(p *CustomHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudfrontDistribution_Origin_CustomHeader_Value(p *CustomHeader, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeCloudfrontDistribution_Origin_CustomOriginConfig(p *CustomOriginConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CustomOriginConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_Origin_CustomOriginConfig_OriginKeepaliveTimeout(v, ctyVal)
		EncodeCloudfrontDistribution_Origin_CustomOriginConfig_OriginProtocolPolicy(v, ctyVal)
		EncodeCloudfrontDistribution_Origin_CustomOriginConfig_OriginReadTimeout(v, ctyVal)
		EncodeCloudfrontDistribution_Origin_CustomOriginConfig_OriginSslProtocols(v, ctyVal)
		EncodeCloudfrontDistribution_Origin_CustomOriginConfig_HttpPort(v, ctyVal)
		EncodeCloudfrontDistribution_Origin_CustomOriginConfig_HttpsPort(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["custom_origin_config"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_Origin_CustomOriginConfig_OriginKeepaliveTimeout(p *CustomOriginConfig, vals map[string]cty.Value) {
	vals["origin_keepalive_timeout"] = cty.IntVal(p.OriginKeepaliveTimeout)
}

func EncodeCloudfrontDistribution_Origin_CustomOriginConfig_OriginProtocolPolicy(p *CustomOriginConfig, vals map[string]cty.Value) {
	vals["origin_protocol_policy"] = cty.StringVal(p.OriginProtocolPolicy)
}

func EncodeCloudfrontDistribution_Origin_CustomOriginConfig_OriginReadTimeout(p *CustomOriginConfig, vals map[string]cty.Value) {
	vals["origin_read_timeout"] = cty.IntVal(p.OriginReadTimeout)
}

func EncodeCloudfrontDistribution_Origin_CustomOriginConfig_OriginSslProtocols(p *CustomOriginConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.OriginSslProtocols {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["origin_ssl_protocols"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_Origin_CustomOriginConfig_HttpPort(p *CustomOriginConfig, vals map[string]cty.Value) {
	vals["http_port"] = cty.IntVal(p.HttpPort)
}

func EncodeCloudfrontDistribution_Origin_CustomOriginConfig_HttpsPort(p *CustomOriginConfig, vals map[string]cty.Value) {
	vals["https_port"] = cty.IntVal(p.HttpsPort)
}

func EncodeCloudfrontDistribution_OriginGroup(p *OriginGroup, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OriginGroup {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_OriginGroup_OriginId(v, ctyVal)
		EncodeCloudfrontDistribution_OriginGroup_FailoverCriteria(v.FailoverCriteria, ctyVal)
		EncodeCloudfrontDistribution_OriginGroup_Member(v.Member, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["origin_group"] = cty.SetVal(valsForCollection)
}

func EncodeCloudfrontDistribution_OriginGroup_OriginId(p *OriginGroup, vals map[string]cty.Value) {
	vals["origin_id"] = cty.StringVal(p.OriginId)
}

func EncodeCloudfrontDistribution_OriginGroup_FailoverCriteria(p *FailoverCriteria, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FailoverCriteria {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_OriginGroup_FailoverCriteria_StatusCodes(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["failover_criteria"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_OriginGroup_FailoverCriteria_StatusCodes(p *FailoverCriteria, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.StatusCodes {
		colVals = append(colVals, cty.IntVal(value))
	}
	vals["status_codes"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_OriginGroup_Member(p *Member, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Member {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_OriginGroup_Member_OriginId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["member"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_OriginGroup_Member_OriginId(p *Member, vals map[string]cty.Value) {
	vals["origin_id"] = cty.StringVal(p.OriginId)
}

func EncodeCloudfrontDistribution_Restrictions(p *Restrictions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Restrictions {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_Restrictions_GeoRestriction(v.GeoRestriction, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["restrictions"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_Restrictions_GeoRestriction(p *GeoRestriction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoRestriction {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_Restrictions_GeoRestriction_Locations(v, ctyVal)
		EncodeCloudfrontDistribution_Restrictions_GeoRestriction_RestrictionType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_restriction"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_Restrictions_GeoRestriction_Locations(p *GeoRestriction, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Locations {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["locations"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_Restrictions_GeoRestriction_RestrictionType(p *GeoRestriction, vals map[string]cty.Value) {
	vals["restriction_type"] = cty.StringVal(p.RestrictionType)
}

func EncodeCloudfrontDistribution_ViewerCertificate(p *ViewerCertificate, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ViewerCertificate {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_ViewerCertificate_AcmCertificateArn(v, ctyVal)
		EncodeCloudfrontDistribution_ViewerCertificate_CloudfrontDefaultCertificate(v, ctyVal)
		EncodeCloudfrontDistribution_ViewerCertificate_IamCertificateId(v, ctyVal)
		EncodeCloudfrontDistribution_ViewerCertificate_MinimumProtocolVersion(v, ctyVal)
		EncodeCloudfrontDistribution_ViewerCertificate_SslSupportMethod(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["viewer_certificate"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_ViewerCertificate_AcmCertificateArn(p *ViewerCertificate, vals map[string]cty.Value) {
	vals["acm_certificate_arn"] = cty.StringVal(p.AcmCertificateArn)
}

func EncodeCloudfrontDistribution_ViewerCertificate_CloudfrontDefaultCertificate(p *ViewerCertificate, vals map[string]cty.Value) {
	vals["cloudfront_default_certificate"] = cty.BoolVal(p.CloudfrontDefaultCertificate)
}

func EncodeCloudfrontDistribution_ViewerCertificate_IamCertificateId(p *ViewerCertificate, vals map[string]cty.Value) {
	vals["iam_certificate_id"] = cty.StringVal(p.IamCertificateId)
}

func EncodeCloudfrontDistribution_ViewerCertificate_MinimumProtocolVersion(p *ViewerCertificate, vals map[string]cty.Value) {
	vals["minimum_protocol_version"] = cty.StringVal(p.MinimumProtocolVersion)
}

func EncodeCloudfrontDistribution_ViewerCertificate_SslSupportMethod(p *ViewerCertificate, vals map[string]cty.Value) {
	vals["ssl_support_method"] = cty.StringVal(p.SslSupportMethod)
}

func EncodeCloudfrontDistribution_CustomErrorResponse(p *CustomErrorResponse, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CustomErrorResponse {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_CustomErrorResponse_ResponseCode(v, ctyVal)
		EncodeCloudfrontDistribution_CustomErrorResponse_ResponsePagePath(v, ctyVal)
		EncodeCloudfrontDistribution_CustomErrorResponse_ErrorCachingMinTtl(v, ctyVal)
		EncodeCloudfrontDistribution_CustomErrorResponse_ErrorCode(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["custom_error_response"] = cty.SetVal(valsForCollection)
}

func EncodeCloudfrontDistribution_CustomErrorResponse_ResponseCode(p *CustomErrorResponse, vals map[string]cty.Value) {
	vals["response_code"] = cty.IntVal(p.ResponseCode)
}

func EncodeCloudfrontDistribution_CustomErrorResponse_ResponsePagePath(p *CustomErrorResponse, vals map[string]cty.Value) {
	vals["response_page_path"] = cty.StringVal(p.ResponsePagePath)
}

func EncodeCloudfrontDistribution_CustomErrorResponse_ErrorCachingMinTtl(p *CustomErrorResponse, vals map[string]cty.Value) {
	vals["error_caching_min_ttl"] = cty.IntVal(p.ErrorCachingMinTtl)
}

func EncodeCloudfrontDistribution_CustomErrorResponse_ErrorCode(p *CustomErrorResponse, vals map[string]cty.Value) {
	vals["error_code"] = cty.IntVal(p.ErrorCode)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DefaultCacheBehavior {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_CachedMethods(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_Compress(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_MaxTtl(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_TrustedSigners(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_AllowedMethods(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_DefaultTtl(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_FieldLevelEncryptionId(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_MinTtl(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_SmoothStreaming(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_TargetOriginId(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_ViewerProtocolPolicy(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues(v.ForwardedValues, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_LambdaFunctionAssociation(v.LambdaFunctionAssociation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["default_cache_behavior"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_CachedMethods(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CachedMethods {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cached_methods"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_Compress(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	vals["compress"] = cty.BoolVal(p.Compress)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_MaxTtl(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	vals["max_ttl"] = cty.IntVal(p.MaxTtl)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_TrustedSigners(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TrustedSigners {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["trusted_signers"] = cty.ListVal(colVals)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_AllowedMethods(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedMethods {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_methods"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_DefaultTtl(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	vals["default_ttl"] = cty.IntVal(p.DefaultTtl)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_FieldLevelEncryptionId(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	vals["field_level_encryption_id"] = cty.StringVal(p.FieldLevelEncryptionId)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_MinTtl(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	vals["min_ttl"] = cty.IntVal(p.MinTtl)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_SmoothStreaming(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	vals["smooth_streaming"] = cty.BoolVal(p.SmoothStreaming)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_TargetOriginId(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	vals["target_origin_id"] = cty.StringVal(p.TargetOriginId)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_ViewerProtocolPolicy(p *DefaultCacheBehavior, vals map[string]cty.Value) {
	vals["viewer_protocol_policy"] = cty.StringVal(p.ViewerProtocolPolicy)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues(p *ForwardedValues, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedValues {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_Headers(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_QueryString(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_QueryStringCacheKeys(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_Cookies(v.Cookies, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_values"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_Headers(p *ForwardedValues, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Headers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["headers"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_QueryString(p *ForwardedValues, vals map[string]cty.Value) {
	vals["query_string"] = cty.BoolVal(p.QueryString)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_QueryStringCacheKeys(p *ForwardedValues, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.QueryStringCacheKeys {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["query_string_cache_keys"] = cty.ListVal(colVals)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_Cookies(p *Cookies, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Cookies {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_Cookies_Forward(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_Cookies_WhitelistedNames(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cookies"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_Cookies_Forward(p *Cookies, vals map[string]cty.Value) {
	vals["forward"] = cty.StringVal(p.Forward)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_ForwardedValues_Cookies_WhitelistedNames(p *Cookies, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.WhitelistedNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["whitelisted_names"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_LambdaFunctionAssociation(p *LambdaFunctionAssociation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LambdaFunctionAssociation {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_LambdaFunctionAssociation_LambdaArn(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_LambdaFunctionAssociation_EventType(v, ctyVal)
		EncodeCloudfrontDistribution_DefaultCacheBehavior_LambdaFunctionAssociation_IncludeBody(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lambda_function_association"] = cty.SetVal(valsForCollection)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_LambdaFunctionAssociation_LambdaArn(p *LambdaFunctionAssociation, vals map[string]cty.Value) {
	vals["lambda_arn"] = cty.StringVal(p.LambdaArn)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_LambdaFunctionAssociation_EventType(p *LambdaFunctionAssociation, vals map[string]cty.Value) {
	vals["event_type"] = cty.StringVal(p.EventType)
}

func EncodeCloudfrontDistribution_DefaultCacheBehavior_LambdaFunctionAssociation_IncludeBody(p *LambdaFunctionAssociation, vals map[string]cty.Value) {
	vals["include_body"] = cty.BoolVal(p.IncludeBody)
}

func EncodeCloudfrontDistribution_LoggingConfig(p *LoggingConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LoggingConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_LoggingConfig_Bucket(v, ctyVal)
		EncodeCloudfrontDistribution_LoggingConfig_IncludeCookies(v, ctyVal)
		EncodeCloudfrontDistribution_LoggingConfig_Prefix(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["logging_config"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_LoggingConfig_Bucket(p *LoggingConfig, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeCloudfrontDistribution_LoggingConfig_IncludeCookies(p *LoggingConfig, vals map[string]cty.Value) {
	vals["include_cookies"] = cty.BoolVal(p.IncludeCookies)
}

func EncodeCloudfrontDistribution_LoggingConfig_Prefix(p *LoggingConfig, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeCloudfrontDistribution_InProgressValidationBatches(p *CloudfrontDistributionObservation, vals map[string]cty.Value) {
	vals["in_progress_validation_batches"] = cty.IntVal(p.InProgressValidationBatches)
}

func EncodeCloudfrontDistribution_Status(p *CloudfrontDistributionObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeCloudfrontDistribution_DomainName(p *CloudfrontDistributionObservation, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeCloudfrontDistribution_HostedZoneId(p *CloudfrontDistributionObservation, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}

func EncodeCloudfrontDistribution_LastModifiedTime(p *CloudfrontDistributionObservation, vals map[string]cty.Value) {
	vals["last_modified_time"] = cty.StringVal(p.LastModifiedTime)
}

func EncodeCloudfrontDistribution_CallerReference(p *CloudfrontDistributionObservation, vals map[string]cty.Value) {
	vals["caller_reference"] = cty.StringVal(p.CallerReference)
}

func EncodeCloudfrontDistribution_TrustedSigners(p *TrustedSigners, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TrustedSigners {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_TrustedSigners_Enabled(v, ctyVal)
		EncodeCloudfrontDistribution_TrustedSigners_Items(v.Items, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["trusted_signers"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_TrustedSigners_Enabled(p *TrustedSigners, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeCloudfrontDistribution_TrustedSigners_Items(p *Items, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Items {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudfrontDistribution_TrustedSigners_Items_AwsAccountNumber(v, ctyVal)
		EncodeCloudfrontDistribution_TrustedSigners_Items_KeyPairIds(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["items"] = cty.ListVal(valsForCollection)
}

func EncodeCloudfrontDistribution_TrustedSigners_Items_AwsAccountNumber(p *Items, vals map[string]cty.Value) {
	vals["aws_account_number"] = cty.StringVal(p.AwsAccountNumber)
}

func EncodeCloudfrontDistribution_TrustedSigners_Items_KeyPairIds(p *Items, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.KeyPairIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["key_pair_ids"] = cty.SetVal(colVals)
}

func EncodeCloudfrontDistribution_Arn(p *CloudfrontDistributionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCloudfrontDistribution_Etag(p *CloudfrontDistributionObservation, vals map[string]cty.Value) {
	vals["etag"] = cty.StringVal(p.Etag)
}