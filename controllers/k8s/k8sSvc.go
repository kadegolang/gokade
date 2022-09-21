package k8s

import "gokade1/services/k8s"

type SvcController struct {
	k8sController
}

func (c *SvcController) QuerySvc() {
	c.Data["svc"] = k8s.SvcService.QuerySvc()
	// q := c.GetString("q")
	// c.Data["q"] = q
	c.TplName = "k8s/query/svc.html"
}
