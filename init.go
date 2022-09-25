package main

type student struct {
	id         int
	name       string
	address    string
	occupation string
	reason     string
}

var students = []student{}

func init() {
	students = []student{
		{id: 1, name: "Farrel Athaillah Putra", address: "Semarang", occupation: "Programmer", reason: "Golang is the best"},
		{id: 2, name: "Adi Sudirtayasa", address: "Semarang", occupation: "Programmer", reason: "Golang is the best"},
		{id: 3, name: "Adrian Metanoia Gawa", address: "Semarang", occupation: "Programmer", reason: "Golang is the best"},
		{id: 4, name: "Agung Chumaidi", address: "Semarang", occupation: "Programmer", reason: "Golang is the best"},
		{id: 5, name: "Andrean Pradana", address: "Semarang", occupation: "Programmer", reason: "Golang is the best"},
	}
}
