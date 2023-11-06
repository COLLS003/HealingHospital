package symptoms

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

func Create(router *gin.RouterGroup) {
	router.POST("/create", CreateSymptoms)
	router.GET("/read/:id", ReadSingleSymptoms)
	router.PUT("/update/:id", UpdateSymptoms)
	router.DELETE("/delete/:id", DeleteSymptoms)
	router.GET("/list", SymptomssList)
}

func CreateSymptoms(c *gin.Context) {
	modelValidator := NewSymptomsModelValidator()
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	if err := SaveSingleSymptoms(&modelValidator.SymptomsModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.Set("my_Symptoms_model", modelValidator.SymptomsModel)
	serializer := NewSymptomsSerializer(c, modelValidator.SymptomsModel)
	c.JSON(http.StatusCreated, gin.H{"Symptoms": serializer.Response()})

	fmt.Println("Symptoms saved ...")
}

func ReadSingleSymptoms(c *gin.Context) {
	SymptomsID := c.Param("id")
	SymptomsIDUint, err := strconv.ParseUint(SymptomsID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Symptoms ID"})
		return
	}

	SymptomsModel, err := GetSymptomsByID(uint(SymptomsIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Symptoms", err))
		return
	}

	serializer := NewSymptomsSerializer(c, SymptomsModel)
	c.JSON(http.StatusOK, gin.H{"Symptoms": serializer.Response()})
}

func UpdateSymptoms(c *gin.Context) {
	SymptomsID := c.Param("id")
	SymptomsIDUint, err := strconv.ParseUint(SymptomsID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Symptoms ID"})
		return
	}
	SymptomsModel, err := GetSymptomsByID(uint(SymptomsIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Symptoms", err))
		return
	}

	// Bind and update SymptomsModel with new data
	modelValidator := NewSymptomsModelValidatorFillWith(SymptomsModel)
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Call UpdateSingleSymptoms function with the SymptomsModel and updated data
	if err := UpdateSingleSymptoms(&SymptomsModel, modelValidator.SymptomsModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	serializer := NewSymptomsSerializer(c, SymptomsModel)
	c.JSON(http.StatusOK, gin.H{"Symptoms": serializer.Response()})
}

func DeleteSymptoms(c *gin.Context) {
	SymptomsID := c.Param("id")
	SymptomsIDUint, err := strconv.ParseUint(SymptomsID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Symptoms ID"})
		return
	}
	SymptomsModel, err := GetSymptomsByID(uint(SymptomsIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Symptoms", err))
		return
	}

	if err := DeleteSingleSymptoms(&SymptomsModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Symptoms deleted successfully"})
}

func SymptomssList(c *gin.Context) {
	SymptomssModels, err := GetAllSymptomss()
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Symptomss", err))
		return
	}
	serializer := NewSymptomssSerializer(c, SymptomssModels)
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"Symptomss": response})
}
