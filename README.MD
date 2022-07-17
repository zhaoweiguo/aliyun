
# 简介

此项目用于与 aliyun 有关的小工具、小项目开发

# 钉钉

## message
可以通过命令行，往钉钉群的 webhook 中发消息。

- [x] 实现一个可用版本

## plugin-drone

实现一个用于 drone 的钉钉 webhook 插件

- 最新版本: 0.1.3
- 示例: [plugin drone](./dingtalk/apps/plugin-drone/README.MD)
- Docker镜像：参见项目 [docker image](./dockerfiles/Dockerfile.plugin-drone-dingtalk)

- 基本消息功能完成说明
  - [x] text 格式
  - [x] link 格式
  - [x] markdown 格式
  - [x] raw 格式
  - [ ] actionCard 格式
  - [ ] feedCard 格式
- 周边功能
  - [x] 签名验证功能
  - [ ] At 功能

- 相关参考文档
  - 自定义机器人接入：https://open.dingtalk.com/document/group/custom-robot-access
  - 自定义机器人安全设置: https://open.dingtalk.com/document/robots/customize-robot-security-settings


# OSS


## plugin-drone

- Docker镜像：参见项目 [docker image](./dockerfiles/Dockerfile.plugin-drone-oss)






