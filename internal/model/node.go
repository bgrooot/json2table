package model

import "fmt"

type Node struct {
	Key      string
	Value    []any
	Metadata map[string]any
}

func (node *Node) AddChild(key string) *Node {
	child := &Node{Key: key, Value: []any{}}
	node.Value = append(node.Value, child)
	return child
}

func (node *Node) AddValue(value string) *Node {
	node.Value = append(node.Value, value)
	return node
}

func (node *Node) PutMetadata(key string, value any) {
	if node.Metadata == nil {
		node.Metadata = map[string]any{}
	}

	node.Metadata[key] = value
}

func (node *Node) Print() {
	fmt.Printf("Key: %s, Value: %v\n", node.Key, node.Value)
	for _, val := range node.Value {
		if nd, ok := val.(*Node); ok {
			nd.Print()
		}
	}
}
