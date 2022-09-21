package k8s

import "gokade1/services/k8s"

type DaemonsetController struct {
	k8sController
}

func (c *DaemonsetController) Query() {
	c.Data["svc"] = k8s.DaemonsetService.Query()
	// c.Data["q"] = q
	c.TplName = "k8s/query/daemonset.html"
}
