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
	consulapi "github.com/hashicorp/consul/api"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	consuliov1alpha1 "github.com/hashicorp/consul-controller/api/v1alpha1"
)

// ServiceDefaultReconciler reconciles a ServiceDefault object
type ServiceDefaultReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	Consul consulapi.Client
}

// +kubebuilder:rbac:groups=consul.io,resources=servicedefaults,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=consul.io,resources=servicedefaults/status,verbs=get;update;patch

func (r *ServiceDefaultReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := r.Log.WithValues("servicedefault", req.NamespacedName)
	svcDefault := consuliov1alpha1.ServiceDefault{}

	err := r.Client.Get(ctx, req.NamespacedName, &svcDefault)
	if errors.IsNotFound(err) {
		logger.Error(err, "could not retrieve Service Default")
		return ctrl.Result{}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}

	// check to see if consul has service default with the same name
	entry, _, err := r.Consul.ConfigEntries().Get(svcDefault.Kind, svcDefault.Name, &consulapi.QueryOptions{})
	//if a config entry with this name does not exist
	if err != nil && strings.Contains(err.Error(), "404") {
		//create the config entry
		consulConfigEntry := svcDefault.ToConsul()
		_, _, err := r.Consul.ConfigEntries().Set(consulConfigEntry, &consulapi.WriteOptions{})
		if err != nil {
			return ctrl.Result{}, err
		}
	} else if err != nil {
		//something went wrong and we should probably exit with a sensible error
	} else {
		if !svcDefault.MatchesConsul(entry) {
			_, _, err := r.Consul.ConfigEntries().Set(svcDefault.ToConsul(), &consulapi.WriteOptions{})
			if err != nil {
				return ctrl.Result{}, nil
			}
		}
	}

	//handles deletes

	return ctrl.Result{}, nil
}

func (r *ServiceDefaultReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&consuliov1alpha1.ServiceDefault{}).
		Complete(r)
}
