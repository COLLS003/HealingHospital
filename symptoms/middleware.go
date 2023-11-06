package symptoms

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

// A helper to write Symptoms_id and Symptoms_model to the context
func UpdateContextSymptomsModel(c *gin.Context, my_Symptoms_id uint) {
	var mySymptomsModel SymptomsModel
	if my_Symptoms_id != 0 {
		db := database.GetConnection()
		db.First(&mySymptomsModel, my_Symptoms_id)
	}
	c.Set("my_Symptoms_id", my_Symptoms_id)
	c.Set("my_Symptoms_model", mySymptomsModel)
}
