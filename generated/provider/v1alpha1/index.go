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
	"context"
	"reflect"

	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane-contrib/terraform-runtime/pkg/client"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/provider-aws/apis/v1beta1"
	xpaws "github.com/crossplane/provider-aws/pkg/clients"
	"github.com/pkg/errors"
	"github.com/zclconf/go-cty/cty"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	kubeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group                         = "aws.terraform-plugin.crossplane.io"
	Version                       = "v1alpha1"
	errProviderNotRetrieved       = "provider could not be retrieved"
	errProviderSecretNotRetrieved = "secret referred in provider could not be retrieved"
	errProviderSecretNil          = "cannot find Secret reference on Provider"
	ProviderName                  = "aws"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}
	// Provider type metadata.
	ProviderKind             = reflect.TypeOf(ProviderConfig{}).Name()
	ProviderGroupKind        = schema.GroupKind{Group: Group, Kind: ProviderKind}.String()
	ProviderKindAPIVersion   = ProviderKind + "." + SchemeGroupVersion.String()
	ProviderGroupVersionKind = SchemeGroupVersion.WithKind(ProviderKind)
)

func initializeProvider(ctx context.Context, mr resource.Managed, ropts *client.RuntimeOptions, kube kubeclient.Client) (*client.Provider, error) {
	pc := &ProviderConfig{}
	if err := kube.Get(ctx, types.NamespacedName{Name: mr.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, "cannot get referenced Provider")
	}

	t := resource.NewProviderConfigUsageTracker(kube, &v1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mr); err != nil {
		return nil, errors.Wrap(err, "cannot track ProviderConfig usage")
	}

	creds, err := providerConfigCredentials(ctx, kube, pc)
	cfg := populateConfig(pc, creds)

	p, err := client.NewProvider(ProviderName, ropts.PluginPath)
	if err != nil {
		return p, err
	}
	err = p.Configure(cfg)
	return p, err
}

// providerConfigCredentials encapsulates the flow for finding credentials based on the credential source
// TODO: if/when this project merges with provider-aws, we'll want to refactor UseProviderConfig to separate
// retrieving credentials and constructing an aws.Config object, because the generated code doesn't want an aws.Config object
func providerConfigCredentials(ctx context.Context, c kubeclient.Client, pc *ProviderConfig) (aws.Credentials, error) {
	switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
	case runtimev1alpha1.CredentialsSourceSecret:
		csr := pc.Spec.Credentials.SecretRef
		if csr == nil {
			return aws.Credentials{}, errors.New("no credentials secret referenced")
		}
		s := &corev1.Secret{}
		if err := c.Get(ctx, types.NamespacedName{Namespace: csr.Namespace, Name: csr.Name}, s); err != nil {
			return aws.Credentials{}, errors.Wrap(err, "cannot get credentials secret")
		}
		creds, err := xpaws.CredentialsIDSecret(s.Data[csr.Key], xpaws.DefaultSection)
		if err != nil {
			return aws.Credentials{}, errors.Wrap(err, "cannot parse credentials secret")
		}
		return creds, nil
	default:
		return aws.Credentials{}, errors.Errorf("credentials source %s is not currently supported", s)
	}
}

// Note that this config still needs to have null values filled in with the correct structure
func populateConfig(p *ProviderConfig, credentials aws.Credentials) map[string]cty.Value {
	merged := make(map[string]cty.Value)
	merged["region"] = cty.StringVal(p.Spec.Region)
	merged["access_key"] = cty.StringVal(credentials.AccessKeyID)
	merged["secret_key"] = cty.StringVal(credentials.SecretAccessKey)

	// This might be gcp specific, leaving commented while testing
	/*
		batching := make(map[string]cty.Value)
		batching["enable_batching"] = cty.BoolVal(false)
		batching["send_after"] = cty.StringVal("3s")
		batchList := []cty.Value{cty.ObjectVal(batching)}
		merged["batching"] = cty.ListVal(batchList)
	*/

	return merged
}

func GetProviderInit() *plugin.ProviderInit {
	schemeBuilder := &scheme.Builder{GroupVersion: SchemeGroupVersion}
	schemeBuilder.Register(&ProviderConfig{}, &ProviderConfigList{})
	return &plugin.ProviderInit{
		SchemeBuilder: schemeBuilder,
		Initializer:   initializeProvider,
	}
}
