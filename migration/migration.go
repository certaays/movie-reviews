package migration

import (
	"fmt"
	"learn-go/database"
	"learn-go/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
