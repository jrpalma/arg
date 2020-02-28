package arg

type ArgSet struct {
	args []ArgType
}

func (as *ArgSet) Add(argType ArgType) {
	as.args = append(as.args, argType)
}
