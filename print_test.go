package u4g

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type MyS struct {
	Name string
	Desc string
}

func TestPrintlnJSON(t *testing.T) {
	PrintlnJSON("111", 222, &MyS{
		Name: "nnn",
		Desc: "ddd",
	})
}

func TestSprintJSON(t *testing.T) {
	var diff string

	diff = cmp.Diff(SprintJSON("111"), "111")
	if diff != "" {
		t.Fatalf("diff: %s", diff)
	}

	diff = cmp.Diff(SprintJSON(&MyS{
		Name: "nnn",
		Desc: "ddd",
	}), "{\n\t\"Name\": \"nnn\",\n\t\"Desc\": \"ddd\"\n}")
	if diff != "" {
		t.Fatalf("diff: %s", diff)
	}
}
