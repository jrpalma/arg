package arg

import (
	"testing"
)

func TestFlag(t *testing.T) {
	f := &flag{}
	f.name = "flag"
	f.help = "help"
	f.count = uint(3)
	f.req = true

	if f.Name() != "flag" {
		t.Errorf("The unexpected flag name %v", f.Name())
	}
	if f.Help() != "help" {
		t.Errorf("The unexpected flag help %v", f.Help())
	}
	if !f.Required() {
		t.Errorf("The flag should be required")
	}

	if f.Count() != 3 {
		t.Errorf("The count should be 3")
	}

	if f.Type() != unknownFlag {
		t.Errorf("Expected unknown flag type")
	}
}
