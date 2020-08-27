//go:generate go run ./generate

package bypass

import (
	"github.com/go-rod/rod"
)

// PageE creates a stealth page that can't be detected as bot.
func PageE(b *rod.Browser) (*rod.Page, error) {
	p, err := b.Page("")
	if err != nil {
		return nil, err
	}

	_, err = p.EvalOnNewDocument(JS)
	if err != nil {
		return nil, err
	}

	err = p.SetUserAgent(nil)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Page creates a stealth page that can't be detected as bot.
func Page(b *rod.Browser) *rod.Page {
	p, err := PageE(b)
	if err != nil {
		panic(err)
	}
	return p
}
