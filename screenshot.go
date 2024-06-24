// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

const WIDTH = int64(1024)
const HEIGHT = int64(768)
const USER_AGENT = "Felipe Mateus - Get Image"

func ScreenshotUrl(url string, dest_path string) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(USER_AGENT),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel = chromedp.NewContext(
		ctx,
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(url, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(dest_path, buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote \"%s\"", dest_path)
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),

		chromedp.EmulateViewport(WIDTH, HEIGHT),
		chromedp.FullScreenshot(res, quality),
	}
}
