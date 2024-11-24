package constant

type Command string

const (
	HelpCmd     Command = "-h"
	VersionCmd  Command = "-v"
	ChatCmd     Command = "-C"
	ShutdownCmd Command = "-shutdown"
	RandCmd     Command = "-rand"
)
