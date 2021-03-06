package view

import (
	"../model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(fileName string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(fileName)),
	}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
