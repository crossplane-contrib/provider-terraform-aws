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
	k := kube.(*WafregionalRegexPatternSet)
	p := prov.(*WafregionalRegexPatternSet)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeWafregionalRegexPatternSet_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeWafregionalRegexPatternSet_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeWafregionalRegexPatternSet_RegexPatternStrings(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeWafregionalRegexPatternSet_Id(k *WafregionalRegexPatternSetParameters, p *WafregionalRegexPatternSetParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeWafregionalRegexPatternSet_Name(k *WafregionalRegexPatternSetParameters, p *WafregionalRegexPatternSetParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeWafregionalRegexPatternSet_RegexPatternStrings(k *WafregionalRegexPatternSetParameters, p *WafregionalRegexPatternSetParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.RegexPatternStrings, p.RegexPatternStrings) {
		p.RegexPatternStrings = k.RegexPatternStrings
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}