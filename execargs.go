package arg

import (
	"strconv"
	"strings"
)

// ExecArgs represent the flag arguments passed to
// a Cmd's Exec function.
type ExecArgs interface {
	// GetFlag gets the flag value given by name.
	// Values are passed as pointers and can be any
	// of the following types: string, bool, int64
	// uint64, float64, []string, []int64, []uint64,
	// or []float64.
	// Returns false if the flag name does not exist
	// or if the value type cannot be used.
	GetFlag(name string, value interface{}) bool
}

type execargs struct {
	kvp map[string]string
}

func (ea *execargs) GetFlag(name string, val interface{}) bool {
	var stat bool
	switch v := val.(type) {
	case *bool:
		stat = ea.getBool(name, v)
	case *string:
		stat = ea.getString(name, v)
	case *int64:
		stat = ea.getInt64(name, v)
	case *uint64:
		stat = ea.getUint64(name, v)
	case *float64:
		stat = ea.getFloat64(name, v)
	case *[]string:
		stat = ea.getStringSlice(name, v)
	case *[]int64:
		stat = ea.getInt64Slice(name, v)
	case *[]uint64:
		stat = ea.getUint64Slice(name, v)
	case *[]float64:
		stat = ea.getFloat64Slice(name, v)
	default:
		stat = false
	}
	return stat
}
func (ea *execargs) getString(name string, val *string) bool {
	if s, ok := ea.kvp[name]; ok {
		*val = s
		return true
	}
	return false
}
func (ea *execargs) getBool(name string, val *bool) bool {
	if s, ok := ea.kvp[name]; ok {
		if b, err := strconv.ParseBool(s); err == nil {
			*val = b
			return true
		}
	}
	return false
}
func (ea *execargs) getInt64(name string, val *int64) bool {
	if s, ok := ea.kvp[name]; ok {
		if i, err := strconv.ParseInt(s, 0, 64); err == nil {
			*val = i
			return true
		}
	}
	return false
}
func (ea *execargs) getUint64(name string, val *uint64) bool {
	if s, ok := ea.kvp[name]; ok {
		if ui, err := strconv.ParseUint(s, 0, 64); err == nil {
			*val = ui
			return true
		}
	}
	return false
}
func (ea *execargs) getFloat64(name string, val *float64) bool {
	if s, ok := ea.kvp[name]; ok {
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			*val = f
			return true
		}
	}
	return false
}
func (ea *execargs) getStringSlice(name string, val *[]string) bool {
	slice := ea.getSlice(name)
	if len(slice) == 0 {
		return false
	}
	*val = slice
	return true
}
func (ea *execargs) getFloat64Slice(name string, val *[]float64) bool {
	var floats []float64
	slice := ea.getSlice(name)
	for _, v := range slice {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			floats = append(floats, f)
		}
	}
	if len(floats) == 0 {
		return false
	}
	*val = floats
	return true
}
func (ea *execargs) getInt64Slice(name string, val *[]int64) bool {
	var ints []int64
	slice := ea.getSlice(name)
	for _, v := range slice {
		if i, err := strconv.ParseInt(v, 0, 64); err == nil {
			ints = append(ints, i)
		}
	}
	if len(ints) == 0 {
		return false
	}
	*val = ints
	return true
}
func (ea *execargs) getUint64Slice(name string, val *[]uint64) bool {
	var ints []uint64
	slice := ea.getSlice(name)
	for _, v := range slice {
		if i, err := strconv.ParseUint(v, 0, 64); err == nil {
			ints = append(ints, i)
		}
	}
	if len(ints) == 0 {
		return false
	}
	*val = ints
	return true
}
func (ea *execargs) getSlice(name string) []string {
	if s, ok := ea.kvp[name]; ok {
		return strings.Split(s, " ")
	}
	return nil
}
