package webhook

/*
{
    "msgtype": "actionCard",
    "actionCard": {
        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
        "text": "![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)
 ### 乔布斯 20 年前想打造的苹果咖啡厅
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
        "btnOrientation": "0",
        "singleTitle" : "阅读全文",
        "singleURL" : "https://www.dingtalk.com/"
    }
}
*/

/*
{
    "msgtype": "actionCard",
    "actionCard": {
        "title": "我 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
        "text": "![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
        "btnOrientation": "0",
        "btns": [
            {
                "title": "内容不错",
                "actionURL": "https://www.dingtalk.com/"
            },
            {
                "title": "不感兴趣",
                "actionURL": "https://www.dingtalk.com/"
            }
        ]
    }
}
*/


type MsgActionCard struct {
	Title		string `json:"title"`
	Text		string `json:"text"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle	string `json:"singleTitle"`
	SingleURL	string `json:"singleURL"`
	Btns		[]Btn  `json:"btns"`
}

type Btn struct {
	Title 		string `json:"title"`
	ActionURL	string `json:"actionURL"`
}


func newActionCardMsgSingle(title, text, btnOrientation, singleTitle, singleURL string) Message {
	return Message{
		MsgType: MsgTypeActionCard,
		ActionCard: MsgActionCard{
			Title:          title,
			Text:           text,
			BtnOrientation: btnOrientation,
			SingleTitle:    singleTitle,
			SingleURL:      singleURL,
		},
	}
}

func newActionCardMsgMulti(title, text, btnOrientation string, subtitles, actionUrls []string) Message {
	// todo
	return Message{
		MsgType: MsgTypeActionCard,
	}
}



