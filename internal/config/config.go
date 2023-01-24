package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Settings struct {
	Debug bool `default:"false"`
}

var settings Settings

func init() {
	if err := envconfig.Process("edr", &settings); err != nil {
		log.Fatalln("cannot apply configuration:", err)
	}

	log.Infof("Using configuration %+v", settings)

	if settings.Debug {
		log.SetLevel(log.DebugLevel)
	}
}

func Get() *Settings {
	return &settings
}
