package config

import (
	"fmt"
	"os"
	"strings"

	ini "gopkg.in/ini.v1"
)

// getenv returns the value of the given env variable "key"
// and when its value is empty, it returns the default "fallback" value.
func getenv(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		value = fallback
	}

	return value
}

// Read the configuration file and assign values from .env
func configure() bool {
	var path = os.Getenv("GOPATH") + "/src/github.com/trellis/ranking-service/.env"

	if exists(path) {
		// read config as ini file
		cfg, err := ini.InsensitiveLoad(path)
		if err != nil {
			fmt.Println(err)
		}

		// read the default section (not using sections in config)
		s, err := cfg.GetSection("")
		if err != nil {
			fmt.Println(err)
		}

		// set the key/value pairs on the env to be read by Viper
		for k, v := range s.KeysHash() {
			os.Setenv(strings.ToUpper(k), v)
		}

		return true
	}

	return false
}

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}
