package k8s 


import "gokade1/services/k8s"

type SecretsController struct {
	k8sController
}

func (c *SecretsController) Query() {
	c.Data["secrets"] = k8s.SecretService.Query()
	// c.Data["q"] = q
	c.TplName = "k8s/query/secret.html"
}