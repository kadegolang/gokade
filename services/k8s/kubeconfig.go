package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sService struct {
}

func (s *k8sService) GetClient() (*kubernetes.Clientset, error) {
	k8sconfig := "conf/kube.conf"
	config, err := clientcmd.BuildConfigFromFlags("", k8sconfig)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
