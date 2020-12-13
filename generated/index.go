package generated

import (
	"github.com/crossplane-contrib/provider-terraform-aws/generated/provider/v1alpha1"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

const ProviderReferenceName string = "aws"

func Index(idxr *plugin.Indexer) {
	for _, impl := range ResourceImplementations {
		idxr.Overlay(impl)
	}
}

func ProviderInit() *plugin.ProviderInit {
	return v1alpha1.GetProviderInit()
}
