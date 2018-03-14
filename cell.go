package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	lines  []string
	Width  int
	Height int
	Align  AlignType
	pos    int
	table  *Table
}

func NewCell(table *Table, pos int, data interface{}) *Cell {
	s := fmt.Sprintf("%v", data)
	c := &Cell{}
	c.table = table
	c.pos = pos
	c.lines = strings.Split(s, "\n")
	c.Align = LEFT
	c.calcMaxWidth()
	c.calcHeight()
	return c
}

func (c *Cell) Render(index int) string {
	var ret string
	line := c.lines[index]
	width := c.table.WidthOfColumns[c.pos]
	switch c.Align {
	case LEFT:
		ret = Ljust(line, width, " ")
	case RIGHT:
		ret = Rjust(line, width, " ")
	case CENTER:
		fallthrough
	default:
		ret = Center(line, width, " ")
	}
	return " " + ret + " "
}

func (c *Cell) calcHeight() {
	c.Height = len(c.lines)
}

func (c *Cell) calcMaxWidth() {
	var ret int
	for _, line := range c.lines {
		length := len(line)
		if length > ret {
			ret = length
		}
	}
	c.Width = ret
}
