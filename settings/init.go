package settings

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var DataSettings Settings

type Settings struct {
	Port string
	PrimaryDB MySql `yaml:"primary_db"`
}

func init() {

	// load file config
	file, err := os.Open("./config.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// close file config
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	// load config to config variable
	configDecoded := yaml.NewDecoder(file)
	err = configDecoded.Decode(&DataSettings)
	if err != nil {
		fmt.Println("File config is not valid")
		os.Exit(0)
	}
}