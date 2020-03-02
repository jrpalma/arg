package arg

import (
	"testing"
)

func TestCmd(t *testing.T) {
	c := &Cmd{}

	c.ReqInt64('a', "age", "")
	c.ReqUint64('p', "points", "")
	c.Option('h', "help", "")
	c.ReqBool('f', "force", "")
	c.ReqBool('c', "count", "")
	c.ReqString('b', "branch", "")
	c.ReqFloat64('s', "salary", "")
	c.ReqEnum('t', "status",
		[]string{"enable", "disabled"}, "")

	if len(c.opts) != 8 {
		t.Errorf("There should be 8 options")
	}
}
