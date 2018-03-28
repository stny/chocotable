package chocotable

func ExampleTable() {
	table := &Table{}
	table.AddRow([]string{"1", "Hello"})
	table.AddRow([]string{"2", "World!"})
	table.AddRow([2]string{"3", "ham egg tomate"})
	table.AddRow([]string{"4", "アイウエオ"})
	table.AddRow([]interface{}{5, "Lorem Ipsum"})
	table.AddRow(struct {
		ID   int
		Name string
	}{6, "Golang"})
	table.AddRow(&[2]string{"7", "Gopher"})
	table.AddRow(&struct {
		ID   int
		Name string
	}{8, "zzz"})
	table.Println()
	// Output:
	// ┌───┬────────────────┐
	// │ 1 │ Hello          │
	// ├───┼────────────────┤
	// │ 2 │ World!         │
	// ├───┼────────────────┤
	// │ 3 │ ham egg tomate │
	// ├───┼────────────────┤
	// │ 4 │ アイウエオ     │
	// ├───┼────────────────┤
	// │ 5 │ Lorem Ipsum    │
	// ├───┼────────────────┤
	// │ 6 │ Golang         │
	// ├───┼────────────────┤
	// │ 7 │ Gopher         │
	// ├───┼────────────────┤
	// │ 8 │ zzz            │
	// └───┴────────────────┘
}

func ExampleLineBreak() {
	table := &Table{}
	table.AddRow([]string{"1", "Hello\nWorld"})
	table.Println()
	// Output:
	// ┌───┬───────┐
	// │ 1 │ Hello │
	// │   │ World │
	// └───┴───────┘
}

func ExampleAlignRight() {
	table := &Table{Align: RIGHT}
	table.AddRow([]string{"1", "Hello World"})
	table.AddRow([]string{"2", ":)"})
	table.Println()
	// Output:
	// ┌───┬─────────────┐
	// │ 1 │ Hello World │
	// ├───┼─────────────┤
	// │ 2 │          :) │
	// └───┴─────────────┘
}
