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
	"github.com/zclconf/go-cty/cty"
)

func EncodeGlacierVault(r GlacierVault) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlacierVault_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlacierVault_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlacierVault_AccessPolicy(r.Spec.ForProvider, ctyVal)
	EncodeGlacierVault_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlacierVault_Notification(r.Spec.ForProvider.Notification, ctyVal)
	EncodeGlacierVault_Arn(r.Status.AtProvider, ctyVal)
	EncodeGlacierVault_Location(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlacierVault_Name(p GlacierVaultParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlacierVault_Tags(p GlacierVaultParameters, vals map[string]cty.Value) {
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
	EncodeGlacierVault_Notification_SnsTopic(p, ctyVal)
	EncodeGlacierVault_Notification_Events(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["notification"] = cty.ListVal(valsForCollection)
}

func EncodeGlacierVault_Notification_SnsTopic(p Notification, vals map[string]cty.Value) {
	vals["sns_topic"] = cty.StringVal(p.SnsTopic)
}

func EncodeGlacierVault_Notification_Events(p Notification, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.SetVal(colVals)
}

func EncodeGlacierVault_Arn(p GlacierVaultObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeGlacierVault_Location(p GlacierVaultObservation, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}