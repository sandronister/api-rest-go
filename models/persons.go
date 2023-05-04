package models

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var People []Person
