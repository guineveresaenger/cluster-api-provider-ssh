package machine

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/cluster-api/pkg/util"
)

type MachineStatus *clusterv1.Machine

const (
	InstanceStatusAnnotationKey = "instance-status"
)

// Get the status of the instance identified by the given machine
func (a *Actuator) status(m *clusterv1.Machine) (MachineStatus, error) {
	if a.v1Alpha1Client == nil {
		return nil, nil
	}
	currentMachine, err := util.GetMachineIfExists(a.v1Alpha1Client.Machines(m.Namespace), m.ObjectMeta.Name)
	if err != nil {
		return nil, err
	}

	if currentMachine == nil {
		// The current status no longer exists because the matching CRD has been deleted (or does not exist yet ie. bootstrapping)
		return nil, nil
	}
	return a.machineStatus(currentMachine)
}

// Gets the state of the instance stored on the given machine CRD
func (a *Actuator) machineStatus(m *clusterv1.Machine) (MachineStatus, error) {
	if m.ObjectMeta.Annotations == nil {
		return nil, nil
	}

	annot := m.ObjectMeta.Annotations[InstanceStatusAnnotationKey]
	if annot == "" {
		return nil, nil
	}

	serializer := json.NewSerializer(json.DefaultMetaFactory, a.scheme, a.scheme, false)
	var status clusterv1.Machine
	_, _, err := serializer.Decode([]byte(annot), &schema.GroupVersionKind{Group: "", Version: "cluster.k8s.io/v1alpha1", Kind: "Machine"}, &status)
	if err != nil {
		return nil, fmt.Errorf("decoding failure: %v", err)
	}

	return MachineStatus(&status), nil
}
