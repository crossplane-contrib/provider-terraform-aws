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
	k := kube.(*DirectoryServiceConditionalForwarder)
	p := prov.(*DirectoryServiceConditionalForwarder)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDirectoryServiceConditionalForwarder_RemoteDomainName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDirectoryServiceConditionalForwarder_DirectoryId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDirectoryServiceConditionalForwarder_DnsIps(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeDirectoryServiceConditionalForwarder_RemoteDomainName(k *DirectoryServiceConditionalForwarderParameters, p *DirectoryServiceConditionalForwarderParameters, md *plugin.MergeDescription) bool {
	if k.RemoteDomainName != p.RemoteDomainName {
		p.RemoteDomainName = k.RemoteDomainName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDirectoryServiceConditionalForwarder_DirectoryId(k *DirectoryServiceConditionalForwarderParameters, p *DirectoryServiceConditionalForwarderParameters, md *plugin.MergeDescription) bool {
	if k.DirectoryId != p.DirectoryId {
		p.DirectoryId = k.DirectoryId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDirectoryServiceConditionalForwarder_DnsIps(k *DirectoryServiceConditionalForwarderParameters, p *DirectoryServiceConditionalForwarderParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.DnsIps, p.DnsIps) {
		p.DnsIps = k.DnsIps
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}