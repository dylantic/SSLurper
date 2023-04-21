package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfiguration struct {
	Server struct {
		DebugMode bool   `yaml:"debugmode"`
		Socket    string `yaml:"socket"`
	} `yaml:"server"`
	Database struct {
		Type     string `yaml:"type"`
		Path     string `yaml:"path"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

var Config AppConfiguration

func init() {
	readFile(&Config)
}

func readFile(cfg *AppConfiguration) {
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
