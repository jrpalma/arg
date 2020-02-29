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

// CmdArgs represents the set of command arguments
type CmdArgs struct {
	args []*cmdarg
}

// Add a command argument with the given name.
// Valid types are: StringArg, BoolArg, Int64Arg, Uint64Arg, and Float64Arg.
func (cas *CmdArgs) Add(name string, cmdArgType CmdArgType) {
	cas.args = append(cas.args, &cmdarg{name: name, typ: cmdArgType})

}

type cmdarg struct {
	name string
	typ  CmdArgType
}
