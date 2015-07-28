# zlist

一个方便查看多个网站上热点信息的应用。

## 网页版

[zlist.whiteworld.me](http://zlist.whiteworld.me/)

## 主要依赖的库

[zlistutil](https://github.com/zlisthq/zlistutil)

## 运行部署

    # 方式0
    go get github.com/zlisthq/zlist
    zlist
    # 方式1: 使用官方 zlist 镜像，没有 Redis 缓存功能 
    docker run -p 8080:8080 whiteworld/zlist
    # 方式2: 使用 Reids 缓存服务
    git clone git@github.com:zlisthq/zlist.git
    cd zlist
    docker-compose up

## TODO

- Custom UI
- Add Tests
- ~~Add Cache(Redis)~~
- Use Gin
