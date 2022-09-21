package k8s

import (
	"context"

	"github.com/astaxie/beego"
	corev1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type namespaceService struct {
	k8sService
}

func (s *namespaceService) Query() []corev1.Namespace {
	clientset, err := s.GetClient()
	if err != nil {
		beego.Error(err)
		return []corev1.Namespace{}
	}
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		beego.Error(err)
		return []corev1.Namespace{}
	}
	return namespaceList.Items
}

var NamespaceService = new(namespaceService)

