package configuration

import (
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	IngestionSource  string
	IngestionChannel string
	IngestionSink    string
}

var onceConfiguration sync.Once
var ConfigurationInstance *Configuration

func New() *Configuration {
	onceConfiguration.Do(func() {
		ConfigurationInstance = &Configuration{}

		ConfigurationInstance.IngestionSource = viper.GetString("go-channels-example.ingestion.source")
		ConfigurationInstance.IngestionChannel = viper.GetString("cago-channels-examplebify.ingestion.channel")
		ConfigurationInstance.IngestionSink = viper.GetString("go-channels-example.ingestion.sink")

		log.WithFields(log.Fields{
			"IngestionSource":  ConfigurationInstance.IngestionSource,
			"IngestionChannel": ConfigurationInstance.IngestionChannel,
			"IngestionSink":    ConfigurationInstance.IngestionSink,
		}).Info("configuration loaded")
	})

	return ConfigurationInstance
}
