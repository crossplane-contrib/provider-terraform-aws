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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*ApiGatewayUsagePlanKey)
	p := prov.(*ApiGatewayUsagePlanKey)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApiGatewayUsagePlanKey_KeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayUsagePlanKey_KeyType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayUsagePlanKey_UsagePlanId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayUsagePlanKey_Name(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayUsagePlanKey_Value(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayUsagePlanKey_KeyId(k *ApiGatewayUsagePlanKeyParameters, p *ApiGatewayUsagePlanKeyParameters, md *plugin.MergeDescription) bool {
	if k.KeyId != p.KeyId {
		p.KeyId = k.KeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayUsagePlanKey_KeyType(k *ApiGatewayUsagePlanKeyParameters, p *ApiGatewayUsagePlanKeyParameters, md *plugin.MergeDescription) bool {
	if k.KeyType != p.KeyType {
		p.KeyType = k.KeyType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayUsagePlanKey_UsagePlanId(k *ApiGatewayUsagePlanKeyParameters, p *ApiGatewayUsagePlanKeyParameters, md *plugin.MergeDescription) bool {
	if k.UsagePlanId != p.UsagePlanId {
		p.UsagePlanId = k.UsagePlanId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApiGatewayUsagePlanKey_Name(k *ApiGatewayUsagePlanKeyObservation, p *ApiGatewayUsagePlanKeyObservation, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		k.Name = p.Name
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApiGatewayUsagePlanKey_Value(k *ApiGatewayUsagePlanKeyObservation, p *ApiGatewayUsagePlanKeyObservation, md *plugin.MergeDescription) bool {
	if k.Value != p.Value {
		k.Value = p.Value
		md.StatusUpdated = true
		return true
	}
	return false
}