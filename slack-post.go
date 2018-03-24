package main

import (
	"flag"
	"fmt"

	"github.com/kkimu/slack-post-cli/model"
)

var (
	textOpt    = flag.String("m", "default", "message text")
	channelOpt = flag.String("c", "test", "help message for s option")
)

func main() {
	flag.Parse()

	slack, err := model.NewSlack()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	slack.Payload.Text = *textOpt
	slack.Payload.Channel = *channelOpt

	if err = slack.Post(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
