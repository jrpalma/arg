package arg

import (
	"sort"
	"strings"
)

func newPrefix(s string) *prefix {
	prefix := &prefix{}
	tokens := strings.Split(s, " ")

	for _, token := range tokens {
		trimmed := strings.TrimSpace(token)
		if len(trimmed) == 0 {
			continue
		}
		prefix.items = append(prefix.items, trimmed)
	}

	prefix.str = strings.Join(prefix.items, " ")
	prefix.cmds = make(map[string]*Cmd)
	prefix.count = len(prefix.items)
	prefix.len = len(prefix.str)

	return prefix
}

type prefix struct {
	count int
	len   int
	str   string
	items []string
	cmds  map[string]*Cmd
}

func (p *prefix) equal(rhs *prefix) bool {
	if p.count != rhs.count {
		return false
	}
	if p.len != rhs.len {
		return false
	}
	if p.str != rhs.str {
		return false
	}
	return true
}

func (p *prefix) addCmd(cmd *Cmd) {
	p.cmds[cmd.Name] = cmd
}

func (p *prefix) getCmd(name string) (*Cmd, bool) {
	cmd, ok := p.cmds[name]
	return cmd, ok
}

func (p *prefix) sortedCmds() []*Cmd {
	var cmds []*Cmd
	var names []string
	for name := range p.cmds {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, cmdName := range names {
		if cmd, ok := p.cmds[cmdName]; ok {
			cmds = append(cmds, cmd)
		}

	}
	return cmds
}
