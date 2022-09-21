package web

import (
	"context"

	// appv1 "k8s.io/api/apps/v1"
	// corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" //重命名
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func DeleteRescource() {
	kubecofig := "kube/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", kubecofig) // 从本机加载kubeconfig配置文件，因此第一个参数为空字符串
	clientset, _ := kubernetes.NewForConfig(config)            // 实例化clientset对象

	namespace := "default"
	podapp, servicename := "nginx", "nginx-service"

	clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), podapp, metav1.DeleteOptions{}) //name 获取  app name=appname  获取完之后才能更改它

	clientset.CoreV1().Services(namespace).Delete(context.TODO(), servicename, metav1.DeleteOptions{})

}
