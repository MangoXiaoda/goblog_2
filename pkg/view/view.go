package view

import (
	"goblog_2/pkg/logger"
	"goblog_2/pkg/route"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

func Render(w io.Writer, data interface{}, tplFiles ...string) {
	// 1 设置模板相对路径
	viewDir := "resources/views/"

	// 2. 遍历传参文件列表 Slice，设置正确的路径，支持 dir.filename 语法糖
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 3 所有布局模板文件 Slice
	layoutFiles, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
	logger.LogError(err)

	// 4 合并所有文件
	allFiles := append(layoutFiles, tplFiles...)

	// 5 解析所有模板文件
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	// 6 渲染模板，将所有文章的数据传输进去
	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}