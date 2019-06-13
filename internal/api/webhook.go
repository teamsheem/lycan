package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/teamsheem/lycan/internal"
	"github.com/teamsheem/lycan/internal/models"
)

type Webhook struct {
	R *gin.Engine
}

func NewWebhook(r *gin.Engine) Webhook {
	return Webhook{R: r}
}

func (w *Webhook) Handle() {
	w.R.POST("/webhook", postMergeWebhook)
}

//============API Handlers============\\

//type mergeWebhook struct {
//	ObjectAttributes models.Merge `json:"object_attributes" bindings:"required"`
//}

func postMergeWebhook(c *gin.Context) {
	var json models.Merge
	h := c.GetHeader("X-Gitlab-Event")
	//fmt.Println(h)
	if h != "Merge Request Hook" {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "invalid merge webhook",
		})
		return
	}

	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "bad json"})
	}

	if json.ObjectAttributes.State == "merged" {
		fl := internal.Flynn{}
		appName := json.ObjectAttributes.LastCommit.Id
		c, err := fl.CreateApp(appName)
		if err != nil {
			fmt.Println(err, string(c))
			return
		}
		fmt.Println(string(c))
		m, err := fl.CreateMysqlDB(appName)
		if err != nil {
			fmt.Println(err, string(m))
			return
		}

		fmt.Println(string(m))
	}

	c.JSON(201, gin.H{
		"success": true,
	})
}