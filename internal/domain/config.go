package domain

type Config struct {
	NzbPath    string `yaml:"nzb_path"`
	ServerPort string `yaml:"server_port" default:"8080"`
	Usenet     Usenet `yaml:"usenet"`
}

type Usenet struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	Group          string `yaml:"group"`
	SSL            bool   `yaml:"ssl"`
	MaxConnections int    `yaml:"max_connections"`
}
