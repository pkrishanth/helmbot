package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-chat-bot/bot/slack"
)

const (
	tokenStr    string = "HELMBOT_SLACK_TOKEN"
	channelsStr string = "HELMBOT_SLACK_CHANNELS_IDS"
	adminsStr   string = "HELMBOT_SLACK_ADMINS_NICKNAMES"
	commandsStr string = "HELMBOT_SLACK_VALID_COMMANDS"
)

var (
	hb *Helmbot
)

func main() {

	if err := validateEnvVars(); err != nil {
		fmt.Printf("Helmbot cannot run. Error: %s\n", err.Error())
		return
	}

	hb = &Helmbot{
		token:    os.Getenv(tokenStr),
		admins:   stringToMap(os.Getenv(adminsStr), ","),
		channels: stringToMap(os.Getenv(channelsStr), ","),
		commands: stringToMap(os.Getenv(commandsStr), ","),
	}

	fmt.Println(hb.admins)

	slack.Run(hb.token)
}

func validateEnvVars() error {
	if os.Getenv(tokenStr) == "" {
		return errors.New(fmt.Sprintf("% env var not defined", tokenStr))
	}
	if os.Getenv(channelsStr) == "" {
		return errors.New(fmt.Sprintf("% env var not defined", channelsStr))
	}
	if os.Getenv(adminsStr) == "" {
		return errors.New(fmt.Sprintf("% env var not defined", adminsStr))
	}
	if os.Getenv(commandsStr) == "" {
		return errors.New(fmt.Sprintf("% env var not defined", commandsStr))
	}

	return nil
}

func stringToMap(s string, sep string) map[string]bool {
	ss := strings.Split(s, sep)
	m := make(map[string]bool)
	for _, word := range ss {
		m[word] = true
	}

	return m

}
