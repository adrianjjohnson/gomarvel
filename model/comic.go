package model

type ComicResult struct {
	Data ComicData `json:"data,omitempty"`
}

type ComicData struct {
	Offset  int     `json:"offset,omitempty"`
	Total   int     `json:"total,omitempty"`
	Limit   int     `json:"limit,omitempty"`
	Count   int     `json:"count,omitempty"`
	Results []Comic `json:"results,omitempty"`
}

type Comic struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
