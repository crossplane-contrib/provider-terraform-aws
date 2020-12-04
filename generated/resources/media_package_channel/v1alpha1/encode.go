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

func EncodeMediaPackageChannel(r MediaPackageChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeMediaPackageChannel_Description(r.Spec.ForProvider, ctyVal)
	EncodeMediaPackageChannel_Id(r.Spec.ForProvider, ctyVal)
	EncodeMediaPackageChannel_Tags(r.Spec.ForProvider, ctyVal)
	EncodeMediaPackageChannel_ChannelId(r.Spec.ForProvider, ctyVal)
	EncodeMediaPackageChannel_HlsIngest(r.Status.AtProvider.HlsIngest, ctyVal)
	EncodeMediaPackageChannel_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeMediaPackageChannel_Description(p MediaPackageChannelParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeMediaPackageChannel_Id(p MediaPackageChannelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMediaPackageChannel_Tags(p MediaPackageChannelParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeMediaPackageChannel_ChannelId(p MediaPackageChannelParameters, vals map[string]cty.Value) {
	vals["channel_id"] = cty.StringVal(p.ChannelId)
}

func EncodeMediaPackageChannel_HlsIngest(p []HlsIngest, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeMediaPackageChannel_HlsIngest_IngestEndpoints(v.IngestEndpoints, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["hls_ingest"] = cty.ListVal(valsForCollection)
}

func EncodeMediaPackageChannel_HlsIngest_IngestEndpoints(p []IngestEndpoints, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeMediaPackageChannel_HlsIngest_IngestEndpoints_Password(v, ctyVal)
		EncodeMediaPackageChannel_HlsIngest_IngestEndpoints_Url(v, ctyVal)
		EncodeMediaPackageChannel_HlsIngest_IngestEndpoints_Username(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ingest_endpoints"] = cty.ListVal(valsForCollection)
}

func EncodeMediaPackageChannel_HlsIngest_IngestEndpoints_Password(p IngestEndpoints, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeMediaPackageChannel_HlsIngest_IngestEndpoints_Url(p IngestEndpoints, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}

func EncodeMediaPackageChannel_HlsIngest_IngestEndpoints_Username(p IngestEndpoints, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeMediaPackageChannel_Arn(p MediaPackageChannelObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}