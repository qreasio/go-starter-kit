package config

import (
	"io/ioutil"
	"os"

	"github.com/qreasio/go-starter-kit/pkg/log"
	"gopkg.in/yaml.v2"
)

// Config holds data for application configuration
type Config struct {
	Server *Server   `yaml:"server,omitempty"`
	DB     *Database `yaml:"database,omitempty"`
}

// Database holds data for database configuration
type Database struct {
	Dsn string `yaml:"dsn,omitempty"`
}

// Server holds data for server configuration
type Server struct {
	Port string `yaml:"port,omitempty"`
}

// Load returns config from yaml and environment variables.
func Load(file string, logger log.Logger) (*Config, error) {
	// default config
	c := Config{}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// add code here to load from env variables and override the value from yaml, if any
	if os.Getenv("APP_DSN") != "" {
		c.DB.Dsn = os.Getenv("APP_DSN")
	}
	if os.Getenv("APP_PORT") != "" {
		c.Server.Port = os.Getenv("APP_PORT")
	}

	return &c, err
}
