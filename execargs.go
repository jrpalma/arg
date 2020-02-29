package arg

import ()

// ExecArgs represent the flag arguments passed to
// a Cmd's Exec function.
type ExecArgs interface {
	// GetFlag gets the command flag value given by name.
	// Values are passed as pointers and can be any
	// of the following types: string, bool, int64
	// uint64, float64, []string, []int64, []uint64,
	// or []float64.
	// Returns false if the flag name does not exist
	// or if the value type cannot be used.
	GetFlag(name string, value interface{}) bool

	// GetArg gets command argument value given by index.
	// Values are passed as pointers and can be any
	// of the following types: string, bool, int64
	// uint64, or float64.
	// Returns false if the argument index is out of range
	// or if the value type cannot be used.
	GetArg(index int, value interface{}) bool
}

type execargs struct {
	flags map[string]string
	args  []string
}

func (ea *execargs) GetArg(index int, val interface{}) bool {
	if index < 0 || index >= len(ea.args) {
		return false
	}

	var stat bool
	argValue := ea.args[index]

	switch v := val.(type) {
	case *string:
		stat = true
		*v = argValue
	case *bool:
		stat = getBool(argValue, v)
	case *int64:
		stat = getInt64(argValue, v)
	case *uint64:
		stat = getUint64(argValue, v)
	case *float64:
		stat = getFloat64(argValue, v)
	default:
		stat = false
	}
	return stat
}
func (ea *execargs) GetFlag(name string, val interface{}) bool {
	flagValue, flagExist := ea.flags[name]
	if !flagExist {
		return false
	}

	var stat bool
	switch v := val.(type) {
	case *string:
		stat = true
		*v = flagValue
	case *bool:
		stat = getBool(flagValue, v)
	case *int64:
		stat = getInt64(flagValue, v)
	case *uint64:
		stat = getUint64(flagValue, v)
	case *float64:
		stat = getFloat64(flagValue, v)
	case *[]string:
		stat = getStringSlice(flagValue, v, ",")
	case *[]int64:
		stat = getInt64Slice(flagValue, v, ",")
	case *[]uint64:
		stat = getUint64Slice(flagValue, v, ",")
	case *[]float64:
		stat = getFloat64Slice(flagValue, v, ",")
	default:
		stat = false
	}
	return stat
}
