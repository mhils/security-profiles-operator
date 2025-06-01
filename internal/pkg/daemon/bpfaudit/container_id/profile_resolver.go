package container_id

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ProfileResolver struct {
	logger    logr.Logger
	cache     map[Profile]ContainerId
	nodeName  string
	clientset *kubernetes.Clientset
}

func NewProfileResolver(clientset *kubernetes.Clientset, nodeName string, logger logr.Logger) *ProfileResolver {
	return &ProfileResolver{
		logger:    logger,
		cache:     make(map[Profile]ContainerId),
		nodeName:  nodeName,
		clientset: clientset,
	}
}

func (b *ProfileResolver) Get(ctx context.Context, profile Profile) (ContainerId, error) {
	containerId, ok := b.cache[profile]
	if !ok {

		pods, err := b.clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{
			FieldSelector: "spec.nodeName=" + b.nodeName,
		})
		if err != nil {
			return "", fmt.Errorf("resolve profile to container id: %w", err)
		}
		if pods == nil {
			return "", errors.New("no pods found in cluster")
		}

		panic("TODO")

		/*
			for p := range pods.Items {
				pod := &pods.Items[p]
				//nolint:gocritic // We explicitly do not want to append to the same slice
				statuses := append(pod.Status.InitContainerStatuses, pod.Status.ContainerStatuses...)
				for c := range statuses {

						containerStatus := statuses[c]
						fullContainerID := containerStatus.ContainerID
						containerName := containerStatus.Name

						containerID := util.ContainerIDRegex.FindString(fullContainerID)
						if containerID == "" {
							b.logger.Error(err,
								"Unable to parse container ID from container status available in pod",
								"fullContainerID", fullContainerID,
								"podName", pod.Name,
								"containerName", containerName,
							)
							continue
						}

						for _, annotation := range []string{
							config.SeccompProfileRecordBpfAnnotationKey,
							config.ApparmorProfileRecordBpfAnnotationKey,
						} {
							key := annotation + containerName
							profile, ok := pod.Annotations[key]
							if ok && profile != "" {
								b.logger.Info(
									"Cache this profile found in cluster",
									"profile", profile,
									"containerID", containerID,
									"podName", pod.Name,
									"containerName", containerName,
								)
								b.containerIDToProfileMap.Insert(containerID, profile)
							}
						}

				}
			}
		*/
	}

	return containerId, nil
}
