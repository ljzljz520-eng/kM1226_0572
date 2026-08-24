package input

import "testing"

func TestParseCommand(t *testing.T) {
	command, err := Parse(" move 4 -2 ")
	if err != nil || command.Name != "move" || len(command.Args) != 2 {
		t.Fatalf("unexpected command: %+v %v", command, err)
	}
	dx, err := IntArg(command, 0)
	if err != nil || dx != 4 {
		t.Fatalf("unexpected argument: %d %v", dx, err)
	}
}

func TestParseRejectsEmpty(t *testing.T) {
	if _, err := Parse(" "); err == nil {
		t.Fatal("expected parse error")
	}
}
