### bee工具的使用

#### new命令
新建一个web项目，会在运行的目录下创建一个文件夹。
如：在lab002下运行`bee new lab001`，就会在lab002里面创建lab001文件夹，这个就是lab001就是新建的项目

#### run命令
在app目录下运行`bee run`，这样项目就会启动。虽然我们在目录下运行`go run main.go`也是可以启动项目的。
但是`bee run`命令会监听文件的改变，不用每次修改完重新编译运行了。

#### 参考资料
https://beego.me/docs/install/bee.md