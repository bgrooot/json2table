package core

import (
	"fmt"

	"github.com/bgrooot/json2table/internal/model"
	"github.com/bgrooot/json2table/pkg/json2table/config"
	"github.com/iancoleman/orderedmap"
)

func Parsing(omArr []*orderedmap.OrderedMap, cfg config.Config) []model.Node {
	var nodeArr []model.Node

	if cfg.DivideArray.IsTrue() {
		for _, om := range omArr {
			nodeArr = append(nodeArr, parsing([]*orderedmap.OrderedMap{om}, cfg))
		}
	} else {
		nodeArr = append(nodeArr, parsing(omArr, cfg))
	}

	return nodeArr
}

func parsing(omArr []*orderedmap.OrderedMap, cfg config.Config) model.Node {
	root := model.Node{}
	for _, om := range omArr {
		for _, key := range om.Keys() {
			value, exists := om.Get(key)
			if exists {
				child := root.AddChild(key)
				traverse(key, value, child, cfg)
			}
		}
	}

	return root
}

func traverse(key string, value any, node *model.Node, cfg config.Config) {
	switch val := value.(type) {
	case string:
		node.AddValue(val)

	case bool:
		node.AddValue(fmt.Sprintf("%t", val))

	case orderedmap.OrderedMap:
		if len(val.Keys()) == 0 {
			node.AddValue("")
		}

		for _, k := range val.Keys() {
			child := node.AddChild(k)
			v, _ := val.Get(k)
			traverse(k, v, child, cfg)
		}

	case map[string]any:
		for k, v := range val {
			child := node.AddChild(k)
			traverse(k, v, child, cfg)
		}

	case []any:
		if len(val) == 0 {
			node.AddValue("")
		}

		for _, v := range val {
			traverse(key, v, node, cfg)
		}

	default:
		node.AddValue(fmt.Sprintf("%v", val))
	}
}
