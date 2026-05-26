package main

import (
	"ERP_APPS/config"
	role "ERP_APPS/internal"
	"ERP_APPS/internal/role/delivery"
	"ERP_APPS/internal/role/repository"
	"ERP_APPS/internal/role/usecase"
	"log"

	"github.com/gin-gonic/gin"

	"ERP_APPS/pkg/database"
)
func main(){
	cfg, err := config.LoadConfig()
	if err!= nil{
		log.Fatal(err)
	}
	db, err := database.InitPostgres(&cfg)
	if err != nil{
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&role.Role{}); err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	repo := repository.NewRoleRepository(db)
	usecase := usecase.NewRoleUsecase(repo) 
	api := r.Group("/api/v1")
	delivery.NewRoleHandler(api, usecase)
	log.Println("server running at :8080")
	log.Default().Fatal(r.Run(":" + cfg.APP_PORT))


}