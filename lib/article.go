package lib

import (
	"fmt"
	"os"
	"strings"

	"github.com/gernest/front"
	"github.com/russross/blackfriday"
)

type Article struct {
	Id          string
	Title       string
	MainImage   string
	Description string
	Body        string
	Created     string
	Updated     string
	Publish     string
	Slug        string
}

func (a *Article) GetArticle(id string) {
	a.Id = id
}

func (a *Article) LoadBody(slug string) {
	file, err := os.ReadFile("data/articles/" + slug)
	if err != nil {
		panic(err)
	}
	frontMatter := string(file[:10])
	fmt.Println(frontMatter)
	common := blackfriday.MarkdownCommon(file)
	a.Body = string(common)
}

type ArticleList struct {
	Articles []Article
}

func LoadArticles() ArticleList {
	var articles ArticleList
	articles.Articles = make([]Article, 0)
	files, _ := os.ReadDir("data/articles")
	m := front.NewMatter()
	m.Handle("---", front.JSONHandler)
	for _, f := range files {
		filePath := "data/articles/" + f.Name()
		// readfile
		file, err := os.ReadFile(filePath)
		body := string(file)
		f, _, err := m.Parse(strings.NewReader(body))
		if err != nil {
			panic(err)
		}
		article := Article{
			Title:       f["title"].(string),
			Description: f["description"].(string),
			Slug:        f["slug"].(string),
			Publish:     f["publishedAt"].(string),
			Id:          f["id"].(string),
		}
		articles.Articles = append(articles.Articles, article)
	}
	return articles
}
