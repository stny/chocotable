package main

import (
	"fmt"
	"strings"
)

const (
	HORIZONTAL         = "─"
	VERTICAL           = "│"
	DOWNRIGHT          = "┌"
	DOWNLEFT           = "┐"
	UPLEFT             = "┘"
	UPRIGHT            = "└"
	VERTICALRIGHT      = "├"
	DOWNHORIZONTAL     = "┬"
	VERTICALLEFT       = "┤"
	UPHORIZONTAL       = "┴"
	VERTICALHORIZONTAL = "┼"
)

var (
	TOPBORDERS    = []string{DOWNRIGHT, HORIZONTAL, DOWNHORIZONTAL, DOWNLEFT}
	MIDDLEBORDERS = []string{VERTICALRIGHT, HORIZONTAL, VERTICALHORIZONTAL, VERTICALLEFT}
	BOTTOMBORDERS = []string{UPRIGHT, HORIZONTAL, UPHORIZONTAL, UPLEFT}
)

type Table struct {
	Rows           []*Row
	WidthOfColumns []int
	buf            strings.Builder
}

func (table *Table) AddRow(row interface{}) {
	if row == nil {
		return
	}
	table.Rows = append(table.Rows, NewRow(table, row))
}

func (table *Table) Render() string {
	table.calcMaxWidthOfColumns()

	table.buf.WriteString(table.render_border(TOPBORDERS))

	for i, row := range table.Rows {
		table.buf.WriteString(row.Render())

		if i != len(table.Rows)-1 {
			table.buf.WriteString(table.render_border(MIDDLEBORDERS))
		}
	}
	table.buf.WriteString(table.render_border(BOTTOMBORDERS))

	return table.buf.String()
}

func (table *Table) Reset() {
	table.Rows = nil
	table.buf.Reset()
}

func (table *Table) Println() {
	fmt.Println(table.Render())
}

func (table *Table) calcMaxWidthOfColumns() {
	table.WidthOfColumns = make([]int, len(table.Rows[0].Cells))
	for _, row := range table.Rows {
		for i, w := range row.WidthOfColumns {
			if w > table.WidthOfColumns[i] {
				table.WidthOfColumns[i] = w
			}
		}
	}
}

func (table *Table) render_border(style []string) string {
	s := style[0]
	for i, n := range table.WidthOfColumns {
		s = s + strings.Repeat(style[1], n+2)
		if i != len(table.WidthOfColumns)-1 {
			s = s + style[2]
		}
	}
	s = s + style[3] + "\n"
	return s
}
