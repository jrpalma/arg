package arg

// FlagSet represent a set of flags for a . The flags
// can be required or optional.
type FlagSet struct {
	kvp map[string]*flag
}

// ReqString adds a required string flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) ReqString(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: stringFlag, req: true})
}

// ReqBool adds a required bool flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) ReqBool(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: boolFlag, req: true})
}

// ReqInt64 adds a required int64 flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) ReqInt64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: int64Flag, req: true})
}

// ReqUint64 adds a required uint64 flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) ReqUint64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: uint64Flag, req: true})
}

// ReqFloat64 adds a required float64 flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed,  the help string will be used to document the flag.
func (fs *FlagSet) ReqFloat64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: float64Flag, req: true})
}

// ReqEnum adds a required enumeration flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// Enum is the list of the valid values that can be used.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) ReqEnum(name string, short string, enum []string, help string) {
	fs.add(&flag{name: name, short: short, enum: enum, help: help, typ: enumFlag, req: true})
}

// ReqArgs adds a required number of named arguments with name to the set.
// The argument short represent's the flag name shorthand letter.
// Names are the names of the ordered arguments for the this flag.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) ReqArgs(name string, short string, names []string, help string) {
	fs.add(&flag{name: name, short: short, names: names, help: help, typ: argsFlag, req: true})
}

// OptString adds an optional string flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) OptString(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: stringFlag})
}

// OptBool adds an optional bool flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) OptBool(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: boolFlag})
}

// OptInt64 adds an optional int64 flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) OptInt64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: int64Flag})
}

// OptUint64 adds an optional uint64 flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) OptUint64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: uint64Flag})
}

// OptFloat64 adds an optional float64 flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) OptFloat64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: float64Flag})
}

// OptEnum adds an optional enum flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// Enum is the list of the valid values that can be used.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) OptEnum(name string, short string, enum []string, help string) {
	fs.add(&flag{name: name, short: short, enum: enum, help: help, typ: enumFlag})
}

// OptArgs adds an optional number of named arguments to a flag with name to the set.
// The argument short represent's the flag name shorthand letter.
// Names are the names of the ordered arguments for the this flag.
// If usage is displayed, the help string will be used to document the flag.
func (fs *FlagSet) OptArgs(name string, short string, names []string, help string) {
	fs.add(&flag{name: name, short: short, names: names, help: help, typ: argsFlag})
}
func (fs *FlagSet) add(f *flag) {
	if fs.kvp == nil {
		fs.kvp = make(map[string]*flag)
	}
	fs.kvp[f.name] = f
}
func (fs *FlagSet) hasReq() bool {
	for _, item := range fs.kvp {
		if item.req {
			return true
		}
	}
	return false
}
func (fs *FlagSet) getFlags() map[string]*flag {
	flags := make(map[string]*flag)
	for _, item := range fs.kvp {
		copy := &flag{}
		*copy = *item
		flags[copy.name] = copy
	}
	return flags
}
