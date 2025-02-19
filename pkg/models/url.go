package models

type URL struct {
	Original string
	Short    string
}

type URLRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	Short string `json:"short"`
}
