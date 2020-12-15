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
	k := kube.(*ApiGatewayApiKey)
	p := prov.(*ApiGatewayApiKey)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApiGatewayApiKey_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayApiKey_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayApiKey_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayApiKey_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayApiKey_Value(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayApiKey_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayApiKey_LastUpdatedDate(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayApiKey_CreatedDate(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayApiKey_Tags(k *ApiGatewayApiKeyParameters, p *ApiGatewayApiKeyParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayApiKey_Description(k *ApiGatewayApiKeyParameters, p *ApiGatewayApiKeyParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayApiKey_Enabled(k *ApiGatewayApiKeyParameters, p *ApiGatewayApiKeyParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayApiKey_Name(k *ApiGatewayApiKeyParameters, p *ApiGatewayApiKeyParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayApiKey_Value(k *ApiGatewayApiKeyParameters, p *ApiGatewayApiKeyParameters, md *plugin.MergeDescription) bool {
	if k.Value != p.Value {
		p.Value = k.Value
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApiGatewayApiKey_Arn(k *ApiGatewayApiKeyObservation, p *ApiGatewayApiKeyObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApiGatewayApiKey_LastUpdatedDate(k *ApiGatewayApiKeyObservation, p *ApiGatewayApiKeyObservation, md *plugin.MergeDescription) bool {
	if k.LastUpdatedDate != p.LastUpdatedDate {
		k.LastUpdatedDate = p.LastUpdatedDate
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApiGatewayApiKey_CreatedDate(k *ApiGatewayApiKeyObservation, p *ApiGatewayApiKeyObservation, md *plugin.MergeDescription) bool {
	if k.CreatedDate != p.CreatedDate {
		k.CreatedDate = p.CreatedDate
		md.StatusUpdated = true
		return true
	}
	return false
}