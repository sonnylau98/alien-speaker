package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/sonnylau98/alien-speaker/proxy/homedir"

	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	configPath string
)

type Config struct {
	ListenAddr string `json:"listen"`
	RemoteAddr string `json:"remote"`
}

func init() {
	home, _ := homedir.Dir()
	configFilename := ".alienspeaker.json"
	if len(os.Args) == 2 {
		configFilename = os.Args[1]
	}
	configPath = path.Join(home, configFilename)
}

func (config *Config) SaveConfig() {
	configJson, _ := json.MarshalIndent(config, "", "    ")
	err := ioutil.WriteFile(configPath, configJson, 0644)
	if err != nil {
		fmt.Errorf("Saving %s failed: %s", configPath, err)
	}
	log.Printf("Succesfully saved %s\n", configPath)
}

func (config *Config) ReadConfig() {
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("Read config from file %s\n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("Opening %s failed: %s", configPath, err)
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("illegal format JSON config file:\n%s", file.Name())
		}
	}
}
