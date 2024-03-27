package migration

import (
	"fmt"
	"gofiberind/database"
	"gofiberind/database/migration"
)

func RunMigrate() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success to Migrate")
}
