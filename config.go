package dashing

import (
	"encoding/json"
    "fmt"
	"os"
)

type Config struct {
    DashingPath string
	DefaultDashboard string
    HackToken string
    GlobalToken string
    WidgetTokens map[string]string
}

func NewConfig(path string) *Config {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	config := new(Config)
	if err := decoder.Decode(&config); err != nil {
        fmt.Printf("Parsing config file error, %s", err.Error())
    }
	return config
}
