package models

type Persona struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var Personas []Persona
