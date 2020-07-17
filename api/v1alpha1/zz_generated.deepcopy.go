// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposeConfig) DeepCopyInto(out *ExposeConfig) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]ExposePath, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposeConfig.
func (in *ExposeConfig) DeepCopy() *ExposeConfig {
	if in == nil {
		return nil
	}
	out := new(ExposeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposePath) DeepCopyInto(out *ExposePath) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposePath.
func (in *ExposePath) DeepCopy() *ExposePath {
	if in == nil {
		return nil
	}
	out := new(ExposePath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshGatewayConfig) DeepCopyInto(out *MeshGatewayConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshGatewayConfig.
func (in *MeshGatewayConfig) DeepCopy() *MeshGatewayConfig {
	if in == nil {
		return nil
	}
	out := new(MeshGatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDefaults) DeepCopyInto(out *ServiceDefaults) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDefaults.
func (in *ServiceDefaults) DeepCopy() *ServiceDefaults {
	if in == nil {
		return nil
	}
	out := new(ServiceDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceDefaults) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDefaultsList) DeepCopyInto(out *ServiceDefaultsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceDefaults, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDefaultsList.
func (in *ServiceDefaultsList) DeepCopy() *ServiceDefaultsList {
	if in == nil {
		return nil
	}
	out := new(ServiceDefaultsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceDefaultsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDefaultsSpec) DeepCopyInto(out *ServiceDefaultsSpec) {
	*out = *in
	out.MeshGateway = in.MeshGateway
	in.Expose.DeepCopyInto(&out.Expose)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDefaultsSpec.
func (in *ServiceDefaultsSpec) DeepCopy() *ServiceDefaultsSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceDefaultsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDefaultsStatus) DeepCopyInto(out *ServiceDefaultsStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDefaultsStatus.
func (in *ServiceDefaultsStatus) DeepCopy() *ServiceDefaultsStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceDefaultsStatus)
	in.DeepCopyInto(out)
	return out
}
