package model

type AddArticleReq struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ViewArticleReq struct {
	Author  string
	Keyword string
}
