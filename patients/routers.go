package patients

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"healing.hospital/database"
)

func Create(router *gin.RouterGroup) {
	router.POST("/create", CreatePatient)
	router.POST("/login", Login)
	router.GET("/read/:id", ReadSinglePatient)
	router.GET("/reset/:email", PasswordReset)
	router.PUT("/update/:id", UpdatePatient)
	router.DELETE("/delete/:id", DeletePatient)
	router.GET("/list", PatientsList)
}

func CreatePatient(c *gin.Context) {
	modelValidator := NewPatientModelValidator()
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	if err := SaveSinglePatient(&modelValidator.PatientModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.Set("my_Patient_model", modelValidator.PatientModel)
	serializer := NewPatientSerializer(c, modelValidator.PatientModel)
	c.JSON(http.StatusCreated, gin.H{"Patient": serializer.Response()})

	fmt.Println("Patient saved ...")
}

func Login(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Find Patient by phone
	PatientModel, err := FindSinglePatient(&PatientModel{Email: loginValidator.Patient.Email})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusForbidden, database.NewError("login", errors.New("Not Registered email or invalid password")))
		} else {
			c.JSON(http.StatusInternalServerError, database.NewError("login", err))
		}
		return
	}

	// Check password
	if err := PatientModel.checkPassword(loginValidator.Patient.Password); err != nil {
		c.JSON(http.StatusForbidden, database.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	UpdateContextPatientModel(c, PatientModel.ID)
	serializer := PatientSerializer{c, PatientModel}
	c.JSON(http.StatusOK, gin.H{"Patient": serializer.Response()})
}

func ReadSinglePatient(c *gin.Context) {
	PatientID := c.Param("id")
	PatientIDUint, err := strconv.ParseUint(PatientID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Patient ID"})
		return
	}

	PatientModel, err := GetPatientByID(uint(PatientIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Patient", err))
		return
	}

	serializer := NewPatientSerializer(c, PatientModel)
	c.JSON(http.StatusOK, gin.H{"Patient": serializer.Response()})
}

func UpdatePatient(c *gin.Context) {
	PatientID := c.Param("id")
	PatientIDUint, err := strconv.ParseUint(PatientID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Patient ID"})
		return
	}
	PatientModel, err := GetPatientByID(uint(PatientIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Patient", err))
		return
	}

	// Bind and update PatientModel with new data
	modelValidator := NewPatientModelValidatorFillWith(PatientModel)
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Call UpdateSinglePatient function with the PatientModel and updated data
	if err := UpdateSinglePatient(&PatientModel, modelValidator.PatientModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	serializer := NewPatientSerializer(c, PatientModel)
	c.JSON(http.StatusOK, gin.H{"Patient": serializer.Response()})
}

func DeletePatient(c *gin.Context) {
	PatientID := c.Param("id")
	PatientIDUint, err := strconv.ParseUint(PatientID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Patient ID"})
		return
	}
	PatientModel, err := GetPatientByID(uint(PatientIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Patient", err))
		return
	}

	if err := DeleteSinglePatient(&PatientModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

func PatientsList(c *gin.Context) {
	PatientsModels, err := GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Patients", err))
		return
	}
	serializer := NewPatientsSerializer(c, PatientsModels)
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"Patients": response})
}

// additional functions added by colls_Codes at sep 6th at 16:33 for password reset using cradlevoices
func PasswordReset(c *gin.Context) {
	email := c.Param("email")
	// PatientIDUint, err := strconv.ParseString(email, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address .."})
	// 	return
	// }

	PatientModel, err := GetPatientByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Patient", err))
		return
	}
	//generate random password something like a otp
	randomPswd := RandomOtpGenerator()
	fmt.Println(randomPswd)

	serializer := NewPatientSerializer(c, PatientModel)
	c.JSON(http.StatusOK, gin.H{"Patient": serializer.Response()})
}

// random password generator
func RandomOtpGenerator() (otp int) {
	// Generate a random integer between 100000 and 999999 (inclusive)
	otpd := rand.Intn(900000) + 100000
	return otpd
}
