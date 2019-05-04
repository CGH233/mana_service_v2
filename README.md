## Mana Service 

[![Build Status](https://travis-ci.org/asynccnu/mana_service_v2.svg?branch=master)](https://travis-ci.org/asynccnu/mana_service_v2)

匣子管理 && 元信息 && 部门信息 && 常用网站 服务

### 开发

> 首先把仓库 clone 到 $GOPATH/github.com/asynccnu 下

需要本地 Redis + Mongo，都是默认端口。Mongo admin 的账号：密码 mongoadmin:secret。

可以用 [docker 快速起一个 Mongo](https://hub.docker.com/_/mongo)。注意要创建 admin 账号。

然后在目录下 `make && ./main` 起服务

### Progress

- [x] iOS 元信息（单测完成）
- [x] Banner（两端公用）（单测完成）
- [x] 部门信息
- [x] 常用网站
- [x] 消息通知（出现在 App 里的临时 Notification）
- [x] 用户反馈（单测完成）
- [ ] 安卓校历
- [ ] 安卓闪屏
- [ ] 安卓检查更新
- [ ] 安卓产品列表

### Change log



### Maintainer

+ [zindex](https://github.com/zxc0328)
+ [CGH](https://github.com/CGH233)
