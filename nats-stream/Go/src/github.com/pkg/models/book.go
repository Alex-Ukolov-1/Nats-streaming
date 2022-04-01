package models

type Book struct{
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"authot"`
	Desc   string `json:"desc"`
}

