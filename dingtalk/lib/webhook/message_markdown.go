package webhook

/*
{
     "msgtype": "markdown",
     "markdown": {
         "title":"杭州天气",
         "text": "#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n"
     },
      "at": {
          "atMobiles": [
              "150XXXXXXXX"
          ],
          "atUserIds": [
              "user123"
          ],
          "isAtAll": false
      }
 }
*/

type MsgMarkdown struct {
	Title 	string `json:"title"`
	Text	string `json:"text"`
}

func newMarkdownMsg (title, text string) Message {
	return Message{
		MsgType: MsgTypeMarkdown,
		Markdown: MsgMarkdown{title, text},
	}
}




