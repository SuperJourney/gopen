# GOPEN

AICPT 大赛 后端代码 (开源)

```bash
docker-compose up -d
```

**项目配置文件**

golang/config.toml

[ChatGPT]
ApiToken = ""

**Doc**

[https://localhost:7211/swagger/index.html]()

相关路由：

app资源

```
// app 应用场景
router.GET("/apps", appCtrl.GetApps)
router.GET("/apps/:app_id", appCtrl.GetApp)
router.POST("/apps", appCtrl.CreateApp)
router.PUT("/apps/:app_id", appCtrl.UpdateApp)
router.DELETE("/apps/:app_id", appCtrl.DeleteApp)
```

// Attr资源

```
//attr 应用属性
router.GET("/apps/:app_id/attrs/", attCtrl.GetAttrs)
router.GET("/apps/:app_id/attrs/:attr_id", attCtrl.GetAttr)
router.POST("/apps/:app_id/attrs", attCtrl.CreateAttr)
router.POST("/apps/:app_id/chat_attrs", attCtrl.CreateChatAttr)
router.PUT("/apps/:app_id/attrs/:attr_id", attCtrl.UpdateAttr)
router.PUT("/apps/:app_id/chat_attrs/:attr_id", attCtrl.UpdateChatAttr)
router.DELETE("/apps/:app_id/attrs/:attr_id", attCtrl.DeleteAttr)
```

// chatgpt 

```
router.POST("/gpt/:attr_id/chat-completion", chatCtrl.Request)
router.POST("/gpt/:attr_id/chat-completion/stream", chatCtrl.SteamRequest)
```

// sd (stable diffsion)

```
router.POST("/sd/:attr_id/img2img", chatCtrl.ImgToImg)
router.POST("/sd/:attr_id/img2imgurl", chatCtrl.ImgToImgUrl)
```


## circleci

 依赖 docker

## 项目目录

golang: 应用管理以及chatgpt接入

nginx: proxy

pages: 待完善

python: 粘合sd接口

### 开发配置：

vscode  devcontainer
