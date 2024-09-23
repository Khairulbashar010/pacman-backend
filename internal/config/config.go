package config

import (
    "log"

    "github.com/spf13/viper"
    "github.com/joho/godotenv"
)

type Config struct {
    Port         string `mapstructure:"port"`
    Environment  string `mapstructure:"environment"`
    
    // Database settings
    DBHost       string `mapstructure:"db_host"`
    DBPort       string `mapstructure:"db_port"`
    DBUser       string `mapstructure:"db_user"`
    DBPassword   string `mapstructure:"db_password"`
    DBName       string `mapstructure:"db_name"`

    // JWT secret
    JWTSecret    string `mapstructure:"jwt_secret"`

    // SMTP settings
    SMTPHost     string `mapstructure:"smtp_host"`
    SMTPPort     string `mapstructure:"smtp_port"`
    SMTPUser     string `mapstructure:"smtp_user"`
    SMTPPassword string `mapstructure:"smtp_password"`
    SMTPFrom     string `mapstructure:"smtp_from"`
}

var AppConfig *Config

func LoadConfig() {
    // Load environment variables from the .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found. Continuing without it.")
    }

    // Initialize Viper to read from environment variables
    viper.AutomaticEnv()

    // Set default values and read from environment variables
    AppConfig = &Config{
        Port:         viper.GetString("PORT"),
        Environment:  viper.GetString("ENVIRONMENT"),

        // Database
        DBHost:       viper.GetString("DB_HOST"),
        DBPort:       viper.GetString("DB_PORT"),
        DBUser:       viper.GetString("DB_USER"),
        DBPassword:   viper.GetString("DB_PASSWORD"),
        DBName:       viper.GetString("DB_NAME"),

        // JWT
        JWTSecret:    viper.GetString("JWT_SECRET"),

        // SMTP settings
        SMTPHost:     viper.GetString("SMTP_HOST"),
        SMTPPort:     viper.GetString("SMTP_PORT"),
        SMTPUser:     viper.GetString("SMTP_USER"),
        SMTPPassword: viper.GetString("SMTP_PASSWORD"),
        SMTPFrom:     viper.GetString("SMTP_FROM"),
    }

    log.Println("Configuration loaded successfully")
}

func GetConfig() *Config {
    return AppConfig
}
