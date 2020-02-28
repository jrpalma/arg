package arg

// CmdFlagSet represent a set of flags for a Cmd. The flags
// can be required or optional.
type CmdFlagSet struct {
	kvp map[string]*flag
}

// ReqString adds a required string flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) ReqString(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: stringFlag, req: true})
}

// ReqBool adds a required bool flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) ReqBool(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: boolFlag, req: true})
}

// ReqInt64 adds a required int64 flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) ReqInt64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: int64Flag, req: true})
}

// ReqUint64 adds a required uint64 flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) ReqUint64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: uint64Flag, req: true})
}

// ReqFloat64 adds a required float64 flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed,  the help string will be used to document the flag.
func (fs *CmdFlagSet) ReqFloat64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: float64Flag, req: true})
}

// ReqEnum adds a required enumeration flag with name to the set.
// The argument short represent's the flag name short version.
// Enum is the list of the valid values that can be used.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) ReqEnum(name string, short string, enum []string, help string) {
	fs.add(&flag{name: name, short: short, enum: enum, help: help, typ: enumFlag, req: true})
}

// ReqArgs adds a required number arguments with name to the set.
// The argument short represent's the flag name short version.
// Count is the number of arguments required for a flag.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) ReqArgs(name string, short string, count uint, help string) {
	fs.add(&flag{name: name, short: short, count: count, help: help, typ: argsFlag, req: true})
}

// OptString adds an optional string flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) OptString(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: stringFlag})
}

// OptBool adds an optional bool flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) OptBool(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: boolFlag})
}

// OptInt64 adds an optional int64 flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) OptInt64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: int64Flag})
}

// OptUint64 adds an optional uint64 flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) OptUint64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: uint64Flag})
}

// OptFloat64 adds an optional float64 flag with name to the set.
// The argument short represent's the flag name short version.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) OptFloat64(name string, short string, help string) {
	fs.add(&flag{name: name, short: short, help: help, typ: float64Flag})
}

// OptEnum adds an optional enum flag with name to the set.
// The argument short represent's the flag name short version.
// Enum is the list of the valid values that can be used.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) OptEnum(name string, short string, enum []string, help string) {
	fs.add(&flag{name: name, short: short, enum: enum, help: help, typ: enumFlag})
}

// OptArgs adds an optional number arguments to a flag with name to the set.
// The argument short represent's the flag name short version.
// Count is the number of arguments required for a flag.
// If usage is displayed, the help string will be used to document the flag.
func (fs *CmdFlagSet) OptArgs(name string, short string, count uint, help string) {
	fs.add(&flag{name: name, short: short, count: count, help: help, typ: argsFlag})
}
func (fs *CmdFlagSet) add(f *flag) {
	if fs.kvp == nil {
		fs.kvp = make(map[string]*flag)
	}
	fs.kvp[f.Name()] = f
}
func (fs *CmdFlagSet) hasReq() bool {
	for _, item := range fs.kvp {
		if item.Required() {
			return true
		}
	}
	return false
}
func (fs *CmdFlagSet) getFlags() map[string]*flag {
	flags := make(map[string]*flag)
	for _, item := range fs.kvp {
		copy := &flag{}
		*copy = *item
		flags[copy.Name()] = copy
	}
	return flags
}
