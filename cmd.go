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

	//private stuff
	opts map[rune]*option
}

// Option adds an option. Short is a one letter option.
// Long is the long version of the option. Help is displayed
// when the usage message is printed.
func (c *Cmd) Option(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help})
}

// ReqString adds a required string option. Short is a one letter option.
// Long is the long version of the option. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqString(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: String})
}

// ReqBool adds a required bool option. Short is a one letter option.
// Long is the long version of the option. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqBool(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Bool})
}

// ReqInt64 adds a required int64 option. Short is a one letter option.
// Long is the long version of the option. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqInt64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Int64})
}

// ReqUint64 adds a required uint64 option. Short is a one letter option.
// Long is the long version of the option. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqUint64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Uint64})
}

// ReqFloat64 adds a required float64 option. Short is a one letter option.
// Long is the long version of the option. Help is displayed
// when the usage message is printed.
func (c *Cmd) ReqFloat64(short rune, long string, help string) {
	c.addOption(&option{short: short, long: long, help: help,
		required: true, dataType: Float64})
}

// ReqEnum adds a required enum option. Short is a one letter option.
// Long is the long version of the option. Valid are the valid values
// for this option. Help is displayed when the usage message is printed.
func (c *Cmd) ReqEnum(short rune, long string, valid []string, help string) {
	c.addOption(&option{short: short, long: long, help: help, valid: valid,
		required: true, dataType: enum})
}

func (c *Cmd) addOption(o *option) {
	if c.opts == nil {
		c.opts = make(map[rune]*option)
	}
	c.opts[o.short] = o
}
