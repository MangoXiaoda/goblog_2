package article

import (
	"goblog_2/pkg/model"
	"goblog_2/pkg/types"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}
