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
	k := kube.(*GlueCatalogDatabase)
	p := prov.(*GlueCatalogDatabase)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeGlueCatalogDatabase_CatalogId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGlueCatalogDatabase_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGlueCatalogDatabase_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGlueCatalogDatabase_LocationUri(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGlueCatalogDatabase_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGlueCatalogDatabase_Parameters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGlueCatalogDatabase_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeGlueCatalogDatabase_CatalogId(k *GlueCatalogDatabaseParameters, p *GlueCatalogDatabaseParameters, md *plugin.MergeDescription) bool {
	if k.CatalogId != p.CatalogId {
		p.CatalogId = k.CatalogId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGlueCatalogDatabase_Description(k *GlueCatalogDatabaseParameters, p *GlueCatalogDatabaseParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGlueCatalogDatabase_Id(k *GlueCatalogDatabaseParameters, p *GlueCatalogDatabaseParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGlueCatalogDatabase_LocationUri(k *GlueCatalogDatabaseParameters, p *GlueCatalogDatabaseParameters, md *plugin.MergeDescription) bool {
	if k.LocationUri != p.LocationUri {
		p.LocationUri = k.LocationUri
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGlueCatalogDatabase_Name(k *GlueCatalogDatabaseParameters, p *GlueCatalogDatabaseParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeGlueCatalogDatabase_Parameters(k *GlueCatalogDatabaseParameters, p *GlueCatalogDatabaseParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Parameters, p.Parameters) {
		p.Parameters = k.Parameters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeGlueCatalogDatabase_Arn(k *GlueCatalogDatabaseObservation, p *GlueCatalogDatabaseObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}