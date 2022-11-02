package main

import (
	"path"
	"path/filepath"
	"runtime/debug"
	"vessl/handlers"
	"vessl/modules/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	debug.SetGCPercent(20)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./webapp/dist/index.html")
		} else {
			c.File("./webapp/dist/" + path.Join(dir, file))
		}
	})

	r.GET("/users/json", helpers.ValidateApiKey(), handlers.GetUsersHandler)
	r.POST("/users/json", helpers.ValidateApiKey(), handlers.UpdateUserHandler)
	r.GET("/users/add", helpers.ValidateApiKey(), handlers.AddUserHandler)
	r.POST("/users/:Id", helpers.ValidateApiKey(), handlers.DeleteUserHandler)
	r.POST("/users/validate", handlers.ValidateUserHandler)
	r.GET("/system/hostnetwork", helpers.ValidateApiKey(), handlers.GetNetworkInfoHandler)
	r.GET("/system/hoststats", helpers.ValidateApiKey(), handlers.GetHostStatsHandler)
	r.POST("/system/hostnetwork", helpers.ValidateApiKey(), handlers.SetNetworkInfoHandler)
	r.POST("/system/restart", helpers.ValidateApiKey(), handlers.RestartHostHandler)
	r.POST("/system/shutdown", helpers.ValidateApiKey(), handlers.ShutDownHostHandler)
	r.GET("/containers/json/:networkName", helpers.ValidateApiKey(), handlers.GetContainersHandler)
	r.GET("/containers/repository/:UserId", helpers.ValidateApiKey(), handlers.GetAppRepositoryHandler)
	r.GET("/containers/info", helpers.ValidateApiKey(), handlers.GetDockerServerInfoHandler)
	r.GET("/containers/cstats", helpers.ValidateApiKey(), handlers.GetCompleteStatsHandler)
	r.GET("/containers/:Id/logs", helpers.ValidateApiKey(), handlers.GetLogsHandler)
	r.POST("/containers/install", helpers.ValidateApiKey(), handlers.InstallContainerHandler)
	r.POST("/containers/:Id/start", helpers.ValidateApiKey(), handlers.StartContainerHandler)
	r.POST("/containers/:Id/stop", helpers.ValidateApiKey(), handlers.StopContainerHandler)
	r.POST("/containers/:Id/restart", helpers.ValidateApiKey(), handlers.RestartContainerHandler)
	r.POST("/containers/:Id/remove", helpers.ValidateApiKey(), handlers.RemoveContainerHandler)
	r.GET("/containers/:Id/stats", helpers.ValidateApiKey(), handlers.GetContainerStatsHandler)
	r.GET("/images/json", helpers.ValidateApiKey(), handlers.GetImagesHandler)
	r.POST("/images/:Id/remove", helpers.ValidateApiKey(), handlers.RemoveImageHandler)
	r.GET("/volumes/json", helpers.ValidateApiKey(), handlers.GetVolumesHandler)
	r.POST("/volumes/:Id/remove", helpers.ValidateApiKey(), handlers.RemoveVolumeHandler)
	r.GET("/networks/json", helpers.ValidateApiKey(), handlers.GetNetworksHandler)
	r.POST("/networks/:Id/create", helpers.ValidateApiKey(), handlers.CreateNetworkHandler)
	r.POST("/networks/:Id/remove", helpers.ValidateApiKey(), handlers.RemoveNetworkHandler)

	err := r.RunTLS(":4443", "./certs/cg-edge.crt", "./certs/cg-edge.key")
	if err != nil {
		panic(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
