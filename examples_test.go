package bypass_test

import (
	"fmt"
	"time"

	"github.com/go-rod/bypass"
	"github.com/go-rod/rod"
)

func Example_main() {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()

	// You can also use bypass.JS directly without rod
	fmt.Printf("js file size: %d\n\n", len(bypass.JS))

	page := bypass.MustPage(browser)

	page.MustNavigate("https://bot.sannysoft.com")

	printReport(page)

	/*
		Output:

		js file size: 112395

		User Agent (Old): Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36

		WebDriver (New): missing (passed)

		WebDriver Advanced: passed

		Chrome (New): present (passed)

		Permissions (New): denied

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
		fmt.Printf("\t\t%s: %s\n\n", cells[0].MustProperty("textContent"), cells[1].MustProperty("textContent"))
	}

	page.MustScreenshot("")
}
