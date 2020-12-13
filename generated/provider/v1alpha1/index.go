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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane-contrib/terraform-runtime/pkg/client"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
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

	t := resource.NewProviderConfigUsageTracker(kube, &ProviderConfigUsage{})
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

// TODO: we are stubbing out empty values for a few fields required by the provider, otherwise
// decide whether we want to surface these in the provider
// TODO: x2 -- ideally we would generate the provider from the terraform schema. This should
// be doable with the existing block generation code in the codegen translator package,
// but will need a bunch of boilerplate and wiring, so punting for now.
func populateConfig(p *ProviderConfig, credentials aws.Credentials) cty.Value {
	merged := make(map[string]cty.Value)
	merged["region"] = cty.StringVal(p.Spec.Region)
	merged["access_key"] = cty.StringVal(credentials.AccessKeyID)
	merged["secret_key"] = cty.StringVal(credentials.SecretAccessKey)
	merged["assume_role"] = cty.ListValEmpty(assumeRoleObjectType())
	merged["ignore_tags"] = cty.ListValEmpty(ignoreTagsObjectType())
	merged["endpoints"] = cty.SetValEmpty(endpointSetElementType())

	merged["token"] = cty.NullVal(cty.String)
	merged["allowed_account_ids"] = cty.SetValEmpty(cty.String)
	merged["forbidden_account_ids"] = cty.SetValEmpty(cty.String)
	merged["insecure"] = cty.NullVal(cty.Bool)
	merged["max_retries"] = cty.NullVal(cty.Number)
	merged["profile"] = cty.NullVal(cty.String)
	merged["s3_force_path_style"] = cty.NullVal(cty.Bool)
	merged["shared_credentials_file"] = cty.NullVal(cty.String)
	merged["skip_credentials_validation"] = cty.NullVal(cty.Bool)
	merged["skip_get_ec2_platforms"] = cty.NullVal(cty.Bool)
	merged["skip_metadata_api_check"] = cty.NullVal(cty.Bool)
	merged["skip_region_validation"] = cty.NullVal(cty.Bool)
	merged["skip_requesting_account_id"] = cty.NullVal(cty.Bool)
	merged["token"] = cty.NullVal(cty.String)

	return cty.ObjectVal(merged)
}

func populateConfigClone(p *ProviderConfig, credentials aws.Credentials) cty.Value {
	merged := make(map[string]cty.Value)
	merged["region"] = cty.StringVal(p.Spec.Region)
	merged["access_key"] = cty.StringVal(credentials.AccessKeyID)
	merged["secret_key"] = cty.StringVal(credentials.SecretAccessKey)
	assumeRole := make(map[string]cty.Value)
	assumeRole["duration_seconds"] = cty.NullVal(cty.Number)
	assumeRole["external_id"] = cty.NullVal(cty.String)
	assumeRole["policy"] = cty.NullVal(cty.String)
	assumeRole["policy_arns"] = cty.SetValEmpty(cty.String)
	assumeRole["session_name"] = cty.NullVal(cty.String)
	assumeRole["tags"] = cty.EmptyObjectVal
	assumeRole["transitive_tag_keys"] = cty.SetValEmpty(cty.String)
	merged["assume_role"] = cty.ListVal([]cty.Value{cty.ObjectVal(assumeRole)})
	merged["endpoints"] = cty.ListVal([]cty.Value{cty.EmptyObjectVal})
	ignoreTags := make(map[string]cty.Value)
	ignoreTags["keys"] = cty.ListValEmpty(cty.String)
	ignoreTags["key_prefixes"] = cty.ListValEmpty(cty.String)
	merged["ignore_tags"] = cty.ListVal([]cty.Value{cty.ObjectVal(ignoreTags)})

	merged["token"] = cty.NullVal(cty.String)
	merged["allowed_account_ids"] = cty.ListValEmpty(cty.String)
	merged["forbidden_account_ids"] = cty.ListValEmpty(cty.String)
	merged["insecure"] = cty.NullVal(cty.Bool)
	merged["max_retries"] = cty.NullVal(cty.Number)
	merged["profile"] = cty.NullVal(cty.String)
	merged["s3_force_path_style"] = cty.NullVal(cty.Bool)
	merged["shared_credentials_file"] = cty.NullVal(cty.String)
	merged["skip_credentials_validation"] = cty.BoolVal(false)
	merged["skip_get_ec2_platforms"] = cty.BoolVal(false)
	merged["skip_metadata_api_check"] = cty.BoolVal(false)
	merged["skip_region_validation"] = cty.BoolVal(false)
	merged["skip_requesting_account_id"] = cty.BoolVal(false)
	merged["token"] = cty.StringVal("")

	return cty.ObjectVal(merged)
}

func GetProviderInit() *plugin.ProviderInit {
	schemeBuilder := &scheme.Builder{GroupVersion: SchemeGroupVersion}
	schemeBuilder.Register(&ProviderConfig{}, &ProviderConfigList{})
	schemeBuilder.Register(&ProviderConfigUsage{}, &ProviderConfigUsageList{})
	return &plugin.ProviderInit{
		SchemeBuilder: schemeBuilder,
		Initializer:   initializeProvider,
	}
}

func endpointSetElementType() cty.Type {
	objDesc := map[string]cty.Type{
		"accessanalyzer":           cty.String,
		"acm":                      cty.String,
		"acmpca":                   cty.String,
		"amplify":                  cty.String,
		"apigateway":               cty.String,
		"applicationautoscaling":   cty.String,
		"applicationinsights":      cty.String,
		"appmesh":                  cty.String,
		"appstream":                cty.String,
		"appsync":                  cty.String,
		"athena":                   cty.String,
		"autoscaling":              cty.String,
		"autoscalingplans":         cty.String,
		"backup":                   cty.String,
		"batch":                    cty.String,
		"budgets":                  cty.String,
		"cloud9":                   cty.String,
		"cloudformation":           cty.String,
		"cloudfront":               cty.String,
		"cloudhsm":                 cty.String,
		"cloudsearch":              cty.String,
		"cloudtrail":               cty.String,
		"cloudwatch":               cty.String,
		"cloudwatchevents":         cty.String,
		"cloudwatchlogs":           cty.String,
		"codeartifact":             cty.String,
		"codebuild":                cty.String,
		"codecommit":               cty.String,
		"codedeploy":               cty.String,
		"codepipeline":             cty.String,
		"cognitoidentity":          cty.String,
		"cognitoidp":               cty.String,
		"configservice":            cty.String,
		"cur":                      cty.String,
		"dataexchange":             cty.String,
		"datapipeline":             cty.String,
		"datasync":                 cty.String,
		"dax":                      cty.String,
		"devicefarm":               cty.String,
		"directconnect":            cty.String,
		"dlm":                      cty.String,
		"dms":                      cty.String,
		"docdb":                    cty.String,
		"ds":                       cty.String,
		"dynamodb":                 cty.String,
		"ec2":                      cty.String,
		"ecr":                      cty.String,
		"ecs":                      cty.String,
		"efs":                      cty.String,
		"eks":                      cty.String,
		"elasticache":              cty.String,
		"elasticbeanstalk":         cty.String,
		"elastictranscoder":        cty.String,
		"elb":                      cty.String,
		"emr":                      cty.String,
		"es":                       cty.String,
		"firehose":                 cty.String,
		"fms":                      cty.String,
		"forecast":                 cty.String,
		"fsx":                      cty.String,
		"gamelift":                 cty.String,
		"glacier":                  cty.String,
		"globalaccelerator":        cty.String,
		"glue":                     cty.String,
		"guardduty":                cty.String,
		"greengrass":               cty.String,
		"iam":                      cty.String,
		"identitystore":            cty.String,
		"imagebuilder":             cty.String,
		"inspector":                cty.String,
		"iot":                      cty.String,
		"iotanalytics":             cty.String,
		"iotevents":                cty.String,
		"kafka":                    cty.String,
		"kinesis":                  cty.String,
		"kinesisanalytics":         cty.String,
		"kinesisanalyticsv2":       cty.String,
		"kinesisvideo":             cty.String,
		"kms":                      cty.String,
		"lakeformation":            cty.String,
		"lambda":                   cty.String,
		"lexmodels":                cty.String,
		"licensemanager":           cty.String,
		"lightsail":                cty.String,
		"macie":                    cty.String,
		"macie2":                   cty.String,
		"managedblockchain":        cty.String,
		"marketplacecatalog":       cty.String,
		"mediaconnect":             cty.String,
		"mediaconvert":             cty.String,
		"medialive":                cty.String,
		"mediapackage":             cty.String,
		"mediastore":               cty.String,
		"mediastoredata":           cty.String,
		"mq":                       cty.String,
		"neptune":                  cty.String,
		"networkmanager":           cty.String,
		"opsworks":                 cty.String,
		"organizations":            cty.String,
		"outposts":                 cty.String,
		"personalize":              cty.String,
		"pinpoint":                 cty.String,
		"pricing":                  cty.String,
		"qldb":                     cty.String,
		"quicksight":               cty.String,
		"ram":                      cty.String,
		"rds":                      cty.String,
		"redshift":                 cty.String,
		"resourcegroups":           cty.String,
		"resourcegroupstaggingapi": cty.String,
		"route53":                  cty.String,
		"route53domains":           cty.String,
		"route53resolver":          cty.String,
		"s3":                       cty.String,
		"s3control":                cty.String,
		"sagemaker":                cty.String,
		"sdb":                      cty.String,
		"secretsmanager":           cty.String,
		"securityhub":              cty.String,
		"serverlessrepo":           cty.String,
		"servicecatalog":           cty.String,
		"servicediscovery":         cty.String,
		"servicequotas":            cty.String,
		"ses":                      cty.String,
		"shield":                   cty.String,
		"sns":                      cty.String,
		"sqs":                      cty.String,
		"ssm":                      cty.String,
		"ssoadmin":                 cty.String,
		"stepfunctions":            cty.String,
		"storagegateway":           cty.String,
		"sts":                      cty.String,
		"swf":                      cty.String,
		"synthetics":               cty.String,
		"transfer":                 cty.String,
		"waf":                      cty.String,
		"wafregional":              cty.String,
		"wafv2":                    cty.String,
		"worklink":                 cty.String,
		"workmail":                 cty.String,
		"workspaces":               cty.String,
		"xray":                     cty.String,
		//"signer":                   cty.String,
		//"mwaa":                     cty.String,
		//"s3outposts":               cty.String,
		//"timestreamwrite":          cty.String,
		//"networkfirewall":          cty.String,
		//"ecrpublic":                cty.String,
		//"codestarconnections":      cty.String,
		//"codestarnotifications":    cty.String,
	}

	return cty.Object(objDesc)
}

func assumeRoleObjectType() cty.Type {
	return cty.Object(map[string]cty.Type{
		"duration_seconds":    cty.Number,
		"external_id":         cty.String,
		"policy":              cty.String,
		"policy_arns":         cty.Set(cty.String),
		"role_arn":            cty.String,
		"session_name":        cty.String,
		"tags":                cty.Map(cty.String),
		"transitive_tag_keys": cty.Set(cty.String),
	})
}

func ignoreTagsObjectType() cty.Type {
	return cty.Object(map[string]cty.Type{
		"keys":         cty.Set(cty.String),
		"key_prefixes": cty.Set(cty.String),
	})
}
