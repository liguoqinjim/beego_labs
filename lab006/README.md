### 路由设置
一些路由的设置方式

#### 注意点
##### 基本路由
```
beego.Get("/baseGet", func(ctx *context.Context) {
    ctx.Output.Body([]byte("hello world"))
})
```

不用控制器，直接在函数里面处理了

##### 固定路由
```
beego.Router("/user", &controllers.UserController{})
```

访问url的时候调用控制器

##### 正则路由
```
beego.Router("/reg1/?:id", &controllers.Reg1Controller{})
beego.Router("/reg2/:id", &controllers.Reg2Controller{})
```

第一个路由，可以匹配到/reg1或者/reg/1，但是第二个路由是比配不到/reg2的。因为没有和正则匹配到。
其他的一些正则可以看参考资料2里面，有很详细的

##### 自定义方法及 RESTful 规则
```
beego.Router("/rest/list", &controllers.RestController{}, "*:ListFood")
```

这个的意思就是访问这个url的时候，会去调用控制器里面的ListFood方法。`*:ListFood`里面的`*`表示匹配所有的method都执行这个方法。
当然还可以指定method调用哪个方法，见参考资料

##### 自动匹配
```
beego.AutoRouter(&controllers.AutoController{})
```

自动匹配的时候，路由会自动生成这个controller和method对应的路由。

##### 注解路由
```
// @router /all/:key [get]
func (c *CMSController) AllBlock() {
	c.Ctx.Output.Body([]byte("AllBlock"))
}
```

同时还要include这个控制器
```
//注解路由
beego.Include(&controllers.CMSController{})
```

生成的路由会放在`/routers/commentsRouter.go`里面

#### 参考资料
1. https://beego.me/docs/mvc/controller/router.md
2. http://www.cnblogs.com/hezhixiong/p/4602671.html