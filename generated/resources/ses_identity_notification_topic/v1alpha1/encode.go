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
	r, ok := mr.(*SesIdentityNotificationTopic)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SesIdentityNotificationTopic.")
	}
	return EncodeSesIdentityNotificationTopic(*r), nil
}

func EncodeSesIdentityNotificationTopic(r SesIdentityNotificationTopic) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSesIdentityNotificationTopic_Id(r.Spec.ForProvider, ctyVal)
	EncodeSesIdentityNotificationTopic_Identity(r.Spec.ForProvider, ctyVal)
	EncodeSesIdentityNotificationTopic_IncludeOriginalHeaders(r.Spec.ForProvider, ctyVal)
	EncodeSesIdentityNotificationTopic_NotificationType(r.Spec.ForProvider, ctyVal)
	EncodeSesIdentityNotificationTopic_TopicArn(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeSesIdentityNotificationTopic_Id(p SesIdentityNotificationTopicParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSesIdentityNotificationTopic_Identity(p SesIdentityNotificationTopicParameters, vals map[string]cty.Value) {
	vals["identity"] = cty.StringVal(p.Identity)
}

func EncodeSesIdentityNotificationTopic_IncludeOriginalHeaders(p SesIdentityNotificationTopicParameters, vals map[string]cty.Value) {
	vals["include_original_headers"] = cty.BoolVal(p.IncludeOriginalHeaders)
}

func EncodeSesIdentityNotificationTopic_NotificationType(p SesIdentityNotificationTopicParameters, vals map[string]cty.Value) {
	vals["notification_type"] = cty.StringVal(p.NotificationType)
}

func EncodeSesIdentityNotificationTopic_TopicArn(p SesIdentityNotificationTopicParameters, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}