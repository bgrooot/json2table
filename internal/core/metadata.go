package core

import (
	"github.com/bgrooot/json2table/internal/model"
	"github.com/bgrooot/json2table/pkg/json2table/config"
)

func SetUp(nodeArr []model.Node, cfg config.Config) {
	for i := range nodeArr {
		n := &(nodeArr)[i]
		maxDepth := getMaxDepth(n, 0)
		setRelativeDepth(n, 0, maxDepth)
		setChildCount(n, cfg)
	}
}

func getMaxDepth(node *model.Node, depth int) int {
	maxDepth := depth
	for _, val := range node.Value {
		if nd, ok := val.(*model.Node); ok {
			dep := getMaxDepth(nd, depth+1)
			if dep > maxDepth {
				maxDepth = dep
			}
		}
	}

	return maxDepth
}

func setRelativeDepth(node *model.Node, depth int, maxDepth int) {
	node.PutMetadata("relativeDepth", maxDepth-depth+1)
	for _, val := range node.Value {
		if nd, ok := val.(*model.Node); ok {
			setRelativeDepth(nd, depth+1, maxDepth)
		}
	}
}

func setChildCount(node *model.Node, cfg config.Config) {
	sum := 0
	if len(node.Value) > 0 {
		for _, val := range node.Value {
			if nd, ok := val.(*model.Node); ok {
				setChildCount(nd, cfg)
				rowSpan, exist := nd.Metadata["childCount"]
				if exist {
					sum += rowSpan.(int)
				}
			} else {
				sum++
			}
		}
	}

	node.PutMetadata("childCount", sum)
}
