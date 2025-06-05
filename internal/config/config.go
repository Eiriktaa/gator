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

// allows for the config file to be in mulitple location
// defaults to homedir
func findConfigFileExists() (filepath string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("unable to find home dir aborting")
	}
	validPaths := []string{
		homedir + "/" + baseConfigFileName,
		"./" + baseConfigFileName,
	}
	for _, path := range validPaths {
		fmt.Println(path)
		if _, err := os.Stat(path); err == nil {
			fmt.Println(path)
			return path
		}
	}
	log.Fatal("Unable to find configuration, aborting")
	return ""
}

func LoadConfiguration() Config {
	filepath := findConfigFileExists()
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
func (c *Config) SetUser(user string) {
	c.Current_user_name = user
}

func (c Config) PrintSelf() {
	fmt.Println("URL :", c.DB_url)
	fmt.Println("Current :", c.Current_user_name)
}
