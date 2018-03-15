package chocotable

import (
	"reflect"
	"strings"
)

type Row struct {
	Cells          []*Cell
	WidthOfColumns []int
	Height         int
}

func NewRow(table *Table, row interface{}) *Row {
	r := &Row{}
	r.AddCells(table, row)
	r.calcMaxHeight()
	r.calcWidthOfColumns()
	return r
}

func (r *Row) AddCells(table *Table, row interface{}) {
	cellParts := extract(row)
	r.Cells = make([]*Cell, len(cellParts))
	for i, c := range cellParts {
		r.Cells[i] = NewCell(table, i, c)
	}
}

func (r *Row) Render() string {
	var s strings.Builder
	for i := 0; i < r.Height; i++ {
		s.WriteString(VERTICAL)
		for _, cell := range r.Cells {
			s.WriteString(cell.Render(i))
			s.WriteString(VERTICAL)
		}
		s.WriteString("\n")
	}
	return s.String()
}

func extract(row interface{}) []interface{} {
	v := reflect.ValueOf(row)
	kind := v.Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		ret := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			ret[i] = v.Index(i).Interface()
		}
		return ret
	default:
		panic("Not Implemented yet")
	}
	return []interface{}{1, 1}
}

func (r *Row) calcWidthOfColumns() {
	ret := make([]int, len(r.Cells))
	for i, cell := range r.Cells {
		ret[i] = cell.Width
	}
	r.WidthOfColumns = ret
}

func (r *Row) calcMaxHeight() {
	var ret int
	for _, cell := range r.Cells {
		if cell.Height > ret {
			ret = cell.Height
		}
	}
	r.Height = ret
}
