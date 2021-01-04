package u4g

import (
	"encoding/json"
	"fmt"
)

func PrintlnJSON(vals ...interface{}) {
	for _, v := range vals {
		c := ""
		if s, ok := v.(string); ok {
			c = s
		} else {
			j, err := json.MarshalIndent(v, "", "\t")
			if err != nil {
				fmt.Println("PrintlnJSON error:", err)
				continue
			}
			c = string(j)
		}
		fmt.Println(c)
	}
}

func SprintJSON(v interface{}) string {
	c := ""
	if s, ok := v.(string); ok {
		c = s
	} else {
		j, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			fmt.Println("PrintlnJSON error:", err)
		} else {
			c = string(j)
		}
	}
	return c
}
