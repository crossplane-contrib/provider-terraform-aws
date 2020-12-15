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
	k := kube.(*CurReportDefinition)
	p := prov.(*CurReportDefinition)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCurReportDefinition_Format(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_RefreshClosedReports(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_ReportName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_S3Region(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_TimeUnit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_AdditionalArtifacts(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_Compression(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_S3Bucket(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_S3Prefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_AdditionalSchemaElements(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCurReportDefinition_ReportVersioning(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeCurReportDefinition_Format(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.Format != p.Format {
		p.Format = k.Format
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCurReportDefinition_Id(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCurReportDefinition_RefreshClosedReports(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.RefreshClosedReports != p.RefreshClosedReports {
		p.RefreshClosedReports = k.RefreshClosedReports
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCurReportDefinition_ReportName(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.ReportName != p.ReportName {
		p.ReportName = k.ReportName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCurReportDefinition_S3Region(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.S3Region != p.S3Region {
		p.S3Region = k.S3Region
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCurReportDefinition_TimeUnit(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.TimeUnit != p.TimeUnit {
		p.TimeUnit = k.TimeUnit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCurReportDefinition_AdditionalArtifacts(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.AdditionalArtifacts, p.AdditionalArtifacts) {
		p.AdditionalArtifacts = k.AdditionalArtifacts
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCurReportDefinition_Compression(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.Compression != p.Compression {
		p.Compression = k.Compression
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCurReportDefinition_S3Bucket(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.S3Bucket != p.S3Bucket {
		p.S3Bucket = k.S3Bucket
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCurReportDefinition_S3Prefix(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.S3Prefix != p.S3Prefix {
		p.S3Prefix = k.S3Prefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCurReportDefinition_AdditionalSchemaElements(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.AdditionalSchemaElements, p.AdditionalSchemaElements) {
		p.AdditionalSchemaElements = k.AdditionalSchemaElements
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCurReportDefinition_ReportVersioning(k *CurReportDefinitionParameters, p *CurReportDefinitionParameters, md *plugin.MergeDescription) bool {
	if k.ReportVersioning != p.ReportVersioning {
		p.ReportVersioning = k.ReportVersioning
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}