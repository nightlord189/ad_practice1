package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//Config - main app's config
type Config struct {
	DbAddress  string
	DbPort     int
	DbName     string
	DbUser     string
	DbPassword string
	SQLPath    string
}

//Load - загружает конфиг по указанному пути
func Load(path string) *Config {
	config := Config{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("LoadConfig error: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err2 := decoder.Decode(&config)
	if err2 != nil {
		fmt.Println("LoadConfig error: ", err2)
	}
	return &config
}
