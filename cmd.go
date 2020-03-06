package arg

import (
	"sort"
)

// Cmd represents a command that will be executed by a Parser.
// The command's Exec method is executed if a Parser can match
// the command's prefix, name, options, and operands.
type Cmd struct {
	// The command's prpefix. The prefix is the string that comes
	//before the command name.
	Prefix string
	// The command's name. The name must be unique per prefix.
	// If two commands with the same prefix and name are added
	// to a Parser object, the later will replace the the earlier.
	Name string
	// The command's help string. The command will display the help
	// if usage is displayed.
	Help string
	// The function that gets executed if the arguments match the command.
	Exec func(ExecArgs) error

	operands  map[int]*operand
	shortOpts map[rune]*option
	longOpts  map[string]*option
}

// Option adds an option that does not required a parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) Option(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help})
}

// ReqString adds a required option with a string parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) ReqString(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: String})
}

// ReqBool adds a required option with a bool parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) ReqBool(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Bool})
}

// ReqInt64 adds a required option with a int64 parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) ReqInt64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Int64})
}

// ReqUint64 adds a required option with a uint64 parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) ReqUint64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Uint64})
}

// ReqFloat64 adds a required option with a float64 parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) ReqFloat64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Float64})
}

// ReqEnum adds a required enum option. Valid are the valid values
// that are acceped by this option.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) ReqEnum(short rune, long string, valid []string, help string) {
	c.addOption(&option{short: short, long: long, help: help, valid: valid,
		required: true, dataType: enum})
}

// OptString adds a optional option with a string parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) OptString(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: String})
}

// OptBool adds a optional option with a bool parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) OptBool(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: Bool})
}

// OptInt64 adds a optional option with a int64 parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) OptInt64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: Int64})
}

// OptUint64 adds a optional option with a uint64 parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) OptUint64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: Uint64})
}

// OptFloat64 adds a optional option with a float64 parameter.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) OptFloat64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		dataType: Float64})
}

// OptEnum adds a optinal enum option. Valid are the valid values
// that are acceped by this option.
// Short is a one letter option. Short is not used if '0'.
// Long is the long name for the option. Long is not used if empty.
// Help is displayed when the usage message is printed.
func (c *Cmd) OptEnum(short rune, long string, valid []string, help string) {
	c.addOption(&option{short: short, long: long, help: help, valid: valid,
		dataType: enum})
}

// Operand adds an operand with the give position, name, and data type.
// Position starts from zero on wards. The operand is named by name.
// The operand data type can be specified. Valid data types are: String,
// Bool, Int64, Uint64, and Float64
func (c *Cmd) Operand(position int, name string, dataType DataType) {
	c.addOperand(&operand{position: position, name: name, dataType: dataType})
}

func (c *Cmd) addOption(o *option) {
	if c.shortOpts == nil {
		c.longOpts = make(map[string]*option)
		c.shortOpts = make(map[rune]*option)
	}

	//We will only add them to the command's
	//options if they are not zero value
	if o.short != 0 {
		c.shortOpts[o.short] = o
	}
	if o.long != "" {
		c.longOpts[o.long] = o
	}
}

func (c *Cmd) longOption(name string) (*option, bool) {
	if o, ok := c.longOpts[name]; ok {
		return o, true
	}
	return nil, false
}

func (c *Cmd) shortOption(name rune) (*option, bool) {
	o, ok := c.shortOpts[name]
	return o, ok
}

func (c *Cmd) getRequiredNames() map[string]struct{} {
	names := make(map[string]struct{})
	for longName, opt := range c.longOpts {
		if opt.required {
			names[longName] = struct{}{}
		}
	}
	for shortRune, opt := range c.shortOpts {
		if opt.required {
			names[string(shortRune)] = struct{}{}
		}
	}
	return names
}
func (c *Cmd) longOptions() map[string]*option {
	return c.longOpts
}
func (c *Cmd) sortedShortOptions() []*option {
	var runes []rune
	var names []string
	var opts []*option
	for _, name := range c.shortOpts {
		names = append(names, string(name.short))
	}

	sort.Strings(names)
	for _, name := range names {
		rs := getRunes(name)
		if len(rs) > 0 {
			runes = append(runes, rs[0])
		}
	}

	for _, short := range runes {
		if opt, ok := c.shortOpts[short]; ok {
			opts = append(opts, opt)
		}
	}

	return opts
}
func (c *Cmd) sortedOperands() []*operand {
	var pos []int
	var ops []*operand
	for p := range c.operands {
		pos = append(pos, p)
	}
	sort.Ints(pos)
	for p := range pos {
		if op, ok := c.operands[p]; ok {
			ops = append(ops, op)
		}
	}
	return ops
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
