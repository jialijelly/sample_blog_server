package models

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DBConfig struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Configuration struct {
	Server ServerConfig `json:"server"`
	DB     DBConfig     `json:"db"`
}
