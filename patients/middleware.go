package patients

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

// A helper to write Patient_id and Patient_model to the context
func UpdateContextPatientModel(c *gin.Context, my_Patient_id uint) {
	var myPatientModel PatientModel
	if my_Patient_id != 0 {
		db := database.GetConnection()
		db.First(&myPatientModel, my_Patient_id)
	}
	c.Set("my_Patient_id", my_Patient_id)
	c.Set("my_Patient_model", myPatientModel)
}
