package model

type User struct {
	ID     int    `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Space  Space  `json:"space"`
}
