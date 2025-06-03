package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const baseConfigFileName string = ".gatorconfig.json"
const configFilePermissons int = 0777

type Config struct {
	DB_url             string
	Current_user_name  string
	configFileLocation string
}

func findCofigFileExists() (filepath string) {
	//Todo: allow for more configs file locations
	validPaths := []string{"./" + baseConfigFileName}
	for _, path := range validPaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	log.Fatal("Unable to find configuration, aborting")
	return ""
}

func LoadConfiguration() Config {
	filepath := findCofigFileExists()
	body, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Unable to read config file, %v", err)
	}
	var config Config
	if err := json.Unmarshal(body, &config); err != nil {
		log.Fatalf("Failed to unmarshal config file %v", err)
	}
	config.configFileLocation = filepath
	return config
}

func (c Config) WriteToFile() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.configFileLocation, data, os.FileMode(configFilePermissons))
	if err != nil {
		return err
	}
	return nil

}
func (c Config) SetUser(user string) {
	c.Current_user_name = user
}

func (c Config) PrintSelf() {
	fmt.Println("URL :", c.DB_url)
	fmt.Println("Current :", c.Current_user_name)
}
