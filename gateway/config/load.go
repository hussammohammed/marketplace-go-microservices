package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Load(cfgName string, cfgPaths ...string) error {
	// set config file type and name
	split := strings.Split(cfgName, ".")
	// if configuratin file name has extension then take if otherwise then the defult extenstion (yaml)
	if len(split) > 1 {
		fileExtenstion := split[len(split)-1]
		viper.SetConfigType(fileExtenstion)
		viper.SetConfigName(strings.Join(split[:len(split)-1], "."))
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName(cfgName)
	}
	// search for configurations at specified paths
	for _, p := range append(cfgPaths, "./config/") {
		if p != "" {
			viper.AddConfigPath(p)
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config %v", err)
	}
	return nil
}
