package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/oswaldom-code/chucknorris/pkg/log"

	"github.com/spf13/viper"
)

const (
	DevEnvironment = "dev"
)

type AppInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type DBConfig struct {
	User               string `json:"user,omitempty"`
	Password           string `json:"password,omitempty"`
	Host               string `json:"host,omitempty"`
	Port               int    `json:"port,omitempty"`
	Database           string `json:"database,omitempty"`
	MaxOpenConnections int    `json:"-"`
	SSLMode            string `json:"sslmode,omitempty"`
}

type S3Config struct {
	BucketName   string `json:"bucket,omitempty"`
	AWSAccessKey string `json:"access_key,omitempty"`
	AWSSecretKey string `json:"secret_access_key,omitempty"`
	Region       string `json:"region,omitempty"`
	Endpoint     string `json:"endpoint,omitempty"`
}

type ServerConfig struct {
	Host   string
	Port   int
	Scheme string
}

type LoggingConfig struct {
	Level        string
	ErrorLogFile string
}

type EmailConfig struct {
	User       string
	Password   string
	Port       int
	SMTPServer string
	Host       string
}

type EnvironmentConfig struct {
	Environment string
}

type AuthConfig struct {
	Secret string
}

// GetProjectPath returns the current project path
func GetProjectPath() string {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Warn("Warning, cannot get current path")
		return ""
	}
	// Traverse back from current directory until service base dir is reach and add to config path
	for !strings.HasSuffix(dir, "chucknorris") && dir != "/" {
		dir, err = filepath.Abs(dir + "/..")
		if err != nil {
			break
		}
	}
	return dir
}

func Load(filename string) {
	viper.SetEnvPrefix("skypet-api")
	viper.SetConfigName(filename)
	viper.AddConfigPath("config")
	viper.AddConfigPath(GetProjectPath() + "/config")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 9000)
	viper.SetDefault("server.scheme", "http")
	viper.SetDefault("storage.db.host",os.Getenv("SKYPET_API_STORAGE_DB_HOST"))
	viper.SetDefault("storage.db.port", os.Getenv("SKYPET_API_STORAGE_DB_PORT"))
	viper.SetDefault("storage.db.database", os.Getenv("SKYPET_API_STORAGE_DB_NAME"))
	viper.SetDefault("storage.db.user", os.Getenv("SKYPET_API_STORAGE_DB_USER"))
	viper.SetDefault("storage.db.password", os.Getenv("SKYPET_API_STORAGE_DB_PASSWORD"))
	viper.SetDefault("storage.db.max_connections", 20)
	viper.SetDefault("storage.db.sslmode", "disable")
	viper.SetDefault("logging.level", "DEBUG")
	viper.SetDefault("logging.errorLogFile", "error.log")
	viper.SetDefault("email.smtp_server", os.Getenv("SMTP_SERVER"))
	viper.SetDefault("email.port", 587)
	viper.SetDefault("email.host", "https://skypet.comm")
	viper.SetDefault("email.user", "skypet@gmail.com")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_", ".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
	log.SetLogLevel(GetLoggingConfig().Level)
}

func GetDBConfig() DBConfig {
	config := DBConfig{
		User:               viper.GetString("storage.db.user"),
		Password:           viper.GetString("storage.db.password"),
		Host:               viper.GetString("storage.db.host"),
		Port:               viper.GetInt("storage.db.port"),
		Database:           viper.GetString("storage.db.database"),
		MaxOpenConnections: viper.GetInt("storage.db.max_connections"),
		SSLMode:            viper.GetString("storage.db.sslmode"),
	}
	return config
}

func GetS3Config() S3Config {
	config := S3Config{
		BucketName:   viper.GetString("storage.s3.bucket"),
		AWSAccessKey: viper.GetString("storage.s3.access_key"),
		AWSSecretKey: viper.GetString("storage.s3.secret_access_key"),
		Region:       viper.GetString("storage.s3.region"),
		Endpoint:     viper.GetString("storage.s3.endpoint"),
	}
	return config
}

func GetServerConfig() ServerConfig {
	config := ServerConfig{
		Host:   viper.GetString("server.host"),
		Port:   viper.GetInt("server.port"),
		Scheme: viper.GetString("server.scheme"),
	}
	return config
}

func GetLoggingConfig() LoggingConfig {
	return LoggingConfig{
		Level:        viper.GetString("logging.level"),
		ErrorLogFile: viper.GetString("logging.errorLogFile"),
	}
}

func GetEmailConfig() EmailConfig {
	return EmailConfig{
		User:       viper.GetString("email.user"),
		Password:   viper.GetString("email.password"),
		Port:       viper.GetInt("email.port"),
		SMTPServer: viper.GetString("email.smtp_server"),
		Host:       viper.GetString("email.host"),
	}
}

func GetEnvironmentConfig() EnvironmentConfig {
	return EnvironmentConfig{
		Environment: viper.GetString("environment"),
	}
}

func GetAuthConfig() AuthConfig {
	return AuthConfig{
		Secret: viper.GetString("auth.secret"),
	}
}

func (s ServerConfig) AsUri() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

func (s DBConfig) GetConnectionString() string {
	bStr, err := json.Marshal(s)
	if err != nil {
		log.FatalWithFields("Cannot marshal DSN", log.Fields{"err": err, "dsn": s})
	}
	connectionString := string(bStr)
	connectionString = strings.NewReplacer(":", "=", "{", "", "}", "", `"`, "", ",", " ").Replace(connectionString)
	return connectionString
}
