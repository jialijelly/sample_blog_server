package server

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jialijelly/sample_blog_server/models"
)

const (
	// Database types:
	mysql = "mysql"
)

var DefaultConfiguration = models.Configuration{
	Server: models.ServerConfig{
		Host: "localhost",
		Port: 8080,
	},
	DB: models.DBConfig{
		Type: mysql,
		Host: "localhost",
		Port: 3306,
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
