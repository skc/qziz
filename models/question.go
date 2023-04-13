package main

import "time"

type Question struct {
	ID            string    `json:"id"`
	Text          string    `json:"text"`
	Answers       []string  `json:"answers"`
	CorrectAnswer []string  `json:"correctAnswer"`
	Tags          []string  `json:"tags"`
	PublishedAt   time.Time `json:"publishedAt"`
	PublishedBy   string    `json:"publishedBy"`
}
