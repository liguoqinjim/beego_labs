### session

#### 注意
##### 开启session
`beego.BConfig.WebConfig.Session.SessionOn = true`

或者

在配置文件里面加上，`sessionon = true`

##### session的使用
```
func (c *UserController) Get() {
    v := c.GetSession("li")
    if v == nil {
        c.SetSession("li", int(1))
        c.Data["num"] = 0
    } else {
        c.SetSession("li", v.(int)+1)
        c.Data["num"] = v.(int)
    }

    c.TplName = "user.tpl"
}
```

#### 参考资料
https://beego.me/docs/mvc/controller/session.md