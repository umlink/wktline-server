# Wktline 项目管理服务端
> 温馨提示：本项目只用作个人学习，不可商用，如需商用，请联系作者授权，微信：`kr_39hd`

前端项目地址：[前端github](https://github.com/umlink/wktline-client)

界面内容查看和介绍：[项目详细介绍](https://juejin.cn/post/7410062139275984936)

如何启动项目
1. 基础环境搭建，请移步 [GoFrame](https://goframe.org/)
2. 安装 mysql8
3. 获取 sql 文件，添加作者微信：`kr_39hd`
4. 打开配置文件 `hack` -> `config.yaml`，修改 mysql 地址，此配置文件用与生成本地代码使用
5. 打开配置文件 `manifest` -> `config` -> `config.yaml`，修改 database 中的 mysql 地址，项目启动时使用该配置
6. 设置 `upload`，默认为七牛云

启动命令（热更新）
```base
gf run main.go
```