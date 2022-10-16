package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

func LoadTomlConfig(filePath string) error {
	config = NewDefaultConfig()

	if _, err := toml.DecodeFile(filePath, config); err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}
	return nil
}

func LoadEnvConfig() error {
	config = NewDefaultConfig()

	if err := env.Parse(config); err != nil {
		return err
	}
	return nil
}

func localGlobal() {
	conn, err := config.MySQL.getDBConn()
	if err != nil {
		panic(err)
	}
	db = conn
}
