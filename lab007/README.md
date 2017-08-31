### 获取参数

#### 注意
##### 获取参数
```
info := c.GetString("info")
```

##### 直接解析到struct
```
<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <form action="/info" method="post">
        名字:<input name="username" type="text"><br>
        年龄:<input name="age" type="text"><br>
        邮箱:<input name="email" type="text"><br>
        <input type="submit" value="提交">
    </form>
</body>
</html>
```

解析到struct
```
type info struct {
	Id    string `-`
	Name  string `form:"username"`
	Age   int    `form:"age"`
	Email string `form:"email"`
}

func (c *InfoController) Post() {
	info := info{}
	if err := c.ParseForm(&info); err != nil {
		c.Ctx.Output.Body([]byte(fmt.Sprintf("parseForm error:%v", err)))
	} else {
		c.Ctx.Output.Body([]byte(fmt.Sprintf("info:%+v", info)))
	}
}
```

#### 参考资料
https://beego.me/docs/mvc/controller/params.md