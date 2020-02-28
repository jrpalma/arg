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

func TestGetFlag(t *testing.T) {
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

	flags := make(map[string]string)
	flags["string"] = "string"
	flags["bool"] = "true"
	flags["int64"] = "-3"
	flags["uint64"] = "3"
	flags["float64"] = "3.0"
	flags["enum"] = "enabled"
	flags["args"] = "1 2"
	flags["sf64"] = "3.0 2.0"
	flags["si64"] = "3 -2"
	flags["sui64"] = "3 2"
	ca := &execargs{flags: flags}

	if !ca.GetFlag("string", &s) {
		t.Errorf("Failed to get arg string")
	}
	if !ca.GetFlag("bool", &b) {
		t.Errorf("Failed to get arg bool")
	}
	if !ca.GetFlag("int64", &i64) {
		t.Errorf("Failed to get arg int64")
	}
	if !ca.GetFlag("uint64", &ui64) {
		t.Errorf("Failed to get arg uint64")
	}
	if !ca.GetFlag("float64", &f64) {
		t.Errorf("Failed to get arg float64")
	}
	if !ca.GetFlag("enum", &e) {
		t.Errorf("Failed to get arg enum")
	}
	if !ca.GetFlag("args", &slice) {
		t.Errorf("Failed to get arg slice")
	}
	if !ca.GetFlag("sf64", &sf64) {
		t.Errorf("Failed to get arg sf64")
	}
	if !ca.GetFlag("si64", &si64) {
		t.Errorf("Failed to get arg si64")
	}
	if !ca.GetFlag("sui64", &sui64) {
		t.Errorf("Failed to get arg sui64")
	}
	if ca.GetFlag("sui64", &invalid) {
		t.Errorf("Failed to get arg sui64")
	}
	if ca.GetFlag("unknown", nil) {
		t.Errorf("Get should failed")
	}
}
