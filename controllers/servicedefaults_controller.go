/*


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

package controllers

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	capi "github.com/hashicorp/consul/api"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/hashicorp/consul-controller/api/v1alpha1"
)

// ServiceDefaultsReconciler reconciles a ServiceDefaults object
type ServiceDefaultsReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	Consul *capi.Client
}

// +kubebuilder:rbac:groups=consul.hashicorp.com,resources=servicedefaults,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=consul.hashicorp.com,resources=servicedefaults/status,verbs=get;update;patch

func (r *ServiceDefaultsReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := r.Log.WithValues("servicedefault", req.NamespacedName)
	svcDefaults := v1alpha1.ServiceDefaults{}

	err := r.Client.Get(ctx, req.NamespacedName, &svcDefaults)
	if errors.IsNotFound(err) {
		return ctrl.Result{}, nil
	} else if err != nil {
		logger.Error(err, "failed to retrieve Service Default")
		return ctrl.Result{}, err
	}

	if svcDefaults.ObjectMeta.DeletionTimestamp.IsZero() {

	}

	// check to see if consul has service default with the same name
	entry, _, err := r.Consul.ConfigEntries().Get(svcDefaults.Kind, svcDefaults.Name, &capi.QueryOptions{})
	//if a config entry with this name does not exist
	if err != nil && strings.Contains(err.Error(), "404") {
		//create the config entry
		consulConfigEntry := svcDefaults.ToConsul()
		_, _, err := r.Consul.ConfigEntries().Set(consulConfigEntry, &capi.WriteOptions{})
		if err != nil {
			return ctrl.Result{}, err
		}
	} else if err != nil {
		//something went wrong and we should probably exit with a sensible error
	} else {
		if !svcDefaults.MatchesConsul(entry) {
			_, _, err := r.Consul.ConfigEntries().Set(svcDefaults.ToConsul(), &capi.WriteOptions{})
			if err != nil {
				return ctrl.Result{}, nil
			}
		}
	}

	return ctrl.Result{}, nil
}

func (r *ServiceDefaultsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ServiceDefaults{}).
		Complete(r)
}
