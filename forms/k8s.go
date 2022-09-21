package forms

import (
	"fmt"
	"strconv"
	"strings"

	coreV1 "k8s.io/api/core/v1"
	appsV1 "k8s.io/api/apps/v1"
)

type DeploymentCreateForm struct {
	Namespace string `form:"namespace"`
	Name      string `form:"name"`
	Replicas  int32  `form:"replicas"`
	Image     string `form:"image"`
	Labels    string `form:"labels"`
	Ports     string `form:"ports"`
}

func (f *DeploymentCreateForm) GetLabels() map[string]string {
	labelsMap := make(map[string]string)
	labels := strings.Split(f.Labels, "\n")
	for _, label := range labels {
		values := strings.SplitN(label, ":", 2)
		if len(values) != 2 {
			continue
		}
		labelsMap[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
	}
	return labelsMap
}

func (f *DeploymentCreateForm) GetPorts() []coreV1.ContainerPort {
	portList := make([]coreV1.ContainerPort, 0, 5)
	ports := strings.Split(f.Ports, "\n")
	for _, port := range ports {
		values := strings.SplitN(port, ",", 3)
		if len(values) != 3 {
			continue
		}
		intPort, err := strconv.Atoi(values[1])
		if err != nil {
			continue
		}
		protocol := coreV1.ProtocolTCP
		if strings.Compare(strings.ToLower(values[0]), "tcp") != 0 {
			protocol = coreV1.ProtocolUDP
		}
		portList = append(portList, coreV1.ContainerPort{
			Name:          strings.TrimSpace(values[2]),
			ContainerPort: int32(intPort),
			Protocol:      protocol,
		})
	}

	return portList
}

func (f *DeploymentCreateForm) GetSelectors() map[string]string {
	selectors := f.GetLabels()
	selectors["app"] = f.Name
	return selectors
}

func (f *DeploymentCreateForm) GetImageName() string {
	// 全部为应为字母数字和:
	pos := strings.Index(f.Image, ":")
	if pos > 0 {
		return f.Image[:pos]
	}
	return f.Image
}

type JobModifyForm struct {
	ID     int    `form:"id"`
	Key    string `form:"key"`
	Remark string `form:"remark"`
	Node   int    `form:"node"`
}

type DeploymentModifyForm struct {
	Namespace string `form:"namespace"`
	Name      string `form:"name"`
	Replicas  int    `form:"replicas"`
	Image     string `form:"image"`
	Labels    string `form:"labels"`
	Expose    string `form:"expose"`
}

func (c *DeploymentModifyForm) Exposes() []coreV1.ContainerPort {
	/*
		expose
		name:port:protocol
	*/
	ports := []coreV1.ContainerPort{}
	for _, line := range strings.Split(c.Expose, "\n") {
		nodes := strings.SplitN(strings.TrimSpace(line), ":", 3)
		if len(nodes) == 3 {
			if port, err := strconv.Atoi(nodes[1]); err == nil {
				ports = append(ports, coreV1.ContainerPort{
					Name:          nodes[0],
					ContainerPort: int32(port),
					Protocol:      coreV1.Protocol(strings.ToUpper(nodes[2])),
				})
			}

		}
	}
	return ports
}

func (f *DeploymentModifyForm) FromModel(deployment *appsV1.Deployment) {
	if deployment == nil {
		return
	}
	exposes := []string{}
	for _, port := range deployment.Spec.Template.Spec.Containers[0].Ports {
		exposes = append(exposes, fmt.Sprintf("%s:%d:%s", port.Name, port.ContainerPort, port.Protocol))
	}

	f.Namespace = deployment.Namespace
	f.Name = deployment.Name
	f.Image = deployment.Spec.Template.Spec.Containers[0].Image
	f.Expose = strings.Join(exposes, "\n")
	f.Replicas = int(*deployment.Spec.Replicas)
}
