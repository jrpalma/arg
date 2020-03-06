package arg

// Cmd represents a command that will be executed by a Parser.
// The command's Exec method is executed if a Parser can match
// the command's prefix, name, and flags.
type Cmd struct {
	// The command's prpefix. The prefix is the string that comes
	//before the command name and its flags.
	Prefix string
	// The command's name. The name must be unique per prefix.
	// If two commands with the same prefix and name are addOptioned
	// to a Parser object, the later will replace the the earlier.
	Name string
	// The command's help string. The command will display the help
	// string to document the command.
	Help string
	// The function that gets executed if the arguments match the command.
	Exec func(ExecArgs) error
	// The command's flags. Flags can be required or optional.

	operands  map[int]*operand
	shortOpts map[rune]string
	longOpts  map[string]*option
}

// Option adds an option. Short is a one letter option.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) Option(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help})
}

// ReqString adds a required string flag. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqString(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: String})
}

// ReqBool adds a required bool flag. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqBool(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Bool})
}

// ReqInt64 adds a required int64 flag. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqInt64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Int64})
}

// ReqUint64 adds a required uint64 flag. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqUint64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Uint64})
}

// ReqFloat64 adds a required float64 flag. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqFloat64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Float64})
}

// ReqEnum adds a required enum flag. Short is a one letter flag.
// Long is the long version of the flag. Valid are the valid values
// for this option. Help is displayed when the usage message is printed.
func (c *Cmd) ReqEnum(short rune, long string, valid []string, help string) {
	c.addOption(&option{short: short, long: long, help: help, valid: valid,
		required: true, dataType: enum})
}

// OptString adds an optional string flag. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) OptString(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: String})
}

// OptBool adds an optional bool flag. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) OptBool(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: Bool})
}

// OptInt64 adds an optional int64 flag. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) OptInt64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: Int64})
}

// OptUint64 adds an optional uint64 flag. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) OptUint64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: Uint64})
}

// OptFloat64 adds an optional float64 option. Short is a one letter flag.
// Long is the long version of the flag. Help is displayed
// when the usage message is printed.
func (c *Cmd) OptFloat64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: Float64})
}

// OptEnum adds an optional enum flag. Short is a one letter flag.
// Long is the long version of the flag. Valid are the valid values
// for this flag. Help is displayed when the usage message is printed.
func (c *Cmd) OptEnum(short rune, long string, valid []string, help string) {
	c.addOption(&option{short: short, long: long, help: help, valid: valid,
		dataType: enum})
}

// Operand adds an operand with the give position, name, and data type.
// Position starts from zero on wards. The operand is named by name.
// The operand data type can be specified.
func (c *Cmd) Operand(position int, name string, dataType DataType) {
	c.addOperand(&operand{position: position, name: name, dataType: dataType})
}

func (c *Cmd) addOption(o *option) {
	if c.shortOpts == nil {
		c.longOpts = make(map[string]*option)
		c.shortOpts = make(map[rune]string)
	}
	c.shortOpts[o.short] = o.long
	c.longOpts[o.long] = o
}

func (c *Cmd) longOption(name string) (*option, bool) {
	if o, ok := c.longOpts[name]; ok {
		return o, true
	}
	return nil, false
}

func (c *Cmd) shortOption(name rune) (*option, bool) {
	if longName, ok := c.shortOpts[name]; ok {
		o, exist := c.longOpts[longName]
		return o, exist
	}
	return nil, false
}

func (c *Cmd) getRequiredLongNames() map[string]struct{} {
	names := make(map[string]struct{})
	for longName, opt := range c.longOpts {
		if opt.required {
			names[longName] = struct{}{}
		}
	}
	return names
}

func (c *Cmd) addOperand(o *operand) {
	if c.operands == nil {
		c.operands = make(map[int]*operand)
	}
	c.operands[o.position] = o
}

func (c *Cmd) getOperands() map[int]operand {
	m := make(map[int]operand)
	for k, v := range c.operands {
		m[k] = *v
	}
	return m
}

type option struct {
	required bool
	short    rune
	long     string
	help     string
	dataType DataType
	valid    []string
	arg      string
}

type operand struct {
	position int
	name     string
	dataType DataType
}
