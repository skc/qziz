package main

import (
	"github.com/gin-gonic/gin"
	"qziz/handlers"
	"qziz/models"
)

var questions []models.Question
var questionsHandler *handlers.QuestionsHandler

func init() {
	questions = make([]models.Question, 0)
	questionsHandler = handlers.NewQuestionHandler(questions)
}

func main() {
	router := gin.Default()
	router.POST("/questions", questionsHandler.AddNewQuestion)
	router.GET("/questions", questionsHandler.ListQuestions)
	router.Run()
}
