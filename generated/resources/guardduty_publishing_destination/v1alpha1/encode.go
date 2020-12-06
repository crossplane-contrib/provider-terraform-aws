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

func EncodeGuarddutyPublishingDestination(r GuarddutyPublishingDestination) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyPublishingDestination_DestinationArn(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyPublishingDestination_DestinationType(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyPublishingDestination_DetectorId(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyPublishingDestination_Id(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyPublishingDestination_KmsKeyArn(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyPublishingDestination_DestinationArn(p GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	vals["destination_arn"] = cty.StringVal(p.DestinationArn)
}

func EncodeGuarddutyPublishingDestination_DestinationType(p GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	vals["destination_type"] = cty.StringVal(p.DestinationType)
}

func EncodeGuarddutyPublishingDestination_DetectorId(p GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	vals["detector_id"] = cty.StringVal(p.DetectorId)
}

func EncodeGuarddutyPublishingDestination_Id(p GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGuarddutyPublishingDestination_KmsKeyArn(p GuarddutyPublishingDestinationParameters, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}