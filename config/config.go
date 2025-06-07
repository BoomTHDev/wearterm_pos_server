package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
	}

	Server struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit    int           `mapstructure:"bodyLimit" validate:"required"`
		TimeOut      time.Duration `mapstructure:"timeout" validate:"required"`
	}

	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

// func getEnvAsInt(name string, defaultVal int) int {
// 	valStr := os.Getenv(name)
// 	if valStr == "" {
// 		return defaultVal
// 	}
// 	val, err := strconv.Atoi(valStr)
// 	if err != nil {
// 		return defaultVal
// 	}
// 	return val
// }

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// // Set default values
		// viper.SetDefault("database.port", 5432)
		// viper.SetDefault("database.sslmode", "disable")
		// viper.SetDefault("database.schema", "public")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		// // Manually set the port from environment if it exists
		// if port := os.Getenv("DATABASE_PORT"); port != "" {
		// 	if portNum, err := strconv.Atoi(port); err == nil {
		// 		viper.Set("database.port", portNum)
		// 	}
		// }
	})

	// // Create a default config first
	// configInstance = &Config{
	// 	Server: &Server{
	// 		Port:         8080,
	// 		AllowOrigins: []string{"*"},
	// 		BodyLimit:    10,
	// 		TimeOut:      30 * time.Second,
	// 	},
	// 	Database: &Database{
	// 		Host:     os.Getenv("DATABASE_HOST"),
	// 		Port:     getEnvAsInt("DATABASE_PORT", 5432),
	// 		User:     os.Getenv("DATABASE_USER"),
	// 		Password: os.Getenv("DATABASE_PASSWORD"),
	// 		DBName:   os.Getenv("DATABASE_DBNAME"),
	// 		SSLMode:  os.Getenv("DATABASE_SSLMODE"),
	// 		Schema:   os.Getenv("DATABASE_SCHEMA"),
	// 	},
	// }

	// Initialize configInstance before unmarshaling
	configInstance = &Config{
		Server:   &Server{},
		Database: &Database{},
	}

	// Unmarshal from config file if exists
	if err := viper.Unmarshal(configInstance); err != nil {
		panic(err)
	}

	validating := validator.New()

	if err := validating.Struct(configInstance); err != nil {
		panic(err)
	}

	return configInstance
}
