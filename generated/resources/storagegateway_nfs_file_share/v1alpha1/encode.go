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
	r, ok := mr.(*StoragegatewayNfsFileShare)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a StoragegatewayNfsFileShare.")
	}
	return EncodeStoragegatewayNfsFileShare(*r), nil
}

func EncodeStoragegatewayNfsFileShare(r StoragegatewayNfsFileShare) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeStoragegatewayNfsFileShare_RequesterPays(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_Id(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_KmsKeyArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_LocationArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_ReadOnly(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_Tags(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_DefaultStorageClass(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_ObjectAcl(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_ClientList(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_GatewayArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_GuessMimeTypeEnabled(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_KmsEncrypted(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_Squash(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeStoragegatewayNfsFileShare_CacheAttributes(r.Spec.ForProvider.CacheAttributes, ctyVal)
	EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults(r.Spec.ForProvider.NfsFileShareDefaults, ctyVal)
	EncodeStoragegatewayNfsFileShare_Arn(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_Path(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewayNfsFileShare_FileshareId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeStoragegatewayNfsFileShare_RequesterPays(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["requester_pays"] = cty.BoolVal(p.RequesterPays)
}

func EncodeStoragegatewayNfsFileShare_Id(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeStoragegatewayNfsFileShare_KmsKeyArn(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeStoragegatewayNfsFileShare_LocationArn(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["location_arn"] = cty.StringVal(p.LocationArn)
}

func EncodeStoragegatewayNfsFileShare_ReadOnly(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["read_only"] = cty.BoolVal(p.ReadOnly)
}

func EncodeStoragegatewayNfsFileShare_RoleArn(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeStoragegatewayNfsFileShare_Tags(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeStoragegatewayNfsFileShare_DefaultStorageClass(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["default_storage_class"] = cty.StringVal(p.DefaultStorageClass)
}

func EncodeStoragegatewayNfsFileShare_ObjectAcl(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["object_acl"] = cty.StringVal(p.ObjectAcl)
}

func EncodeStoragegatewayNfsFileShare_ClientList(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ClientList {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["client_list"] = cty.SetVal(colVals)
}

func EncodeStoragegatewayNfsFileShare_GatewayArn(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["gateway_arn"] = cty.StringVal(p.GatewayArn)
}

func EncodeStoragegatewayNfsFileShare_GuessMimeTypeEnabled(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["guess_mime_type_enabled"] = cty.BoolVal(p.GuessMimeTypeEnabled)
}

func EncodeStoragegatewayNfsFileShare_KmsEncrypted(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["kms_encrypted"] = cty.BoolVal(p.KmsEncrypted)
}

func EncodeStoragegatewayNfsFileShare_Squash(p StoragegatewayNfsFileShareParameters, vals map[string]cty.Value) {
	vals["squash"] = cty.StringVal(p.Squash)
}

func EncodeStoragegatewayNfsFileShare_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeStoragegatewayNfsFileShare_Timeouts_Create(p, ctyVal)
	EncodeStoragegatewayNfsFileShare_Timeouts_Delete(p, ctyVal)
	EncodeStoragegatewayNfsFileShare_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeStoragegatewayNfsFileShare_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeStoragegatewayNfsFileShare_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeStoragegatewayNfsFileShare_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeStoragegatewayNfsFileShare_CacheAttributes(p CacheAttributes, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeStoragegatewayNfsFileShare_CacheAttributes_CacheStaleTimeoutInSeconds(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cache_attributes"] = cty.ListVal(valsForCollection)
}

func EncodeStoragegatewayNfsFileShare_CacheAttributes_CacheStaleTimeoutInSeconds(p CacheAttributes, vals map[string]cty.Value) {
	vals["cache_stale_timeout_in_seconds"] = cty.NumberIntVal(p.CacheStaleTimeoutInSeconds)
}

func EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults(p NfsFileShareDefaults, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults_DirectoryMode(p, ctyVal)
	EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults_FileMode(p, ctyVal)
	EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults_GroupId(p, ctyVal)
	EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults_OwnerId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["nfs_file_share_defaults"] = cty.ListVal(valsForCollection)
}

func EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults_DirectoryMode(p NfsFileShareDefaults, vals map[string]cty.Value) {
	vals["directory_mode"] = cty.StringVal(p.DirectoryMode)
}

func EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults_FileMode(p NfsFileShareDefaults, vals map[string]cty.Value) {
	vals["file_mode"] = cty.StringVal(p.FileMode)
}

func EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults_GroupId(p NfsFileShareDefaults, vals map[string]cty.Value) {
	vals["group_id"] = cty.NumberIntVal(p.GroupId)
}

func EncodeStoragegatewayNfsFileShare_NfsFileShareDefaults_OwnerId(p NfsFileShareDefaults, vals map[string]cty.Value) {
	vals["owner_id"] = cty.NumberIntVal(p.OwnerId)
}

func EncodeStoragegatewayNfsFileShare_Arn(p StoragegatewayNfsFileShareObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeStoragegatewayNfsFileShare_Path(p StoragegatewayNfsFileShareObservation, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeStoragegatewayNfsFileShare_FileshareId(p StoragegatewayNfsFileShareObservation, vals map[string]cty.Value) {
	vals["fileshare_id"] = cty.StringVal(p.FileshareId)
}