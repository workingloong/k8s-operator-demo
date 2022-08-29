/*
Copyright 2022.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	// JobCreated mean the jobs has been accepted by the cluster,
	// but one or more of pods has not been started.
	JobCreated string = "Created"

	// JobRunning means all Pods of the jos have been successfully
	// scheduled and launched
	JobRunning string = "Running"

	// JobSucceed means all Pods of the Job is terminated. The training
	// completes without error
	JobSucceeded string = "Succeed"

	//JobFailed means the training has failed
	JobFailed string = "Failed"
)

// DemoJobSpec defines the desired state of DemoJob
type DemoJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Command string                    `json:"command,omitempty"`
	Image   string                    `json:"image,omitempty"`
	Envs    map[string]*corev1.EnvVar `json:"envs,omitempty"`
	Worker  *ReplicaSpec              `json:"worker,omitempty"`
}

// DemoJobStatus defines the observed state of DemoJob
type DemoJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Phase string `json:"phase,omitempty"`
}

type ReplicaSpec struct {
	// Count is the desired number of replicas
	Count int32 `json:"count,omitempty"`
	// Resource is the resource of each replica
	Resource *ResourceSpec `json:"resource,omitempty"`
}

type ResourceSpec struct {
	// CPU is CPU cores of a replica
	CPU int32 `json:"cpu,omitempty"`
	// Memory is the memory with MB of a replica
	Memory int32 `json:"memory,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DemoJob is the Schema for the demojobs API
type DemoJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoJobSpec   `json:"spec,omitempty"`
	Status DemoJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DemoJobList contains a list of DemoJob
type DemoJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DemoJob{}, &DemoJobList{})
}
