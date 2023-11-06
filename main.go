package main

import (
	"colls.labs.claire/illness"
	illnesssymptom "colls.labs.claire/illnessSymptom"
	"colls.labs.claire/specialist"
	"colls.labs.claire/subillness"
	"colls.labs.claire/symptoms"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"healing.hospital/database"
	"healing.hospital/patients"
)

// models migration
func Migration(database *gorm.DB) {
	patients.AutoMigrate()
	specialist.AutoMigrate()
	illness.AutoMigrate()
	illnesssymptom.AutoMigrate()
	subillness.AutoMigrate()
	symptoms.AutoMigrate()

}

func main() {
	// Initialize the database connection
	database := database.Init()
	Migration(database)

	router := gin.Default()

	// CORS configuration :updated on august 28 17:16 pm
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	sun := router.Group("/Api")
	patients.Create(sun.Group("/patients"))
	specialist.Create(sun.Group("/specialist"))
	illness.Create(sun.Group("/illness"))
	illnesssymptom.Create(sun.Group("/illnesssymptom"))
	subillness.Create(sun.Group("/subillness"))
	symptoms.Create(sun.Group("/symptoms"))

	// Run the server
	// if err := router.Run(":3000"); err != nil {
	// 	panic(err)
	// }

	if err := router.Run("0.0.0.0:3000"); err != nil {
		panic(err)
	}
}
