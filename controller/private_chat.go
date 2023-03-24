package controller

import (
	"fmt"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

func PrivateChatEvents(botId int64, packet *OPQBot.FriendMsgPack) {

	logrus.Info(fmt.Sprintf("Receive [ %s ] from [ %d ]", packet.Content, packet.FromUin))

	// Pretreatment
	inst, err := ParseWithPrefix(packet.Content)
	cmd := inst.Content
	if err != nil {
		return
	}

	// Ping.
	if cmd == "Ping!" {
		Bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeFriend,
			ToUserUid:  packet.FromUin,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "Pong!",
			},
		})
	}
}
