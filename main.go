package main

import (
	"gorm_crud/configs"
	"gorm_crud/database"
	"gorm_crud/models"
	"gorm_crud/repositories"
	"log"
)

func main() {

	dbUser, dbPassword, dbName := "root", "", "gorm_crud"
	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)

	if err != nil {
		log.Fatalln(err)
	}

	// ping to database
	err = db.DB().Ping()

	// error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	// migration
	db.AutoMigrate(&models.Contact{})

	defer db.Close()		

	contactRepository := repositories.NewContactRepository(db)
			
	route := configs.SetupRoutes(contactRepository)


	route.Run(":8000")
}
