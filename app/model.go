package app

type Student struct{
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Group   string `json:"group"`
}