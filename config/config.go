package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
)

type Config struct {
	Port         string
	MongoURI     string
	MongoTimeout int
	Duration     time.Duration
	OtpDuration  time.Duration
	Secret       string
}

func Load() Config {

	var cfg Config

	/*port, ok := os.LookupEnv("SERVICE_PORT")
	if !ok {
		port = "8080"
	}*/

	port := os.Getenv("PORT")
	cfg.Port = port

	mongoUri := os.Getenv("MONGODB_URI")

	/*mongoUri, ok := os.LookupEnv("MONGODB_URI")
	if !ok {
		mongoUri = "mongodb://user:password@localhost:27017"
	}*/
	cfg.MongoURI = mongoUri
	cfg.MongoTimeout = 10
	cfg.Secret = os.Getenv("SECRET")
	cfg.OtpDuration = 10 * time.Minute

	return cfg
}
