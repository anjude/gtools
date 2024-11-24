package constant

type Command string

const (
	HelpCmd     Command = "-h"
	VersionCmd  Command = "-v"
	ChatCmd     Command = "-c"
	ShutdownCmd Command = "-shutdown"
	RandCmd     Command = "-rand"
)

func (c Command) String() string {
	return string(c)
}
