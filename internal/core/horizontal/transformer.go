package horizontal

import (
	"fmt"

	"github.com/bgrooot/json2table/internal/model"
	"github.com/bgrooot/json2table/internal/util"
	"github.com/bgrooot/json2table/pkg/json2table/config"
)

type Transformer struct{}

func (t *Transformer) Transform(nodeArr []model.Node, config config.Config) []model.Table {
	var tableArr []model.Table

	for _, node := range nodeArr {
		table := model.Table{}
		t.traverse(&node, &table)

		tableArr = append(tableArr, table)
	}

	return tableArr
}

func (t *Transformer) traverse(node *model.Node, table *model.Table) {
	queue := []*model.Node{node}

	for len(queue) > 0 {
		breadth := len(queue)
		for i := 0; i < breadth; i++ {
			node := queue[0]
			queue = queue[1:]

			addKey(node, table)
			addValue(node, table)
			enqueueChildren(node, &queue)
		}

		if len(queue) > 0 {
			table.AddRow()
		}
	}
}

func addKey(node *model.Node, table *model.Table) {
	if node.Key != "" {
		cell := table.AddKeyToLastRow(*node)
		colSpan, exist := node.Metadata["childCount"]
		if exist {
			cell.ColSpan = colSpan.(int)
		}
	}
}

func addValue(node *model.Node, table *model.Table) {
	if valueOnly, ok := node.Metadata["valueOnly"].(bool); ok && valueOnly {
		cell := table.AddValueToLastRow(fmt.Sprintf("%v", node.Value[0]))
		rowSpan, exist := node.Metadata["relativeDepth"]
		if exist {
			cell.RowSpan = rowSpan.(int)
		}
	}
}

func enqueueChildren(node *model.Node, queue *[]*model.Node) {
	for _, value := range node.Value {
		if nd, ok := value.(*model.Node); ok {
			*queue = append(*queue, nd)
		} else {
			if _, ok := node.Metadata["valueOnly"].(bool); !ok {
				valueOnly := model.Node{Value: []any{value}, Metadata: util.DeepCopy(node.Metadata)}
				valueOnly.Metadata["valueOnly"] = true
				*queue = append(*queue, &valueOnly)
			}
		}
	}
}
