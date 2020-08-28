/*
Copyright 2020 The Knative Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package apiserversource

import (
	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1alpha2 "knative.dev/eventing/pkg/client/informers/externalversions/sources/v1alpha2"
	factory "knative.dev/eventing/pkg/client/injection/informers/factory"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(
		withInformer,
		metav1.GroupVersionResource{
			Group:    "sources.knative.dev",
			Version:  "v1alpha2",
			Resource: "apiserversources",
		})
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Sources().V1alpha2().ApiServerSources()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1alpha2.ApiServerSourceInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch knative.dev/eventing/pkg/client/informers/externalversions/sources/v1alpha2.ApiServerSourceInformer from context.")
	}
	return untyped.(v1alpha2.ApiServerSourceInformer)
}
