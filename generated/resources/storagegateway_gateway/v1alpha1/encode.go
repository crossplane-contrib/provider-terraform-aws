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
	r, ok := mr.(*StoragegatewayGateway)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a StoragegatewayGateway.")
	}
	return EncodeStoragegatewayGateway(*r), nil
}

func EncodeStoragegatewayGateway(r StoragegatewayGateway) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeStoragegatewayGateway_TapeDriveType(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_AverageDownloadRateLimitInBitsPerSec(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_AverageUploadRateLimitInBitsPerSec(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_SmbSecurityStrategy(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_Tags(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_GatewayType(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_Id(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_ActivationKey(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_CloudwatchLogGroupArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_GatewayIpAddress(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_GatewayTimezone(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_GatewayName(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_GatewayVpcEndpoint(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_MediumChangerType(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_SmbGuestPassword(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayGateway_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeStoragegatewayGateway_SmbActiveDirectorySettings(r.Spec.ForProvider.SmbActiveDirectorySettings, ctyVal)
	EncodeStoragegatewayGateway_Arn(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewayGateway_GatewayId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeStoragegatewayGateway_TapeDriveType(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["tape_drive_type"] = cty.StringVal(p.TapeDriveType)
}

func EncodeStoragegatewayGateway_AverageDownloadRateLimitInBitsPerSec(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["average_download_rate_limit_in_bits_per_sec"] = cty.NumberIntVal(p.AverageDownloadRateLimitInBitsPerSec)
}

func EncodeStoragegatewayGateway_AverageUploadRateLimitInBitsPerSec(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["average_upload_rate_limit_in_bits_per_sec"] = cty.NumberIntVal(p.AverageUploadRateLimitInBitsPerSec)
}

func EncodeStoragegatewayGateway_SmbSecurityStrategy(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["smb_security_strategy"] = cty.StringVal(p.SmbSecurityStrategy)
}

func EncodeStoragegatewayGateway_Tags(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeStoragegatewayGateway_GatewayType(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["gateway_type"] = cty.StringVal(p.GatewayType)
}

func EncodeStoragegatewayGateway_Id(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeStoragegatewayGateway_ActivationKey(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["activation_key"] = cty.StringVal(p.ActivationKey)
}

func EncodeStoragegatewayGateway_CloudwatchLogGroupArn(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["cloudwatch_log_group_arn"] = cty.StringVal(p.CloudwatchLogGroupArn)
}

func EncodeStoragegatewayGateway_GatewayIpAddress(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["gateway_ip_address"] = cty.StringVal(p.GatewayIpAddress)
}

func EncodeStoragegatewayGateway_GatewayTimezone(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["gateway_timezone"] = cty.StringVal(p.GatewayTimezone)
}

func EncodeStoragegatewayGateway_GatewayName(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["gateway_name"] = cty.StringVal(p.GatewayName)
}

func EncodeStoragegatewayGateway_GatewayVpcEndpoint(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["gateway_vpc_endpoint"] = cty.StringVal(p.GatewayVpcEndpoint)
}

func EncodeStoragegatewayGateway_MediumChangerType(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["medium_changer_type"] = cty.StringVal(p.MediumChangerType)
}

func EncodeStoragegatewayGateway_SmbGuestPassword(p StoragegatewayGatewayParameters, vals map[string]cty.Value) {
	vals["smb_guest_password"] = cty.StringVal(p.SmbGuestPassword)
}

func EncodeStoragegatewayGateway_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeStoragegatewayGateway_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeStoragegatewayGateway_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeStoragegatewayGateway_SmbActiveDirectorySettings(p SmbActiveDirectorySettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeStoragegatewayGateway_SmbActiveDirectorySettings_Password(p, ctyVal)
	EncodeStoragegatewayGateway_SmbActiveDirectorySettings_Username(p, ctyVal)
	EncodeStoragegatewayGateway_SmbActiveDirectorySettings_DomainName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["smb_active_directory_settings"] = cty.ListVal(valsForCollection)
}

func EncodeStoragegatewayGateway_SmbActiveDirectorySettings_Password(p SmbActiveDirectorySettings, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeStoragegatewayGateway_SmbActiveDirectorySettings_Username(p SmbActiveDirectorySettings, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeStoragegatewayGateway_SmbActiveDirectorySettings_DomainName(p SmbActiveDirectorySettings, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeStoragegatewayGateway_Arn(p StoragegatewayGatewayObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeStoragegatewayGateway_GatewayId(p StoragegatewayGatewayObservation, vals map[string]cty.Value) {
	vals["gateway_id"] = cty.StringVal(p.GatewayId)
}