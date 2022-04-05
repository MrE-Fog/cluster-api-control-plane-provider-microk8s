package controllers

import (
	"context"

	clusterv1beta1 "github.com/AlexsJones/cluster-api-control-plane-provider-microk8s/api/v1beta1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func (r *MicroK8sControlPlaneReconciler) ensureNodesBooted(ctx context.Context,
	mcp *clusterv1beta1.MicroK8sControlPlane, cluster *clusterv1.Cluster, machines []clusterv1.Machine) error {
	// client, err := r.talosconfigForMachines(ctx, mcp, machines...)
	// if err != nil {
	// 	return err
	// }

	// defer client.Close() //nolint:errcheck

	// ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	// nodesBootStarted := map[string]struct{}{}
	// nodesBootStopped := map[string]struct{}{}

	// err = client.EventsWatch(ctx, func(ch <-chan talosclient.Event) {
	// 	defer cancel()

	// 	for event := range ch {
	// 		if msg, ok := event.Payload.(*machineapi.SequenceEvent); ok {
	// 			if msg.GetSequence() == "boot" { // can't use runtime constants as they're in `internal/`
	// 				switch msg.GetAction() { //nolint:exhaustive
	// 				case machineapi.SequenceEvent_START:
	// 					nodesBootStarted[event.Node] = struct{}{}
	// 				case machineapi.SequenceEvent_STOP:
	// 					nodesBootStopped[event.Node] = struct{}{}
	// 				}
	// 			}
	// 		}
	// 	}
	// }, talosclient.WithTailEvents(-1))

	// if err != nil {
	// 	unwrappedErr := err

	// 	for {
	// 		if s, ok := status.FromError(unwrappedErr); ok && s.Code() == codes.DeadlineExceeded {
	// 			// ignore deadline exceeded as we've just exhausted events list
	// 			err = nil

	// 			break
	// 		}

	// 		unwrappedErr = errors.Unwrap(unwrappedErr)
	// 		if unwrappedErr == nil {
	// 			break
	// 		}
	// 	}
	// }

	// if err != nil && !errors.Is(err, context.Canceled) {
	// 	return err
	// }

	// nodesNotFinishedBooting := []string{}

	// // check for nodes which have Boot/Start event, but no Boot/Stop even
	// // if the node is up long enough, Boot/Start even might get out of the window,
	// // so we can't check such nodes reliably
	// for node := range nodesBootStarted {
	// 	if _, ok := nodesBootStopped[node]; !ok {
	// 		nodesNotFinishedBooting = append(nodesNotFinishedBooting, node)
	// 	}
	// }

	// sort.Strings(nodesNotFinishedBooting)

	// if len(nodesNotFinishedBooting) > 0 {
	// 	return fmt.Errorf("nodes %q are still in boot sequence", nodesNotFinishedBooting)
	// }

	return nil
}