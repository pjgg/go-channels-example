package channel

import (
	"sync"

	"github.com/pjgg/go-channels-example/configuration"
	"github.com/pjgg/go-channels-example/domain"
)

type ChannelType int

const (
	DIRECT_GRAPH ChannelType = 1 + iota
)

var channelTypeName = [...]string{
	"DIRECT_GRAPH",
}

func (channelType ChannelType) String() string {
	return channelTypeName[channelType-1]
}

type channel struct{}

type ChannelInterface interface {
	Handler(inboundEvents <-chan domain.Events) <-chan domain.Events
}

var onceChannel sync.Once
var channelInstance channel

// Factory method
func New(configuration *configuration.Configuration) ChannelInterface {

	onceChannel.Do(func() {
		switch configuration.IngestionChannel {
		case DIRECT_GRAPH.String():
			// TODO custom implementation
			break
		default:
			// TODO an other implementation
			break
		}
	})

	return &channelInstance
}

func (channel *channel) Handler(journey <-chan domain.Events) <-chan domain.Events {
	out := make(chan domain.Events)
	go func() {
		for trip := range journey {
			out <- trip
		}
		close(out)
	}()
	return out
}
