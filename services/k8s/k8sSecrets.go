package k8s

import (
	"context"

	"github.com/astaxie/beego"
	corev1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type secretService struct {
	k8sService
}

func (s *secretService) Query() []corev1.Secret {
	clientset, _ := s.GetClient()
	secretList, err := clientset.CoreV1().Secrets("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		beego.Error(err)
		return []corev1.Secret{}
	}
	return secretList.Items
}

var SecretService = new(secretService)
