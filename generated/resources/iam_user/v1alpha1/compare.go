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
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube xpresource.Managed, prov xpresource.Managed) plugin.MergeDescription {
	k := kube.(*IamUser)
	p := prov.(*IamUser)
	md := plugin.MergeDescription{}

	if k.Status.AtProvider.Arn != p.Status.AtProvider.Arn {
		k.Status.AtProvider.Arn = p.Status.AtProvider.Arn
		md.StatusUpdated = true
	}
	if k.Status.AtProvider.UniqueId != p.Status.AtProvider.UniqueId {
		k.Status.AtProvider.UniqueId = p.Status.AtProvider.UniqueId
		md.StatusUpdated = true
	}
	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}

	return md
}
