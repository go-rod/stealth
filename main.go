//go:generate go run ./generate

package bypass

import (
	"github.com/go-rod/rod"
)

// Page creates a stealth page that can't be detected as bot.
func Page(b *rod.Browser) (*rod.Page, error) {
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

// MustPage creates a stealth page that can't be detected as bot.
func MustPage(b *rod.Browser) *rod.Page {
	p, err := Page(b)
	if err != nil {
		panic(err)
	}
	return p
}
