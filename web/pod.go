package web

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" //重命名
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Podlist() {
	path := "conf/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", path)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	PodList, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	fmt.Println(err)

	for _, Podlist := range PodList.Items {
		fmt.Printf("pod:%s",Podlist.Namespace)
		fmt.Println(Podlist.Name)
	}
}
