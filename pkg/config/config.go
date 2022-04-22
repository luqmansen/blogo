package config

import (
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	filename "github.com/keepeye/logrus-filename"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type Configuration struct {
	Host string `env:"SERVER_HOST"`
	Port string `env:"SERVER_PORT"`

	DatabaseURI string        `env:"POSTGRES_URI"`
	Timeout     time.Duration `env:"POSTGRES_TIMEOUT"`

	GithubClientId     string `env:"GITHUB_CLIENT_ID"`
	GithubClientSecret string `env:"GITHUB_CLIENT_SECRET"`
}

func LoadConfig() *Configuration {

	if os.Getenv("DEPLOY") != "PROD" {
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stdout)
	}

	// add filename in logging
	filenameHook := filename.NewHook()
	filenameHook.Field = "SRC"
	log.AddHook(filenameHook)

	formatter := &prefixed.TextFormatter{
		ForceColors:     true,
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	formatter.SetColorScheme(&prefixed.ColorScheme{
		PrefixStyle:    "blue+b",
		TimestampStyle: "white+h",
	})
	log.SetFormatter(formatter)

	conf := &Configuration{}
	if err := godotenv.Load(); err != nil {
		log.Errorf("Failed to load .env for DEV, %s", err.Error())
	}

	if err := env.Parse(conf); err != nil {
		log.Fatalln("unable to parse environment variables: ", err)
	}
	return conf
}
