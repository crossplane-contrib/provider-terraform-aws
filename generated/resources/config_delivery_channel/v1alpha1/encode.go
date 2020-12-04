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

func EncodeConfigDeliveryChannel(r ConfigDeliveryChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeConfigDeliveryChannel_Id(r.Spec.ForProvider, ctyVal)
	EncodeConfigDeliveryChannel_Name(r.Spec.ForProvider, ctyVal)
	EncodeConfigDeliveryChannel_S3BucketName(r.Spec.ForProvider, ctyVal)
	EncodeConfigDeliveryChannel_S3KeyPrefix(r.Spec.ForProvider, ctyVal)
	EncodeConfigDeliveryChannel_SnsTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeConfigDeliveryChannel_SnapshotDeliveryProperties(r.Spec.ForProvider.SnapshotDeliveryProperties, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeConfigDeliveryChannel_Id(p ConfigDeliveryChannelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeConfigDeliveryChannel_Name(p ConfigDeliveryChannelParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeConfigDeliveryChannel_S3BucketName(p ConfigDeliveryChannelParameters, vals map[string]cty.Value) {
	vals["s3_bucket_name"] = cty.StringVal(p.S3BucketName)
}

func EncodeConfigDeliveryChannel_S3KeyPrefix(p ConfigDeliveryChannelParameters, vals map[string]cty.Value) {
	vals["s3_key_prefix"] = cty.StringVal(p.S3KeyPrefix)
}

func EncodeConfigDeliveryChannel_SnsTopicArn(p ConfigDeliveryChannelParameters, vals map[string]cty.Value) {
	vals["sns_topic_arn"] = cty.StringVal(p.SnsTopicArn)
}

func EncodeConfigDeliveryChannel_SnapshotDeliveryProperties(p SnapshotDeliveryProperties, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeConfigDeliveryChannel_SnapshotDeliveryProperties_DeliveryFrequency(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["snapshot_delivery_properties"] = cty.ListVal(valsForCollection)
}

func EncodeConfigDeliveryChannel_SnapshotDeliveryProperties_DeliveryFrequency(p SnapshotDeliveryProperties, vals map[string]cty.Value) {
	vals["delivery_frequency"] = cty.StringVal(p.DeliveryFrequency)
}