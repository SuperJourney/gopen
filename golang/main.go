package main

import (
	"fmt"

	_ "github.com/SuperJourney/gopen/api"
	"github.com/SuperJourney/gopen/config"
	_ "github.com/SuperJourney/gopen/docs"
	"github.com/SuperJourney/gopen/infra"
	"github.com/SuperJourney/gopen/vars"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title OPENAI
// @version 1.0
// @description This is a API documentation for OPENAI.
// @host 192.168.3.246:7211 // 家用ip地址，如果希望远程访问需要改一下
// @host localhost:8080
// @BasePath /api
func main() {
	infra.Engine.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")                                // 允许任意域名访问
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // 允许的 HTTP 方法
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")     // 允许的请求头
		c.Next()
	})
	// infra.Engine.GET("/hello", HelloWorld)
	infra.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// unsafe ; for debug
	infra.Engine.GET("/config", func(ctx *gin.Context) {
		// ret, _ := json.Marshal(config.LoadConfig())
		ctx.JSON(200, vars.Setting)
	})

	// config 热加载
	watcher := configWatch()
	defer watcher.Close()

	infra.Engine.Run(":8080")
}

func configWatch() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	err = watcher.Add(config.ConfigPath)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Config file modified. Reloading...")
					vars.Setting = config.LoadConfig()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()
	return watcher
}
