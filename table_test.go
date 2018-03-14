package main

func ExampleTable() {
	table := &Table{}
	table.AddRow([]string{"1", "Hello"})
	table.AddRow([]string{"2", "World!"})
	table.Println()
	// Output:
	// ┌───┬────────┐
	// │ 1 │ Hello  │
	// ├───┼────────┤
	// │ 2 │ World! │
	// └───┴────────┘
}
