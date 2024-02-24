package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Database          *Database `mapstructure:"DATABASE" json:"DATABASE"`
	ServerAddress     string    `mapstructure:"SERVER_ADDRESS" json:"SERVER_ADDRESS"`
	TokenSymmetricKey string    `mapstructure:"TOKEN_SYMMETRIC_KEY" json:"TOKEN_SYMMETRIC_KEY"`
	Twilio            *Twilio   `mapstructure:"TWILIO" json:"TWILIO"`
	Email             *Email    `mapstructure:"EMAIL" json:"EMAIL"`
	Redis             *Redis    `mapstructure:"REDIS_SERVER" json:"REDIS_SERVER"`
	OAuth             *OAuth    `mapstructure:"OAUTH" json:"OAUTH"`
}

type Database struct {
	DBDriver string `mapstructure:"DB_DRIVER" json:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE" json:"DB_SOURCE"`
}

type Twilio struct {
	AccountSid string `mapstructure:"ACCOUNT_SID" json:"ACCOUNT_SID"`
	AuthToken  string `mapstructure:"AUTH_TOKEN" json:"AUTH_TOKEN"`
	ServiceSid string `mapstructure:"SERVICE_SID" json:"SERVICE_SID"`
}

type Email struct {
	SenderName     string `mapstructure:"SENDER_NAME" json:"SENDER_NAME"`
	SenderAddress  string `mapstructure:"SENDER_ADDRESS" json:"SENDER_ADDRESS"`
	SenderPassword string `mapstructure:"SENDER_PASSWORD" json:"SENDER_PASSWORD"`
}

type Redis struct {
	Address  string `mapstructure:"ADDRESS" json:"ADDRESS"`
	Password string `mapstructure:"PASSWORD" json:"PASSWORD"`
	Db       string `mapstructure:"DB" json:"DB"`
}

type OAuth struct {
	WebClientId     string `mapstructure:"GOOGLE_WEB_CLIENT_ID" json:"GOOGLE_WEB_CLIENT_ID"`
	WebClientSecret string `mapstructure:"GOOGLE_WEB_CLIENT_SECRET" json:"GOOGLE_WEB_CLIENT_SECRET"`
	AndroidClientId string `mapstructure:"GOOGLE_ANDROID_CLIENT_ID" json:"GOOGLE_ANDROID_CLIENT_ID"`
	IOSClientId     string `mapstructure:"GOOGLE_IOS_CLIENT_ID" json:"GOOGLE_IOS_CLIENT_ID"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

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

// OLD Config
// type Config struct {
// 	DBServer            string `mapstructure:"DB_DRIVER"`
// 	DBSource            string `mapstructure:"DB_SOURCE"`
// 	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
// 	TokenSymmetricKey   string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
// 	TwilioAccountSid    string `mapstructure:"TWILIO_ACCOUNT_SID"`
// 	TwilioAuthToken     string `mapstructure:"TWILIO_AUTH_TOKEN"`
// 	TwilioServiceSid    string `mapstructure:"TWILIO_SERVICE_SID"`
// 	SenderName          string `mapstructure:"SENDER_NAME"`
// 	SenderEmailAddress  string `mapstructure:"SENDER_EMAIL_ADDRESS"`
// 	SemderEmailPassword string `mapstructure:"SENDER_PASSWORD"`
// }

// OLD Config using app.env file
// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return
// }
