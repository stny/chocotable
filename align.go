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
	if runewidth.StringWidth(s) >= width {
		return s
	}
	return strings.Repeat(pad, width-runewidth.StringWidth(s)) + s
}

func Ljust(s string, width int, pad string) string {
	if runewidth.StringWidth(s) >= width {
		return s
	}
	return s + strings.Repeat(pad, width-runewidth.StringWidth(s))
}

func Center(s string, width int, pad string) string {
	if runewidth.StringWidth(s) >= width {
		return s
	}
	margin := width - runewidth.StringWidth(s)
	leftMargin := margin/2 + (margin & width & 1)
	return strings.Repeat(pad, leftMargin) + s + strings.Repeat(pad, margin-leftMargin)
}
