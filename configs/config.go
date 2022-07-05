package configs

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BindAddr int `json:"bind_addr,omitempty"`

	HostDB     string `json:"host_db,omitempty"`
	PortDB     int    `json:"port_db,omitempty"`
	NameDB     string `json:"name_db,omitempty"`
	UserDB     string `json:"user_db,omitempty"`
	PasswordDB string `json:"password_db,omitempty"`
}

func LoadConfig(path string) (*Config, error) {
	config := Config{
		BindAddr: 8080,

		HostDB:     "localhost",
		PortDB:     5432,
		NameDB:     "postgres",
		UserDB:     "postgres",
		PasswordDB: "postgres",
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
