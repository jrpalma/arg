package arg

import (
	"testing"
)

func TestInvalidOperand(t *testing.T) {
	var dst []string
	ea := newExecArgs()

	if ea.GetOperand(-1, nil) {
		t.Errorf("Should fail with invalid index")
	}
	if ea.GetOperand(0, &dst) {
		t.Errorf("Should fail with invalid dst")
	}
}
func TestGetOperand(t *testing.T) {
	var s string
	var b bool
	var i64 int64
	var ui64 uint64
	var f64 float64

	ea := newExecArgs()
	ea.setOperand(0, "string")
	ea.setOperand(1, "true")
	ea.setOperand(2, "-3")
	ea.setOperand(3, "3")
	ea.setOperand(4, "3.14")

	if !ea.GetOperand(0, &s) {
		t.Errorf("Failed to get arg string")
	}
	if !ea.GetOperand(1, &b) {
		t.Errorf("Failed to get arg bool")
	}
	if !ea.GetOperand(2, &i64) {
		t.Errorf("Failed to get arg int64")
	}
	if !ea.GetOperand(3, &ui64) {
		t.Errorf("Failed to get arg uint64")
	}
	if !ea.GetOperand(4, &f64) {
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
	var invalid *execArgs
	ea := newExecArgs()

	ea.addFlag('a', "string", "string")
	ea.addFlag('b', "bool", "true")
	ea.addFlag('c', "int64", "-3")
	ea.addFlag('d', "uint64", "3")
	ea.addFlag('e', "float64", "3.0")
	ea.addFlag('f', "enum", "enabled")
	ea.addFlag('g', "args", "1 2")
	ea.addFlag('h', "sf64", "3.0 2.0")
	ea.addFlag('i', "si64", "3 -2")
	ea.addFlag('j', "sui64", "3 2")

	if !ea.GetFlag("string", &s) {
		t.Errorf("Failed to get arg string")
	}
	if !ea.GetFlag("bool", &b) {
		t.Errorf("Failed to get arg bool")
	}
	if !ea.GetFlag("int64", &i64) {
		t.Errorf("Failed to get arg int64")
	}
	if !ea.GetFlag("uint64", &ui64) {
		t.Errorf("Failed to get arg uint64")
	}
	if !ea.GetFlag("float64", &f64) {
		t.Errorf("Failed to get arg float64")
	}
	if !ea.GetFlag("enum", &e) {
		t.Errorf("Failed to get arg enum")
	}
	if !ea.GetFlag("args", &slice) {
		t.Errorf("Failed to get arg slice")
	}
	if !ea.GetFlag("sf64", &sf64) {
		t.Errorf("Failed to get arg sf64")
	}
	if !ea.GetFlag("si64", &si64) {
		t.Errorf("Failed to get arg si64")
	}
	if !ea.GetFlag("sui64", &sui64) {
		t.Errorf("Failed to get arg sui64")
	}
	if ea.GetFlag("sui64", &invalid) {
		t.Errorf("Failed to get arg sui64")
	}
	if ea.GetFlag("unknown", nil) {
		t.Errorf("Get should failed")
	}
}
