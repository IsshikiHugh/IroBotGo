package codeBin

import (
	"IroBot/utils/crawler"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/chromedp/chromedp"
)

var binPasteUrl string = "http://pastie.org/pastes/create"
var binGetUrl string = "http://pastie.org"
var binDomain string = "pastie.org"

func PasteCode(pl string, code string) (string, error) {
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

func Preview(url string) ([]byte, error) {
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
