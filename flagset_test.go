package arg

import (
	"testing"
)

func TestInvalidFlagSet(t *testing.T) {
	var fs2 FlagSet
	fs := getFlagSet()

	_, ok := fs2.kvp["bad"]
	if ok {
		t.Errorf("Flag should not exist")
	}
	set := fs.getFlags()
	if len(set) != 14 {
		t.Errorf("Incorrect number of flags")
	}
}

func TestExist(t *testing.T) {
	fs := getFlagSet()

	_, exist := fs.kvp["rs"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["rb"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["ri64"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["rui64"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["rf64"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["renum"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["os"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["ob"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["oi64"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["oui64"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["of64"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
	_, exist = fs.kvp["oenum"]
	if !exist {
		t.Errorf("Flag does not exist")
	}
}

func TestFlagSet(t *testing.T) {
	fs := getFlagSet()

	f, _ := fs.kvp["rs"]
	if f.Type() != stringFlag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["rb"]
	if f.Type() != boolFlag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["ri64"]
	if f.Type() != int64Flag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["rui64"]
	if f.Type() != uint64Flag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["rf64"]
	if f.Type() != float64Flag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["renum"]
	if f.Type() != enumFlag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["os"]
	if f.Type() != stringFlag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["ob"]
	if f.Type() != boolFlag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["oi64"]
	if f.Type() != int64Flag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["oui64"]
	if f.Type() != uint64Flag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["of64"]
	if f.Type() != float64Flag {
		t.Errorf("Flag is incorrect type")
	}
	f, _ = fs.kvp["oenum"]
	if f.Type() != enumFlag {
		t.Errorf("Flag is incorrect type")
	}
}

func getFlagSet() *FlagSet {
	fs := &FlagSet{}
	fs.ReqString("rs", "", "rs")
	fs.ReqBool("rb", "", "rb")
	fs.ReqInt64("ri64", "", "ri64")
	fs.ReqUint64("rui64", "", "rui64")
	fs.ReqFloat64("rf64", "", "rf64")
	fs.ReqEnum("renum", "", []string{"e", "n", "u", "m"}, "renum")

	fs.ReqArgs("rags", "", 3, "rags")
	fs.OptString("os", "", "os")
	fs.OptBool("ob", "", "ob")
	fs.OptInt64("oi64", "", "oi64")
	fs.OptUint64("oui64", "", "oui64")
	fs.OptFloat64("of64", "", "of64")
	fs.OptEnum("oenum", "", []string{"e", "n", "u", "m"}, "oenum")
	fs.OptArgs("oags", "", 3, "oags")
	return fs
}
