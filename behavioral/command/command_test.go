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

	if len(scheduler.queue) != 2 {
		t.Errorf("Expected queue of 2 commands, got %d", len(scheduler.queue))
	}

	scheduler.Run()

	if len(scheduler.queue) != 0 {
		t.Errorf("Expected queue to be empty after run, got %d", len(scheduler.queue))
	}

	if !cmd1.Executed {
		t.Error("Expected cmd1 to be executed")
	}

	if !cmd2.Executed {
		t.Error("Expected cmd2 to be executed")
	}
}
