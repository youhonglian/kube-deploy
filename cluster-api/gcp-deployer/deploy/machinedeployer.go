package deploy

import (
	clusterv1 "k8s.io/kube-deploy/cluster-api/pkg/apis/cluster/v1alpha1"
	"k8s.io/kube-deploy/cluster-api/pkg/controller/machine"
)

// Provider-specific machine logic the deployer needs.
type machineDeployer interface {
	machine.Actuator
	GetIP(machine *clusterv1.Machine) (string, error)
	GetKubeConfig(master *clusterv1.Machine) (string, error)

	// Create and start the machine controller. The list of initial
	// machines don't have to be reconciled as part of this function, but
	// are provided in case the function wants to refer to them (and their
	// ProviderConfigs) to know how to configure the machine controller.
	// Not idempotent.
	CreateMachineController(cluster *clusterv1.Cluster, initialMachines []*clusterv1.Machine) error
	PostDelete(cluster *clusterv1.Cluster, machines []*clusterv1.Machine) error
}
