package illnesssymptom

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

func Create(router *gin.RouterGroup) {
	router.POST("/create", CreatellnessSymptoms)
	router.GET("/read/:id", ReadSinglellnessSymptoms)
	router.PUT("/update/:id", UpdatellnessSymptoms)
	router.DELETE("/delete/:id", DeletellnessSymptoms)
	router.GET("/list", llnessSymptomssList)
}

func CreatellnessSymptoms(c *gin.Context) {
	modelValidator := NewIllnessSymptomsModelValidator()
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	if err := SaveSinglellnessSymptoms(&modelValidator.IllnessSymptomsModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.Set("my_llnessSymptoms_model", modelValidator.IllnessSymptomsModel)
	serializer := NewllnessSymptomsSerializer(c, modelValidator.IllnessSymptomsModel)
	c.JSON(http.StatusCreated, gin.H{"llnessSymptoms": serializer.Response()})

	fmt.Println("llnessSymptoms saved ...")
}

func ReadSinglellnessSymptoms(c *gin.Context) {
	llnessSymptomsID := c.Param("id")
	llnessSymptomsIDUint, err := strconv.ParseUint(llnessSymptomsID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid llnessSymptoms ID"})
		return
	}

	IllnessSymptomsModel, err := GetllnessSymptomsByID(uint(llnessSymptomsIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("llnessSymptoms", err))
		return
	}

	serializer := NewllnessSymptomsSerializer(c, IllnessSymptomsModel)
	c.JSON(http.StatusOK, gin.H{"llnessSymptoms": serializer.Response()})
}

func UpdatellnessSymptoms(c *gin.Context) {
	llnessSymptomsID := c.Param("id")
	llnessSymptomsIDUint, err := strconv.ParseUint(llnessSymptomsID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid llnessSymptoms ID"})
		return
	}
	IllnessSymptomsModel, err := GetllnessSymptomsByID(uint(llnessSymptomsIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("llnessSymptoms", err))
		return
	}

	// Bind and update IllnessSymptomsModel with new data
	modelValidator := NewIllnessSymptomsModelValidatorFillWith(IllnessSymptomsModel)
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Call UpdateSinglellnessSymptoms function with the IllnessSymptomsModel and updated data
	if err := UpdateSinglellnessSymptoms(&IllnessSymptomsModel, modelValidator.IllnessSymptomsModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	serializer := NewllnessSymptomsSerializer(c, IllnessSymptomsModel)
	c.JSON(http.StatusOK, gin.H{"llnessSymptoms": serializer.Response()})
}

func DeletellnessSymptoms(c *gin.Context) {
	llnessSymptomsID := c.Param("id")
	llnessSymptomsIDUint, err := strconv.ParseUint(llnessSymptomsID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid llnessSymptoms ID"})
		return
	}
	IllnessSymptomsModel, err := GetllnessSymptomsByID(uint(llnessSymptomsIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("llnessSymptoms", err))
		return
	}

	if err := DeleteSinglellnessSymptoms(&IllnessSymptomsModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "llnessSymptoms deleted successfully"})
}

func llnessSymptomssList(c *gin.Context) {
	llnessSymptomssModels, err := GetAllllnessSymptomss()
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("llnessSymptomss", err))
		return
	}
	serializer := NewllnessSymptomssSerializer(c, llnessSymptomssModels)
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"llnessSymptomss": response})
}
