package illnesssymptom

import (
	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

type IllnessSymptomsModelValidator struct {
	IllnessSymptoms struct {
		IllnessID  string `form:"illnessID" json:"illnessID" binding:"required"`
		SymptomsID string `form:"symptomID" json:"symptomID" binding:"required"`
	} `json:"illnessSymptoms"`
	IllnessSymptomsModel IllnessSymptomsModel `json:"-"`
}

// type IllnessSymptomsModelValidator struct {
// 	llnessSymptoms struct {
// 		Name        string `form:"name" json:"name" binding:"required,min=4"`
// 		Description string `form:"description" json:"description" binding:"required"`
// 	} `json:"Symptoms"`
// 	SymptomsModel SymptomsModel `json:"-"`
// }

// Bind binds the request data to the IllnessSymptomsModelValidator.
func (v *IllnessSymptomsModelValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, v)
	if err != nil {
		return err
	}

	v.IllnessSymptomsModel.IllnessID = v.IllnessSymptoms.IllnessID
	v.IllnessSymptomsModel.SymptomsID = v.IllnessSymptoms.SymptomsID

	return nil
}

// NewIllnessSymptomsModelValidator creates a new IllnessSymptomsModelValidator instance.
func NewIllnessSymptomsModelValidator() IllnessSymptomsModelValidator {
	return IllnessSymptomsModelValidator{}
}

// NewIllnessSymptomsModelValidatorFillWith creates a new IllnessSymptomsModelValidator instance and fills it with llnessSymptoms model data.
func NewIllnessSymptomsModelValidatorFillWith(IllnessSymptomsModel IllnessSymptomsModel) IllnessSymptomsModelValidator {
	return IllnessSymptomsModelValidator{
		IllnessSymptoms: struct {
			IllnessID  string `form:"illnessID" json:"illnessID" binding:"required"`
			SymptomsID string `form:"symptomID" json:"symptomID" binding:"required"`
		}{
			IllnessID:  IllnessSymptomsModel.IllnessID,
			SymptomsID: IllnessSymptomsModel.SymptomsID,
		},
	}
}

// LoginValidator represents a validator for llnessSymptoms login.
