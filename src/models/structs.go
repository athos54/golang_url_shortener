package models

type UrlShorten struct {
	ShortPath string `json:"short_path"`
	LongURL   string `json:"long_url"`
}
