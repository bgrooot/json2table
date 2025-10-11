package model

import "fmt"

type Table struct {
	Row []Row
}

type Row struct {
	Cell []Cell
}

type Cell struct {
	Value   string
	Type    string
	RowSpan int
	ColSpan int
}

func (t *Table) AddRow() {
	t.Row = append(t.Row, Row{})
}

func (t *Table) LastRow() *Row {
	return &t.Row[len(t.Row)-1]
}

func (t *Table) AddKeyToLastRow(node Node) *Cell {
	cell := Cell{Value: node.Key, Type: "key"}
	return t.AddCellToLastRow(&cell)
}

func (t *Table) AddValueToLastRow(value string) *Cell {
	cell := Cell{Value: value, Type: "value"}
	return t.AddCellToLastRow(&cell)
}

func (t *Table) AddCellToLastRow(cell *Cell) *Cell {
	t.LastRow().Cell = append(t.LastRow().Cell, *cell)
	return &t.LastRow().Cell[len(t.LastRow().Cell)-1]
}

func (t *Table) Print() {
	for _, r := range t.Row {
		for _, c := range r.Cell {
			fmt.Printf("[Type: %s, Value: %s]", c.Type, c.Value)
		}
		fmt.Println()
	}
}
