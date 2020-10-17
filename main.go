//go:generate go run ./generate

package bypass

import (
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// Page creates a stealth page that can't be detected as bot.
func Page(b *rod.Browser) (*rod.Page, error) {
	p, err := b.Page(proto.TargetCreateTarget{})
	if err != nil {
		return nil, err
	}

	_, err = p.EvalOnNewDocument(JS)
	if err != nil {
		return nil, err
	}

	ua, err := NormalizeUA(b)
	if err != nil {
		return nil, err
	}

	err = p.SetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: ua,
	})
	if err != nil {
		return nil, err
	}

	return p, nil
}

// MustPage creates a stealth page that can't be detected as bot.
func MustPage(b *rod.Browser) *rod.Page {
	p, err := Page(b)
	if err != nil {
		panic(err)
	}
	return p
}

// NormalizeUA normalize the default user-agent, as a head mode browser.
func NormalizeUA(b *rod.Browser) (string, error) {
	v, err := proto.BrowserGetVersion{}.Call(b)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(v.UserAgent, "HeadlessChrome/", "Chrome/"), nil
}
