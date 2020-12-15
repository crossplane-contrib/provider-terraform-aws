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
	k := kube.(*AppsyncFunction)
	p := prov.(*AppsyncFunction)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeAppsyncFunction_FunctionVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppsyncFunction_ResponseMappingTemplate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppsyncFunction_ApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppsyncFunction_DataSource(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppsyncFunction_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppsyncFunction_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppsyncFunction_RequestMappingTemplate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppsyncFunction_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppsyncFunction_FunctionId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeAppsyncFunction_FunctionVersion(k *AppsyncFunctionParameters, p *AppsyncFunctionParameters, md *plugin.MergeDescription) bool {
	if k.FunctionVersion != p.FunctionVersion {
		p.FunctionVersion = k.FunctionVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppsyncFunction_ResponseMappingTemplate(k *AppsyncFunctionParameters, p *AppsyncFunctionParameters, md *plugin.MergeDescription) bool {
	if k.ResponseMappingTemplate != p.ResponseMappingTemplate {
		p.ResponseMappingTemplate = k.ResponseMappingTemplate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppsyncFunction_ApiId(k *AppsyncFunctionParameters, p *AppsyncFunctionParameters, md *plugin.MergeDescription) bool {
	if k.ApiId != p.ApiId {
		p.ApiId = k.ApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppsyncFunction_DataSource(k *AppsyncFunctionParameters, p *AppsyncFunctionParameters, md *plugin.MergeDescription) bool {
	if k.DataSource != p.DataSource {
		p.DataSource = k.DataSource
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppsyncFunction_Description(k *AppsyncFunctionParameters, p *AppsyncFunctionParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppsyncFunction_Name(k *AppsyncFunctionParameters, p *AppsyncFunctionParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppsyncFunction_RequestMappingTemplate(k *AppsyncFunctionParameters, p *AppsyncFunctionParameters, md *plugin.MergeDescription) bool {
	if k.RequestMappingTemplate != p.RequestMappingTemplate {
		p.RequestMappingTemplate = k.RequestMappingTemplate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeAppsyncFunction_Arn(k *AppsyncFunctionObservation, p *AppsyncFunctionObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeAppsyncFunction_FunctionId(k *AppsyncFunctionObservation, p *AppsyncFunctionObservation, md *plugin.MergeDescription) bool {
	if k.FunctionId != p.FunctionId {
		k.FunctionId = p.FunctionId
		md.StatusUpdated = true
		return true
	}
	return false
}