package config

import "fmt"

type Logging struct {
	LogLevel   string `envconfig:"LOG_LEVEL" default:"info"`
	LogConsole bool   `envconfig:"LOG_CONSOLE" default:"false"`
}

type Security struct {
	RSAPublicKey string `split_words:"true"`
}

type DB struct {
	Database       string `required:"true" split_words:"true"`
	Host           string `required:"true" split_words:"true"`
	Port           uint16 `default:"5432" split_words:"true"`
	User           string `required:"true" split_words:"true"`
	Password       string `required:"true" split_words:"true"`
	MaxConnections int32  `default:"16" split_words:"true"`
	LogLevel       string `default:"none" split_words:"true"`
	Hostname       string `default:"go-application" envconfig:"HOSTNAME" split_words:"true"`
}

type HTTP struct {
	Host    string `default:""`
	Port    string `default:"8085"`
	Debug   bool   `default:"false"`
	EnvName string `default:"" split_words:"true"`
	//AuthServer string `required:"true" split_words:"true"`
}

type Config struct {
	Logging  Logging
	Security Security
	DB       DB
	HTTP     HTTP
}

func (c *Config) GetConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?pool_max_conns=%d&application_name=%s&&timezone=UTC",
		c.DB.User,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.Database,
		c.DB.MaxConnections,
		c.DB.Hostname,
	)
}

func (c *Config) GetLogLevel() string {
	return c.Logging.LogLevel
}

func (c *Config) GetLogConsole() bool {
	return c.Logging.LogConsole
}

func (c *Config) GetRSAPublicKey() string {
	return c.Security.RSAPublicKey
}

func (c *Config) GetServerInitString() string {
	return fmt.Sprintf("%s:%s", c.HTTP.Host, c.HTTP.Port)
}
