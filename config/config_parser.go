package config


import "github.com/BurntSushi/toml"

var APIConfig = Config{
}

func Parse(configFile string) (err error) {
	_, err = toml.DecodeFile(configFile, &APIConfig)
	return
}