package arg

// CmdOptions represent a set of options for a Cmd. The options
// can be required or optional.
type CmdOptions struct {
	opts map[string]*option
}

// ReqString adds a required string option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) ReqString(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: stringOption, required: true})
}

// ReqBool adds a required bool option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) ReqBool(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: boolOption, required: true})
}

// ReqInt64 adds a required int64 option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) ReqInt64(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: int64Option, required: true})
}

// ReqUint64 adds a required uint64 option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) ReqUint64(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: uint64Option, required: true})
}

// ReqFloat64 adds a required float64 option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed,  the help string will be used to document the option.
func (co *CmdOptions) ReqFloat64(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: float64Option, required: true})
}

// ReqEnum adds a required enumeration option.
// The argument short represent's the option name shorthand letter.
// Enum is the list of the valid values that can be used.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) ReqEnum(name string, short string, enum []string, help string) {
	co.add(&option{name: name, short: short, enum: enum, help: help, optionType: enumOption, required: true})
}

// ReqArgs adds a required number of named arguments with name to the set.
// The argument short represent's the option name shorthand letter.
// Names are the names of the ordered arguments for the this option.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) ReqArgs(name string, short string, names []string, help string) {
	co.add(&option{name: name, short: short, args: names, help: help, optionType: argsOption, required: true})
}

// OptString adds an optional string option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) OptString(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: stringOption})
}

// OptBool adds an optional bool option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) OptBool(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: boolOption})
}

// OptInt64 adds an optional int64 option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) OptInt64(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: int64Option})
}

// OptUint64 adds an optional uint64 option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) OptUint64(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: uint64Option})
}

// OptFloat64 adds an optional float64 option.
// The argument short represent's the option name shorthand letter.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) OptFloat64(name string, short string, help string) {
	co.add(&option{name: name, short: short, help: help, optionType: float64Option})
}

// OptEnum adds an optional enum option.
// The argument short represent's the option name shorthand letter.
// Enum is the list of the valid values that can be used.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) OptEnum(name string, short string, enum []string, help string) {
	co.add(&option{name: name, short: short, enum: enum, help: help, optionType: enumOption})
}

// OptArgs adds an optional number of named arguments to a option.
// The argument short represent's the option name shorthand letter.
// Names are the names of the ordered arguments for the this option.
// If usage is displayed, the help string will be used to document the option.
func (co *CmdOptions) OptArgs(name string, short string, names []string, help string) {
	co.add(&option{name: name, short: short, args: names, help: help, optionType: argsOption})
}
func (co *CmdOptions) add(f *option) {
	if co.opts == nil {
		co.opts = make(map[string]*option)
	}
	co.opts[f.name] = f
}
func (co *CmdOptions) hasRequired() bool {
	for _, item := range co.opts {
		if item.required {
			return true
		}
	}
	return false
}
func (co *CmdOptions) getOptions() map[string]*option {
	options := make(map[string]*option)
	for _, item := range co.opts {
		copy := &option{}
		*copy = *item
		options[copy.name] = copy
	}
	return options
}
