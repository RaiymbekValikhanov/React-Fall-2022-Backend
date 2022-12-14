package main

import (
	"project-backend/config"
	"project-backend/handler"
	"project-backend/store/orm"

	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig("config.yaml")
	log, _ := zap.NewDevelopment()
	defer log.Sync()

	db := orm.NewDB()
	if err := db.Connect(cfg); err != nil {
		panic(err)
	}

	h := handler.NewHandler(cfg, log, db)

	srv := h.InitRouter()
	srv.Run(cfg.HttpPort)
}