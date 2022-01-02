package model

type KnowledgeEntryVersion struct {
	KnowledgeEntryVersionId string `json:"id"`
	Title                   string `json:"title"`
	Content                 string `json:"content"`
	Editor                  int    `json:"editor"`
	EntryId                 int    `json:"entryid"`
}
