package arg

// CmdArgType represents a command's argument type
type CmdArgType int

const (
	// StringArg represent a command string argument
	StringArg CmdArgType = iota
	// BoolArg represent a command bool argument
	BoolArg
	// Int64Arg represent a command int64 argument
	Int64Arg
	// Uint64Arg represent a command uint64 argument
	Uint64Arg
	// Float64Arg represent a command float64 argument
	Float64Arg
)

// CmdArgSet represents the set of command arguments
type CmdArgSet struct {
	args []CmdArgType
}

// SetTypes sets the type of arguments used by a command.
// Valid types are: StringArg, BoolArg, Int64Arg, Uint64Arg, and Float64Arg.
func (cas *CmdArgSet) SetTypes(types ...CmdArgType) {
	for _, argType := range types {
		cas.args = append(cas.args, argType)
	}
}
