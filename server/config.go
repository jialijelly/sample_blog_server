package server

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jialijelly/sample_blog_server/models"
)

var DefaultConfiguration = models.Configuration{
	Server: models.ServerConfig{
		Host: "",
		Port: 8080,
	},
}

func LoadConfigFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	byteFile, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteFile, &DefaultConfiguration)
	if err != nil {
		return err
	}
	return nil
}
