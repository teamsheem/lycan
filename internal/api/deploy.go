package api

import "github.com/gin-gonic/gin"

type Deploy struct {
	R *gin.Engine
}

func NewDeploy(r *gin.Engine) Deploy {
	return Deploy{R: r}
}

func (d *Deploy) Handle() {
	d.R.POST("/deploy", PostDeploy)
}

//============API Handlers============\\

type CreateDeployRequest struct {
	UrlAddress string `json:"url_address"`
	EnvVars map[string]string
}

func PostDeploy(c *gin.Context) {
	
}