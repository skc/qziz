package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"qziz/models"
	"time"
)

type QuestionsHandler struct {
	questions []models.Question
}

func NewQuestionHandler(questions []models.Question) *QuestionsHandler {
	return &QuestionsHandler{
		questions: questions,
	}
}

func (h *QuestionsHandler) ListQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, h.questions)
}

func (h *QuestionsHandler) AddNewQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	question.ID = xid.New().String()
	question.PublishedAt = time.Now()
	question.PublishedBy = "057724817"
	h.questions = append(h.questions, question)
	c.JSON(http.StatusOK, question)
}

func (h *QuestionsHandler) UpdateQuestion(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	index := -1
	for i, q := range h.questions {
		if q.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Question not found",
		})
	}
	h.questions[index] = question
	c.JSON(http.StatusOK, question)
}

func (h *QuestionsHandler) DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i, q := range h.questions {
		if q.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Question not found",
		})
	}
	h.questions = append(h.questions[:index], h.questions[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Question has been deleted",
	})
}
