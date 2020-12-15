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
	k := kube.(*LambdaLayerVersion)
	p := prov.(*LambdaLayerVersion)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeLambdaLayerVersion_CompatibleRuntimes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_S3Key(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_SourceCodeHash(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_LayerName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_LicenseInfo(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_Filename(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_S3Bucket(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_S3ObjectVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_Version(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_LayerArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_SourceCodeSize(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaLayerVersion_CreatedDate(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeLambdaLayerVersion_CompatibleRuntimes(k *LambdaLayerVersionParameters, p *LambdaLayerVersionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.CompatibleRuntimes, p.CompatibleRuntimes) {
		p.CompatibleRuntimes = k.CompatibleRuntimes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaLayerVersion_S3Key(k *LambdaLayerVersionParameters, p *LambdaLayerVersionParameters, md *plugin.MergeDescription) bool {
	if k.S3Key != p.S3Key {
		p.S3Key = k.S3Key
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaLayerVersion_SourceCodeHash(k *LambdaLayerVersionParameters, p *LambdaLayerVersionParameters, md *plugin.MergeDescription) bool {
	if k.SourceCodeHash != p.SourceCodeHash {
		p.SourceCodeHash = k.SourceCodeHash
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaLayerVersion_Description(k *LambdaLayerVersionParameters, p *LambdaLayerVersionParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaLayerVersion_LayerName(k *LambdaLayerVersionParameters, p *LambdaLayerVersionParameters, md *plugin.MergeDescription) bool {
	if k.LayerName != p.LayerName {
		p.LayerName = k.LayerName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaLayerVersion_LicenseInfo(k *LambdaLayerVersionParameters, p *LambdaLayerVersionParameters, md *plugin.MergeDescription) bool {
	if k.LicenseInfo != p.LicenseInfo {
		p.LicenseInfo = k.LicenseInfo
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaLayerVersion_Filename(k *LambdaLayerVersionParameters, p *LambdaLayerVersionParameters, md *plugin.MergeDescription) bool {
	if k.Filename != p.Filename {
		p.Filename = k.Filename
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaLayerVersion_S3Bucket(k *LambdaLayerVersionParameters, p *LambdaLayerVersionParameters, md *plugin.MergeDescription) bool {
	if k.S3Bucket != p.S3Bucket {
		p.S3Bucket = k.S3Bucket
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaLayerVersion_S3ObjectVersion(k *LambdaLayerVersionParameters, p *LambdaLayerVersionParameters, md *plugin.MergeDescription) bool {
	if k.S3ObjectVersion != p.S3ObjectVersion {
		p.S3ObjectVersion = k.S3ObjectVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLambdaLayerVersion_Version(k *LambdaLayerVersionObservation, p *LambdaLayerVersionObservation, md *plugin.MergeDescription) bool {
	if k.Version != p.Version {
		k.Version = p.Version
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLambdaLayerVersion_LayerArn(k *LambdaLayerVersionObservation, p *LambdaLayerVersionObservation, md *plugin.MergeDescription) bool {
	if k.LayerArn != p.LayerArn {
		k.LayerArn = p.LayerArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLambdaLayerVersion_SourceCodeSize(k *LambdaLayerVersionObservation, p *LambdaLayerVersionObservation, md *plugin.MergeDescription) bool {
	if k.SourceCodeSize != p.SourceCodeSize {
		k.SourceCodeSize = p.SourceCodeSize
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLambdaLayerVersion_Arn(k *LambdaLayerVersionObservation, p *LambdaLayerVersionObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLambdaLayerVersion_CreatedDate(k *LambdaLayerVersionObservation, p *LambdaLayerVersionObservation, md *plugin.MergeDescription) bool {
	if k.CreatedDate != p.CreatedDate {
		k.CreatedDate = p.CreatedDate
		md.StatusUpdated = true
		return true
	}
	return false
}