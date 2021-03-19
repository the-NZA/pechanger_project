package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	"github.com/the-NZA/pechanger_project/backend/internal/app/pechanger"
)

var (
	configPath string
)

// Author: Sergey Shakotko
// Setup command line flags
func init() {
	flag.StringVar(&configPath, "config", "backend/config/pechanger_config.json", "Path to config file")
}

func main() {
	flag.Parse()

	config := pechanger.NewConfig()
	jsonFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	configBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	s := pechanger.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
