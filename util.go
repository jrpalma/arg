package arg

import (
	"strconv"
	"strings"
)

func getBool(src string, dst interface{}) bool {
	destination, correctType := dst.(*bool)
	if !correctType {
		return false
	}
	parsedValue, err := strconv.ParseBool(src)
	if err != nil {
		return false
	}
	*destination = parsedValue
	return true
}
func getInt64(src string, dst interface{}) bool {
	destination, correctType := dst.(*int64)
	if !correctType {
		return false
	}
	parsedValue, err := strconv.ParseInt(src, 0, 64)
	if err != nil {
		return false
	}
	*destination = parsedValue
	return true
}
func getUint64(src string, dst interface{}) bool {
	destination, correctType := dst.(*uint64)
	if !correctType {
		return false
	}
	parsedValue, err := strconv.ParseUint(src, 0, 64)
	if err != nil {
		return false
	}
	*destination = parsedValue
	return true
}
func getFloat64(src string, dst interface{}) bool {
	destination, correctType := dst.(*float64)
	if !correctType {
		return false
	}
	parsedValue, err := strconv.ParseFloat(src, 64)
	if err != nil {
		return false
	}
	*destination = parsedValue
	return true
}
func getStringSlice(src string, dst interface{}, sep string) bool {
	destination, correctType := dst.(*[]string)
	if !correctType {
		return false
	}
	*destination = strings.Split(src, sep)
	return true
}
func getFloat64Slice(src string, dst interface{}, sep string) bool {
	destination, correctType := dst.(*[]float64)
	if !correctType {
		return false
	}
	for _, item := range strings.Split(src, sep) {
		var value float64
		if getFloat64(item, &value) {
			*destination = append(*destination, value)
		}
	}
	return true
}
func getInt64Slice(src string, dst interface{}, sep string) bool {
	destination, correctType := dst.(*[]int64)
	if !correctType {
		return false
	}
	for _, item := range strings.Split(src, sep) {
		var value int64
		if getInt64(item, &value) {
			*destination = append(*destination, value)
		}
	}
	return true
}
func getUint64Slice(src string, dst interface{}, sep string) bool {
	destination, correctType := dst.(*[]uint64)
	if !correctType {
		return false
	}
	for _, item := range strings.Split(src, sep) {
		var value uint64
		if getUint64(item, &value) {
			*destination = append(*destination, value)
		}
	}
	return true
}
