package models

import "gorm.io/gorm"

//Book model
type Book struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}

//QueryInfo
type QueryInfo struct {
	gorm.Model

	SQL         string  `json:"text"`
	Type        string  `json:"type"`
	CostSeconds string  `json:"performance"`
}
