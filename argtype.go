package arg

type ArgType int

const (
	StringArg ArgType = iota
	BoolArg
	Int64Arg
	Uint64Arg
	Float64Arg
)
