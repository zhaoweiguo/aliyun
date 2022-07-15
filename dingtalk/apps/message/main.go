package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"

)

var Version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "DingTalk Message Tool"
	app.Usage = "Sending message to DingTalk Group"
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
			Name:   "group.access_token",
			Usage:  "dingtalk group access_token",
			EnvVar: "DINGTALK_GROUP_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "group.secret",
			Usage:  "dingtalk group secret",
			EnvVar: "DINGTALK_GROUP_SECRET",
		},
		cli.StringFlag{
			Name:   "message.type",
			Usage:  "dingtalk message type: text, link, markdown, actioncard, feedcard",
			EnvVar: "DINGTALK_MESSAGE_TYPE",
		},
		cli.StringFlag{
			Name:   "message.title",
			Usage:  "dingtalk message title",
			EnvVar: "DINGTALK_MESSAGE_TITLE",
		},

	}

	if _, err := os.Stat("./.env"); err == nil {
		godotenv.Overload("./.env")
	}

	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}

}

func run(c *cli.Context) {
	log.Println("run ...")
}
