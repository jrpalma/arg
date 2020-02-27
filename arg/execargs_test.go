package arg

import (
	"testing"
)

func TestExecArgs(t *testing.T) {
	kvp := make(map[string]string)
	kvp["string"] = "string"
	kvp["bool"] = "true"
	kvp["int64"] = "-3"
	kvp["uint64"] = "3"
	kvp["float64"] = "3.0"
	kvp["enum"] = "enabled"
	kvp["args"] = "1 2"
	kvp["sf64"] = "3.0 2.0"
	kvp["si64"] = "3 -2"
	kvp["sui64"] = "3 2"

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

	ca := &execargs{kvp: kvp}
	if !ca.getString("string", &s) && s != "string" {
		t.Errorf("Failed to get arg string")
	}
	if !ca.getBool("bool", &b) && b != true {
		t.Errorf("Failed to get arg bool")
	}
	if !ca.getInt64("int64", &i64) && i64 != int64(-3) {
		t.Errorf("Failed to get arg int64")
	}
	if !ca.getUint64("uint64", &ui64) && ui64 != uint64(3) {
		t.Errorf("Failed to get arg uint64")
	}
	if !ca.getFloat64("float64", &f64) && f64 != 3.0 {
		t.Errorf("Failed to get arg f64")
	}
	if !ca.getString("enum", &e) && e != "disabled" {
		t.Errorf("Failed to get arg enum")
	}
	if !ca.getStringSlice("args", &slice) && len(slice) != 2 {
		t.Errorf("Failed to get arg args")
	}
	if !ca.getFloat64Slice("sf64", &sf64) && len(sf64) != 2 {
		t.Errorf("Failed to get arg sf64")
	}
	if !ca.getInt64Slice("si64", &si64) && len(si64) != 2 {
		t.Errorf("Failed to get arg si64")
	}
	if !ca.getUint64Slice("sui64", &sui64) && len(sui64) != 2 {
		t.Errorf("Failed to get arg sui64")
	}

	if !ca.Get("string", &s) {
		t.Errorf("Failed to get arg string")
	}
	if !ca.Get("bool", &b) {
		t.Errorf("Failed to get arg bool")
	}
	if !ca.Get("int64", &i64) {
		t.Errorf("Failed to get arg int64")
	}
	if !ca.Get("uint64", &ui64) {
		t.Errorf("Failed to get arg uint64")
	}
	if !ca.Get("float64", &f64) {
		t.Errorf("Failed to get arg float64")
	}
	if !ca.Get("enum", &e) {
		t.Errorf("Failed to get arg enum")
	}
	if !ca.Get("args", &slice) {
		t.Errorf("Failed to get arg slice")
	}
	if !ca.Get("sf64", &sf64) {
		t.Errorf("Failed to get arg sf64")
	}
	if !ca.Get("si64", &si64) {
		t.Errorf("Failed to get arg si64")
	}
	if !ca.Get("sui64", &sui64) {
		t.Errorf("Failed to get arg sui64")
	}
	if ca.Get("unknown", nil) {
		t.Errorf("Get should failed")
	}

	if ca.getString("unknown", nil) {
		t.Errorf("Should fail")
	}
	if ca.getBool("unknown", nil) {
		t.Errorf("Should fail")
	}
	if ca.getInt64("unknown", nil) {
		t.Errorf("Should fail")
	}
	if ca.getUint64("unknown", nil) {
		t.Errorf("Should fail")
	}
	if ca.getFloat64("unknown", nil) {
		t.Errorf("Should fail")
	}
	if ca.getStringSlice("unknown", nil) {
		t.Errorf("Should fail")
	}
	if ca.getFloat64Slice("unknown", nil) {
		t.Errorf("Should fail")
	}
	if ca.getInt64Slice("unknown", nil) {
		t.Errorf("Should fail")
	}
	if ca.getUint64Slice("unknown", nil) {
		t.Errorf("Should fail")
	}
}
