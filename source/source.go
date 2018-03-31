package source

import (
	"sync"
	"time"

	"github.com/pjgg/go-channels-example/configuration"
	"github.com/pjgg/go-channels-example/domain"
)

type SourceType int

const (
	DEFAULT SourceType = 1 + iota
)

var sourceTypeName = [...]string{
	"DEFAULT",
}

func (sourceType SourceType) String() string {
	return sourceTypeName[sourceType-1]
}

type source struct{}

type SourceInterface interface {
	Handler(events ...string) <-chan domain.Events
}

var onceSource sync.Once
var sourceInstance source

// Factory method
func New(configuration *configuration.Configuration) SourceInterface {

	onceSource.Do(func() {
		switch configuration.IngestionSource {
		case DEFAULT.String():
			// TODO custom implementation
			break
		default:
			// TODO an other implementation
			break
		}

	})

	return &sourceInstance
}

func (source *source) Handler(events ...string) <-chan domain.Events {

	out := make(chan domain.Events)
	go func() {
		for _, trip := range events {
			journey := new(domain.Journey).FromJson([]byte(trip))

			select {
			case <-time.Tick(time.Duration(journey.Time) * time.Millisecond):
				out <- new(domain.Journey).FromJson([]byte(trip))
				break
			}

		}

		close(out)
	}()
	return out

}
