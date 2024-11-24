package chat

import (
	"bufio"
	"fmt"
	"github.com/anjude/gtools/handler/base"
	"github.com/anjude/gtools/third_party/wenxin"
	"math/rand"
	"os"
)

var _ base.IHandler = Handler{}

type Handler struct {
	Command string
	Desc    string
}

func (c Handler) GetCommand() (command string) {
	return c.Command
}

func (c Handler) GetDesc() (desc string) {
	return c.Desc
}

func (c Handler) GetArgs(args []string) (curArgs []string, nextArgs []string) {
	return args, nextArgs
}

func (c Handler) Handle(args []string) {
	fmt.Println(`[Please enter 'quit' to exit or enter any other input to continue the conversation.]`)
	userId := fmt.Sprintf("%d", rand.Intn(10000000))
	if len(args) != 0 {
		reply, err := wenxin.GetWenXinReply(args[0], userId)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(reply)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n> ")
		scanner.Scan()
		text := scanner.Text()
		if text == "quit" {
			break
		}
		reply, err := wenxin.GetWenXinReply(text, userId)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(reply)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Goodbye!")
}
