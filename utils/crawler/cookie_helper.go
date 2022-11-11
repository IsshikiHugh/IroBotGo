package crawler

import (
	"context"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// Set cookies as a task action
func SetCookies(domain, name, value string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
		err := network.SetCookie(name, value).
			WithExpires(&expr).
			WithDomain(domain).
			WithHTTPOnly(true).
			Do(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
