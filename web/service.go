package web

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" //重命名
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateService() {
	k8sconfig := "conf/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", k8sconfig) // 从本机加载kubeconfig配置文件，因此第一个参数为空字符串
	clientset, _ := kubernetes.NewForConfig(config)            // 实例化clientset对象

	namespace := "default"

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx-service",
			Labels: map[string]string{
				"env": "dev",
			},
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Selector: map[string]string{
				"env": "dev",
				"app": "nginx",
			},
			Ports: []corev1.ServicePort{
				{
					Name:     "http",
					Port:     80,
					Protocol: corev1.ProtocolTCP,
				},
			},
		},
	}

	service, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	fmt.Println(service)
	fmt.Println()
	fmt.Println(err)

}
