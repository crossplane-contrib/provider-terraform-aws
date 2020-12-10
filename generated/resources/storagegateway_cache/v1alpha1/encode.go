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
	r, ok := mr.(*StoragegatewayCache)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a StoragegatewayCache.")
	}
	return EncodeStoragegatewayCache(*r), nil
}

func EncodeStoragegatewayCache(r StoragegatewayCache) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeStoragegatewayCache_DiskId(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCache_GatewayArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCache_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeStoragegatewayCache_DiskId(p StoragegatewayCacheParameters, vals map[string]cty.Value) {
	vals["disk_id"] = cty.StringVal(p.DiskId)
}

func EncodeStoragegatewayCache_GatewayArn(p StoragegatewayCacheParameters, vals map[string]cty.Value) {
	vals["gateway_arn"] = cty.StringVal(p.GatewayArn)
}

func EncodeStoragegatewayCache_Id(p StoragegatewayCacheParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}