package model

type KnowledgeEntryVersion struct {
	ID      int            `json:"id"`
	Title   string         `json:"title"`
	Content string         `json:"content"`
	Editor  string         `json:"editor"`
	Entry   KnowledgeEntry `json:"entry"`
}
