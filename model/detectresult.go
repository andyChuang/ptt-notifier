package model

import (
	"fmt"
)

type DetectResult struct {
	BoardName string
	Keyword   string
	Title     string
	Id        string
	Url       string
	Author    string
}

func (dr *DetectResult) Display() string {
	return fmt.Sprintf(
		"不眠不休盯著板，總算幫你等到啦。板名：%s, 關鍵字：%s, 標題：%s, 作者：%s, 網址：%s",
		dr.BoardName, dr.Keyword, dr.Title, dr.Author, dr.Url)
}
