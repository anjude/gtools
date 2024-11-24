package rand

import (
	"fmt"
	"github.com/anjude/gtools/handler/base"
	"math/rand"
	"strconv"
)

var _ base.IHandler = Handler{}

type Handler struct {
	Command string
	Desc    string
}

func (s Handler) GetCommand() (command string) {
	return s.Command
}

func (s Handler) GetDesc() (desc string) {
	return s.Desc
}

func (s Handler) GetArgs(args []string) (curArgs []string, nextArgs []string) {
	return args, nextArgs
}

func (s Handler) Handle(args []string) {
	// args: 随机字符串位数，默认10；随机字符串类型，默认数字（使用字符集不一样）
	numberCharSet := "0123456789"
	stringCharSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	allCharSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+-=[]{}|;':,.<>/?`~"
	length := int64(10)
	usedCharSet := stringCharSet

	// 2. 随机字符串类型
	if len(args) >= 1 {
		length, _ = strconv.ParseInt(args[0], 10, 64)
	}
	if len(args) >= 2 {
		switch args[1] {
		case "number":
			usedCharSet = numberCharSet
		case "string":
			usedCharSet = stringCharSet
		case "all":
			usedCharSet = allCharSet
		}
	}

	// 3. 生成随机字符串
	randStr := make([]byte, length)
	for i := 0; i < int(length); i++ {
		randStr[i] = usedCharSet[rand.Intn(len(usedCharSet))]
	}

	// 4. 输出
	fmt.Println(string(randStr))
}
