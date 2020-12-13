package generated

import (
	iam "github.com/crossplane-contrib/provider-terraform-aws/generated/resources/iam_user/v1alpha1"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

// TODO: output this list somewhere in the codegen pipeline
var ResourceImplementations = []*plugin.Implementation{
	//iam.ServiceAccountImplementation(),
	iam.Implementation(),
}
