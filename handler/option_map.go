package handler

import (
	"fmt"
	"github.com/anjude/gtools/constant"
	"github.com/anjude/gtools/handler/base"
	"github.com/anjude/gtools/handler/chat"
	"github.com/anjude/gtools/handler/rand"
	"github.com/anjude/gtools/handler/shutdown"
	"github.com/anjude/gtools/handler/version"
)

// OptionMap 在这里维护指令集
var OptionMap = map[constant.Command]base.IHandler{
	constant.HelpCmd: Handler{
		Command: "-h",
		Desc:    "show help info",
	},
	constant.VersionCmd: version.Handler{
		Command: "-v",
		Desc:    "show version info",
	},
	constant.ChatCmd: chat.Handler{
		Command: "-q",
		Desc:    "chat with the bot",
	},
	constant.ShutdownCmd: shutdown.Handler{
		Command: "-shutdown",
		Desc:    "shutdown this computer",
	},
	constant.RandCmd: rand.Handler{
		Command: "-rand [len] [type]",
		Desc:    "rand string/number",
	},
}

func Execute(args []string) {
	if len(args) == 0 {
		return
	}
	handler, ok := OptionMap[constant.Command(args[0])]
	if !ok {
		fmt.Printf("command [%v] not found\n", args[0])
		return
	}
	curArgs, _ := handler.GetArgs(args[1:])
	handler.Handle(curArgs)
	//Execute(nextArgs)
}
