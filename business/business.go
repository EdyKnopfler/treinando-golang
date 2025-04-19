package business

import (
	"encoding/json"
	"fmt"
)

func BusinessRule(jsonString string) (result map[string]any) {
	// A so much complex business rule
	err := json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		panic("Please pass a valid JSON or bad things will happen")
	}

	return
}

func PrintMap(mapData map[string]any, indent string) {
	for key, value := range mapData {
		switch val := value.(type) {
		case map[string]any:
			fmt.Printf("%s%s:\n", indent, key)
			PrintMap(val, indent+"  ")
		case []any:
			fmt.Printf("%s%s:\n", indent, key)
			for i, item := range val {
				fmt.Printf("%s  [%d]:\n", indent, i)
				if itemMap, ok := item.(map[string]any); ok {
					PrintMap(itemMap, indent+"    ")
				} else {
					fmt.Printf("%s    %v\n", indent, item)
				}
			}
		default:
			fmt.Printf("%s%s: %v\n", indent, key, val)
		}
	}
}
