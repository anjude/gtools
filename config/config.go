package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os/user"
)

var BotConf Config

type ChatGPTConfig struct {
	ApiKey     string `mapstructure:"apiKey"`
	Proxy      string `mapstructure:"proxy"`
	ChatRounds int    `mapstructure:"chatRounds"`
}

type Config struct {
	ChatGPT ChatGPTConfig `mapstructure:"chatGPTConfig"`
	Version string        `mapstructure:"version"`
}

func InitConfig() error {
	currentUser, _ := user.Current()
	configFile := currentUser.HomeDir + "/.terminalx1/config.yaml"
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read config file err: %v\n", err)
		setDefaultConfig()
		return nil
	}

	if err := viper.Unmarshal(&BotConf); err != nil {
		fmt.Printf("unmarshal config file err: %v\n", err)
		return err
	}
	return nil
}

func setDefaultConfig() {
	log.Printf("use default config")
	BotConf = Config{
		ChatGPT: ChatGPTConfig{
			ApiKey:     "",
			Proxy:      "",
			ChatRounds: 3,
		},
		Version: "0.0.1",
	}
}
