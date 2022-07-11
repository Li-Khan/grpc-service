package configs

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type Config struct {
	BindAddr int `json:"bind_addr,omitempty"`

	HostDB     string `json:"host_db,omitempty"`
	PortDB     int    `json:"port_db,omitempty"`
	NameDB     string `json:"name_db,omitempty"`
	UserDB     string `json:"user_db,omitempty"`
	PasswordDB string `json:"password_db,omitempty"`

	ClientAddr int `json:"client_addr,omitempty"`
}

const cfgPath string = "./configs/config.json"

var cfg *Config
var onceCfg sync.Once

func GetConfig() *Config {
	onceCfg.Do(func() {
		var err error

		cfg, err = LoadConfig(cfgPath)
		if err != nil {
			panic(err)
		}
	})
	return cfg
}

func LoadConfig(path string) (*Config, error) {
	config := Config{
		BindAddr: 8080,

		HostDB:     "localhost",
		PortDB:     5432,
		NameDB:     "postgres",
		UserDB:     "postgres",
		PasswordDB: "postgres",

		ClientAddr: 8080,
	}

	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
