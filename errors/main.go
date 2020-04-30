package main

import (
	"fmt"
	"os"
	"log"
	"github.com/pkg/errors" //This provides Wrap function on error
)

type Config struct {
 //configuration fields go here
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file") //This includes the description in your error
	}

	defer file.Close()

	cfg := &Config{}
	//Parse file here

	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONGLY, 0644) //openfile for append
	if err != nil {
		return
	}

	log.SetOutput(out) // set the output of the logging package to the log file created above
}

func main() {
	setupLogging()
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	//Normal Operation
	fmt.Println(cfg)
}