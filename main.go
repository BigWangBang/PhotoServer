package main

import (
	"PhotoKitServer/config"
	"PhotoKitServer/photo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	config.ReadConfig()
	//开启定时任务
	go config.InitCron()

	router.GET("/photokit/hot", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"code":   http.StatusOK,
				"status": "获取成功",
			},
			"result": photo.HotPhoto,
		})
	})
	router.GET("/photokit/all", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"code":   http.StatusOK,
				"status": "获取成功",
			},
			"result": photo.AllGroupPhoto,
		})
	})
	router.GET("/photokit/version-info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"code":   http.StatusOK,
				"status": "获取成功",
			},
			"result": config.VersionInfo,
		})
	})
	_ = router.Run(":8080")
}
