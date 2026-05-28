package main

import (
	"ERP_APPS/config"

	admin "ERP_APPS/internal/admin"
	adminDelivery "ERP_APPS/internal/admin/delivery"
	adminRepository "ERP_APPS/internal/admin/repository"
	adminUsecase "ERP_APPS/internal/admin/usecase"

	role "ERP_APPS/internal"
	roleDelivery "ERP_APPS/internal/role/delivery"
	roleRepository "ERP_APPS/internal/role/repository"
	roleUsecase "ERP_APPS/internal/role/usecase"

	"ERP_APPS/pkg/database"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitPostgres(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// migrate
	if err := db.AutoMigrate(
		&role.Role{},
		&admin.Admin{},
	); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// ROLE MODULE
	roleRepo := roleRepository.NewRoleRepository(db)
	roleUC := roleUsecase.NewRoleUsecase(roleRepo)

	// ADMIN MODULE
	adminRepo := adminRepository.NewAdminRepository(db)
	adminUC := adminUsecase.NewAdminUsecase(adminRepo)

	api := r.Group("/api/v1")

	// register handler
	roleDelivery.NewRoleHandler(api, roleUC)
	adminDelivery.NewAdminHandler(api, adminUC)

	log.Println("server running at :" + cfg.APP_PORT)

	log.Fatal(r.Run(":" + cfg.APP_PORT))
}