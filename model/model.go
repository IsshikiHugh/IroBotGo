package model

import "github.com/mcoo/OPQBot"

type BotEnvironment struct {
	Manager *OPQBot.BotManager
	Conf    *Configuration
}
type Configuration struct {
	Basic struct {
		QidStr  string `yaml:"AccountId"`
		Qid     int64
		MQidStr string `yaml:"MasterId"`
		MQid    int64
		Url     string `yaml:"ServerUrl"`
		Retry   int    `yaml:"MaxRetry"`
		Key     string `yaml:"CommandKey"`
	} `yaml:"Basic"`
}

type Instruction struct {
	OptionName string // The option name.
	HasArg     bool   // Whether the option has arguments. (We only support one argument for convenience.)
	Args       string // The arguments.
	Content    string // The content of the message.
}

type BotData struct {
	SendToType int   // The send-to-type.
	TargetId   int64 // Message target id. (user id or group id)
}
