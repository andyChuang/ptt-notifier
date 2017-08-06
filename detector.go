package pttnotifier

import (
	"github.com/andychuang/pttnotifier/crawler"
	"github.com/andychuang/pttnotifier/model"
	"fmt"
	"strings"
	"context"
	"time"
	"log"
)

type Detector struct {
	BoardName string
	Keywords  []string
	Frequency string
	Crawler   crawler.Crawler
	SummitCh  chan *model.DetectResult
	ErrCh     chan error
	SummitID  []string
}

func NewDetector(boardName, frequency string, keywords []string, crawler crawler.Crawler, targetCh chan *model.DetectResult, errCh chan error) (*Detector, error) {
	return &Detector{
		BoardName: boardName,
		Keywords:  keywords,
		Frequency: frequency,
		Crawler:   crawler,
		SummitCh:  targetCh,
		ErrCh:     errCh,
	}, nil
}

func (d *Detector) Run() {
	freq, err := time.ParseDuration(d.Frequency + "s")
	if err != nil {
		// Parse outside the detector
		log.Println(fmt.Errorf("Invalid frequency %s", d.Frequency))
		return
	}
	for {
		articles, err := d.Crawler.Crawl(nil)
		if err != nil {
			log.Println(fmt.Errorf("Crawling failed with board name: %s. Skiped. %s", d.BoardName, err.Error()))
			return
		}
		d.detect(articles)

		ctx, _ := context.WithTimeout(context.Background(), freq)
		select {
		case <-ctx.Done():
		}
	}
}

func (d *Detector) detect(articles []*model.Article) {
	for _, article := range articles {
		for _, keyword := range d.Keywords {
			if d.contains(*article, keyword) && !d.isSummit(article.ID) {
				d.SummitID = append(d.SummitID, article.ID)
				result := &model.DetectResult{
					BoardName: article.Board,
					Keyword:   keyword,
					Title:     article.Title,
					Id:        article.ID,
					Author:    article.Author,
					Url:       fmt.Sprintf(model.PTT_WEB_URL_PATTERN, article.Board, article.ID),
				}
				d.SummitCh <- result
			}
		}
	}
}

func (d *Detector) contains(article model.Article, keyword string) bool {
	// TODO: article.Content is always empty... will report issue to author
	return strings.Contains(article.Content, keyword) || strings.Contains(article.Title, keyword)
}

func (d *Detector) isSummit(id string) bool {
	for _, summitId := range d.SummitID {
		if summitId == id {
			log.Printf("Already summit. %s", id)
			return true
		}
	}
	return false
}
