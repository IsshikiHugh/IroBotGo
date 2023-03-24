package controller

import (
	"IroBot/model"
	"IroBot/units/emojiHelper"
	"IroBot/units/languageHelper"
	"IroBot/units/programerHelper"
	"fmt"
	"strings"

	"github.com/mcoo/OPQBot"
)

func GroupChatEvents(botQQ int64, packet *OPQBot.GroupMsgPack) {
	// Ignore message if the sender is the bot.
	if packet.FromUserID == Bot.Conf.Basic.Qid {
		return
	}

	if true {
		// Pretreatment
		var (
			inst model.Instruction
			err  error
		)

		if packet.MsgType == "AtMsg" {
			atMsg, _ := OPQBot.ParserGroupAtMsg(*packet)
			inst, err = ParseAtMsg(atMsg.Content)

		} else {
			inst, err = ParseWithPrefix(packet.Content)
		}
		if err != nil {
			return
		}
		// logrus.Info(fmt.Sprintf("Receive (%s)[ %s ] from [ %d ]", packet.MsgType, packet.Content, packet.FromGroupID))

		// Choose the option
		switch inst.OptionName {
		case "menu":
			msg := fmt.Sprintf("🥰 你好！我是 IroBot！目前我支持这些功能！\n")
			msg += fmt.Sprintf("🔑「menu 」\n      👉 查看帮助手册；\n")
			msg += fmt.Sprintf("🔑「whatis <declaration> 」\n      👉 C* 语言类型解释；\n")
			msg += fmt.Sprintf("🔑「trans[<lang>] <sentence> 」\n      👉 翻译为某种语言；\n")
			msg += fmt.Sprintf("🔑「trans-help 」\n      👉 查看有哪些语言可以翻译；\n")
			msg += fmt.Sprintf("🔑「mix <emoji>+<emoji> 」\n      👉 合成两个 emoji；\n")
			msg += fmt.Sprintf("你可以使用「@我/%s <func><[<arg>]> <content>」来使用这些功能！\n", Bot.Conf.Basic.Key)
			msg += fmt.Sprintf("例如「%s trans[en] 你好！」或「@IroBot menu」", Bot.Conf.Basic.Key)

			Bot.Manager.Send(OPQBot.SendMsgPack{
				SendToType: OPQBot.SendToTypeGroup,
				ToUserUid:  packet.FromGroupID,
				Content: OPQBot.SendTypeTextMsgContent{
					Content: msg,
				},
			})
		case "say":
			// Test only!
			// TODO: make it simple
			msg := strings.TrimSpace(strings.TrimPrefix(inst.Content, "say "))
			if packet.FromUserID != Bot.Conf.Basic.MQid {
				return
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

	}
}
