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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*GlacierVault)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GlacierVault.")
	}
	return EncodeGlacierVault(*r), nil
}

func EncodeGlacierVault(r GlacierVault) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlacierVault_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlacierVault_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlacierVault_AccessPolicy(r.Spec.ForProvider, ctyVal)
	EncodeGlacierVault_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlacierVault_Notification(r.Spec.ForProvider.Notification, ctyVal)
	EncodeGlacierVault_Location(r.Status.AtProvider, ctyVal)
	EncodeGlacierVault_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeGlacierVault_Name(p GlacierVaultParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlacierVault_Tags(p GlacierVaultParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGlacierVault_AccessPolicy(p GlacierVaultParameters, vals map[string]cty.Value) {
	vals["access_policy"] = cty.StringVal(p.AccessPolicy)
}

func EncodeGlacierVault_Id(p GlacierVaultParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlacierVault_Notification(p Notification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlacierVault_Notification_Events(p, ctyVal)
	EncodeGlacierVault_Notification_SnsTopic(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["notification"] = cty.ListVal(valsForCollection)
}

func EncodeGlacierVault_Notification_Events(p Notification, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.SetVal(colVals)
}

func EncodeGlacierVault_Notification_SnsTopic(p Notification, vals map[string]cty.Value) {
	vals["sns_topic"] = cty.StringVal(p.SnsTopic)
}

func EncodeGlacierVault_Location(p GlacierVaultObservation, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeGlacierVault_Arn(p GlacierVaultObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}