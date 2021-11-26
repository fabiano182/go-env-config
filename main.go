package main

import (
	"fmt"

	"github.com/fabiano182/go-env-config/config"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

func main() {
	//From a file
	cfg_file := config.ConfigHelper("config/config.yml")

	fmt.Printf("Host: %v, Port: %v \n", cfg_file.Server.Host, cfg_file.Server.Port)

	//From Environments Variables
	cfg_env := config.ConfigHelperEnv()

	fmt.Printf("User: %v, Home: %v \n", cfg_env.Env.User, cfg_env.Env.Home)
}
