# tiktok-demo-backend
![GitHub](https://img.shields.io/github/license/yunyandz/tiktok-demo-backend?style=for-the-badge)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/yunyandz/tiktok-demo-backend/Docker%20Image%20CI?style=for-the-badge)
![GitHub top language](https://img.shields.io/github/languages/top/yunyandz/tiktok-demo-backend?style=for-the-badge)


## 抖音项目服务端demo

字节青训营后端专场抖音项目。
来自4396-云研顶针组。
基于gin/gorm/s3的抖音项目服务端。

🎉*该项目在字节青训营后端专场--基础班抖音项目中获得了* **第八名** *以及* **超级码力奖** *的好成绩！！开香槟咯*🎊🎊🎊

## 运行
``` bash
git clone https://github.com/yunyandz/tiktok-demo-backend.git
cd tiktok-demo-backend
docker compose up -d
```

## 关于项目结构
项目采用了典型的三层结构，以达到更好的组织和管理。

整体部分采用了依赖注入，使得模块的耦合性更低，更容易编写测试用例和维护。

```
.
├── README.md
├── api.json
├── cmd
│   └── main.go //程序入口
├── internal //内部模块
│   ├── config //配置文件
│   │   └── config.go
│   ├── constant //常量
│   │   ├── cache.go
│   │   ├── s3.go
│   │   └── video.go
│   ├── controller //controller层
│   │   ├── comment.go
│   │   ├── controller.go
│   │   ├── favorite.go
│   │   ├── feed.go
│   │   ├── publish.go
│   │   ├── relation.go
│   │   └── user.go
│   ├── dao //数据访问层
│   │   ├── mysql
│   │   │   └── mysql.go
│   │   └── redis
│   │       └── redis.go
│   ├── errorx //错误归档
│   │   └── error.go
│   ├── httpserver //http服务器
│   │   ├── http.go
│   │   ├── middleware
│   │   │   └── jwt.go
│   │   └── router
│   │       └── router.go
│   ├── jwtx //jwt相关
│   │   ├── jwt.go
│   │   └── jwt_test.go
│   ├── kafka //kafka相关
│   │   ├── consumer.go
│   │   └── producer.go
│   ├── logger //zap日志相关
│   │   └── logger.go
│   ├── model //模型层
│   │   ├── comment.go
│   │   ├── user.go
│   │   └── video.go
│   ├── s3 //s3相关
│   │   ├── mock
│   │   │   └── mock_s3.go
│   │   └── s3.go
│   ├── scryptx //密码加解密相关
│   │   └── scrypt.go
│   ├── service //服务层
│   │   ├── comment.go
│   │   ├── feed.go
│   │   ├── publish.go
│   │   ├── publish_test.go
│   │   ├── response.go
│   │   ├── service.go
│   │   ├── user.go
│   │   ├── user_test.go
│   │   └── video.go
│   └── util //工具类
│       ├── cover.go //封面图片处理,使用ffmpeg
│       ├── cover_test.go
│       ├── url.go
│       └── util.go
└── public
    ├── bear.jpg
    ├── bear.mp4
    └── data
```
