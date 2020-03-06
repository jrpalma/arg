package arg

import (
	"strings"
	"testing"
)

func TestGetBool(t *testing.T) {
	var dst bool
	if getBool("x", nil) {
		t.Errorf("Should fail with invalid src")
	}
	if getBool("true", nil) {
		t.Errorf("Should fail with invalid dst")
	}
	if getBool("x", &dst) {
		t.Errorf("Should fail with invalid src")
	}
	if !getBool("true", &dst) {
		t.Errorf("Should not fail with vaild src and dst")
	}
}
func TestGetInt64(t *testing.T) {
	var dst int64
	if getInt64("x", nil) {
		t.Errorf("Should fail with invalid src")
	}
	if getInt64("-3", nil) {
		t.Errorf("Should fail with invalid dst")
	}
	if getInt64("x", &dst) {
		t.Errorf("Should fail with invalid src")
	}
	if !getInt64("-3", &dst) {
		t.Errorf("Should not fail with vaild src and dst")
	}
}
func TestGetUint64(t *testing.T) {
	var dst uint64
	if getUint64("x", nil) {
		t.Errorf("Should fail with invalid src")
	}
	if getUint64("3", nil) {
		t.Errorf("Should fail with invalid dst")
	}
	if getUint64("x", &dst) {
		t.Errorf("Should fail with invalid src")
	}
	if !getUint64("3", &dst) {
		t.Errorf("Should not fail with vaild src and dst")
	}
}
func TestGetFloat64(t *testing.T) {
	var dst float64
	if getFloat64("x", nil) {
		t.Errorf("Should fail with invalid src")
	}
	if getFloat64("3", nil) {
		t.Errorf("Should fail with invalid dst")
	}
	if getFloat64("x", &dst) {
		t.Errorf("Should fail with invalid src")
	}
	if !getFloat64("3", &dst) {
		t.Errorf("Should not fail with vaild src and dst")
	}
}
func TestGetFloat64Slice(t *testing.T) {
	list := strings.Split("1,2,3", ",")
	var dst []float64
	if getFloat64Slice(list, nil) {
		t.Errorf("Should fail with invalid dst")
	}
	if !getFloat64Slice(list, &dst) {
		t.Errorf("Should not fail with vaild src and dst")
	}
	if len(dst) != 3 {
		t.Errorf("String slice should be of length 3")
	}
}
func TestGetInt64Slice(t *testing.T) {
	list := strings.Split("1,2,3", ",")
	var dst []int64
	if getInt64Slice(list, nil) {
		t.Errorf("Should fail with invalid dst")
	}
	if !getInt64Slice(list, &dst) {
		t.Errorf("Should not fail with vaild src and dst")
	}
	if len(dst) != 3 {
		t.Errorf("String slice should be of length 3")
	}
}
func TestGetUint64Slice(t *testing.T) {
	list := strings.Split("1,2,3", ",")
	var dst []uint64
	if getUint64Slice(list, nil) {
		t.Errorf("Should fail with invalid dst")
	}
	if !getUint64Slice(list, &dst) {
		t.Errorf("Should not fail with vaild src and dst")
	}
	if len(dst) != 3 {
		t.Errorf("String slice should be of length 3")
	}
}
