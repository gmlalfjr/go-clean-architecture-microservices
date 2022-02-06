package web

type ArticleCreateResponse struct {
	Title string `json:"title"`
	Text string `json:"text"`
	Status string `json:"status"`
}