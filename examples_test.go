package stealth_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func init() {
	launcher.NewBrowser().MustGet()
}

func Example_main() {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()

	// You can also use stealth.JS directly without rod
	fmt.Printf("js file size: %d\n\n", len(stealth.JS))

	page := stealth.MustPage(browser)

	page.MustNavigate("https://bot.sannysoft.com")

	printReport(page)

	/*
		Output:

		js file size: 161797

		User Agent (Old): true

		WebDriver (New): missing (passed)

		WebDriver Advanced: passed

		Chrome (New): present (passed)

		Permissions (New): prompt

		Plugins Length (Old): 3

		Plugins is of type PluginArray: passed

		Languages (Old): en-US,en

		WebGL Vendor: Intel Inc.

		WebGL Renderer: Intel Iris OpenGL Engine

		Hairline Feature: missing

		Broken Image Dimensions: 16x16
	*/
}

func printReport(page *rod.Page) {
	el := page.MustElement("#broken-image-dimensions.passed")
	for _, row := range el.MustParents("table").First().MustElements("tr:nth-child(n+2)") {
		cells := row.MustElements("td")
		key := cells[0].MustProperty("textContent")
		if strings.HasPrefix(key.String(), "User Agent") {
			fmt.Printf("\t\t%s: %t\n\n", key, !strings.Contains(cells[1].MustProperty("textContent").String(), "HeadlessChrome/"))
		} else {
			fmt.Printf("\t\t%s: %s\n\n", key, cells[1].MustProperty("textContent"))
		}
	}

	page.MustScreenshot("")
}
