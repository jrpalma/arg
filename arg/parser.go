package arg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// ErrInvalidArgs Error to signal invalid arguments
var ErrInvalidArgs = errors.New("Invalid Arguments")

// NewParser creates a new Parser object. A Parser object is used
// to parse the arguments and execute the command that matches
// the arguments. A Parser will write to the writer on a failure
// if the writer is not nil.
func NewParser(writer io.Writer) *Parser {
	parser := &Parser{
		pfxs:   make(map[string]*prefix),
		output: writer,
	}

	return parser
}

// Parser represents a parser object used to parse the arguments and
// execute the command that matches the arguments.
type Parser struct {
	args   []string
	output io.Writer
	pfxs   map[string]*prefix
}

// AddCmd adds a command that will execute if the Parser's Parse
// function can match command against the arguments.
func (p *Parser) AddCmd(cmd *Cmd) {
	cmdPfx := newPrefix(cmd.Prefix)
	pfx, ok := p.pfxs[cmdPfx.str]

	if !ok {
		p.pfxs[cmdPfx.str] = cmdPfx
		cmdPfx.addCmd(cmd)
		return
	}

	pfx.addCmd(cmd)
}

// Parse parses the arguments and execute the appropriate command.
// It is assumed that args contains the name of the executable as
// the first element in the args slice. A Parser will return
// ErrInvalidArgs if len(args) < 2.
// A Parser will print to the writer the commands usage, or the any
// non-nil error that is returned by a comand.
//
// Argument exitOnFailure
//
//
// A Parser's Parse function will exit the program if: "the parser
// cannot find a command to execute, or the command failed to
// execute and returns a non nil error".
//
// A Parser's Parse function will return: "ErrInvalidARgs if the
// parser cannot find a command to execute, or the command's
// returned error"
func (p *Parser) Parse(exitOnError bool, args []string) error {
	var cmd *Cmd
	var prefix string
	var models []model
	var cmdArgs map[string]string

	if len(args) < 2 {
		return ErrInvalidArgs
	}

	execName := args[0]
	execArgs := args[1:]

	for i, token := range execArgs {
		models = append(models, model{
			pfx:  newPrefix(prefix),
			args: execArgs[i+1:],
			name: token,
		})
		prefix += " " + token
	}

	for _, model := range models {
		pfx, ok := p.pfxs[model.pfx.str]

		if !ok || !pfx.equal(model.pfx) {
			continue
		}

		match, ok := pfx.getCmd(model.name)
		if !ok {
			continue
		}

		if len(model.args) == 0 && !match.Flags.hasReq() {
			cmd = match
			break
		}

		cmdArgs = model.getCmdArgs(match)
		if len(cmdArgs) == 0 {
			continue
		}

		cmd = match
		break
	}

	if cmd == nil {
		p.usage(execName)
		if !exitOnError {
			return ErrInvalidArgs
		}
		os.Exit(1)
	}

	err := cmd.Exec(&execargs{kvp: cmdArgs})
	if err != nil {
		p.print(err.Error())
		if !exitOnError {
			return err
		}
		os.Exit(1)
	}

	return nil
}
func (p *Parser) usage(exec string) {
	p.print("Usage: %v\n", exec)
	for _, prefixKey := range p.sortedPrefixes(p.pfxs) {
		prefix := p.pfxs[prefixKey]
		for _, cmdKey := range p.sortedCmds(prefix.cmds) {
			cmd := prefix.cmds[cmdKey]
			p.cmdUsage(prefix, cmd)
		}
	}

}
func (p *Parser) cmdUsage(pfx *prefix, cmd *Cmd) {
	flagHelp := ""
	allFlags := cmd.Flags.getFlags()
	reqFlags := make(map[string]*flag)
	optFlags := make(map[string]*flag)
	headLine := fmt.Sprintf("%v %v", pfx.str, cmd.Name)

	for _, flag := range allFlags {
		if flag.Required() {
			reqFlags[flag.Name()] = flag
		} else {
			optFlags[flag.Name()] = flag
		}
	}

	for _, reqFlagName := range p.sortedFlags(reqFlags) {
		reqFlag := reqFlags[reqFlagName]
		headLine += fmt.Sprintf(" %v", reqFlag.Name())
		flagHelp += fmt.Sprintf("\t%v: %v", reqFlag.Name(), reqFlag.Help())
	}

	if len(optFlags) > 0 {
		keys := p.sortedFlags(optFlags)
		headLine += fmt.Sprintf(" [%v]", strings.Join(keys, " "))
		for _, optFlagName := range p.sortedFlags(optFlags) {
			optFlag := optFlags[optFlagName]
			flagHelp += fmt.Sprintf("\t%v: %v\n", optFlag.Name(), optFlag.Help())
		}
	}

	p.print(headLine + "\n")
	if flagHelp != "" {
		p.print(flagHelp + "\n\n")
	}
}
func (p *Parser) print(msg string, args ...interface{}) {
	if p.output != nil {
		p.output.Write([]byte(fmt.Sprintf(msg, args...)))
	}
}
func (p *Parser) sortedFlags(flags map[string]*flag) []string {
	keys := make([]string, 0, len(flags))
	for key := range flags {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
func (p *Parser) sortedPrefixes(flags map[string]*prefix) []string {
	keys := make([]string, 0, len(flags))
	for key := range flags {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
func (p *Parser) sortedCmds(flags map[string]*Cmd) []string {
	keys := make([]string, 0, len(flags))
	for key := range flags {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
