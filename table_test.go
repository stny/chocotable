package chocotable

func ExampleTable() {
	table := &Table{}
	table.AddRow([]string{"1", "Hello"})
	table.AddRow([]string{"2", "World!"})
	table.AddRow([2]string{"3", "ham egg tomate"})
	table.Println()
	// Output:
	// ┌───┬────────────────┐
	// │ 1 │ Hello          │
	// ├───┼────────────────┤
	// │ 2 │ World!         │
	// ├───┼────────────────┤
	// │ 3 │ ham egg tomate │
	// └───┴────────────────┘
}
