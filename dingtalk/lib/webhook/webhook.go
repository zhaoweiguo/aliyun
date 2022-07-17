package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 自定义机器人接入：https://open.dingtalk.com/document/group/custom-robot-access

type IWebHook interface {
	SendRawMessage(msg string) error
	SendMessage(msg Message) error
	SendTextMsg(content string) error
	SendLinkMsg(title, text, picUrl, messageUrl string) error
	SendMarkdownMsg(title, text string) error
	SendActionCardSingle(title, text, btnOrientation, singleTitle, singleURL string) error
	SendActionCardMulti(title, text, btnOrientation string, titles, actionUrls []string) error
}

type WebHook struct {
	apiUrl      string
	secret      string
	debug 		bool
}

// https://oapi.dingtalk.com/robot/send?access_token=cac7ba30a71caxxxx
func NewWebHook(accessToken, secret string, debug bool) *WebHook {
	baseUrl := "https://oapi.dingtalk.com/robot/send?access_token="
	return &WebHook{
		apiUrl: baseUrl + accessToken,
		secret: secret,
		debug: debug,
	}
}

func (w WebHook) SendMessage(msg Message) error {
	bs, _ := json.Marshal(msg)
	if w.debug {
		log.Println(string(bs))
	}
	return w.SendRawMessage(bs)
}

func (w WebHook) SendRawMessage(msg []byte) error {
	if w.debug {
		log.Println(w)
	}
	if w.secret != "" {
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		sign := getSign(timestamp, w.secret)
		w.apiUrl = fmt.Sprintf("%s&timestamp=%d&sign=%s", w.apiUrl, timestamp, sign)
	}

	resp, err := http.Post(w.apiUrl, "application/json", bytes.NewReader(msg))
	if nil != err {
		return errors.New("request error: " + err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if 200 != resp.StatusCode {
		return fmt.Errorf("response status error: %d", resp.StatusCode)
	}
	if w.debug {
		log.Println(string(body))
	}

	var result Result
	err = json.Unmarshal(body, &result)
	if nil != err {
		return errors.New("response error: need a json: " + err.Error())
	}
	if 0 != result.Errorcode {
		return fmt.Errorf("response error: {code: %d, msg: %s}", result.Errorcode, result.Errmsg)
	}
	return nil

}

func (w WebHook) SendTextMsg(content string) error {
	message := newTextMsg(content)
	return w.SendMessage(message)
}

func (w WebHook) SendLinkMsg(title, text, picUrl, messageUrl string) error {
	message := newLinkMsg(title, text, picUrl, messageUrl)
	return w.SendMessage(message)
}

func (w WebHook) SendMarkdownMsg(title, text string) error {
	message := newMarkdownMsg(title, text)
	return w.SendMessage(message)
}

func (w WebHook) SendActionCardSingle(title, text, btnOrientation, singleTitle, singleURL string) error {
	message := newActionCardMsgSingle(title, text, btnOrientation, singleTitle, singleURL)
	return w.SendMessage(message)
}

func (w WebHook) SendActionCardMulti(title, text, btnOrientation string, titles, actionUrls []string) error {
	message := newActionCardMsgMulti(title, text, btnOrientation, titles, actionUrls)
	return w.SendMessage(message)
}

// 设备类型: text, link, markdown, actioncard(1,2), feedcard

type MsgType string
const (
	MsgTypeRaw MsgType = "raw"
	MsgTypeText MsgType = "text"
	MsgTypeLink MsgType = "link"
	MsgTypeMarkdown MsgType = "markdown"
	MsgTypeActionCard MsgType = "actioncard"
	MsgTypeFeedCard MsgType = "feedcard"
)

type Message struct {
	MsgType		MsgType 	`json:"msgtype"`
	Text 		MsgText 	`json:"text"`
	Link 		MsgLink		`json:"link"`
	Markdown 	MsgMarkdown	`json:"markdown"`
	ActionCard MsgActionCard `json:"actionCard"`
	FeedCard 	MsgFeedCard	`json:"feedCard"`
	At 			AT 			`json:"at"`
}












