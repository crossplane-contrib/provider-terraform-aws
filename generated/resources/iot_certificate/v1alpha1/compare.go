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
	k := kube.(*IotCertificate)
	p := prov.(*IotCertificate)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeIotCertificate_Active(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIotCertificate_Csr(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIotCertificate_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIotCertificate_PrivateKey(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIotCertificate_PublicKey(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIotCertificate_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIotCertificate_CertificatePem(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeIotCertificate_Active(k *IotCertificateParameters, p *IotCertificateParameters, md *plugin.MergeDescription) bool {
	if k.Active != p.Active {
		p.Active = k.Active
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIotCertificate_Csr(k *IotCertificateParameters, p *IotCertificateParameters, md *plugin.MergeDescription) bool {
	if k.Csr != p.Csr {
		p.Csr = k.Csr
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIotCertificate_Id(k *IotCertificateParameters, p *IotCertificateParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIotCertificate_PrivateKey(k *IotCertificateObservation, p *IotCertificateObservation, md *plugin.MergeDescription) bool {
	if k.PrivateKey != p.PrivateKey {
		k.PrivateKey = p.PrivateKey
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIotCertificate_PublicKey(k *IotCertificateObservation, p *IotCertificateObservation, md *plugin.MergeDescription) bool {
	if k.PublicKey != p.PublicKey {
		k.PublicKey = p.PublicKey
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIotCertificate_Arn(k *IotCertificateObservation, p *IotCertificateObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIotCertificate_CertificatePem(k *IotCertificateObservation, p *IotCertificateObservation, md *plugin.MergeDescription) bool {
	if k.CertificatePem != p.CertificatePem {
		k.CertificatePem = p.CertificatePem
		md.StatusUpdated = true
		return true
	}
	return false
}