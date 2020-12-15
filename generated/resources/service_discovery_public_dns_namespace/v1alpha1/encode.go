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
	r, ok := mr.(*ServiceDiscoveryPublicDnsNamespace)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ServiceDiscoveryPublicDnsNamespace.")
	}
	return EncodeServiceDiscoveryPublicDnsNamespace(*r), nil
}

func EncodeServiceDiscoveryPublicDnsNamespace(r ServiceDiscoveryPublicDnsNamespace) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeServiceDiscoveryPublicDnsNamespace_Description(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryPublicDnsNamespace_Id(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryPublicDnsNamespace_Name(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryPublicDnsNamespace_Tags(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryPublicDnsNamespace_HostedZone(r.Status.AtProvider, ctyVal)
	EncodeServiceDiscoveryPublicDnsNamespace_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeServiceDiscoveryPublicDnsNamespace_Description(p ServiceDiscoveryPublicDnsNamespaceParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeServiceDiscoveryPublicDnsNamespace_Id(p ServiceDiscoveryPublicDnsNamespaceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeServiceDiscoveryPublicDnsNamespace_Name(p ServiceDiscoveryPublicDnsNamespaceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeServiceDiscoveryPublicDnsNamespace_Tags(p ServiceDiscoveryPublicDnsNamespaceParameters, vals map[string]cty.Value) {
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

func EncodeServiceDiscoveryPublicDnsNamespace_HostedZone(p ServiceDiscoveryPublicDnsNamespaceObservation, vals map[string]cty.Value) {
	vals["hosted_zone"] = cty.StringVal(p.HostedZone)
}

func EncodeServiceDiscoveryPublicDnsNamespace_Arn(p ServiceDiscoveryPublicDnsNamespaceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}