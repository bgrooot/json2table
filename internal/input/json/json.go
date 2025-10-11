package json

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iancoleman/orderedmap"
)

func Unmarshal(jsonStr string) ([]*orderedmap.OrderedMap, error) {
	jsonStr = strings.TrimSpace(jsonStr)
	if len(jsonStr) == 0 {
		return nil, createInvalidJsonErr(fmt.Errorf("empty json string"))
	}

	if strings.HasPrefix(jsonStr, "{") {
		var orderedMap orderedmap.OrderedMap
		if err := json.Unmarshal([]byte(jsonStr), &orderedMap); err != nil {
			return nil, createInvalidJsonErr(err)
		}
		return []*orderedmap.OrderedMap{&orderedMap}, nil
	}

	if strings.HasPrefix(jsonStr, "[") {
		var orderedMapArr []*orderedmap.OrderedMap
		if err := json.Unmarshal([]byte(jsonStr), &orderedMapArr); err != nil {
			return nil, createInvalidJsonErr(err)
		}
		return orderedMapArr, nil
	}

	return nil, createInvalidJsonErr(fmt.Errorf("json must be an object or an array of objects"))
}

func createInvalidJsonErr(err error) error {
	if err == nil {
		return fmt.Errorf("invalid JSON")
	} else {
		return fmt.Errorf("invalid JSON: %w", err)
	}
}
