package arg

// ExecArgs represent the flag arguments passed to
// a Cmd's Exec function.
type ExecArgs interface {
	// HasOption returns true if the option given by name is set.
	HasOption(name string) bool
	// GetFlag gets the command flag value given by name.
	// Value is passed as pointer and can be any
	// of the following types: string, bool, int64
	// uint64, float64, []string, []int64, []uint64,
	// or []float64.
	// Returns false if the flag name does not exist
	// or if the value type cannot be used.
	GetFlag(name string, value interface{}) bool

	// GetOperand gets command operand value given by index.
	// Values are passed as pointers and can be any
	// of the following types: string, bool, int64
	// uint64, or float64.
	// Returns false if the argument index is out of range
	// or if the value type cannot be used.
	GetOperand(index int, value interface{}) bool
}

type execArgs struct {
	longFlags  map[string][]string
	shortKeys  map[rune]string
	longOpts   map[string]struct{}
	shortOpts  map[rune]struct{}
	operValues map[int]string
}

func newExecArgs() *execArgs {
	ea := &execArgs{}
	ea.longFlags = make(map[string][]string)
	ea.shortKeys = make(map[rune]string)
	ea.longOpts = make(map[string]struct{})
	ea.shortOpts = make(map[rune]struct{})
	ea.operValues = make(map[int]string)
	return ea
}

func (ea *execArgs) HasOption(name string) bool {
	runes := getRunes(name)

	if len(runes) == 0 {
		return false
	}

	if len(runes) == 1 {
		if _, ok := ea.shortOpts[runes[0]]; ok {
			return true
		}
	}

	if _, ok := ea.longOpts[name]; ok {
		return true
	}

	return false
}

func (ea *execArgs) addFlag(short rune, long string, value string) {
	ea.longFlags[long] = append(ea.longFlags[long], value)
	ea.shortKeys[short] = long
}
func (ea *execArgs) setOption(short rune, long string) {
	ea.shortOpts[short] = struct{}{}
	ea.longOpts[long] = struct{}{}
}
func (ea *execArgs) setOperand(position int, value string) {
	ea.operValues[position] = value
}

func (ea *execArgs) GetOperand(position int, val interface{}) bool {

	argValue, exist := ea.operValues[position]
	if !exist {
		return false
	}

	var stat bool
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
func (ea *execArgs) GetFlag(name string, val interface{}) bool {

	longKey := name
	runes := getRunes(name)

	if len(runes) == 1 {
		longKey = ea.shortKeys[runes[0]]
	}

	flagValue, flagExist := ea.longFlags[longKey]
	if !flagExist {
		return false
	}

	var stat bool
	switch v := val.(type) {
	case *string:
		stat = true
		*v = flagValue[0]
	case *bool:
		stat = getBool(flagValue[0], v)
	case *int64:
		stat = getInt64(flagValue[0], v)
	case *uint64:
		stat = getUint64(flagValue[0], v)
	case *float64:
		stat = getFloat64(flagValue[0], v)
	case *[]string:
		stat = true
		*v = flagValue
	case *[]int64:
		stat = getInt64Slice(flagValue, v)
	case *[]uint64:
		stat = getUint64Slice(flagValue, v)
	case *[]float64:
		stat = getFloat64Slice(flagValue, v)
	default:
		stat = false
	}
	return stat
}
