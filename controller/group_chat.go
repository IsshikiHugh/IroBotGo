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
}
