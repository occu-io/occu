package settings

import (
	"encoding/json"
	"fmt"
	"os"
	"util/models"
)

func LoadConfiguration(file string) models.Config {
	var config models.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return config
}
