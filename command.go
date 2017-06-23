package main

type Command int

const (
	InitCmd Command = iota
	NewCmd
	GenerateCmd
	NoneCmd
)

func getCommand(name string) Command {
	switch name {
	case "init":
		return InitCmd
	case "new":
		return NewCmd
	case "generate":
		fallthrough
	case "g":
		return GenerateCmd
	default:
		return NoneCmd
	}
	return NoneCmd
}

func RunCommand(args []string) {
	cmd := getCommand(args[0])
	switch cmd {
	case InitCmd:
		InitBlog()
	case NewCmd:
		Create(args[1:])
	case GenerateCmd:
		Generate()
	default:
		Help()
	}
}
