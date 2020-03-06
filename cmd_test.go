package arg

import (
	"testing"
)

func TestReqCmd(t *testing.T) {
	c := &Cmd{}

	c.ReqInt64('a', "age", "")
	c.ReqUint64('p', "points", "")
	c.Option('h', "help", "")
	c.ReqBool('c', "count", "")
	c.ReqString('b', "branch", "")
	c.ReqFloat64('s', "salary", "")
	c.ReqEnum('t', "status",
		[]string{"enable", "disabled"}, "")

	if len(c.shortOpts) != 7 {
		t.Errorf("There should be 7 options")
	}
	if len(c.longOpts) != 7 {
		t.Errorf("There should be 7 options")
	}
}
func TestOptCmd(t *testing.T) {
	c := &Cmd{}

	c.OptInt64('a', "age", "")
	c.OptUint64('p', "points", "")
	c.Option('h', "help", "")
	c.OptBool('f', "force", "")
	c.OptString('b', "branch", "")
	c.OptFloat64('s', "salary", "")
	c.OptEnum('t', "status",
		[]string{"enable", "disabled"}, "")

	if len(c.shortOpts) != 7 {
		t.Errorf("There should be 7 options")
	}
	if len(c.longOpts) != 7 {
		t.Errorf("There should be 7 options")
	}
}
func TestOperand(t *testing.T) {
	c := &Cmd{}

	c.Operand(1, "url", String)
	c.Operand(2, "dir", String)

	if len(c.operands) != 2 {
		t.Errorf("There should be 2 operands")
	}

	for _, op := range c.operands {
		if op.dataType != String {
			t.Errorf("Operand data type should be String")
		}
	}

}
