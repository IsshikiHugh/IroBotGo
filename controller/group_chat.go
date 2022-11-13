package controller

import (
	"IroBot/utils"
	cTypeExplainer "IroBot/utils/c_type_explainer"
	codeBin "IroBot/utils/code_bin"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

func (bot *BotEnvironment) GroupChatEvents(botQQ int64, packet *OPQBot.GroupMsgPack) {

	// Ignore message if the sender is the bot.
	if packet.FromUserID == bot.Conf.Basic.Qid {
		return
	}

	reply, err := OPQBot.ParserGroupReplyMsg(*packet)
	// That is, not a reply.
	if err != nil {
		logrus.Info("Not a reply.")
		// Pretreatment
		// Check if the message is a command.
		var cmd string
		if !utils.IsIroCommand(packet.Content) {
			return
		} else {
			cmd, _ = utils.DecodeIroCommand(packet.Content)
		}

		logrus.Info(fmt.Sprintf("Receive (%s)[ %s ] from [ %d ]", packet.MsgType, packet.Content, packet.FromGroupID))

		// Say
		if strings.HasPrefix(cmd, "say") {
			msg := strings.TrimSpace(strings.TrimPrefix(cmd, "say"))
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

		// 代码剪贴板
		if strings.HasPrefix(cmd, "omg") {
			cmd = strings.TrimPrefix(cmd, "omg")
			pl := "plaintext"
			if strings.HasPrefix(cmd, "[") {
				cmd = strings.TrimPrefix(cmd, "[")
				argvs := strings.Split(cmd, "]")
				if len(argvs) >= 1 && len(argvs[0]) > 0 {
					pl = argvs[0]
				}
				cmd = strings.TrimPrefix(cmd, pl)
				cmd = strings.TrimPrefix(cmd, "]")
				cmd = strings.TrimPrefix(cmd, "\n")
			}
			logrus.Info(fmt.Sprintf("Try to past code in [ %s ]!", pl))
			code := cmd
			url, err := codeBin.PasteCode(pl, code)
			if err != nil {
				logrus.Error("Error happens when paste code: ", err)
				bot.Manager.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeGroup,
					ToUserUid:  packet.FromGroupID,
					Content: OPQBot.SendTypeTextMsgContent{
						Content: "😖 粘贴代码时发生错误！",
					},
				})
				return
			}
			img, err := codeBin.Preview(url)
			if err != nil {
				logrus.Error("Error happens when preview code: ", err)
				bot.Manager.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeGroup,
					ToUserUid:  packet.FromGroupID,
					Content: OPQBot.SendTypeTextMsgContent{
						Content: "📋 " + url + "\n😖 预览生成错误。",
					},
				})
				return
			}
			bot.Manager.Send(OPQBot.SendMsgPack{
				SendToType: OPQBot.SendToTypeGroup,
				ToUserUid:  packet.FromGroupID,
				Content: OPQBot.SendTypePicMsgByBase64Content{
					Content: "📋 " + url,
					Base64:  base64.StdEncoding.EncodeToString(img),
				},
			})
			return
		}

		// C 类型解释
		if strings.HasPrefix(cmd, "whatis ") {
			cmd = strings.TrimPrefix(cmd, "whatis ")
			msg, err := cTypeExplainer.Explain(cmd)
			if err != nil {
				if err.Error() != "invalid syntax" {
					logrus.Error("Error happens when explain sentence: ", err)
					msg = "😖 一时语塞。"
				} else {
					msg = "🤔 看起来这句话并不合法。"
				}
			} else {
				msg = "「" + cmd + "」" + msg
			}
			bot.Manager.Send(OPQBot.SendMsgPack{
				SendToType: OPQBot.SendToTypeGroup,
				ToUserUid:  packet.FromGroupID,
				Content: OPQBot.SendTypeTextMsgContent{
					Content: msg,
				},
			})
		}
	} else {
		logrus.Info("A reply.")
		// Pretreatment
		// Check if the message is a command.
		var cmd string
		if !utils.IsIroCommand(reply.Content) {
			return
		} else {
			cmd, _ = utils.DecodeIroCommand(reply.Content)
		}
		_ = cmd
		logrus.Info(fmt.Sprintf("Receive (%s)[ %s ] from [ %d ]", packet.MsgType, packet.Content, packet.FromGroupID))
	}
}
