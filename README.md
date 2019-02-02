# Edge PProf

在内部环境，对 Go 服务进行聚合性能分析（PProf）的一个边缘应用

![image](https://i.imgur.com/EBMfh6Q.png)

## 环境要求

- go >= 1.10
- graphviz

## 配置

若要修改默认配置，则到 `conf/app.ini` 文件进行修改。该配置文件通过 go-bindata 嵌入到二进制文件中

## 安装

```
$ go get -u github.com/EDDYCJY/edge-pprof

$ go-bindata -pkg=bindata -o pkg/bindata/bindata.go ./conf/app.ini

$ go run main.go
```

### 运行信息

```
$ go run main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /resource/*filepath       --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /resource/*filepath       --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/debug/pprof/profile --> github.com/EDDYCJY/edge-pprof/server.(*Profile).Handle-fm (3 handlers)
[GIN-debug] GET    /api/v1/debug/pprof/heap  --> github.com/EDDYCJY/edge-pprof/server.(*Heap).Handle-fm (3 handlers)
[GIN-debug] GET    /api/v1/debug/pprof/block --> github.com/EDDYCJY/edge-pprof/server.(*Block).Handle-fm (3 handlers)
[GIN-debug] GET    /api/v1/debug/pprof/mutex --> github.com/EDDYCJY/edge-pprof/server.(*Mutex).Handle-fm (3 handlers)
[GIN-debug] Listening and serving HTTP on 0.0.0.0:8080
```

## 例子

1. debug/pprof/profile：`$HOST/api/v1/debug/pprof/profile?service_name=$NAME&service_port=$PORT`

```
{
    "code": 1,
    "msg": "ok",
    "data": {
        "image_url": "$HOST/resource/profile/$NAME.1549087498.svg",
        "pzpb_url": "$HOST/resource/profile/$NAME.1549087498.pb.gz"
    }
}
```
### 返回格式 

#### .pb.gz

可供下载到本地后利用 `go tool pprof` 进行更详细的性能分析

#### .svg

![image](https://i.imgur.com/xwqqqGI.jpg)

2. debug/pprof/heap：`$HOST/api/v1/debug/pprof/heap?service_name=$NAME&service_port=$PORT`

```
{
    "code": 1,
    "msg": "ok",
    "data": {
        "image_url": "$HOST/resource/heap/$NAME.1549087819.svg",
        "pzpb_url": "$HOST/resource/heap/$NAME.1549087819.pb.gz"
    }
}
```

3. debug/pprof/block: `$HOST/api/v1/debug/pprof/block?service_name=$NAME&service_port=$PORT`

4. debug/pprof/mutex: `$HOST/api/v1/debug/pprof/mutex?service_name=$NAME&service_port=$PORT`

## TODO

- Logging
- JWT
- AliyunOss 支持
- PProf 定制化参数
- Makefile
