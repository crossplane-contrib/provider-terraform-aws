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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "api-gateway-resource.terraform-provider-aws.crossplane.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}
)

var (
	Kind                  = "ApiGatewayResource"
	GroupKind             = schema.GroupKind{Group: Group, Kind: Kind}.String()
	KindAPIVersion        = Kind + "." + SchemeGroupVersion.String()
	GroupVersionKind      = SchemeGroupVersion.WithKind(Kind)
	TerraformResourceName = "aws_api_gateway_resource"
)

func Implementation() *plugin.Implementation {
	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	schemeBuilder := &scheme.Builder{GroupVersion: SchemeGroupVersion}
	schemeBuilder.Register(&ApiGatewayResource{}, &ApiGatewayResourceList{})
	return &plugin.Implementation{
		GVK:                      GroupVersionKind,
		TerraformResourceName:    TerraformResourceName,
		SchemeBuilder:            schemeBuilder,
		ReconcilerConfigurer:     &reconcilerConfigurer{},
		ResourceMerger:           &resourceMerger{},
		CtyEncoder:               &ctyEncoder{},
		CtyDecoder:               &ctyDecoder{},
	}
}
