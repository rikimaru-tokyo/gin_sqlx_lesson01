package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rikimaru-tokyo/gin_sqlx_lesson01/controller"
)

func LoadRouter(engine *gin.Engine) {

	// グループ・コントローラを通さず、出力したい場合。
	engine.GET("/health_check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	v1 := engine.Group("/v1")
	{
		// コントローラを通して、出力したい場合。
		v1.GET("/hello", controller.Hello)
		v1.GET("/ranks", controller.RankGet)
		v1.POST("/ranks",controller.PostRank)
		v1.GET("/members",controller.GetMembers)
	}
}
