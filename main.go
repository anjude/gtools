package main

import (
	"fmt"
	"github.com/anjude/gtools/handler"
	"os"
)

func main() {
	// 解析命令行参数
	args := os.Args[1:]
	if len(args) == 0 {
		// 如果没有指定任何参数，则显示帮助信息
		fmt.Println("Usage: gtools [options]")
		fmt.Println("Use -h for more information.")
	}

	// 加载配置
	//if err := config.InitConfig(); err != nil {
	//}

	args = append([]string{"-c"}, "nihao")

	// 执行命令
	handler.Execute(args)
}
