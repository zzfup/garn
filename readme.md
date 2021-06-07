# garn

仿照yarn的命令格式。

自动抓取

```
garn add xxx
```

# 优点
自动抓取 https://pkg.go.dev/ 上的包，用户不用去记录整个包的路径，只需要记得是哪个包的名字，即可添加依赖。

比如我要下载gin

不用这样
```
go get -u -v github.com/gin-gonic/gin
```

只需要执行
```
garn add gin
```

就可以自动抓取到符合条件的包了。
