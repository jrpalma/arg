package arg

import (
	"testing"
)

func TestInvalidCmdOptions(t *testing.T) {
	var fs2 CmdOptions
	fs := getCmdOptions()

	_, ok := fs2.opts["bad"]
	if ok {
		t.Errorf("Option should not exist")
	}
	set := fs.getOptions()
	if len(set) != 14 {
		t.Errorf("Incorrect number of flags")
	}
}

func TestExist(t *testing.T) {
	fs := getCmdOptions()

	_, exist := fs.opts["rs"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["rb"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["ri64"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["rui64"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["rf64"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["renum"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["os"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["ob"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["oi64"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["oui64"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["of64"]
	if !exist {
		t.Errorf("Option does not exist")
	}
	_, exist = fs.opts["oenum"]
	if !exist {
		t.Errorf("Option does not exist")
	}
}

func TestCmdOptions(t *testing.T) {
	fs := getCmdOptions()

	f, _ := fs.opts["rs"]
	if f.optionType != stringOption {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["rb"]
	if f.optionType != boolOption {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["ri64"]
	if f.optionType != int64Option {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["rui64"]
	if f.optionType != uint64Option {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["rf64"]
	if f.optionType != float64Option {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["renum"]
	if f.optionType != enumOption {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["os"]
	if f.optionType != stringOption {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["ob"]
	if f.optionType != boolOption {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["oi64"]
	if f.optionType != int64Option {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["oui64"]
	if f.optionType != uint64Option {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["of64"]
	if f.optionType != float64Option {
		t.Errorf("Option is incorrect optionTypee")
	}
	f, _ = fs.opts["oenum"]
	if f.optionType != enumOption {
		t.Errorf("Option is incorrect optionTypee")
	}
}

func getCmdOptions() *CmdOptions {
	fs := &CmdOptions{}
	fs.ReqString("rs", "", "rs")
	fs.ReqBool("rb", "", "rb")
	fs.ReqInt64("ri64", "", "ri64")
	fs.ReqUint64("rui64", "", "rui64")
	fs.ReqFloat64("rf64", "", "rf64")
	fs.ReqEnum("renum", "", []string{"e", "n", "u", "m"}, "renum")

	fs.ReqArgs("rags", "", []string{"a", "b"}, "rags")
	fs.OptString("os", "", "os")
	fs.OptBool("ob", "", "ob")
	fs.OptInt64("oi64", "", "oi64")
	fs.OptUint64("oui64", "", "oui64")
	fs.OptFloat64("of64", "", "of64")
	fs.OptEnum("oenum", "", []string{"e", "n", "u", "m"}, "oenum")
	fs.OptArgs("oags", "", []string{"a", "b"}, "oags")
	return fs
}
