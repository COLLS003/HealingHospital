package symptoms

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

type SymptomsModelValidator struct {
	Symptoms struct {
		Name        string `form:"name" json:"name" binding:"required,min=4"`
		Description string `form:"description" json:"description" binding:"required"`
	} `json:"Symptoms"`
	SymptomsModel SymptomsModel `json:"-"`
}

// Bind binds the request data to the SymptomsModelValidator.
func (v *SymptomsModelValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, v)
	if err != nil {
		return err
	}

	v.SymptomsModel.Name = v.Symptoms.Name
	v.SymptomsModel.Description = v.Symptoms.Description

	return nil
}

// NewSymptomsModelValidator creates a new SymptomsModelValidator instance.
func NewSymptomsModelValidator() SymptomsModelValidator {
	return SymptomsModelValidator{}
}

// NewSymptomsModelValidatorFillWith creates a new SymptomsModelValidator instance and fills it with Symptoms model data.
func NewSymptomsModelValidatorFillWith(SymptomsModel SymptomsModel) SymptomsModelValidator {
	return SymptomsModelValidator{
		Symptoms: struct {
			Name        string `form:"name" json:"name" binding:"required,min=4"`
			Description string `form:"description" json:"description" binding:"required"`
		}{
			Name:        SymptomsModel.Name,
			Description: SymptomsModel.Description,
		},
	}
}

// LoginValidator represents a validator for Symptoms login.
