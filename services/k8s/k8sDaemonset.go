package k8s

import (
	"context"

	"github.com/astaxie/beego"
	appsv1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type daemonsetService struct {
	k8sService
}

func (s *daemonsetService) Query() []appsv1.DaemonSet {
	clientset, _ := s.GetClient()
	daemonSetList, err := clientset.AppsV1().DaemonSets("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		beego.Error(err)
		return []appsv1.DaemonSet{}
	}
	return daemonSetList.Items
}

var (
	DaemonsetService = new(daemonsetService)
)
