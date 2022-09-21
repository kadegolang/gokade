package web

import (
	"context"
	"fmt"
	"log"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" //重命名
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func K8sClient() {
	k8sconfig := "kube/kube.conf"
	config, err := clientcmd.BuildConfigFromFlags("", k8sconfig) //连接k8s文件
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nodes:")
	for _, node := range nodeList.Items {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
			node.Name,                                    // master
			node.Status.Addresses,                        //  [{InternalIP 10.211.55.9} {Hostname master}]
			node.Status.NodeInfo.BootID,                  //  bccb9c28-e358-4b6e-9ed2-2a51a792d91e
			node.Status.NodeInfo.OSImage,                 //  CentOS Linux 7 (AltArch)
			node.Status.NodeInfo.KubeProxyVersion,        //  v1.15.12
			node.Status.NodeInfo.ContainerRuntimeVersion, //  docker://20.10.17
			node.CreationTimestamp,                       //2022-06-25 13:52:26 +0800 CST
		)
	}

}

func K8sNamespace() {
	k8sconfig := "kube/kube.conf"
	config, err := clientcmd.BuildConfigFromFlags("", k8sconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	fmt.Println("namespace:")
	namespaceList, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, namespace := range namespaceList.Items {
		fmt.Println(namespace.Name, namespace.Status.Phase, namespace.CreationTimestamp)
	}

	fmt.Println("pods:")
	pod, _ := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	for _, pod := range pod.Items {
		fmt.Println(pod.Name, pod.Status.Phase, pod.Namespace)
	}
}

func K8sService() {
	kubeconfig := "kube/kube.conf"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig) // 从本机加载kubeconfig配置文件，因此第一个参数为空字符串
	if err != nil {
		log.Fatal(err)
	}
	clientset, _ := kubernetes.NewForConfig(config) // 实例化clientset对象
	ctx := context.Background()
	service, _ := clientset.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	for _, service := range service.Items {
		fmt.Println(service.Name, service.Spec.ClusterIP, service.Spec.Type)
	}
}

func K8sCreateDeployment() {
	k8sconfig := "kube/kube.conf"
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
									ContainerPort: 8090,
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
