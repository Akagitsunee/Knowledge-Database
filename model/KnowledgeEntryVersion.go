package model

type KnowledgeEntryVersion struct {
	KnowledgeEntryVersionId string `json:"id"`
	Title                   string `json:"title"`
	Content                 string `json:"content"`
	EditorId                int    `json:"editorId"`
	EntryId                 int    `json:"entryid"`
}
