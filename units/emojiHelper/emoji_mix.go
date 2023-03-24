package emojiHelper

import (
	"IroBot/model"

	"github.com/mcoo/OPQBot"
)

func MixEmojiInGroup(bot *model.BotEnvironment, packet *OPQBot.GroupMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromGroupID,
	}
	solveMixEmoji(bot, &bd, &inst)
}

func MixEmojiInChat(bot *model.BotEnvironment, packet *OPQBot.FriendMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromUin,
	}
	solveMixEmoji(bot, &bd, &inst)
}

func solveMixEmoji(bot *model.BotEnvironment, bd *model.BotData, inst *model.Instruction) {

}
