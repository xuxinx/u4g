package u4g

import "testing"

func TestPrintlnJSON(t *testing.T) {
	type MyS struct {
		Name string
		Desc string
	}

	PrintlnJSON("111", 222, &MyS{
		Name: "nnn",
		Desc: "ddd",
	})
}
