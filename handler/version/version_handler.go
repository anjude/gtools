package version

import (
	"fmt"
	"github.com/anjude/gtools/config"
	"github.com/anjude/gtools/handler/base"
)

var _ base.IHandler = Handler{}

type Handler struct {
	Command string
	Desc    string
}

func (v Handler) GetCommand() string {
	return v.Command
}

func (v Handler) GetDesc() string {
	return v.Desc
}

func (v Handler) GetArgs(args []string) (curArgs []string, nextArgs []string) {
	return args, nextArgs
}

func (v Handler) Handle(strings []string) {
	fmt.Println(config.BotConf.Version)
}
