// Copyright © 2018 The Kubernetes Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package machine

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/controller"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	record "k8s.io/client-go/tools/record"
	options "sigs.k8s.io/cluster-api-provider-ssh/cloud/ssh/controllers/machine/options"
	"sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
)

func TestNewActuator(t *testing.T) {
	ConfigWatch, _ := NewConfigWatch("./test_files/fake_machine.yaml")
	server := *options.NewServer("./test_files/fake_machine.yaml")
	fmt.Println(server.CommonConfig.Kubeconfig, "<---- supposed to be a kubeconfig")
	config, _ := controller.GetConfig(server.CommonConfig.Kubeconfig)
	fmt.Println(config)
	client, _ := clientset.NewForConfig(config)
	kubeClient, _ := kubernetes.NewForConfig(config)
	testcases := []struct {
		name    string
		params  ActuatorParams
		want    *Actuator
		wantErr bool
	}{
		{
			name: "why is this hard",
			params: ActuatorParams{
				ClusterClient:            client.ClusterV1alpha1().Clusters(corev1.NamespaceDefault),
				EventRecorder:            &record.FakeRecorder{},
				KubeClient:               kubeClient,
				V1Alpha1Client:           client.ClusterV1alpha1(),
				MachineSetupConfigGetter: ConfigWatch,
			},
			want:    &Actuator{},
			wantErr: false,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewActuator(tc.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("NewActuator() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("NewActuator() = %v, want %v", got, tc.want)
			}
		})
	}
}

// func TestActuator_Create(t *testing.T) {
// 	type fields struct {
// 		clusterClient            client.ClusterInterface
// 		eventRecorder            record.EventRecorder
// 		sshProviderConfigCodec   *v1alpha1.SSHProviderConfigCodec
// 		kubeClient               *kubernetes.Clientset
// 		v1Alpha1Client           client.ClusterV1alpha1Interface
// 		scheme                   *runtime.Scheme
// 		machineSetupConfigGetter SSHClientMachineSetupConfigGetter
// 		kubeadm                  SSHClientKubeadm
// 	}
// 	type args struct {
// 		c *clusterv1.Cluster
// 		m *clusterv1.Machine
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			a := &Actuator{
// 				clusterClient:          tt.fields.clusterClient,
// 				eventRecorder:          tt.fields.eventRecorder,
// 				sshProviderConfigCodec: tt.fields.sshProviderConfigCodec,
// 				kubeClient:             tt.fields.kubeClient,
// 				v1Alpha1Client:         tt.fields.v1Alpha1Client,
// 				scheme:                 tt.fields.scheme,
// 				machineSetupConfigGetter: tt.fields.machineSetupConfigGetter,
// 				kubeadm:                  tt.fields.kubeadm,
// 			}
// 			if err := a.Create(tt.args.c, tt.args.m); (err != nil) != tt.wantErr {
// 				t.Errorf("Actuator.Create() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestActuator_Delete(t *testing.T) {
// 	type fields struct {
// 		clusterClient            client.ClusterInterface
// 		eventRecorder            record.EventRecorder
// 		sshProviderConfigCodec   *v1alpha1.SSHProviderConfigCodec
// 		kubeClient               *kubernetes.Clientset
// 		v1Alpha1Client           client.ClusterV1alpha1Interface
// 		scheme                   *runtime.Scheme
// 		machineSetupConfigGetter SSHClientMachineSetupConfigGetter
// 		kubeadm                  SSHClientKubeadm
// 	}
// 	type args struct {
// 		c *clusterv1.Cluster
// 		m *clusterv1.Machine
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			a := &Actuator{
// 				clusterClient:          tt.fields.clusterClient,
// 				eventRecorder:          tt.fields.eventRecorder,
// 				sshProviderConfigCodec: tt.fields.sshProviderConfigCodec,
// 				kubeClient:             tt.fields.kubeClient,
// 				v1Alpha1Client:         tt.fields.v1Alpha1Client,
// 				scheme:                 tt.fields.scheme,
// 				machineSetupConfigGetter: tt.fields.machineSetupConfigGetter,
// 				kubeadm:                  tt.fields.kubeadm,
// 			}
// 			if err := a.Delete(tt.args.c, tt.args.m); (err != nil) != tt.wantErr {
// 				t.Errorf("Actuator.Delete() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestActuator_Update(t *testing.T) {
// 	type fields struct {
// 		clusterClient            client.ClusterInterface
// 		eventRecorder            record.EventRecorder
// 		sshProviderConfigCodec   *v1alpha1.SSHProviderConfigCodec
// 		kubeClient               *kubernetes.Clientset
// 		v1Alpha1Client           client.ClusterV1alpha1Interface
// 		scheme                   *runtime.Scheme
// 		machineSetupConfigGetter SSHClientMachineSetupConfigGetter
// 		kubeadm                  SSHClientKubeadm
// 	}
// 	type args struct {
// 		c           *clusterv1.Cluster
// 		goalMachine *clusterv1.Machine
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			a := &Actuator{
// 				clusterClient:          tt.fields.clusterClient,
// 				eventRecorder:          tt.fields.eventRecorder,
// 				sshProviderConfigCodec: tt.fields.sshProviderConfigCodec,
// 				kubeClient:             tt.fields.kubeClient,
// 				v1Alpha1Client:         tt.fields.v1Alpha1Client,
// 				scheme:                 tt.fields.scheme,
// 				machineSetupConfigGetter: tt.fields.machineSetupConfigGetter,
// 				kubeadm:                  tt.fields.kubeadm,
// 			}
// 			if err := a.Update(tt.args.c, tt.args.goalMachine); (err != nil) != tt.wantErr {
// 				t.Errorf("Actuator.Update() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestActuator_Exists(t *testing.T) {
// 	type fields struct {
// 		clusterClient            client.ClusterInterface
// 		eventRecorder            record.EventRecorder
// 		sshProviderConfigCodec   *v1alpha1.SSHProviderConfigCodec
// 		kubeClient               *kubernetes.Clientset
// 		v1Alpha1Client           client.ClusterV1alpha1Interface
// 		scheme                   *runtime.Scheme
// 		machineSetupConfigGetter SSHClientMachineSetupConfigGetter
// 		kubeadm                  SSHClientKubeadm
// 	}
// 	type args struct {
// 		c *clusterv1.Cluster
// 		m *clusterv1.Machine
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    bool
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			a := &Actuator{
// 				clusterClient:          tt.fields.clusterClient,
// 				eventRecorder:          tt.fields.eventRecorder,
// 				sshProviderConfigCodec: tt.fields.sshProviderConfigCodec,
// 				kubeClient:             tt.fields.kubeClient,
// 				v1Alpha1Client:         tt.fields.v1Alpha1Client,
// 				scheme:                 tt.fields.scheme,
// 				machineSetupConfigGetter: tt.fields.machineSetupConfigGetter,
// 				kubeadm:                  tt.fields.kubeadm,
// 			}
// 			got, err := a.Exists(tt.args.c, tt.args.m)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Actuator.Exists() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("Actuator.Exists() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
