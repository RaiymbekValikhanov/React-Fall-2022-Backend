package handler

import (
	"net/http"
	cfg "project-backend/config"
	"project-backend/store"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	config *cfg.Config
	logger *zap.Logger
	store  store.Store
}

func NewHandler(config *cfg.Config, logger *zap.Logger, store store.Store) *Handler {
	return &Handler{
		config: config,
		logger: logger,
		store:  store,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.Default()
	s := cookie.NewStore([]byte(sessions.DefaultKey))
	s.Options(sessions.Options{
		SameSite: http.SameSiteNoneMode,
		Secure: true,
		Path: "/",
		HttpOnly: true,
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://192.168.1.110:3000", "https://traffic-rules.onrender.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(sessions.Sessions("session", s))
	r.GET("/health", h.HealthCheck)

	auth := r.Group("/auth")
	{
		auth.POST("/signup", h.Signup)
		auth.POST("/signin", h.Signin)
		auth.POST("/logout", h.LogOut)
		auth.GET("/whoami", h.WhoAmI)
	}

	r.GET("/profile",)
	r.GET("/rules", h.Rules)
	r.GET("/rules/:section", h.RuleSection)

	r.GET("/exams", h.AuthRequired(), h.Exams)
	r.GET("/exams/:section", h.AuthRequired(), h.ExamSection)

	r.GET("/scores", h.Scores)
	r.POST("/score", h.AddScore)

	return r
}
