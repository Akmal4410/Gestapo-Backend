package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBServer            string `mapstructure:"DB_DRIVER"`
	DBSource            string `mapstructure:"DB_SOURCE"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	TwilioAccountSid    string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken     string `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioServiceSid    string `mapstructure:"TWILIO_SERVICE_SID"`
	SenderName          string `mapstructure:"SENDER_NAME"`
	SenderEmailAddress  string `mapstructure:"SENDER_EMAIL_ADDRESS"`
	SemderEmailPassword string `mapstructure:"SENDER_PASSWORD"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func EnvAccountSid() string {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func EnvAuthToken() string {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_AUTH_TOKEN")
}

func EnvServiceSid() string {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_SERVICE_SID")
}
