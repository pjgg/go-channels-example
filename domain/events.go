package domain

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

type Journey struct {
	Id   int `json:"id"`
	Time int `json:"journey_time"`
}

type Events interface {
	FromJson(json []byte) *Journey
	LogJourney()
}

func (journey *Journey) FromJson(inboundJson []byte) *Journey {
	if err := json.Unmarshal(inboundJson, journey); err != nil {
		log.Error(err.Error())
	}
	return journey
}

func (journey *Journey) LogJourney() {

	log.WithFields(log.Fields{
		"Id":           journey.Id,
		"journey_time": journey.Time,
	}).Info("Journey stored")
}
