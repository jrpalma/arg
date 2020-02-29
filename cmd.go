package arg

// Cmd represents a command that will be executed by a Parser.
// The command's Exec method is executed if a Parser can match
// the command's prefix, name, and flags.
type Cmd struct {
	// The command's prpefix. The prefix is the string that comes
	//before the command name and its flags.
	Prefix string
	// The command's name. The name must be unique per prefix.
	// If two commands with the same prefix and name are added
	// to a Parser object, the later will replace the the earlier.
	Name string
	// The command's help string. The command will display the help
	// string to document the command.
	Help string
	// The function that gets executed if the arguments match the command.
	Exec func(ExecArgs) error
	// The command's flags. Flags can be required or optional.
	Flags CmdFlagSet
	// The command's arguments. Commands can have zero or more arguments.
	Args CmdArgSet
}
