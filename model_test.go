package arg

import (
	"testing"
)

func TestAllRequired(t *testing.T) {
	m := model{}
	m.name = "show"
	m.pfx = newPrefix("box user")
	m.args = []string{"id", "3", "city", "houston", "admin", "true", "zipcode", "77777", "minsalary", "50000", "experience", "3", "5"}

	cmd := &Cmd{}
	cmd.Name = "show"
	cmd.Prefix = "box user"
	cmd.Help = "Show user name"
	cmd.Options.ReqInt64("id", "", "The ID of the user to show")
	cmd.Options.ReqString("city", "", "The name of the city")
	cmd.Options.ReqBool("admin", "", "Include admins in the result")
	cmd.Options.ReqUint64("zipcode", "", "The zipcode of the city to match")
	cmd.Options.ReqFloat64("minsalary", "", "The minimum salary to  use")
	cmd.Options.ReqArgs("experience", "", []string{"min", "max"},
		"The <min years> <max years> years of experience")

	//Verify all parsed args
	args := m.getCmdOptions(cmd)
	verify := func(name string, expected string) {
		v, ok := args[name]
		if !ok || v != expected {
			t.Errorf("Invalid %v %v, expected %v", name, v, expected)
		}
	}

	verify("id", "3")
	verify("city", "houston")
	verify("admin", "true")
	verify("zipcode", "77777")
	verify("minsalary", "50000")
	verify("experience", "3,5")

}

func TestInvalid(t *testing.T) {
	m := model{}
	m.name = "show"
	m.pfx = newPrefix("box user")
	m.args = []string{"id", "3"}

	cmd := &Cmd{}
	cmd.Name = "show"
	cmd.Prefix = "box user"
	cmd.Help = "Show user name"
	cmd.Options.ReqInt64("id", "", "The ID of the user to show")
	cmd.Options.ReqString("city", "", "The name of the city")

	args := m.getCmdOptions(cmd)
	if len(args) != 0 {
		t.Errorf("getCmdOptions should have return no args")
	}

	//Missing city flag arg
	m.args = []string{"id", "3", "city"}
	args = m.getCmdOptions(cmd)
	if len(args) != 0 {
		t.Errorf("getCmdOptions should have return no args")
	}

	//Invalid flag
	m.args = []string{"id", "3", "invalidflag", "invaliddarg"}
	args = m.getCmdOptions(cmd)
	if len(args) != 0 {
		t.Errorf("getCmdOptions should have return no args")
	}

	//Invalid number of args
	cmd.Options.ReqArgs("experience", "", []string{"min", "max"},
		"The <min years> <max years> years of experience")
	m.args = []string{"id", "3", "", "3"}
	args = m.getCmdOptions(cmd)
	if len(args) != 0 {
		t.Errorf("getCmdOptions should have return no args")
	}

	ret := m.useOptions(nil, nil)
	if len(ret) != 0 {
		t.Errorf("useOptions should have return no args")
	}
}
