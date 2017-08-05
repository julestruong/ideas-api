package app 

import (
	"encoding/json"
	"fmt"
	"os"
)

type JWTConfig struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

type Config struct {
	JWT JWTConfig `json:"jwt"`
}

func GetConfig() Config {
	file, _ := os.Open("app/config.json")
	decoder := json.NewDecoder(file)

	config := Config{}
	err := decoder.Decode(&config)

	if err != nil {
		fmt.Printf("Error has occured :", err)
	}

	return config
}