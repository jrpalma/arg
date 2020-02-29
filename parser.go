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
	var cmdOptions map[string]string

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

		if len(model.args) == 0 && !match.Options.hasRequired() {
			cmd = match
			break
		}

		cmdOptions = model.getCmdOptions(match)
		if len(cmdOptions) == 0 {
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

	err := cmd.Exec(&execargs{options: cmdOptions})
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
	optionHelp := ""
	allOptions := cmd.Options.getOptions()
	reqOptions := make(map[string]*option)
	optOptions := make(map[string]*option)
	headLine := fmt.Sprintf("%v %v", pfx.str, cmd.Name)

	for _, option := range allOptions {
		if option.required {
			reqOptions[option.name] = option
		} else {
			optOptions[option.name] = option
		}
	}

	for _, reqOptionName := range p.sortedOptions(reqOptions) {
		reqOption := reqOptions[reqOptionName]
		headLine += fmt.Sprintf(" %v", reqOption.name)
		optionHelp += fmt.Sprintf("\t%v: %v", reqOption.name, reqOption.help)
	}

	if len(optOptions) > 0 {
		keys := p.sortedOptions(optOptions)
		headLine += fmt.Sprintf(" [%v]", strings.Join(keys, " "))
		for _, optOptionName := range p.sortedOptions(optOptions) {
			optOption := optOptions[optOptionName]
			optionHelp += fmt.Sprintf("\t%v: %v\n", optOption.name, optOption.help)
		}
	}

	p.print(headLine + "\n")
	if optionHelp != "" {
		p.print(optionHelp + "\n\n")
	}
}
func (p *Parser) print(msg string, args ...interface{}) {
	if p.output != nil {
		p.output.Write([]byte(fmt.Sprintf(msg, args...)))
	}
}
func (p *Parser) sortedOptions(options map[string]*option) []string {
	keys := make([]string, 0, len(options))
	for key := range options {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
func (p *Parser) sortedPrefixes(options map[string]*prefix) []string {
	keys := make([]string, 0, len(options))
	for key := range options {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
func (p *Parser) sortedCmds(options map[string]*Cmd) []string {
	keys := make([]string, 0, len(options))
	for key := range options {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
