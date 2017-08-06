package crawler

import (
	"context"
	"github.com/andychuang/pttnotifier/model"
)

type Crawler interface {
	Crawl(c context.Context) ([]*model.Article, error)
}

func NewCrawler(boardName string) (Crawler, error) {
	return &WebCrawler{
		BoardName: boardName,
	}, nil
}
