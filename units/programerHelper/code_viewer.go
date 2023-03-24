package programerHelper

import (
	"IroBot/model"
	"IroBot/utils/crawler"
	"context"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

func CodeViewerInGroup(bot *model.BotEnvironment, packet *OPQBot.GroupMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromGroupID,
	}
	solveCodeViewer(bot, &bd, &inst)
}

func CodeViewerInChat(bot *model.BotEnvironment, packet *OPQBot.FriendMsgPack, inst model.Instruction) {
	bd := model.BotData{
		SendToType: OPQBot.SendToTypeGroup,
		TargetId:   packet.FromUin,
	}
	solveCodeViewer(bot, &bd, &inst)
}

/*
 * @brief Main functional module.
 */
func solveCodeViewer(bot *model.BotEnvironment, bd *model.BotData, inst *model.Instruction) {
	code := inst.Content
	pl := "plaintext"
	if inst.HasArg {
		pl = inst.Args
	}

	url, err := codeBin(pl, code)
	if err != nil {
		logrus.Error("Error happens when paste code: ", err)
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: bd.SendToType,
			ToUserUid:  bd.TargetId,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "😖 粘贴代码时发生错误！",
			},
		})
		return
	}
	img, err := preview(url)
	if err != nil {
		logrus.Error("Error happens when preview code: ", err)
		bot.Manager.Send(OPQBot.SendMsgPack{
			SendToType: bd.SendToType,
			ToUserUid:  bd.TargetId,
			Content: OPQBot.SendTypeTextMsgContent{
				Content: "📋 " + url + "\n😖 预览生成错误。",
			},
		})
		return
	}
	bot.Manager.Send(OPQBot.SendMsgPack{
		SendToType: bd.SendToType,
		ToUserUid:  bd.TargetId,
		Content: OPQBot.SendTypePicMsgByBase64Content{
			Content: "📋 " + url,
			Base64:  base64.StdEncoding.EncodeToString(img),
		},
	})
	return
}

/* Functional Part */

var binPasteUrl string = "http://pastie.org/pastes/create"
var binGetUrl string = "http://pastie.org"
var binDomain string = "pastie.org"

func codeBin(pl string, code string) (string, error) {
	forms := url.Values{}
	forms.Add("language", pl)
	forms.Add("content", code)

	// Ban redirect, so that we can get the target url.
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.PostForm(binPasteUrl, forms)

	body, _ := ioutil.ReadAll(resp.Body)
	// logrus.Info(string(body))
	suffix := strings.TrimLeft(string(body), "Found. Redirecting to")
	if err != nil {
		return "", err
	}
	if suffix == "/" {
		return "", errors.New("invalid redirection")
	}
	return binGetUrl + suffix, nil
}

func preview(url string) ([]byte, error) {
	var buf []byte
	ctx, cancel := crawler.NewCtxWithSize(960, 4800)
	defer cancel()

	err := chromedp.Run(ctx, chromedp.Tasks{
		crawler.SetCookies(binDomain, "theme", "nord.css"), // Set theme
		chromedp.Navigate(url),
		chromedp.WaitVisible(`document.querySelector("section.code")`, chromedp.ByJSPath),
		getImg(&buf),
	})
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func getImg(buf *[]byte) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		if err = chromedp.Screenshot(`document.querySelector("section.code")`, buf, chromedp.ByJSPath).Do(ctx); err != nil {
			return
		}
		return
	}
}
