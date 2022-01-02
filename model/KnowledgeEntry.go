package model

type KnowledgeEntry struct {
	KnowledgeEntryId string `json:"id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Creator          int    `json:"creator"`
	CommunityId      int    `json:"communityid"`
}
