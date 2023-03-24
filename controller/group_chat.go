package controller

import (
	"IroBot/units/languageHelper"
	"IroBot/units/programerHelper"
	"fmt"
	"strings"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

func GroupChatEvents(botQQ int64, packet *OPQBot.GroupMsgPack) {
	// Ignore message if the sender is the bot.
	if packet.FromUserID == Bot.Conf.Basic.Qid {
		return
	}

	reply, err := OPQBot.ParserGroupReplyMsg(*packet)
	// That is, not a reply.
	if err != nil {
		logrus.Info("Not a reply.")

		// Pretreatment
		inst, err := Parse(packet.Content)
		cmd := inst.Content
		if err != nil {
			return
		}
		logrus.Info(fmt.Sprintf("Receive (%s)[ %s ] from [ %d ]", packet.MsgType, packet.Content, packet.FromGroupID))

		// Choose the option
		switch inst.OptionName {
		case "say":
			// TODO: make it simple
			msg := strings.TrimSpace(strings.TrimPrefix(cmd, "say "))
			if packet.FromUserID != Bot.Conf.Basic.MQid {
				msg = "🥺 不要！"
			} else {
				msg = "🥰 IroBot 也想说 「" + msg + "」"
			}
			Bot.Manager.Send(OPQBot.SendMsgPack{
				SendToType: OPQBot.SendToTypeGroup,
				ToUserUid:  packet.FromGroupID,
				Content: OPQBot.SendTypeTextMsgContent{
					Content: msg,
				},
			})
		case "omg":
			// Code paste bin.
			programerHelper.CodeViewerInGroup(&Bot, packet, inst)
		case "whatis":
			// C Type explainer.
			programerHelper.CTypeExplainerInGroup(&Bot, packet, inst)
		case "trans":
			// Translator.
			languageHelper.TranslatorInGroup(&Bot, packet, inst)
		case "trans-help":
			// List translator support language.
			languageHelper.TranslatorHelpInGroup(&Bot, packet, inst)
		}

	} else {
		logrus.Info("A reply.")
		// Pretreatment
		_ = reply
		// Check if the message is a command.
		logrus.Info(fmt.Sprintf("Receive (%s)[ %s ] from [ %d ]", packet.MsgType, packet.Content, packet.FromGroupID))
	}
}
