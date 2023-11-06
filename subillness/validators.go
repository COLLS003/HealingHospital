package subillness

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

type SubIllnessModelValidator struct {
	SubIllness struct {
		Name      string `form:"name" json:"name" binding:"required,min=4"`
		IllnessID string `form:"illness" json:"illness" binding:"required"`
	} `json:"SubIllness"`
	SubIllnessModel SubIllnessModel `json:"-"`
}

// Bind binds the request data to the SubIllnessModelValidator.
func (v *SubIllnessModelValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, v)
	if err != nil {
		return err
	}

	v.SubIllnessModel.Name = v.SubIllness.Name
	v.SubIllnessModel.IllnessID = v.SubIllness.IllnessID

	return nil
}

// NewSubIllnessModelValidator creates a new SubIllnessModelValidator instance.
func NewSubIllnessModelValidator() SubIllnessModelValidator {
	return SubIllnessModelValidator{}
}

// NewSubIllnessModelValidatorFillWith creates a new SubIllnessModelValidator instance and fills it with SubIllness model data.
func NewSubIllnessModelValidatorFillWith(SubIllnessModel SubIllnessModel) SubIllnessModelValidator {
	return SubIllnessModelValidator{
		SubIllness: struct {
			Name      string `form:"name" json:"name" binding:"required,min=4"`
			IllnessID string `form:"illness" json:"illness" binding:"required"`
		}{
			Name:      SubIllnessModel.Name,
			IllnessID: SubIllnessModel.IllnessID,
		},
	}
}

// LoginValidator represents a validator for SubIllness login.
