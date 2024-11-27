/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/anjude/xc/third_party/wenxin"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "在命令行工具和AI对话",
	Long:  `在命令行工具和AI连续对话，quit退出!`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
