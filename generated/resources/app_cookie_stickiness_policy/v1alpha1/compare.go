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
	k := kube.(*AppCookieStickinessPolicy)
	p := prov.(*AppCookieStickinessPolicy)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeAppCookieStickinessPolicy_LoadBalancer(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppCookieStickinessPolicy_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppCookieStickinessPolicy_CookieName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppCookieStickinessPolicy_LbPort(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeAppCookieStickinessPolicy_LoadBalancer(k *AppCookieStickinessPolicyParameters, p *AppCookieStickinessPolicyParameters, md *plugin.MergeDescription) bool {
	if k.LoadBalancer != p.LoadBalancer {
		p.LoadBalancer = k.LoadBalancer
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppCookieStickinessPolicy_Name(k *AppCookieStickinessPolicyParameters, p *AppCookieStickinessPolicyParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppCookieStickinessPolicy_CookieName(k *AppCookieStickinessPolicyParameters, p *AppCookieStickinessPolicyParameters, md *plugin.MergeDescription) bool {
	if k.CookieName != p.CookieName {
		p.CookieName = k.CookieName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppCookieStickinessPolicy_LbPort(k *AppCookieStickinessPolicyParameters, p *AppCookieStickinessPolicyParameters, md *plugin.MergeDescription) bool {
	if k.LbPort != p.LbPort {
		p.LbPort = k.LbPort
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}