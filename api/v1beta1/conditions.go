package v1beta1

import clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"

// Conditions and condition Reasons for the MicroK8sControlPlane object

const (
	// MachinesReadyCondition reports an aggregate of current status of the machines controlled by the MicroK8sControlPlane.
	MachinesReadyCondition clusterv1.ConditionType = "MachinesReady"
)

const (
	// MachinesBootstrapped is tracking control planes bootstrap status.
	MachinesBootstrapped clusterv1.ConditionType = "MachinesBootstrapped"

	// WaitingForMachinesReason (Severity=Info) documents a MicroK8sControlPlane bootstrap is waiting
	// for all control plane nodes to be created.
	WaitingForMachinesReason = "WaitingForMachines"
)

const (
	// AvailableCondition documents that the first control plane instance has completed MicroK8s boot sequence
	// and so the control plane is available and an API server instance is ready for processing requests.
	AvailableCondition clusterv1.ConditionType = "Available"

	// WaitingForMicroK8sBootReason (Severity=Info) documents a MicroK8sControlPlane object waiting for the first
	// control plane instance to complete MicroK8s boot sequence.
	WaitingForMicroK8sBootReason = "WaitingForMicroK8sBoot"

	// InvalidControlPlaneConfigReason (Severity=Error) documents that controlplane config is invalid and the provider
	// can not proceed with the bootstrap.
	InvalidControlPlaneConfigReason = "InvalidControlPlaneConfig"
)

const (
	// ResizedCondition documents a MicroK8sControlPlane that is resizing the set of controlled machines.
	ResizedCondition clusterv1.ConditionType = "Resized"

	// ScalingUpReason (Severity=Info) documents a MicroK8sControlPlane that is increasing the number of replicas.
	ScalingUpReason = "ScalingUp"

	// ScalingDownReason (Severity=Info) documents a MicroK8sControlPlane that is decreasing the number of replicas.
	ScalingDownReason = "ScalingDown"
)

const (
	// ControlPlaneComponentsHealthyCondition reports the overall status of control plane components
	// implemented as static pods generated by MicroK8s including kube-api-server, kube-controller manager,
	// kube-scheduler and etcd.
	ControlPlaneComponentsHealthyCondition clusterv1.ConditionType = "ControlPlaneComponentsHealthy"

	// ControlPlaneComponentsUnhealthyReason (Severity=Error) documents a control plane component not healthy.
	ControlPlaneComponentsUnhealthyReason = "ControlPlaneComponentsUnhealthy"

	// ControlPlaneComponentsInspectionFailedReason documents a failure in inspecting the control plane component status.
	ControlPlaneComponentsInspectionFailedReason = "ControlPlaneComponentsInspectionFailed"
)

const (
	// EtcdClusterHealthyCondition documents the overall etcd cluster's health.
	EtcdClusterHealthyCondition clusterv1.ConditionType = "EtcdClusterHealthyCondition"

	// EtcdClusterUnhealthyReason (Severity=Error) is set when the etcd cluster is unhealthy.
	EtcdClusterUnhealthyReason = "EtcdClusterUnhealthy"
)

const (
	// MachinesCreatedCondition documents that the machines controlled by the MicroK8sControlPlane are created.
	// When this condition is false, it indicates that there was an error when cloning the infrastructure/bootstrap template or
	// when generating the machine object.
	MachinesCreatedCondition clusterv1.ConditionType = "MachinesCreated"

	// InfrastructureTemplateCloningFailedReason (Severity=Error) documents a MicroK8sControlPlane failing to
	// clone the infrastructure template.
	InfrastructureTemplateCloningFailedReason = "InfrastructureTemplateCloningFailed"

	// BootstrapTemplateCloningFailedReason (Severity=Error) documents a MicroK8sControlPlane failing to
	// clone the bootstrap template.
	BootstrapTemplateCloningFailedReason = "BootstrapTemplateCloningFailed"

	// MachineGenerationFailedReason (Severity=Error) documents a MicroK8sControlPlane failing to
	// generate a machine object.
	MachineGenerationFailedReason = "MachineGenerationFailed"
)
