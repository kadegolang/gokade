package k8s

import (
	"context"

	"github.com/astaxie/beego"
	corev1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type svcService struct {
	k8sService
}

func (s *svcService) QuerySvc() []corev1.Service {
	clientset, _ := s.GetClient()
	serviceList, err := clientset.CoreV1().Services("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		beego.Error(err)
		return []corev1.Service{}
	}

	return serviceList.Items
}

var (
	SvcService = new(svcService)
)
