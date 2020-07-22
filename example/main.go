package main

import (
	"fmt"
	"github.com/go-rod/bypass"
	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().Connect()
	defer browser.Close()

	page := bypass.Page(browser)

	page.Navigate("https://bot.sannysoft.com").WaitLoad()

	page.ScreenshotFullPage("testresult.png")

	fmt.Println("All done, check the screenshot. ✨")
}
