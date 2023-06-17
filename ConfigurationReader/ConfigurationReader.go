package ConfigurationReader

import (
	"encoding/json"
	"os"
)

type ConfigurationReader struct {
	Ip        string `json:"ip"`
	Port      string `json:"port"`
	Ssl       bool   `json:"ssl"`
	LLMAPIURL string `json:"llmapiurl"`
	Keyfile   string `json:"keyfile"`
	Certfile  string `json:"certfile"`
}

// New Will read conf.json and return a new ConfigurationReader struct
func New() *ConfigurationReader {
	f, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Read the file and unmarshal the JSON
	decoder := json.NewDecoder(f)
	// create a new ConfigurationReader struct
	conf := ConfigurationReader{}
	// decode the JSON and write it to the struct
	err = decoder.Decode(&conf)
	// Check for errors
	if err != nil {
		panic(err)
	}
	// Panic if ssl is true but no keyfile or certfile is set
	if conf.Ssl && (conf.Keyfile == "" || conf.Certfile == "") {
		panic("SSL TRUE and Keyfile or Certfile not set")
	}
	// return the struct
	return &conf
}
