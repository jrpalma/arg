package arg

import (
	"testing"
)

func TestInvalidArg(t *testing.T) {
	var dst []string
	var args []string
	args = append(args, "string")
	ca := &execargs{args: args}

	if ca.GetArg(-1, nil) {
		t.Errorf("Should fail with invalid index")
	}
	if ca.GetArg(0, &dst) {
		t.Errorf("Should fail with invalid dst")
	}
}
func TestGetArg(t *testing.T) {
	var s string
	var b bool
	var i64 int64
	var ui64 uint64
	var f64 float64
	var args []string

	args = append(args, "string")
	args = append(args, "true")
	args = append(args, "-3")
	args = append(args, "3")
	args = append(args, "3.14")

	ca := &execargs{args: args}

	if !ca.GetArg(0, &s) {
		t.Errorf("Failed to get arg string")
	}
	if !ca.GetArg(1, &b) {
		t.Errorf("Failed to get arg bool")
	}
	if !ca.GetArg(2, &i64) {
		t.Errorf("Failed to get arg int64")
	}
	if !ca.GetArg(3, &ui64) {
		t.Errorf("Failed to get arg uint64")
	}
	if !ca.GetArg(4, &f64) {
		t.Errorf("Failed to get arg float64")
	}
}

func TestGetOption(t *testing.T) {
	var s string
	var b bool
	var i64 int64
	var ui64 uint64
	var f64 float64
	var e string
	var slice []string
	var sf64 []float64
	var si64 []int64
	var sui64 []uint64
	var invalid *execargs

	opts := make(map[string]string)
	opts["string"] = "string"
	opts["bool"] = "true"
	opts["int64"] = "-3"
	opts["uint64"] = "3"
	opts["float64"] = "3.0"
	opts["enum"] = "enabled"
	opts["args"] = "1 2"
	opts["sf64"] = "3.0 2.0"
	opts["si64"] = "3 -2"
	opts["sui64"] = "3 2"
	ca := &execargs{options: opts}

	if !ca.GetOption("string", &s) {
		t.Errorf("Failed to get arg string")
	}
	if !ca.GetOption("bool", &b) {
		t.Errorf("Failed to get arg bool")
	}
	if !ca.GetOption("int64", &i64) {
		t.Errorf("Failed to get arg int64")
	}
	if !ca.GetOption("uint64", &ui64) {
		t.Errorf("Failed to get arg uint64")
	}
	if !ca.GetOption("float64", &f64) {
		t.Errorf("Failed to get arg float64")
	}
	if !ca.GetOption("enum", &e) {
		t.Errorf("Failed to get arg enum")
	}
	if !ca.GetOption("args", &slice) {
		t.Errorf("Failed to get arg slice")
	}
	if !ca.GetOption("sf64", &sf64) {
		t.Errorf("Failed to get arg sf64")
	}
	if !ca.GetOption("si64", &si64) {
		t.Errorf("Failed to get arg si64")
	}
	if !ca.GetOption("sui64", &sui64) {
		t.Errorf("Failed to get arg sui64")
	}
	if ca.GetOption("sui64", &invalid) {
		t.Errorf("Failed to get arg sui64")
	}
	if ca.GetOption("unknown", nil) {
		t.Errorf("Get should failed")
	}
}
