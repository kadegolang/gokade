package web

import (
	"context"
	"fmt"

	// appv1 "k8s.io/api/apps/v1"
	// corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" //重命名
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func EditReplicas() {
	kubecofig := "kube/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", kubecofig) // 从本机加载kubeconfig配置文件，因此第一个参数为空字符串
	clientset, _ := kubernetes.NewForConfig(config)            // 实例化clientset对象

	namespace := "default"
	name := "nginx"

	GetScale, err := clientset.AppsV1().Deployments(namespace).GetScale(context.TODO(), name, metav1.GetOptions{}) //name 获取  app name=appname  获取完之后才能更改它
	fmt.Println(GetScale)
	fmt.Println()
	fmt.Println(err)

	//修改配置

	GetScale.Spec.Replicas = 1  // GetDeployment.Spec.Replicas = 2
	

	UpdateScale, err := clientset.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), name, GetScale, metav1.UpdateOptions{})
	fmt.Println(UpdateScale)
	fmt.Println()
	fmt.Println(err)

}
