package subillness



import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

// A helper to write SubIllness_id and SubIllness_model to the context
func UpdateContextSubIllnessModel(c *gin.Context, my_SubIllness_id uint) {
	var mySubIllnessModel SubIllnessModel
	if my_SubIllness_id != 0 {
		db := database.GetConnection()
		db.First(&mySubIllnessModel, my_SubIllness_id)
	}
	c.Set("my_SubIllness_id", my_SubIllness_id)
	c.Set("my_SubIllness_model", mySubIllnessModel)
}
