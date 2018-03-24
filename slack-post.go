package main

import (
	"flag"
	"fmt"
	"os/user"

	"github.com/kkimu/slack-post-cli/model"
)

var (
	textOpt    string
	channelOpt string
	configOpt  string
)

func init() {
	usr, _ := user.Current()
	homeDir := usr.HomeDir

	flag.StringVar(&textOpt, "m", "default", "message string")
	flag.StringVar(&channelOpt, "ch", "test", "channel to send")
	flag.StringVar(&configOpt, "config", homeDir+"/slack-post-cli.toml", "config file path")
}

func main() {
	flag.Parse()

	slack, err := model.NewSlack(configOpt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	slack.Payload.Text = textOpt
	slack.Payload.Channel = channelOpt

	if err = slack.Post(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
