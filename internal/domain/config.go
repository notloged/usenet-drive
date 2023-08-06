package domain

type Config struct {
	NzbPath    string `yaml:"nzb_path"`
	ServerPort string `yaml:"server_port" default:"8080"`
}
