package crawler

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

/*
 * @brief Create a new context for the browser.
 */
func NewCtx() (context.Context, context.CancelFunc) {
	ctx, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true),
			chromedp.WindowSize(640, 4096),
		)...,
	)
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	ctx, cancel := chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf),
	)
	return ctx, cancel
}

/*
 * @brief Create a new context for the browser.
 */
func NewCtxWithSize(width, height int) (context.Context, context.CancelFunc) {
	ctx, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true),
			chromedp.WindowSize(width, height),
		)...,
	)
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	ctx, cancel := chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf),
	)
	return ctx, cancel
}
