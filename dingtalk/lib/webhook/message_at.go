package webhook

/*
{
    "at": {
        "atMobiles":[
            "180xxxxxx"
        ],
        "atUserIds":[
            "user123"
        ],
        "isAtAll": false
    },
    "text": {
        "content":"我就是我, @XXX 是不一样的烟火"
    },
    "msgtype":"text"
}
*/

type AT struct {
	AtMobiles	[]string `json:"atMobiles"`
	AtUserIds 	[]string `json:"atUserIds"`
	IsAtAll 	bool	 `json:"isAtAll"`
}


