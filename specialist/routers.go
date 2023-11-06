package specialist

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
	router.POST("/create", CreateSpecialist)
	router.POST("/login", Login)
	router.GET("/read/:id", ReadSingleSpecialist)
	router.GET("/reset/:email", PasswordReset)
	router.PUT("/update/:id", UpdateSpecialist)
	router.DELETE("/delete/:id", DeleteSpecialist)
	router.GET("/list", SpecialistsList)
}

func CreateSpecialist(c *gin.Context) {
	modelValidator := NewSpecialistModelValidator()
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	if err := SaveSingleSpecialist(&modelValidator.SpecialistModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.Set("my_Specialist_model", modelValidator.SpecialistModel)
	serializer := NewSpecialistSerializer(c, modelValidator.SpecialistModel)
	c.JSON(http.StatusCreated, gin.H{"Specialist": serializer.Response()})

	fmt.Println("Specialist saved ...")
}

func Login(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Find Specialist by phone
	SpecialistModel, err := FindSingleSpecialist(&SpecialistModel{Email: loginValidator.Specialist.Email})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusForbidden, database.NewError("login", errors.New("Not Registered email or invalid password")))
		} else {
			c.JSON(http.StatusInternalServerError, database.NewError("login", err))
		}
		return
	}

	// Check password
	if err := SpecialistModel.checkPassword(loginValidator.Specialist.Password); err != nil {
		c.JSON(http.StatusForbidden, database.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	UpdateContextSpecialistModel(c, SpecialistModel.ID)
	serializer := SpecialistSerializer{c, SpecialistModel}
	c.JSON(http.StatusOK, gin.H{"Specialist": serializer.Response()})
}

func ReadSingleSpecialist(c *gin.Context) {
	SpecialistID := c.Param("id")
	SpecialistIDUint, err := strconv.ParseUint(SpecialistID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Specialist ID"})
		return
	}

	SpecialistModel, err := GetSpecialistByID(uint(SpecialistIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Specialist", err))
		return
	}

	serializer := NewSpecialistSerializer(c, SpecialistModel)
	c.JSON(http.StatusOK, gin.H{"Specialist": serializer.Response()})
}

func UpdateSpecialist(c *gin.Context) {
	SpecialistID := c.Param("id")
	SpecialistIDUint, err := strconv.ParseUint(SpecialistID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Specialist ID"})
		return
	}
	SpecialistModel, err := GetSpecialistByID(uint(SpecialistIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Specialist", err))
		return
	}

	// Bind and update SpecialistModel with new data
	modelValidator := NewSpecialistModelValidatorFillWith(SpecialistModel)
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Call UpdateSingleSpecialist function with the SpecialistModel and updated data
	if err := UpdateSingleSpecialist(&SpecialistModel, modelValidator.SpecialistModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	serializer := NewSpecialistSerializer(c, SpecialistModel)
	c.JSON(http.StatusOK, gin.H{"Specialist": serializer.Response()})
}

func DeleteSpecialist(c *gin.Context) {
	SpecialistID := c.Param("id")
	SpecialistIDUint, err := strconv.ParseUint(SpecialistID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Specialist ID"})
		return
	}
	SpecialistModel, err := GetSpecialistByID(uint(SpecialistIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Specialist", err))
		return
	}

	if err := DeleteSingleSpecialist(&SpecialistModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Specialist deleted successfully"})
}

func SpecialistsList(c *gin.Context) {
	SpecialistsModels, err := GetAllSpecialists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Specialists", err))
		return
	}
	serializer := NewSpecialistsSerializer(c, SpecialistsModels)
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"Specialists": response})
}

// additional functions added by colls_Codes at sep 6th at 16:33 for password reset using cradlevoices
func PasswordReset(c *gin.Context) {
	email := c.Param("email")
	// SpecialistIDUint, err := strconv.ParseString(email, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address .."})
	// 	return
	// }

	SpecialistModel, err := GetSpecialistByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Specialist", err))
		return
	}
	//generate random password something like a otp
	randomPswd := RandomOtpGenerator()
	fmt.Println(randomPswd)

	serializer := NewSpecialistSerializer(c, SpecialistModel)
	c.JSON(http.StatusOK, gin.H{"Specialist": serializer.Response()})
}

// random password generator
func RandomOtpGenerator() (otp int) {
	// Generate a random integer between 100000 and 999999 (inclusive)
	otpd := rand.Intn(900000) + 100000
	return otpd
}
