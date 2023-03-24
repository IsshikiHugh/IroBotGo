package languageHelper

import (
	"IroBot/model"
	"os/exec"
	"strings"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

func TranslatorInGroup(bot *model.BotEnvironment, packet *OPQBot.GroupMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromGroupID,
	}
	solveTranslator(bot, &bd, &inst)
}

func TranslatorInChat(bot *model.BotEnvironment, packet *OPQBot.FriendMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromUin,
	}
	solveTranslator(bot, &bd, &inst)
}

func TranslatorHelpInGroup(bot *model.BotEnvironment, packet *OPQBot.GroupMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromGroupID,
	}
	solveTranslatorHelp(bot, &bd, &inst)
}

func TranslatorHelpInChat(bot *model.BotEnvironment, packet *OPQBot.FriendMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromUin,
	}
	solveTranslatorHelp(bot, &bd, &inst)
}

/*
 * @brief Main functional module.
 */
func solveTranslator(bot *model.BotEnvironment, bd *model.BotData, inst *model.Instruction) {
	srcText := inst.Content
	resLang := "english"
	if inst.HasArg {
		resLang = inst.Args
	}
	logrus.Info(inst)

	res, err := trans(srcText, resLang)
	if err != nil {
		logrus.Error("Error happens when translate text: ", err)
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeGroup,
			ToUserUid:  bd.TargetId,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "😖 一时语塞。",
			},
		})
	} else {
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeGroup,
			ToUserUid:  bd.TargetId,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: res,
			},
		})
	}
	return
}

/*
 * @brief Main functional module.
 */
func solveTranslatorHelp(bot *model.BotEnvironment, bd *model.BotData, inst *model.Instruction) {
	ret := listLang()
	ret = strings.ReplaceAll(ret, "\n", ", ")
	if ret == "" {
		logrus.Error("Error happens when getting language supported in translator")
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeGroup,
			ToUserUid:  bd.TargetId,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "😖 一时语塞。",
			},
		})
	} else {
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: OPQBot.SendToTypeGroup,
			ToUserUid:  bd.TargetId,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: ret,
			},
		})
	}
	return
}

/* Functional Part */

/*
 * @brief Translate the sentence into the taregt language.
 * 		  Based on the command line app 'trans'.
 * @param srcSentence: the sentence to be translated.
 * @param resultLang: the target language.
 */
func trans(srcSentence string, resultLang string) (string, error) {
	// Abbreviation.
	switch resultLang {
	case "zh":
		resultLang = "Chinese"
	case "en":
		resultLang = "English"
	case "jp":
		resultLang = "Japanese"
	}

	out, err := exec.Command("trans", "-t", resultLang, "\""+srcSentence+"\"", "-b").Output()
	if err != nil {
		return "", err
	} else {
		ret := format2text(string(out))
		ret = strings.TrimSuffix(ret, "\n")
		return ret, nil
	}
}

/*
 * @brief Remove the format sign in shell.
 * @param formatStr: the original message.
 */
func format2text(formatStr string) string {
	formatStr = strings.ReplaceAll(formatStr, "[1m", "")
	formatStr = strings.ReplaceAll(formatStr, "[4m", "")
	formatStr = strings.ReplaceAll(formatStr, "[22m", "")
	formatStr = strings.ReplaceAll(formatStr, "[24m", "")
	return formatStr
}

/*
 * @brief: List the supported language.
 */
func listLang() string {
	out, err := exec.Command("trans", "-list-languages-english").Output()
	logrus.Info(string(out))
	if err != nil {
		return ""
	} else {
		return string(out)
	}
}
