package controller

import (
	"IroBot/utils"
	codeBin "IroBot/utils/code_bin"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/mcoo/OPQBot"
	"github.com/mcoo/requests"
	"github.com/sirupsen/logrus"
)

func (bot *BotEnvironment) GroupChatEvents(botQQ int64, packet *OPQBot.GroupMsgPack) {
	logrus.Info(fmt.Sprintf("Receive (%s)[ %s ] from [ %d ]", packet.MsgType, packet.Content, packet.FromGroupID))

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

		// Say
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

		// 彩虹屁
		if strings.HasPrefix(cmd, "彩虹屁") {
			msg := ""
			if res, err := requests.Get("https://api.shadiao.pro/chp"); err != nil {
				msg = "😖 放不出来了"
			} else {
				chp := struct {
					Data struct {
						Type string `json:"type"`
						Text string `json:"text"`
					} `json:"data"`
				}{}
				res.Json(&chp)
				msg = chp.Data.Text
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
				cmd = strings.TrimPrefix(cmd, "]\n")
			}
			logrus.Info(fmt.Sprintf("Try to past code in [ %s ]!", pl))
			code := cmd
			url, err := codeBin.PasteCode(pl, code)
			if err != nil {
				bot.Manager.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeGroup,
					ToUserUid:  packet.FromGroupID,
					Content: OPQBot.SendTypeTextMsgContent{
						Content: "😫 粘贴代码时发生错误！",
					},
				})
				return
			}
			img, err := codeBin.Preview(url)
			if err != nil {
				bot.Manager.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeGroup,
					ToUserUid:  packet.FromGroupID,
					Content: OPQBot.SendTypeTextMsgContent{
						Content: "📋 " + url + "\n😫 预览生成错误。",
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
	}
}
