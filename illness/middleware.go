package illness

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

// A helper to write Illness_id and Illness_model to the context
func UpdateContextIllnessModel(c *gin.Context, my_Illness_id uint) {
	var myIllnessModel IllnessModel
	if my_Illness_id != 0 {
		db := database.GetConnection()
		db.First(&myIllnessModel, my_Illness_id)
	}
	c.Set("my_Illness_id", my_Illness_id)
	c.Set("my_Illness_model", myIllnessModel)
}
