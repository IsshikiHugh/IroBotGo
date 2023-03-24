package emojiHelper

import (
	"IroBot/model"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
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
	args := strings.ReplaceAll(inst.Content, " ", "")
	if len(args) != 9 {
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: bd.SendToType,
			ToUserUid:  bd.TargetId,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "🤔 请使用 <emoji>+<emoji> 的形式作为输入。",
			},
		})
		return
	}

	var emojis []rune
	for _, v := range args {
		emojis = append(emojis, v)
	}

	img, err := mixEmoji(emojis)
	if err != nil && err.Error() == "bad combination" {
		logrus.Error("Error happens when mix emoji, the combination is invalid")
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: bd.SendToType,
			ToUserUid:  bd.TargetId,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "🤯 这个组合好像不太行！",
			},
		})
		return
	} else if err != nil {
		logrus.Error("Error happens when mix emoji:", err)
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: bd.SendToType,
			ToUserUid:  bd.TargetId,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "😖 一时语塞。",
			},
		})
		return
	}
	bot.Manager.Send(OPQBot.SendMsgPack{
		SendToType: bd.SendToType,
		ToUserUid:  bd.TargetId,
		Content: OPQBot.SendTypePicMsgByBase64Content{
			Base64: base64.StdEncoding.EncodeToString(img),
		},
	})
}

/* Functional Part */

var baseURL string = "https://www.gstatic.com/android/keyboard/emojikitchen/20201001/" // "u1fxxx/u1fxxx_u1fyyy.png"

/*
 * @brief mix two emojis
 * @param emojis: a slice of two emojis: emojis[0] and emojis[2]
 */
func mixEmoji(emojis []rune) ([]byte, error) {
	logrus.Info(fmt.Sprintf("Try to mix %s and %s.", string(emojis[0]), string(emojis[2])))
	// Translate rune to utf8 string.
	codeX := fmt.Sprintf("u%x", emojis[0])
	codeY := fmt.Sprintf("u%x", emojis[2])

	url := baseURL + codeX + "/" + codeX + "_" + codeY + ".png"

	// Send request.
	logrus.Info(url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		url = baseURL + codeY + "/" + codeY + "_" + codeX + ".png"
		if err != nil {
			return nil, err
		}

		// Send request again.
		logrus.Info("try again: " + url)
		resp, err = http.Get(url)

		if err != nil {
			return nil, err
		}

		if resp.StatusCode == 404 {
			return nil, errors.New("bad combination")
		}
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, err
}
