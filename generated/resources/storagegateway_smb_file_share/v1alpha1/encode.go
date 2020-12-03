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

package v1alpha1func EncodeStoragegatewaySmbFileShare(r StoragegatewaySmbFileShare) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeStoragegatewaySmbFileShare_GuessMimeTypeEnabled(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_KmsEncrypted(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_AdminUserList(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_AuditDestinationArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_Authentication(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_DefaultStorageClass(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_KmsKeyArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_RequesterPays(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_SmbAclEnabled(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_Tags(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_CaseSensitivity(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_Id(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_InvalidUserList(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_ObjectAcl(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_ValidUserList(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_GatewayArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_LocationArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_ReadOnly(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_CacheAttributes(r.Spec.ForProvider.CacheAttributes, ctyVal)
	EncodeStoragegatewaySmbFileShare_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeStoragegatewaySmbFileShare_FileshareId(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_Path(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewaySmbFileShare_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeStoragegatewaySmbFileShare_GuessMimeTypeEnabled(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["guess_mime_type_enabled"] = cty.BoolVal(p.GuessMimeTypeEnabled)
}

func EncodeStoragegatewaySmbFileShare_KmsEncrypted(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["kms_encrypted"] = cty.BoolVal(p.KmsEncrypted)
}

func EncodeStoragegatewaySmbFileShare_AdminUserList(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AdminUserList {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["admin_user_list"] = cty.SetVal(colVals)
}

func EncodeStoragegatewaySmbFileShare_AuditDestinationArn(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["audit_destination_arn"] = cty.StringVal(p.AuditDestinationArn)
}

func EncodeStoragegatewaySmbFileShare_Authentication(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["authentication"] = cty.StringVal(p.Authentication)
}

func EncodeStoragegatewaySmbFileShare_DefaultStorageClass(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["default_storage_class"] = cty.StringVal(p.DefaultStorageClass)
}

func EncodeStoragegatewaySmbFileShare_KmsKeyArn(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeStoragegatewaySmbFileShare_RequesterPays(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["requester_pays"] = cty.BoolVal(p.RequesterPays)
}

func EncodeStoragegatewaySmbFileShare_SmbAclEnabled(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["smb_acl_enabled"] = cty.BoolVal(p.SmbAclEnabled)
}

func EncodeStoragegatewaySmbFileShare_Tags(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeStoragegatewaySmbFileShare_CaseSensitivity(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["case_sensitivity"] = cty.StringVal(p.CaseSensitivity)
}

func EncodeStoragegatewaySmbFileShare_Id(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeStoragegatewaySmbFileShare_InvalidUserList(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.InvalidUserList {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["invalid_user_list"] = cty.SetVal(colVals)
}

func EncodeStoragegatewaySmbFileShare_ObjectAcl(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["object_acl"] = cty.StringVal(p.ObjectAcl)
}

func EncodeStoragegatewaySmbFileShare_ValidUserList(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ValidUserList {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["valid_user_list"] = cty.SetVal(colVals)
}

func EncodeStoragegatewaySmbFileShare_GatewayArn(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["gateway_arn"] = cty.StringVal(p.GatewayArn)
}

func EncodeStoragegatewaySmbFileShare_LocationArn(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["location_arn"] = cty.StringVal(p.LocationArn)
}

func EncodeStoragegatewaySmbFileShare_ReadOnly(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["read_only"] = cty.BoolVal(p.ReadOnly)
}

func EncodeStoragegatewaySmbFileShare_RoleArn(p *StoragegatewaySmbFileShareParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeStoragegatewaySmbFileShare_CacheAttributes(p *CacheAttributes, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CacheAttributes {
		ctyVal = make(map[string]cty.Value)
		EncodeStoragegatewaySmbFileShare_CacheAttributes_CacheStaleTimeoutInSeconds(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cache_attributes"] = cty.ListVal(valsForCollection)
}

func EncodeStoragegatewaySmbFileShare_CacheAttributes_CacheStaleTimeoutInSeconds(p *CacheAttributes, vals map[string]cty.Value) {
	vals["cache_stale_timeout_in_seconds"] = cty.IntVal(p.CacheStaleTimeoutInSeconds)
}

func EncodeStoragegatewaySmbFileShare_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeStoragegatewaySmbFileShare_Timeouts_Delete(p, ctyVal)
	EncodeStoragegatewaySmbFileShare_Timeouts_Update(p, ctyVal)
	EncodeStoragegatewaySmbFileShare_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeStoragegatewaySmbFileShare_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeStoragegatewaySmbFileShare_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeStoragegatewaySmbFileShare_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeStoragegatewaySmbFileShare_FileshareId(p *StoragegatewaySmbFileShareObservation, vals map[string]cty.Value) {
	vals["fileshare_id"] = cty.StringVal(p.FileshareId)
}

func EncodeStoragegatewaySmbFileShare_Path(p *StoragegatewaySmbFileShareObservation, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeStoragegatewaySmbFileShare_Arn(p *StoragegatewaySmbFileShareObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}