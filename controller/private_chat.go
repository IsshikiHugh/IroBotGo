package controller

import (
	"IroBot/utils"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

func (bot *BotEnvironment) PrivateChatEvents(botId int64, packet *OPQBot.FriendMsgPack) {
	var cmd string
	if !utils.IsIroCommand(packet.Content) {
		return
	} else {
		cmd, _ = utils.DecodeIroCommand(packet.Content)
	}

	// Ping.
	if cmd == "Ping!" {
		logrus.Info("Receive [ %s ] from [ %d ]", cmd, packet.FromUin)
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeFriend,
			ToUserUid:  packet.FromUin,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "Pong!",
			},
		})
	}
}
