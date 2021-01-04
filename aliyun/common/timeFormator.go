package common

import (
	"log"
	"time"

	"github.com/aiaoyang/resourceManager/resource"
)

type timeFormat string

const (
	ecsTimeFormat    timeFormat = "2006-01-02T15:04Z"
	certTimeFormat   timeFormat = "2006-01-02"
	domainTimeFormat timeFormat = "2006-01-02 15:04:05"
	rdsTimeFormat    timeFormat = "2006-01-02T15:04:05Z"
)

func parseTime(timeString string, tFormat timeFormat) (s resource.Stat) {
	pTime, err := time.Parse(string(tFormat), timeString)
	if err != nil {
		log.Fatal(err)
	}

	s = resource.Green

	if time.Now().AddDate(0, 1, 0).After(pTime) {
		s = resource.Yellow
	}
	if time.Now().AddDate(0, 0, 7).After(pTime) {
		s = resource.Red
	}
	if time.Now().After(pTime) {
		s = resource.NearDead
	}
	return s
}
