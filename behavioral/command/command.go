package command

import "fmt"

// Command interface represents any executable task.
type Command interface {
	Execute() error
}

// EmailSender is a concrete command.
type EmailSender struct {
	To      string
	Subject string
}

func (c *EmailSender) Execute() error {
	fmt.Printf("Sending email to %s with subject '%s'\n", c.To, c.Subject)
	return nil
}

// ReportGenerator is another concrete command.
type ReportGenerator struct {
	ReportName string
}

func (c *ReportGenerator) Execute() error {
	fmt.Printf("Generating report: %s\n", c.ReportName)
	return nil
}

// Scheduler queues and runs commands.
type Scheduler struct {
	queue []Command
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Add(cmd Command) {
	s.queue = append(s.queue, cmd)
}

func (s *Scheduler) Run() {
	for _, cmd := range s.queue {
		if err := cmd.Execute(); err != nil {
			fmt.Printf("error executing command: %v\n", err)
		}
	}
	s.queue = nil
}

func Run() {
	scheduler := NewScheduler()

	scheduler.Add(&EmailSender{
		To:      "user@email.com",
		Subject: "Email Subject",
	})

	scheduler.Add(&ReportGenerator{
		ReportName: "Report XYZ",
	})

	scheduler.Run()
}
