package crawler

import (
	"context"
	"github.com/andychuang/pttnotifier/model"
	"github.com/julianshen/gopttcrawler"
	"log"
)

const (
	MAX_ARTICLES_PER_PAGE = 20
)

type WebCrawler struct {
	BoardName string
}

func (wc *WebCrawler) Crawl(c context.Context) ([]*model.Article, error) {
	iter, err := wc.initIterator()
	if err != nil {
		return nil, err
	}
	return wc.getLatestArticles(iter), nil
}

func (wc *WebCrawler) initIterator() (gopttcrawler.Iterator, error) {
	articles, err := gopttcrawler.GetArticles(wc.BoardName, 0)
	if err != nil {
		return nil, err
	}
	return articles.Iterator(), nil
}

func (wc *WebCrawler) getLatestArticles(iter gopttcrawler.Iterator) []*model.Article {
	i := 0
	articles := []*model.Article{}
	for {
		if article, e := iter.Next(); e == nil {
			if i >= MAX_ARTICLES_PER_PAGE*2 {
				break
			}
			i++
			articles = append(articles, convertModel(*article))
			log.Printf("%s", article.Title)
		}
	}
	return articles
}

func convertModel(article gopttcrawler.Article) *model.Article {
	return &model.Article{
		ID:       article.ID,
		Board:    article.Board,
		Title:    article.Title,
		Content:  article.Content,
		Author:   article.Author,
		DateTime: article.DateTime,
	}
}
