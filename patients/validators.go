package patients

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

type PatientModelValidator struct {
	Patient struct {
		Name     string `form:"name" json:"name" binding:"required,min=4"`
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=4"`
		Location string `form:"location" json:"location" binding:"required"`
	} `json:"Patient"`
	PatientModel PatientModel `json:"-"`
}

// Bind binds the request data to the PatientModelValidator.
func (v *PatientModelValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, v)
	if err != nil {
		return err
	}

	v.PatientModel.Name = v.Patient.Name
	v.PatientModel.Email = v.Patient.Email
	v.PatientModel.Location = v.Patient.Location

	if v.Patient.Password != database.VOICESLIGHTSECRETPASSWORD {
		v.PatientModel.setPassword(v.Patient.Password)
	}

	return nil
}

// NewPatientModelValidator creates a new PatientModelValidator instance.
func NewPatientModelValidator() PatientModelValidator {
	return PatientModelValidator{}
}

// NewPatientModelValidatorFillWith creates a new PatientModelValidator instance and fills it with Patient model data.
func NewPatientModelValidatorFillWith(PatientModel PatientModel) PatientModelValidator {
	return PatientModelValidator{
		Patient: struct {
			Name     string `form:"name" json:"name" binding:"required,min=4"`
			Email    string `form:"email" json:"email" binding:"required,email"`
			Password string `form:"password" json:"password" binding:"required,min=4"`
			Location string `form:"location" json:"location" binding:"required"`
		}{
			Name:     PatientModel.Name,
			Email:    PatientModel.Email,
			Password: PatientModel.Password,
			Location: PatientModel.Location,
		},
	}
}

// LoginValidator represents a validator for Patient login.
func phone() bool {
	return true
}

type LoginValidator struct {
	Patient struct {
		Email    string `form:"email" json:"email" binding:"required,min=4"`
		Password string `form:"password" json:"password" binding:"required,min=2"`
	} `json:"Patient"`
	PatientModel PatientModel `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, self)
	// err := database.Bind
	if err != nil {
		return err
	}

	self.PatientModel.Email = self.Patient.Email
	return nil
}

// You can put the default value of a Validator here
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
