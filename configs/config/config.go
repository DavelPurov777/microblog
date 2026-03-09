package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port            string
	Salt            string
	PprofEnabled    bool
	LikeQueueBuffer int

	DB     ConfigPgxpool
	DBPool PoolSettings
}

type ConfigPgxpool struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type PoolSettings struct {
	MaxConns          int32
	MinConns          int32
	MaxConnLifeTime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
}

func Load() (Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	// дефолты (чтобы проект стартовал даже если ключей нет)
	viper.SetDefault("port", "8080")
	viper.SetDefault("like_queue.buffer", 100)

	viper.SetDefault("db.sslmode", "disable")
	viper.SetDefault("db.pool.max_conns", int32(10))
	viper.SetDefault("db.pool.min_conns", int32(0))
	viper.SetDefault("db.pool.max_conn_lifetime", "1h")
	viper.SetDefault("db.pool.max_conn_idle_time", "30m")
	viper.SetDefault("db.pool.health_check_period", "1m")

	return Config{
		Port:            viper.GetString("port"),
		Salt:            viper.GetString("salt"),
		PprofEnabled:    os.Getenv("PPROF_ENABLED") == "true",
		LikeQueueBuffer: viper.GetInt("like_queue.buffer"),

		DB: ConfigPgxpool{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},

		DBPool: PoolSettings{
			MaxConns:          viper.GetInt32("db.pool.max_conns"),
			MinConns:          viper.GetInt32("db.pool.min_conns"),
			MaxConnLifeTime:   viper.GetDuration("db.pool.max_conn_lifetime"),
			MaxConnIdleTime:   viper.GetDuration("db.pool.max_conn_idle_time"),
			HealthCheckPeriod: viper.GetDuration("db.pool.health_check_period"),
		},
	}, nil
}
