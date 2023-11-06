package specialist

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

// A helper to write Specialist_id and Specialist_model to the context
func UpdateContextSpecialistModel(c *gin.Context, my_Specialist_id uint) {
	var mySpecialistModel SpecialistModel
	if my_Specialist_id != 0 {
		db := database.GetConnection()
		db.First(&mySpecialistModel, my_Specialist_id)
	}
	c.Set("my_Specialist_id", my_Specialist_id)
	c.Set("my_Specialist_model", mySpecialistModel)
}
