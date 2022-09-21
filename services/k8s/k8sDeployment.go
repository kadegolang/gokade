package k8s

import (
	"context"
	"fmt"
	"gokade1/forms"
	"gokade1/tools"
	"strings"

	"github.com/astaxie/beego"
	appsV1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type deploymentService struct {
	k8sService
}


func (s *deploymentService) Query() []appsV1.Deployment {
	clientset, err := s.GetClient()
	if err != nil {
		beego.Error(err)
		return []appsV1.Deployment{}
	}
	deploymentList, err := clientset.AppsV1().Deployments("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		beego.Error(err)
		return []appsV1.Deployment{}
	}

	// for _, j := range deploymentList.Items {
	// 	fmt.Sprintf("%s\n", j.Spec.Template.Spec.Containers[0].Image)
	// } //遍历的是这里面的数据

	return deploymentList.Items
}

func (s *deploymentService) QueryPod() []corev1.Pod {
	clientset, _ := s.GetClient()
	podlist, err := clientset.CoreV1().Pods("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		beego.Error(err)
		return []corev1.Pod{}
	}
	return podlist.Items
}

func (s *deploymentService) Delete(name string, namespace string) {
	clientset, err := s.GetClient()
	if err != nil {
		beego.Error(err)
		return
	}
	clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (s *deploymentService) DeletePod(name string, namespace string) {
	clientset, err := s.GetClient()
	if err != nil {
		beego.Error(err)
		return
	}
	clientset.CoreV1().Pods(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (s *deploymentService) Create(form *forms.DeploymentCreateForm) {
	fmt.Printf("%#v\n", form)
	clientset, err := s.GetClient()
	if err != nil {
		beego.Error(err)
		return
	}
	var Replicas int32 = 1
	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name:   form.Name,
			Labels: form.GetLabels(),
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: &Replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels: form.GetSelectors(),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Name:   form.Name,
					Labels: form.GetSelectors(),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  form.GetImageName(),
							Image: form.Image,
							Ports: form.GetPorts(),
						},
					},
				},
			},
		},
	}
	_, err = clientset.AppsV1().Deployments(form.Namespace).Create(context.TODO(), deployment, metaV1.CreateOptions{})
	fmt.Println(err)
}

func (s *deploymentService) Modify(namespace, name, image string, ports []corev1.ContainerPort, replicas int) {

	clientset, err := s.GetClient()
	if err != nil {
		return
	}
	fmt.Println(namespace, name, image, ports, replicas)

	names := strings.SplitN(image, ":", 2)
	imageName := names[0]

	deploymentsclient := clientset.AppsV1().Deployments(namespace)

	deployment, err := deploymentsclient.Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return
	}

	deployment.Spec.Replicas = tools.Int32str(int32(replicas))
	deployment.Spec.Template.Spec.Containers[0].Name = imageName
	deployment.Spec.Template.Spec.Containers[0].Image = image
	deployment.Spec.Template.Spec.Containers[0].Ports = ports

	_, err = deploymentsclient.Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	fmt.Println(err)

}

var (
	DeploymentService = new(deploymentService)
)
