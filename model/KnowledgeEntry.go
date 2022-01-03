package model

type KnowledgeEntry struct {
	KnowledgeEntryId string `json:"id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	CreatorId        int    `json:"creatorid"`
	CommunityId      int    `json:"communityid"`
	Deleted          bool   `json:"deleted"`
}
