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

package v1

import (
	"context"
	"encoding/json"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/mutate-v1-pod,mutating=true,failurePolicy=ignore,groups=core,resources=pods,verbs=create;update,versions=v1,name=mpod.kb.io

// This webhook implementation is deeply inspired by controller-runtime examples.
//
// See examples here :
// https://github.com/kubernetes-sigs/controller-runtime/tree/release-0.6/examples/builtins

// log is for logging in this package.
var podlog = logf.Log.WithName("pod-resource")

// PodSchedulerSetter set a default schedulerName to Pods
type PodSchedulerSetter struct {
	client        client.Client
	decoder       *admission.Decoder
	SchedulerName string
}

// PodSchedulerSetter set a default schedulerName to every incoming pods without schedulerName.
func (s *PodSchedulerSetter) Handle(ctx context.Context, req admission.Request) admission.Response {
	podName := req.Name
	if podName == "" {
		podName = "<unknown>"
	}
	podlog.Info("default", "name", podName)

	pod := &corev1.Pod{}

	err := s.decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	if pod.Spec.SchedulerName == "" || pod.Spec.SchedulerName == "default-scheduler" {
		pod.Spec.SchedulerName = s.SchedulerName
	}

	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

// PodSchedulerSetter implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (s *PodSchedulerSetter) InjectDecoder(d *admission.Decoder) error {
	s.decoder = d
	return nil
}
