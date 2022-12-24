package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"project-backend/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) Rules(c *gin.Context) {
	data, err := ioutil.ReadFile(h.config.RulesPath + "sections.json")
	if err != nil {
		go h.logger.Error("read from file error")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, data)
}

func (h *Handler) RuleSection(c *gin.Context) {
	section := c.Param("section")
	
	data, err := ioutil.ReadFile(h.config.RulesPath + fmt.Sprintf("%s.html", section))
	if err != nil {
		go h.logger.Error("read from file error")
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Тут пока пусто..."))
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}

func (h *Handler) Exams(c *gin.Context) {
	data, err := ioutil.ReadFile(h.config.ExamsPath + "sections.json")
	if err != nil {
		go h.logger.Error("read from file error")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, data)
}

func (h *Handler) ExamSection(c *gin.Context) {
	section := c.Param("section")

	data, err := ioutil.ReadFile(h.config.ExamsPath + fmt.Sprintf("%s.json", section))
	if err != nil {
		go h.logger.Error("read from file error")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, data)
}

func (h *Handler) Scores(c *gin.Context) {
	user := sessions.Default(c).Get("user")

	scores, err := h.store.ScoreRepository().ScoresByUser(user.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return 
	}
	
	c.JSON(http.StatusOK, scores)
	return
}

func (h *Handler) AddScore(c *gin.Context) {
	var score model.Score

	if err := c.BindJSON(&score); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	score.UserID = sessions.Default(c).Get("user").(uint)
	if err := h.store.ScoreRepository().Create(&score); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "new score"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "new score"})
	return
}