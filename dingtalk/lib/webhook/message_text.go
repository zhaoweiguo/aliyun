package webhook

/*
{
    "msgtype":"text",
    "text": {
        "content":"我就是我, @XXX 是不一样的烟火"
    },
    "at": {
        "atMobiles":[
            "180xxxxxx"
        ],
        "atUserIds":[
            "user123"
        ],
        "isAtAll": false
    }
}
*/

type MsgText struct {
	Content string `json:"content"`
}

func newTextMsg(content string) Message {
	return Message{
		MsgType: MsgTypeText,
		Text: MsgText{content},
	}
}

