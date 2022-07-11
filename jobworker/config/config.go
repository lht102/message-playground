package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Cfg struct {
	RestPort   int
	MySQLCfg   MySQLConfig
	NATSConfig NATSConfig
}

type NATSConfig struct {
	Protocol string
	Address  string
}

type MySQLConfig struct {
	Username     string
	Password     string
	Protocol     string
	Address      string
	Database     string
	MaxOpenConns int
	MaxIdleConns int
}

func Load() *Cfg {
	v := viper.New()
	v.AutomaticEnv()

	return &Cfg{
		RestPort: v.GetInt("REST_PORT"),
		MySQLCfg: MySQLConfig{
			Username:     v.GetString("MYSQL_USERNAME"),
			Password:     v.GetString("MYSQL_PASSWORD"),
			Protocol:     v.GetString("MYSQL_PROTOCOL"),
			Address:      v.GetString("MYSQL_ADDRESS"),
			Database:     v.GetString("MYSQL_DATABASE"),
			MaxOpenConns: v.GetInt("MYSQL_MAX_OPEN_CONNS"),
			MaxIdleConns: v.GetInt("MYSQL_MAX_IDLE_CONNS"),
		},
		NATSConfig: NATSConfig{
			Protocol: v.GetString("NATS_PROTOCOL"),
			Address:  v.GetString("NATS_ADDRESS"),
		},
	}
}

func GetMySQLDSN(cfg MySQLConfig) string {
	return fmt.Sprintf(
		"%s:%s@%s(%s)/%s?parseTime=true&multiStatements=true",
		cfg.Username,
		cfg.Password,
		cfg.Protocol,
		cfg.Address,
		cfg.Database,
	)
}

func GetNATSURL(cfg NATSConfig) string {
	return fmt.Sprintf(
		"%s://%s",
		cfg.Protocol,
		cfg.Address,
	)
}
