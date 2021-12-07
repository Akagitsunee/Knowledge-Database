package model

type Space struct {
	ID   int    	`json:"id"`
	Name string 	`json:"name"`
	Admin User 		`json:"admin"`
	Users []User	`json:"users"`
}
