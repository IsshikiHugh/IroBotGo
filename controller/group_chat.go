package controller

import (
	"IroBot/units/emojiHelper"
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
		case "menu":
			msg := fmt.Sprintf("🥰 你好！我是 IroBot！目前我支持这些功能！\n")
			msg += fmt.Sprintf("🔑「%s menu 」\n      👉 查看帮助手册；\n", Bot.Conf.Basic.Key)
			msg += fmt.Sprintf("🔑「%s whatis <declaration> 」\n      👉 C* 语言类型解释；\n", Bot.Conf.Basic.Key)
			msg += fmt.Sprintf("🔑「%s trans[<lang>] <sentence> 」\n      👉 翻译为某种语言；\n", Bot.Conf.Basic.Key)
			msg += fmt.Sprintf("🔑「%s trans-help 」\n      👉 查看有哪些语言可以翻译；\n", Bot.Conf.Basic.Key)
			msg += fmt.Sprintf("🔑「%s mix <emoji>+<emoji> 」\n      👉 合成两个 emoji；\n", Bot.Conf.Basic.Key)
			msg += fmt.Sprintf("你可以使用「%s <func><[<arg>]> <content>」来使用这些功能！\n", Bot.Conf.Basic.Key)
			msg += fmt.Sprintf("例如「%s trans[en] 你好！」或「%s menu」", Bot.Conf.Basic.Key, Bot.Conf.Basic.Key)
			Bot.Manager.Send(OPQBot.SendMsgPack{
				SendToType: OPQBot.SendToTypeGroup,
				ToUserUid:  packet.FromGroupID,
				Content: OPQBot.SendTypeTextMsgContent{
					Content: msg,
				},
			})
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
		case "mix":
			// Mix emoji!
			emojiHelper.MixEmojiInGroup(&Bot, packet, inst)
		}

	} else {
		logrus.Info("A reply.")
		// Pretreatment
		_ = reply
		// Check if the message is a command.
		logrus.Info(fmt.Sprintf("Receive (%s)[ %s ] from [ %d ]", packet.MsgType, packet.Content, packet.FromGroupID))
	}
}
