package controller

import (
	"IroBot/utils"
	"fmt"
	"strings"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

func (bot *BotEnvironment) GroupChatEvents(botQQ int64, packet *OPQBot.GroupMsgPack) {
	logrus.Info(fmt.Sprintf("Receive [ %s ] from [ %d ]", packet.Content, packet.FromGroupID))

	// Ignore message if the sender is the bot.
	if packet.FromUserID == bot.Conf.Basic.Qid {
		return
	}
	// Check if the message is a command.
	var cmd string
	if !utils.IsIroCommand(packet.Content) {
		return
	} else {
		cmd, _ = utils.DecodeIroCommand(packet.Content)
	}

	// Echo
	if strings.HasPrefix(cmd, "Echo") {
		msg := strings.TrimSpace(strings.TrimPrefix(cmd, "Echo"))
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeGroup,
			ToUserUid:  packet.FromGroupID,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: msg,
			},
		})
		return
	}

	// Echo
	if strings.HasPrefix(cmd, "Say") {
		msg := strings.TrimSpace(strings.TrimPrefix(cmd, "Say"))
		if packet.FromUserID != bot.Conf.Basic.MQid {
			msg = "可恶的 「" + packet.FromNickName + "」 强迫可怜的 IroBot 说 「" + msg + "」"
		} else {
			msg = "🥰 IroBot 也想说 「" + msg + "」"
		}
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeGroup,
			ToUserUid:  packet.FromGroupID,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: msg,
			},
		})
		return
	}

	// 色色
	if strings.HasPrefix(cmd, "色色！") {
		msg := "😠 不许色色！"

		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeGroup,
			ToUserUid:  packet.FromGroupID,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: msg,
			},
		})
		return
	}
}
