package pttnotifier

import (
	"github.com/andychuang/pttnotifier/crawler"
	"github.com/andychuang/pttnotifier/model"
	"log"
)

type DetectorCenter struct {
	Target      map[string][]string
	DetectorNum int
	Frequency   string
}

func NewDetectorCenter(rawTarget, frequency string, detectorNum int) (*DetectorCenter, error) {
	target, err := parseKeyword(rawTarget)
	if err != nil {
		return nil, err
	}

	return &DetectorCenter{
		Target:      target,
		DetectorNum: detectorNum,
		Frequency:   frequency,
	}, nil
}

func (dc *DetectorCenter) Run() error {
	errCh := make(chan error)
	summitCh := make(chan *model.DetectResult)

	for boardName, keywords := range dc.Target {
		// TODO: Make detector num useful
		crawler, err := crawler.NewCrawler(boardName)
		if err != nil {
			return err
		}
		detector, err := NewDetector(boardName, dc.Frequency, keywords, crawler, summitCh, errCh)
		if err != nil {
			return err
		}

		go detector.Run()
	}
	for {
		select {
		case result := <-summitCh:
			// TODO: push notification
			log.Printf("Push %v", result)
		case err := <-errCh:
			return err
		}
	}
	return nil

}

// Return map[board name]{keyword1, keyword2...}
func parseKeyword(rawData string) (map[string][]string, error) {
	return map[string][]string{
		"Mabinogi": {"贈送"},
	}, nil
}
