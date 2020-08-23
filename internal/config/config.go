package config

import (
	"errors"
	"fmt"
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
	logger.Infof("loading config file : %s \n", file)

	// default config
	c := Config{}
	// load from YAML config file
	if rawcfg, err := ioutil.ReadFile(file); err == nil {
		if err := yaml.Unmarshal(rawcfg, &c); err != nil {
			logger.Errorf("error on json marshall of config file : %s \n", file)
			return nil, err
		}
	} else {
		logger.Errorf("error reading config file : %s \n", file)
		return nil, err
	}

	// load from env variables and override the value from yaml, if any
	if os.Getenv("APP_DSN") != "" {
		c.DB.Dsn = os.Getenv("APP_DSN")
	}
	if os.Getenv("APP_PORT") != "" {
		c.Server.Port = os.Getenv("APP_PORT")
	}

	// construct dsn from separate db env vars if it is still empty
	if c.DB.Dsn == "" {
		c.DB.Dsn = fmt.Sprintf("%s:%s@(%s:3306)/%s?parseTime=true&multiStatements=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"))
	}

	// if dsn still empty, throw error
	if c.DB.Dsn == "" {
		return nil, errors.New("database configuration is missing")
	}

	return &c, nil
}
