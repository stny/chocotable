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
