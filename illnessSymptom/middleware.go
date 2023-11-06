package illnesssymptom

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

// A helper to write llnessSymptoms_id and llnessSymptoms_model to the context
func UpdateContextIllnessSymptomsModel(c *gin.Context, my_llnessSymptoms_id uint) {
	var myIllnessSymptomsModel IllnessSymptomsModel
	if my_llnessSymptoms_id != 0 {
		db := database.GetConnection()
		db.First(&myIllnessSymptomsModel, my_llnessSymptoms_id)
	}
	c.Set("my_llnessSymptoms_id", my_llnessSymptoms_id)
	c.Set("my_llnessSymptoms_model", myIllnessSymptomsModel)
}
