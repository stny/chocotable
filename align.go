package chocotable

import (
	"strings"

	runewidth "github.com/mattn/go-runewidth"
)

type AlignType int

const (
	LEFT AlignType = iota
	CENTER
	RIGHT

	ALIGNDEFAULT = CENTER
)

func Rjust(s string, width int, pad string) string {
	w := runewidth.StringWidth(s)
	if w >= width {
		return s
	}
	return strings.Repeat(pad, width-w) + s
}

func Ljust(s string, width int, pad string) string {
	w := runewidth.StringWidth(s)
	if w >= width {
		return s
	}
	return s + strings.Repeat(pad, width-w)
}

func Center(s string, width int, pad string) string {
	w := runewidth.StringWidth(s)
	if w >= width {
		return s
	}
	margin := width - w
	leftMargin := margin/2 + (margin & width & 1)
	return strings.Repeat(pad, leftMargin) + s + strings.Repeat(pad, margin-leftMargin)
}
