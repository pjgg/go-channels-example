package sink

import (
	"sync"

	"github.com/pjgg/go-channels-example/configuration"
	"github.com/pjgg/go-channels-example/domain"
)

type SinkType int

const (
	FUNKY_LOGGER SinkType = 1 + iota
)

var sinkTypeName = [...]string{
	"FUNKY_LOGGER",
}

func (sinkType SinkType) String() string {
	return sinkTypeName[sinkType-1]
}

type funkyLogger struct{}

type SinkInterface interface {
	Handler(inboundEvents <-chan domain.Events)
}

var onceSink sync.Once
var funkyLoggerInstance funkyLogger

// Factory method
func New(configuration *configuration.Configuration) SinkInterface {

	onceSink.Do(func() {
		switch configuration.IngestionSink {
		case FUNKY_LOGGER.String():
			// TODO custom implementation
			break
		default:
			// TODO an other implementation
			break
		}
	})

	return &funkyLoggerInstance
}

func (logger *funkyLogger) Handler(inboundEvents <-chan domain.Events) {

	for trip := range inboundEvents {
		trip.LogJourney()
	}

}
