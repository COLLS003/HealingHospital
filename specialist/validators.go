package specialist

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

type SpecialistModelValidator struct {
	Specialist struct {
		Name         string `form:"name" json:"name" binding:"required,min=4"`
		Email        string `form:"email" json:"email" binding:"required,email"`
		Password     string `form:"password" json:"password" binding:"required,min=4"`
		Category     string `form:"category" json:"category" binding:"required"`
		Availability string `form:"availability" json:"availability" binding:"required"`
		Location     string `form:"location" json:"location" binding:"required"`
	} `json:"Specialist"`
	SpecialistModel SpecialistModel `json:"-"`
}

// Bind binds the request data to the SpecialistModelValidator.
func (v *SpecialistModelValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, v)
	if err != nil {
		return err
	}

	v.SpecialistModel.Name = v.Specialist.Name
	v.SpecialistModel.Email = v.Specialist.Email
	v.SpecialistModel.Availability = v.Specialist.Availability
	v.SpecialistModel.Category = v.Specialist.Category
	v.SpecialistModel.Location = v.Specialist.Location

	if v.Specialist.Password != database.VOICESLIGHTSECRETPASSWORD {
		v.SpecialistModel.setPassword(v.Specialist.Password)
	}

	return nil
}

// NewSpecialistModelValidator creates a new SpecialistModelValidator instance.
func NewSpecialistModelValidator() SpecialistModelValidator {
	return SpecialistModelValidator{}
}

// NewSpecialistModelValidatorFillWith creates a new SpecialistModelValidator instance and fills it with Specialist model data.
func NewSpecialistModelValidatorFillWith(SpecialistModel SpecialistModel) SpecialistModelValidator {
	return SpecialistModelValidator{
		Specialist: struct {
			Name         string `form:"name" json:"name" binding:"required,min=4"`
			Email        string `form:"email" json:"email" binding:"required,email"`
			Password     string `form:"password" json:"password" binding:"required,min=4"`
			Category     string `form:"category" json:"category" binding:"required"`
			Availability string `form:"availability" json:"availability" binding:"required"`
			Location     string `form:"location" json:"location" binding:"required"`
		}{
			Name:         SpecialistModel.Name,
			Email:        SpecialistModel.Email,
			Password:     SpecialistModel.Password,
			Availability: SpecialistModel.Availability,
			Category:     SpecialistModel.Category,
			Location:     SpecialistModel.Location,
		},
	}
}

// LoginValidator represents a validator for Specialist login.
func phone() bool {
	return true
}

type LoginValidator struct {
	Specialist struct {
		Email    string `form:"email" json:"email" binding:"required,min=4"`
		Password string `form:"password" json:"password" binding:"required,min=2"`
	} `json:"Specialist"`
	SpecialistModel SpecialistModel `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, self)
	// err := database.Bind
	if err != nil {
		return err
	}

	self.SpecialistModel.Email = self.Specialist.Email
	return nil
}

// You can put the default value of a Validator here
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
