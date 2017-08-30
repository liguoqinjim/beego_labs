### 运行入门程序

#### 创建项目
`bee new quickstart`

#### 运行项目
`bee run`

#### 路由设置
```
beego.Router("/", &controllers.MainController{})
```

第一个参数是url，第二个参数是对应的控制器controller

#### 控制器
https://beego.me/docs/quickstart/controller.md

#### View编写
`Official website: <a href="http://{{.Website}}">{{.Website}}</a>`中的
`{{.Website}}`就会读取controller传过来的数据

#### 静态文件处理
https://beego.me/docs/quickstart/static.md

`StaticDir["/static"] = "static"`，URL 前缀和映射的目录

https://beego.me/docs/quickstart/view.md

#### 参考资料
https://beego.me/docs/quickstart/