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

// hand-written overlay for iam decoding

package v1alpha1

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
)

type ctyDecoder struct{}

func (d *ctyDecoder) DecodeCty(previousManaged resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	prev, ok := previousManaged.(*IamUser)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	new.Spec.ForProvider.Name = valMap["name"].AsString()
	new.Spec.ForProvider.Id = valMap["id"].AsString()
	new.Spec.ForProvider.Path = valMap["path"].AsString()
	new.Spec.ForProvider.PermissionsBoundary = valMap["permissions_boundary"].AsString()
	if !valMap["force_destroy"].IsNull() {
		new.Spec.ForProvider.ForceDestroy = valMap["force_destroy"].True()
	}
	new.Status.AtProvider.UniqueId = valMap["unique_id"].AsString()
	new.Status.AtProvider.Arn = valMap["arn"].AsString()
	meta.SetExternalName(new, valMap["id"].AsString())
	tags := make(map[string]string)
	for k, v := range valMap["tags"].AsValueMap() {
		tags[k] = v.AsString()
	}
	new.Spec.ForProvider.Tags = tags
	return new, nil
}
