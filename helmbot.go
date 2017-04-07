package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-chat-bot/bot"
)

type Helmbot struct {
	token    string
	admins   map[string]bool
	channels map[string]bool
	commands map[string]bool
}

const (
	forbiddenUserMessage     string = "%s - ⚠ helm forbidden for user @%s\n"
	forbiddenChannelMessage  string = "%s - ⚠ Channel %s forbidden for user @%s\n"
	forbiddenCommandMessage  string = "%s - ⚠ Command %s forbidden for user @%s\n"
	forbiddenFlagMessage     string = "%s - ⚠ Flag(s) %s forbidden for user @%s\n"
	forbiddenUserResponse    string = "Sorry @%s, but you don't have permission to run this command :confused:"
	forbiddenChannelResponse string = "Sorry @%s, but I'm not allowed to run this command here :zipper_mouth_face:"
	forbiddenCommandResponse string = "Sorry @%s, but I cannot run this command."
	forbiddenFlagResponse    string = "Sorry @%s, but I'm not allowed to run one of your flags."
	okResponse               string = "Roger that!\n@%s, this is the response to your request:\n ```\n%s\n``` "
)

var (
	ignored = map[string]map[string]bool{
		"init": map[string]bool{
			"--canary-image": true,
			" -c":            true,
			"--client-only":  true,
			"--dry-run":      true,
			"-i":             true,
			"--tiller-image": true,
			"--debug":        true,
		},
		"list": map[string]bool{
			"--all":      true,
			"-d":         true,
			"--date":     true,
			"--deleted":  true,
			"--deleting": true,
			"--deployed": true,
			"--failed":   true,
			"-m":         true,
			"--max":      true,
			"-o":         true,
			"--offset":   true,
			"-r":         true,
			"--reverse":  true,
			"-q":         true,
			"--short":    true,
			"--debug":    true,
		},
	}
)

func validateFlags(arguments ...string) error {
	if len(arguments) <= 1 {
		return nil
	}

	for i := 1; i < len(arguments); i++ {
		if ignored[arguments[0]][arguments[i]] {
			return errors.New(fmt.Sprintf("Error: %s is an invalid flag", arguments[i]))
		}

	}

	return nil
}

func helm(command *bot.Cmd) (msg string, err error) {
	t := time.Now()
	time := t.Format(time.RFC3339)
	nickname := command.User.Nick

	fmt.Printf(time)

	output := execute("helm", command.Args...)

	return fmt.Sprintf(okResponse, nickname, output), nil
}

func init() {

	bot.RegisterCommand(
		"helm",
		"helm Slack integration",
		"",
		helm)
}
