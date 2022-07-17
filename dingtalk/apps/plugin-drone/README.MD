


示例-text
```yaml
  - name: 钉钉通知
    image: zhaoweiguo/plugin-drone-dingtalk:0.1.3
    environment:
      DINGTALK_GROUP_ACCESS_TOKEN:
        from_secret: dingding
      DINGTALK_DEBUG: true
      DINGTALK_MESSAGE_TYPE: text
      DINGTALK_MESSAGE_TITLE: <title>
    when:
      status: [ failure, success ]
```
示例-link
```yaml
  - name: 钉钉通知
    image: zhaoweiguo/plugin-drone-dingtalk:0.1.3
    environment:
      DINGTALK_GROUP_ACCESS_TOKEN:
        from_secret: dingding
      DINGTALK_DEBUG: true
      DINGTALK_MESSAGE_TYPE: link
      DINGTALK_MESSAGE_TITLE: <title>
      DINGTALK_MESSAGE_TEXT: <text>
      DINGTALK_MESSAGE_PICURL: https://img1.baidu.com/it/u=4240680084,4077606272&fm=253
      DINGTALK_MESSAGE_URL: https://www.zhaoweiguo.com
    when:
      status: [ failure, success ]
```
示例-markdown
```yaml
  - name: 钉钉通知
    image: zhaoweiguo/plugin-drone-dingtalk:0.1.3
    environment:
      DINGTALK_GROUP_ACCESS_TOKEN:
        from_secret: dingding
      DINGTALK_DEBUG: true
      DINGTALK_MESSAGE_TYPE: markdown
      DINGTALK_MESSAGE_TITLE: 杭州天气
      DINGTALK_MESSAGE_TEXT: "#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n"
    when:
      status: [ failure, success ]
```

示例-raw
```yaml
  - name: 钉钉通知
    image: zhaoweiguo/plugin-drone-dingtalk:0.1.3
    environment:
      DINGTALK_GROUP_ACCESS_TOKEN:
        from_secret: dingding
      DINGTALK_DEBUG: true
      DINGTALK_MESSAGE_TYPE: raw
      DINGTALK_MESSAGE_RAW: |
        {
          "msgtype": "link",
          "link": {
              "title": "GitHubAction",
              "text": "[GitHubAction]dingtalk Action is done!",
              "picUrl": "https://img1.baidu.com/it/u=4240680084,4077606272&fm=253",
              "messageUrl": "https://www.zhaoweiguo.com"
          }
        }
    when:
      status: [ failure, success ]
```


