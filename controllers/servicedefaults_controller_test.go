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

package controllers_test

import (
	"testing"

	"github.com/hashicorp/consul-controller/api/v1alpha1"
	"github.com/hashicorp/consul-controller/controllers"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestServiceDefaultsController_testFoo(t *testing.T) {
	svc := &v1alpha1.ServiceDefaults{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "",
			Namespace: "",
		},
		Spec: v1alpha1.ServiceDefaultsSpec{
			Protocol:    "",
			MeshGateway: v1alpha1.MeshGatewayConfig{},
			Expose:      v1alpha1.ExposeConfig{},
			ExternalSNI: "",
		},
	}

	s := scheme.Scheme
	s.AddKnownTypes(v1alpha1.GroupVersion, svc)

	client := fake.NewFakeClientWithScheme(s, svc)

	r := controllers.ServiceDefaultsReconciler{
		Client: client,
		Log:    nil,
		Scheme: nil,
		Consul: nil,
	}

	resp, err := r.Reconcile(ctrl.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "",
			Name:      "",
		},
	})

	require.False(t, resp.Requeue)

	require.NoError(t, err)

}
