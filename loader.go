package gojsonconfig

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

// Reads -config parameter for a config file parameter, if it is not present, it will attempt to
// read a file named config.json in the same directory that program is executed
// JSON is then unmarshalled into "config"
func Load(config interface{}) {

	defaultConfig := "config.json"

	var configFile string
	flag.StringVar(&configFile, "config", defaultConfig, "Path to JSON config file")
	flag.Parse()

	log.Printf("Config file: %s", configFile)
	configFileContents, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Unable to read config file: %s, Err: %v", configFile, err)
	}

	err = json.Unmarshal(configFileContents, &config)
	if err != nil {
		log.Fatalf("Unable to parse config file: %s. Err: %v", configFile, err)
	}
}
