### 静态文件处理

#### 注意点
##### 默认的路径
static目录是默认注册为静态处理的目录的

##### 自定义路径
```
beego.SetStaticPath("/download1", "down1")
```

`/download1`是url，`down1`是文件路径

#### 参考资料
https://beego.me/docs/quickstart/static.md