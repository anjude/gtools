/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/spf13/cobra"
)

// randCmd represents the rand command
var randCmd = &cobra.Command{
	Use:   "rand",
	Short: "生成随机字符串",
	Long: `10位随机数字 -- xc rand 10 number
20位随机数字/字母 -- xc rand 20 string
10位随机字符 -- xc rand 10 all`,
	Run: func(cmd *cobra.Command, args []string) {
		// args: 随机字符串位数，默认10；随机字符串类型，默认数字（使用字符集不一样）
		numberCharSet := "0123456789"
		stringCharSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		allCharSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+-=[]{}|;':,.<>/?`~"
		length := int64(6)
		usedCharSet := numberCharSet

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
	},
}

func init() {
	rootCmd.AddCommand(randCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//randCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//randCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
