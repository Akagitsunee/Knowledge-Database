package model

type Employee struct {
	EmployeeId string `json:"id"`
	UserId     string `json:"userId"`
	Name       string `json:"name"`
	Disabled   bool   `json:"disabled"`
}
