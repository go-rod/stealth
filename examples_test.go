package bypass_test

import (
	"fmt"

	"github.com/go-rod/bypass"
	"github.com/go-rod/rod"
)

func Example_main() {
	browser := rod.New().Connect()
	defer browser.Close()

	// You can also use bypass.JS directly without rod
	fmt.Println(len(bypass.JS))

	page := bypass.Page(browser)

	page.Navigate("https://bot.sannysoft.com").WaitLoad()

	fmt.Println(page.Element("table").Text())

	// Output:
	// 87552
	// Test Name	Result
	// User Agent
	// (Old)
	// 	Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36
	// WebDriver
	// (New)
	// 	missing (passed)
	// WebDriver Advanced	passed
	// Chrome
	// (New)
	// 	present (passed)
	// Permissions
	// (New)
	// 	denied
	// Plugins Length
	// (Old)
	// 	3
	// Plugins is of type PluginArray	passed
	// Languages
	// (Old)
	// 	en-US,en
	// WebGL Vendor	Intel Inc.
	// WebGL Renderer	Intel Iris OpenGL Engine
	// Hairline Feature	missing
	// Broken Image Dimensions	16x16
}
