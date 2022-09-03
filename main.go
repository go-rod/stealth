//go:generate go run ./generate

// Package stealth is a package for anti-bot-detection with rod
package stealth

import (
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
