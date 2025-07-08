package command

import "testing"

type MockCommand struct {
	Executed bool
}

func (m *MockCommand) Execute() error {
	m.Executed = true
	return nil
}

func TestScheduler_Run(t *testing.T) {
	scheduler := NewScheduler()

	cmd1 := &MockCommand{}
	cmd2 := &MockCommand{}
	scheduler.Add(cmd1)
	scheduler.Add(cmd2)
	scheduler.Run()

	if !cmd1.Executed {
		t.Error("Expected cmd1 to be executed")
	}

	if !cmd2.Executed {
		t.Error("Expected cmd2 to be executed")
	}
}
