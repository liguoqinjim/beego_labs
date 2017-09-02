### 过滤器
在`/user`这个url上面加上一个过滤器，只要session没有值的话就跳转会`/`这个url。
`/login`访问这个url之后，会在session里面设置值

#### 注意点
##### 设置过滤器
```
//过滤器
var FilterLogin = func(ctx *context.Context) {
    //判断session中是否有li这个值
    v := ctx.Input.Session("li")
    if v == nil {
        ctx.Redirect(302, "/")
    } else {
        if _, ok := v.(int); !ok {
            ctx.Redirect(302, "/")
        }
    }
}

//插入过滤器
beego.InsertFilter("/user", beego.BeforeRouter, FilterLogin)
```

`/*`是在所有的路径上都是用过滤器

##### 执行过程
要是过滤器里面用到了session，必须在BeforeStatic之后才能获取，因为session没有在这之前初始化。
其他值的含义见参考资料。

#### 参考资料
https://beego.me/docs/mvc/controller/filter.md