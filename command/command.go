package command

type Command struct {
	command string
}

func newCommand(name string) Command {
	return Command{
		command: name,
	}
}
