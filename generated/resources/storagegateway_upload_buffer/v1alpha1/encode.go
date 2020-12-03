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

package v1alpha1func EncodeStoragegatewayUploadBuffer(r StoragegatewayUploadBuffer) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeStoragegatewayUploadBuffer_DiskId(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayUploadBuffer_GatewayArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayUploadBuffer_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeStoragegatewayUploadBuffer_DiskId(p *StoragegatewayUploadBufferParameters, vals map[string]cty.Value) {
	vals["disk_id"] = cty.StringVal(p.DiskId)
}

func EncodeStoragegatewayUploadBuffer_GatewayArn(p *StoragegatewayUploadBufferParameters, vals map[string]cty.Value) {
	vals["gateway_arn"] = cty.StringVal(p.GatewayArn)
}

func EncodeStoragegatewayUploadBuffer_Id(p *StoragegatewayUploadBufferParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}