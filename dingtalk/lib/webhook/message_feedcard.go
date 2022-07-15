package webhook

/*
{
    "msgtype":"feedCard",
    "feedCard": {
        "links": [
            {
                "title": "时代的火车向前开1",
                "messageURL": "https://www.dingtalk.com/",
                "picURL": "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png"
            },
            {
                "title": "时代的火车向前开2",
                "messageURL": "https://www.dingtalk.com/",
                "picURL": "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png"
            }
        ]
    }
}
*/

type MsgFeedCard struct {
	Links []Link `json:"links"`
}

type Link struct {
	Title 	string `json:"title"`
	PicURL 	string `json:"picURL"`
	MessageURL string `json:"messageURL"`
}


func newFeedCardMsg (title, picUrl, messageUrl string) Message {
	return Message{
		MsgType: MsgTypeFeedCard,
		FeedCard: MsgFeedCard{
			Links: nil,
		},
	}
}


