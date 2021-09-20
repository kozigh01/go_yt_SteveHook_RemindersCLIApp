package client

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type CmdFunc func(string) error
type CmdFuncFactory func() CmdFunc

type BackendHTTPClient interface {
	create(title, message string, duration time.Duration) ([]byte, error)
	edit(id, title, message string, duration time.Duration) ([]byte, error)
	fetch(ids []string) ([]byte, error)
	delete(ids []string) error
	healthy(host string) bool
}

func NewSwitch(uri string) Switch {
	httpClient := NewHTTPClient(uri)
	s := Switch{
		client:        httpClient,
		backendAPIUrl: uri,
		commands:      map[string]CmdFuncFactory{},
	}
	s.commands = map[string]CmdFuncFactory{
		"create": s.create,
		"edit":   s.edit,
		"fetch":  s.fetch,
		"delete": s.delete,
		"health": s.health,
	}
	return s
}

type Switch struct {
	client        BackendHTTPClient
	backendAPIUrl string
	commands      map[string]CmdFuncFactory
}

func (s Switch) Switch() error {
	cmdName := os.Args[1]
	cmd, ok := s.commands[cmdName]
	if !ok {
		return fmt.Errorf("invalid command: %q", cmdName)
	}
	return cmd()(cmdName)
}

func (s Switch) Help() {
	var help string
	for name := range s.commands {
		help += "    " + name + "\t --help\n"
	}
	fmt.Printf("Usage of %s\n <commands> [<args>]\n%s", os.Args[0], help)
}

func (s Switch) create() CmdFunc {
	return func(cmd string) error {
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(createCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(createCmd); err != nil {
			return err
		}

		res, err := s.client.create(*t, *m, *d)
		if err != nil {
			return wrapError("could not create reminder", err)
		}

		fmt.Printf("reminder created successfully:\n%s", string(res))
		return nil
	}
}

func (s Switch) edit() CmdFunc {
	return func(string) error {
		fmt.Println("edit reminder")
		return nil
	}
}

func (s Switch) fetch() CmdFunc {
	return func(string) error {
		fmt.Println("fetch reminder")
		return nil
	}
}

func (s Switch) delete() CmdFunc {
	return func(string) error {
		fmt.Println("delete reminder")
		return nil
	}
}

func (s Switch) health() CmdFunc {
	return func(string) error {
		fmt.Println("health reminder")
		return nil
	}
}

func (s Switch) reminderFlags(f *flag.FlagSet) (*string, *string, *time.Duration) {
	t, m, d := "", "", time.Duration(0)
	f.StringVar(&t, "title", "", "Reminder title")
	f.StringVar(&t, "t", "", "Reminder title")
	f.StringVar(&m, "message", "", "Reminder message")
	f.StringVar(&m, "m", "", "Reminder message")
	f.DurationVar(&d, "duration", time.Duration(0), "Reminder time")
	f.DurationVar(&d, "d", time.Duration(0), "Reminder time")
	return &t, &m, &d
}

func (s Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return wrapError(fmt.Sprintf("could not parse %q command flags", cmd.Name()), err)
	}
	return nil
}

func (s Switch) checkArgs(minArgs int) error {
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}
	if len(os.Args)-2 < minArgs {
		fmt.Printf(
			"incorrect use of %s\n%s %s",
			os.Args[1], os.Args[0], os.Args[1])
		return fmt.Errorf(
			"%s expects at least %d arg(s), %d provided",
			os.Args[1], minArgs, len(os.Args)-2)
	}
	return nil
}
