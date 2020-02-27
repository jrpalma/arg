package arg

import (
	"testing"
)

func TestPrefix(t *testing.T) {

	pfxStr := "box user -t   "

	pfx := newPrefix(pfxStr)
	if pfx.count != 3 {
		t.Errorf("The prefix count should be 3 and not %v", pfx.count)
	}
	if pfx.str != "box user -t" {
		t.Errorf("Incorrect str(%v)", pfx.str)
	}
	if pfx.len != len("box user -t") {
		t.Errorf("Incorrect length  %v", pfx.len)
	}

	pfx2 := newPrefix(pfxStr)
	if !pfx.equal(pfx2) {
		t.Errorf("Prefixes should be equal")
	}
	if pfx.equal(newPrefix("a b")) {
		t.Errorf("Prefixes should not be equal")
	}
	if pfx.equal(newPrefix("box user -tt  ")) {
		t.Errorf("Prefixes should not be equal")
	}
	if pfx.equal(newPrefix("box user -x  ")) {
		t.Errorf("Prefixes should not be equal")
	}

	pfx.addCmd(&Cmd{Prefix: pfxStr, Name: "cmd", Help: "cmd"})
	cmd, ok := pfx.getCmd("cmd")
	if !ok {
		t.Errorf("Prefix should have a command named cmd")
	}

	if cmd.Name != "cmd" || cmd.Help != "cmd" {
		t.Errorf("Prefix has wrong command")
	}

}
