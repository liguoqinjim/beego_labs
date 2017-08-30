### 参数配置
就配置了test模式下的httpport，其他的参数含义见参考资料

#### 注意
```
appname = quickstart
httpport = 8080
runmode = "test"

[dev]
httpport = 8080
[test]
httpport = 7070
```

runmode是test的时候，就会读取test下的配置

#### 截图
![Imgur](http://i.imgur.com/xl5NqtS.png)

#### 参考资料
https://beego.me/docs/mvc/controller/config.md
