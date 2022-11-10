package controller

import (
	"IroBot/utils"

	"github.com/mcoo/OPQBot"
)

func (bot *BotEnvironment) PrivateChatEvents(botId int64, packet *OPQBot.FriendMsgPack) {
	// If not a command.
	if packet.MsgType != "TextMsg" || !utils.IsIroCommand(packet.Content) {
		return
	}

	// Ping.
	if packet.Content == "/Iro Ping?" {
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeFriend,
			ToUserUid:  packet.FromUin,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "Pong!",
			},
		})
	}
}
