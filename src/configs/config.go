package configs

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var Config App

type App struct {
	Host    string `default:"0.0.0.0" envconfig:"HOST"`
	Port    int    `default:"8080" envconfig:"PORT"`
	RunMode string `default:"debug" envconfig:"RUN_MODE"`
	Env     string `default:"debug" envconfig:"ENV"`
	Redis   Redis
	MySql   MySQL
}

func New() (*App, error) {
	if err := envconfig.Process("Prefix", &Config); err != nil {
		return nil, err
	}
	return &Config, nil
}

// AddressListener returns address listener of HTTP server.
func (c *App) AddressListener() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

// Redis represents configuration of Redis caching database.
type Redis struct {
	Host       string `default:"127.0.0.1" envconfig:"REDIS_HOST"`
	Port       int    `default:"6379" envconfig:"REDIS_PORT"`
	Password   string `default:"" envconfig:"REDIS_PASSWORD"`
	Database   int    `default:"0" envconfig:"REDIS_DB"`
	MasterName string `default:"mymaster" envconfig:"REDIS_MASTER_NAME"`
}

// URL return redis connection URL.
func (c Redis) URL() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

// MySQL represents configuration of MySQL database.
type MySQL struct {
	Username string `default:"vin_id" envconfig:"MYSQL_USER"`
	Password string `default:"vin_id" envconfig:"MYSQL_PASS"`
	Host     string `default:"127.0.0.1" envconfig:"MYSQL_HOST"`
	Port     int    `default:"3306" envconfig:"MYSQL_PORT"`
	Database string `default:"gamezone" envconfig:"MYSQL_DB"`
}

// ConnectionString returns connection string of MySQL database.
func (c *MySQL) ConnectionString() string {
	format := "%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8"
	return fmt.Sprintf(format, c.Username, c.Password, c.Host, c.Port, c.Database) + "&loc=Asia%2FHo_Chi_Minh"
}
