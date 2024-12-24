package conf

import (
	"fmt"
	"github.com/ZRothschild/ldp/infrastr/static/config"
	"github.com/spf13/viper"
)

var (
	Conf = new(config.Config)
)

func init() {
	viper.SetConfigName(config.FileName) // name of config file (without extension)
	viper.SetConfigType(config.FileType) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(config.FilePath) // path to look for the config file in
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("conf init viper ReadInConfig fatal error: %w", err))
	}

	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("conf init viper Unmarshal fatal error: %w", err))
	}
}
