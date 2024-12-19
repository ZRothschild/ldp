package conf

import (
	"fmt"
	"github.com/ZRothschild/ldp/infrastr/env"
	"github.com/spf13/viper"
)

var (
	Conf = new(env.Config)
)

func init() {
	viper.SetConfigName(env.ConfigFileName) // name of config file (without extension)
	viper.SetConfigType(env.ConfigFileType) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(env.ConfigFilePath) // path to look for the config file in
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("conf init viper ReadInConfig fatal error: %w", err))
	}

	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("conf init viper Unmarshal fatal error: %w", err))
	}
}
