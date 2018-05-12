package model

type CharacterResult struct {
	Data CharacterData `json:"data,omitempty"`
}

type CharacterData struct {
	Offset  int         `json:"offset,omitempty"`
	Total   int         `json:"total,omitempty"`
	Limit   int         `json:"limit,omitempty"`
	Count   int         `json:"count,omitempty"`
	Results []Character `json:"results,omitempty"`
}

type Character struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
