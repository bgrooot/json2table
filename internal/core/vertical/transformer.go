package vertical

import (
	"fmt"

	"github.com/bgrooot/json2table/internal/model"
	"github.com/bgrooot/json2table/pkg/json2table/config"
)

type Transformer struct{}

func (t *Transformer) Transform(nodeArr []model.Node, config config.Config) []model.Table {
	var tableArr []model.Table

	for _, node := range nodeArr {
		table := model.Table{}
		table.AddRow()

		t.traverse(&node, &table)

		tableArr = append(tableArr, table)
	}

	return tableArr
}

func (t *Transformer) traverse(node *model.Node, table *model.Table) {
	if node.Key != "" {
		cell := table.AddKeyToLastRow(*node)
		rowSpan, exist := node.Metadata["childCount"]
		if exist {
			cell.RowSpan = rowSpan.(int)
		}
	}

	for idx, val := range node.Value {
		if idx > 0 {
			table.AddRow()
		}

		if nd, ok := val.(*model.Node); ok {
			t.traverse(nd, table)
		} else {
			cell := table.AddValueToLastRow(fmt.Sprintf("%v", val))
			colSpan, exist := node.Metadata["relativeDepth"]
			if exist {
				cell.ColSpan = colSpan.(int)
			}
		}
	}
}
