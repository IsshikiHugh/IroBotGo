package programerHelper

import (
	"IroBot/model"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

func CTypeExplainerInGroup(bot *model.BotEnvironment, packet *OPQBot.GroupMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromGroupID,
	}
	solveCTypeExplainer(bot, &bd, &inst)
}

func CTypeExplainerInChat(bot *model.BotEnvironment, packet *OPQBot.FriendMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromUin,
	}
	solveCTypeExplainer(bot, &bd, &inst)
}

/*
 * @brief Main functional module.
 */
func solveCTypeExplainer(bot *model.BotEnvironment, bd *model.BotData, inst *model.Instruction) {
	state := inst.Content
	logrus.Info(inst)

	msg, err := explain(state)
	if err != nil {
		if err.Error() != "invalid syntax" {
			logrus.Error("Error happens when explain C type: ", err)
			msg = "😖 一时语塞。"
		} else {
			msg = "🤔 看起来这句话并不合法。"
		}
	} else {
		msg = "💡「" + state + "」" + msg
	}
	bot.Manager.Send(OPQBot.SendMsgPack{
		SendToType: OPQBot.SendToTypeGroup,
		ToUserUid:  bd.TargetId,
		Content: OPQBot.SendTypeTextMsgContent{
			Content: msg,
		},
	})
	return
}

/* Functional Part */

var explainerUrl string = "https://xwd733f66f.execute-api.us-west-1.amazonaws.com/prod/cdecl_backend?q="

/*
 * @brief Explain the C type.
 * @param sentence: The sentence to explain.
 */
func explain(sentence string) (string, error) {
	// Replace space with %20.
	sentence = strings.Replace(sentence, " ", "%20", -1)

	// logrus.Info(sentence)
	resp, err := http.Get(explainerUrl + sentence)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ret := string(body)
	if ret == "\"syntax error\"" {
		return "", errors.New("invalid syntax")
	} else if strings.HasPrefix(ret, "<html>") {
		// 502 Bad Gateway
		return "", errors.New("502 bad gateway")
	} else if strings.HasPrefix(ret, "\"declare") {
		ret = strings.TrimPrefix(ret, "\"")
		ret = strings.TrimSuffix(ret, "\"")
		return ret, nil
	}
	return "", errors.New("receive unaccepted result")
}
