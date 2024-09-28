package models

type Table struct {
	Id       string `json:"id"`
	Capacity int    `json:"capacity"`
	Occupied bool   `json:"occupied"`
}
