package web

import (
	"context"
	"fmt"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" //重命名
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateDeployment() {
	k8sconfig := "conf/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", k8sconfig)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	clientset, _ := kubernetes.NewForConfig(config)

	namespace := "default"
	var Replicas int32 = 1

	Deployment := &appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx",
			Labels: map[string]string{
				// "app": "ingnx",
				"env": "dev",
			},
		},
		Spec: appv1.DeploymentSpec{
			Replicas: &Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
					"env": "dev",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "nginx",
					Labels: map[string]string{
						"app": "nginx",
						"env": "dev",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								}, //里面是个切片
							},
						},
					},
				},
			},
		},
	} //创建对象

	deploy, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), Deployment, metav1.CreateOptions{})
	fmt.Println(err, deploy)
}
