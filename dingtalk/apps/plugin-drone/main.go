package main

import (
	"errors"
	"fmt"
	"github.com/zhaoweiguo/aliyun/dingtalk/lib/webhook"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"

)

var Version = "0.0.1"

func init()  {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	app := cli.NewApp()
	app.Name = "Drone Plugin: DingTalk Message Tool"
	app.Usage = "Drone Plugin: Sending message to DingTalk Group"
	year := time.Now().Year()
	app.Copyright = fmt.Sprintf("© 2022-%d zhaoweiguo.com", year)
	app.Authors = []cli.Author{
		{
			Name:  "zhaoweiguo",
			Email: "admin@zhaoweiguo.com",
		},
	}
	app.Action = run
	app.Version = Version

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "DINGTALK_DEBUG",
		},
		cli.StringFlag{
			Name:   "group.access_token,access_token",
			Usage:  "dingtalk group access_token",
			EnvVar: "DINGTALK_GROUP_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "group.secret,secret",
			Usage:  "dingtalk group secret",
			EnvVar: "DINGTALK_GROUP_SECRET",
		},
		cli.StringFlag{
			Name:   "message.type,type",
			Usage:  "dingtalk message type: text, link, markdown, actioncard, feedcard",
			EnvVar: "DINGTALK_MESSAGE_TYPE",
		},
		cli.StringFlag{
			Name:   "message.title,title",
			Usage:  "dingtalk message title",
			EnvVar: "DINGTALK_MESSAGE_TITLE",
		},
		cli.StringFlag{
			Name:   "message.text,text",
			Usage:  "dingtalk message text",
			EnvVar: "DINGTALK_MESSAGE_TEXT",
		},
		cli.StringFlag{
			Name:   "message.picUrl,picUrl",
			Usage:  "dingtalk message picUrl",
			EnvVar: "DINGTALK_MESSAGE_PICURL",
		},
		cli.StringFlag{
			Name:   "message.messageUrl,messageUrl",
			Usage:  "dingtalk message messageUrl",
			EnvVar: "DINGTALK_MESSAGE_URL",
		},

	}

	if _, err := os.Stat("./.env"); err == nil {
		godotenv.Overload("./.env")
	}

	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}

}

// 示例: go run main.go --message.title="111" --message.type="text" --group.access_token="cac7ba30a71cafe8751d0585bdxxxxxx
func run(c *cli.Context) {
	debug := c.Bool("debug")

	accessToken := c.String("group.access_token")
	secret := c.String("group.secret")

	messageType := c.String("message.type")

	messageTitle := c.String("message.title")
	messageText := c.String("message.text")
	messagePicUrl := c.String("message.picUrl")
	messageUrl := c.String("message.messageUrl")

	webhooker := webhook.NewWebHook(accessToken, secret, debug)

	var err error
	switch (webhook.MsgType)(messageType) {
	case webhook.MsgTypeText:
		err = webhooker.SendTextMsg(messageTitle)
	case webhook.MsgTypeLink:
		err = webhooker.SendLinkMsg(messageTitle, messageText, messagePicUrl, messageUrl)
	default:
		err = errors.New("not support message type: " + messageType)
	}
	if err != nil {
		panic(err)
	}
}




