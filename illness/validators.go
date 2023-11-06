package illness

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

type IllnessModelValidator struct {
	Illness struct {
		Name        string `form:"name" json:"name" binding:"required,min=4"`
		Description string `form:"description" json:"description" binding:"required"`
	} `json:"Illness"`
	IllnessModel IllnessModel `json:"-"`
}

// Bind binds the request data to the IllnessModelValidator.
func (v *IllnessModelValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, v)
	if err != nil {
		return err
	}

	v.IllnessModel.Name = v.Illness.Name
	v.IllnessModel.Description = v.Illness.Description

	return nil
}

// NewIllnessModelValidator creates a new IllnessModelValidator instance.
func NewIllnessModelValidator() IllnessModelValidator {
	return IllnessModelValidator{}
}

// NewIllnessModelValidatorFillWith creates a new IllnessModelValidator instance and fills it with Illness model data.
func NewIllnessModelValidatorFillWith(IllnessModel IllnessModel) IllnessModelValidator {
	return IllnessModelValidator{
		Illness: struct {
			Name        string `form:"name" json:"name" binding:"required,min=4"`
			Description string `form:"description" json:"description" binding:"required"`
		}{
			Name:        IllnessModel.Name,
			Description: IllnessModel.Description,
		},
	}
}

// LoginValidator represents a validator for Illness login.
