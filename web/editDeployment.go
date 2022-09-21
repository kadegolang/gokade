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

func Editdeployment() {
	kubecofig := "kube/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", kubecofig) // 从本机加载kubeconfig配置文件，因此第一个参数为空字符串
	clientset, _ := kubernetes.NewForConfig(config)            // 实例化clientset对象

	namespace := "default"
	name := "nginx"

	GetDeployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{}) //name 获取  app name=appname  获取完之后才能更改它
	fmt.Println(GetDeployment)
	fmt.Println()
	fmt.Println(err)

	//修改配置
	var replicas int32 = 2
	var images string = "nginx:1.14"
	GetDeployment.Spec.Replicas = &replicas      
	GetDeployment.Spec.Template.Spec.Containers[0].Image = images

	Deployment, err := clientset.AppsV1().Deployments(namespace).Update(context.TODO(), GetDeployment, metav1.UpdateOptions{})
	fmt.Println(Deployment)
	fmt.Println()
	fmt.Println(err)

}
