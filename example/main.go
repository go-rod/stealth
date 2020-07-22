package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"io/ioutil"
)

func main() {
	js, err := fetchBypassJS()
	if err != nil {
		panic(err)
	}

	if len(js) == 0 {
		panic("js empty")
	}

	l := launcher.New().Launch()

	b := rod.New().ControlURL(l).Connect()
	defer b.Close()

	p := b.Page("")

	_, err = proto.PageAddScriptToEvaluateOnNewDocument{
		Source: js,
	}.Call(p)
	if err != nil {
		panic(err)
	}

	// p.Emulate(devices.Pixel2XL)
	// puppeteer:protocol SEND ► {"sessionId":"AEEF67CFDF10EEF9A0013C44E73EF2C1","method":"Network.setUserAgentOverride","params":{"userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3882.0 Safari/537.36","acceptLanguage":"en-US,en","platform":"Win32"},"id":27}
	err = proto.NetworkSetUserAgentOverride{
		UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3882.0 Safari/537.36",
		AcceptLanguage: "en-US,en",
		Platform:       "Win32",
	}.Call(p)
	if err != nil {
		panic(err)
	}

	p.Navigate("https://bot.sannysoft.com")

	p.WaitRequestIdle()()

	p.ScreenshotFullPage("testresult.png")

	fmt.Println("All done, check the screenshot. ✨")

}

func fetchBypassJS() (js string, err error) {
	//resp, err := http.Get("https://cdn.jsdelivr.net/gh/go-rod/bypass@latest/puppeteer-dump/dist/stealth.js")
	//if err != nil {
	//	return
	//}
	//
	//defer resp.Body.Close()
	//b, err := ioutil.ReadAll(resp.Body)

	b, err := ioutil.ReadFile("puppeteer-dump/dist/stealth.js")
	if err != nil {
		return
	}

	js = string(b)
	return
}
