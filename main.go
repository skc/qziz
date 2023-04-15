package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"qziz/handlers"
	"qziz/models"
)

var questions []models.Question
var questionsHandler *handlers.QuestionsHandler

func init() {
	questions = make([]models.Question, 0)
	file, err := os.ReadFile("questions.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal([]byte(file), &questions)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(len(questions))
	questionsHandler = handlers.NewQuestionHandler(questions)
}

func main() {
	router := gin.Default()
	router.POST("/questions", questionsHandler.AddNewQuestion)
	router.GET("/questions", questionsHandler.ListQuestions)
	router.PUT("/questions/:id", questionsHandler.UpdateQuestion)
	router.DELETE("/questions/:id", questionsHandler.DeleteQuestion)
	router.GET("/questions/search", questionsHandler.Search)
	router.Run()
}
