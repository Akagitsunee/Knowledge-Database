package model

type Community struct {
	CommunityId string `json:"id"`
	Name        string `json:"name"`
	AdminId     int    `json:"adminid"`
}
