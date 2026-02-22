package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitConfig() {
	// 12-factor config: Env vars first
	viper.SetEnvPrefix("YOG")
	viper.AutomaticEnv()

	// Default config path: ~/.config/yog/
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		os.Exit(1)
	}

	configDir := filepath.Join(home, ".config", "yog")
	
	// Create config directory if it does not exist
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			fmt.Println("Could not create config directory:", err)
		}
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// File exists but was unreadable
			fmt.Println("Error reading config file:", err)
		}
		// If the file doesn't exist, we just rely on defaults/env vars
	}
}
