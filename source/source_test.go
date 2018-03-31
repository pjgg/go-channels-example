package source

import (
	"os"
	"testing"

	"github.com/pjgg/go-channels-example/channel"
	"github.com/pjgg/go-channels-example/configuration"
	"github.com/pjgg/go-channels-example/sink"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

func (sourceTestSuite *SourceTestSuite) TestorderJourneys() {
	journeys := []string{`{"id": 1, "journey_time": 100}`, `{"id": 2, "journey_time": 110}`, `{"id": 3, "journey_time": 120}`, `{"id": 4, "journey_time": 550}`}

	// source -> channel -> sink
	source := sourceTestSuite.source.Handler(journeys...)
	channel := sourceTestSuite.channel.Handler(source)
	sourceTestSuite.sink.Handler(channel)
}

type SourceTestSuite struct {
	suite.Suite
	source  SourceInterface
	channel channel.ChannelInterface
	sink    sink.SinkInterface
}

func (suite *SourceTestSuite) SetupTest() {}

func (suite *SourceTestSuite) TearDownTest() {}

func TestSourceTestSuite(t *testing.T) {
	configInit()
	sourceTestSuite := new(SourceTestSuite)

	sourceTestSuite.source = New(configuration.ConfigurationInstance)
	sourceTestSuite.channel = channel.New(configuration.ConfigurationInstance)
	sourceTestSuite.sink = sink.New(configuration.ConfigurationInstance)

	suite.Run(t, sourceTestSuite)
}

func configInit() {
	viper.SetConfigName("config")
	configPath, exist := os.LookupEnv("CONFIG_PATH")
	if exist {
		viper.AddConfigPath(configPath)
	}
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	configuration.New()
}
