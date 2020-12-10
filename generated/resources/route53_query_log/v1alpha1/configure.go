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
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/client"
	"github.com/crossplane-contrib/terraform-runtime/pkg/controller"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	ctrl "sigs.k8s.io/controller-runtime"
)

type reconcilerConfigurer struct{}

// ConfigureReconciler adds a controller that reconciles the autogenerated managed.Resources in this package
func (c *reconcilerConfigurer) ConfigureReconciler(mgr ctrl.Manager, l logging.Logger, idx *plugin.Index, pool *client.ProviderPool) error {
	name := managed.ControllerName(GroupKind)
	r := managed.NewReconciler(mgr,
		resource.ManagedKind(GroupVersionKind),
		managed.WithInitializers(),
		managed.WithTimeout(time.Duration(3600*time.Second)),
		managed.WithExternalConnecter(&controller.Connector{KubeClient: mgr.GetClient(), PluginIndex: idx, Logger: l, Pool: pool}),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&Route53QueryLog{}).
		Complete(r)
}
