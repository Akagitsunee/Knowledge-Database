package model

type KnowledgeEntry struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Creator string `json:"creator"`
	Space   Space  `json:"space"`
}
