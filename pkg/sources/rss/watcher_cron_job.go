/*
Copyright 2018 The Knative Authors
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

package main

import (
	"k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeCronJob(namespace, name, image, target, cronSchedule, rssURL string) *v1beta1.CronJob {
	labels := map[string]string{
		"receive-adapter": "rss",
		"rssURL":          rssURL,
	}
	return &v1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: v1beta1.CronJobSpec{
			Schedule:          cronSchedule,
			ConcurrencyPolicy: v1beta1.ReplaceConcurrent,
			JobTemplate: v1beta1.JobTemplateSpec{
				Spec: v1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Image: image,
									Env: []corev1.EnvVar{
										{
											Name:  "TARGET",
											Value: target,
										},
										{
											Name:  "RSS_URL",
											Value: rssURL,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
