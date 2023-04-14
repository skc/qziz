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
