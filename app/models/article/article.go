package article

import (
	"goblog_2/app/models"
	"goblog_2/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}
