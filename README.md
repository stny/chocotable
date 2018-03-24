# chocotable

[![Build Status](https://travis-ci.org/stny/chocotable.svg?branch=master)](https://travis-ci.org/stny/chocotable)
[![codecov](https://codecov.io/gh/stny/chocotable/branch/master/graph/badge.svg)](https://codecov.io/gh/stny/chocotable)
[![Go Report Card](https://goreportcard.com/badge/github.com/stny/chocotable)](https://goreportcard.com/report/github.com/stny/chocotable)

![choco](https://user-images.githubusercontent.com/71077/37500072-b7955fcc-290a-11e8-9f00-cb6394386538.jpg)

## Example

```go
tbl := &chocotable.Table{}
tbl.AddRow([]string{"1", "Hello"})
tbl.AddRow([]string{"2", "World!"})
tbl.AddRow([2]string{"3", "ham egg tomate"})
tbl.AddRow([]string{"4", "アイウエオ"})
tbl.AddRow([]interface{}{5, "Lorem Ipsum"})
tbl.AddRow(struct {
	ID   int
	Name string
}{6, "Golang"})
s := tbl.Render()
// or
tbl.Println()
```

### Output:

```
┌───┬────────────────┐
│ 1 │ Hello          │
├───┼────────────────┤
│ 2 │ World!         │
├───┼────────────────┤
│ 3 │ ham egg tomate │
├───┼────────────────┤
│ 4 │ アイウエオ     │
├───┼────────────────┤
│ 5 │ Lorem Ipsum    │
├───┼────────────────┤
│ 6 │ Golang         │
└───┴────────────────┘
```
