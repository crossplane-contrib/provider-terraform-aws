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
	r, ok := mr.(*DirectoryServiceLogSubscription)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DirectoryServiceLogSubscription.")
	}
	return EncodeDirectoryServiceLogSubscription(*r), nil
}

func EncodeDirectoryServiceLogSubscription(r DirectoryServiceLogSubscription) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDirectoryServiceLogSubscription_LogGroupName(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceLogSubscription_DirectoryId(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceLogSubscription_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeDirectoryServiceLogSubscription_LogGroupName(p DirectoryServiceLogSubscriptionParameters, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeDirectoryServiceLogSubscription_DirectoryId(p DirectoryServiceLogSubscriptionParameters, vals map[string]cty.Value) {
	vals["directory_id"] = cty.StringVal(p.DirectoryId)
}

func EncodeDirectoryServiceLogSubscription_Id(p DirectoryServiceLogSubscriptionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}