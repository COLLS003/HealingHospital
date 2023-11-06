package illness

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

func Create(router *gin.RouterGroup) {
	router.POST("/create", CreateIllness)
	router.GET("/read/:id", ReadSingleIllness)
	router.GET("/read/symptoms/:id", ReadSymptoms)
	router.PUT("/update/:id", UpdateIllness)
	router.DELETE("/delete/:id", DeleteIllness)
	router.GET("/list", IllnesssList)
}

func CreateIllness(c *gin.Context) {
	modelValidator := NewIllnessModelValidator()
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	if err := SaveSingleIllness(&modelValidator.IllnessModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.Set("my_Illness_model", modelValidator.IllnessModel)
	serializer := NewIllnessSerializer(c, modelValidator.IllnessModel)
	c.JSON(http.StatusCreated, gin.H{"Illness": serializer.Response()})

	fmt.Println("Illness saved ...")
}

func ReadSingleIllness(c *gin.Context) {
	IllnessID := c.Param("id")
	IllnessIDUint, err := strconv.ParseUint(IllnessID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Illness ID"})
		return
	}

	IllnessModel, err := GetIllnessByID(uint(IllnessIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Illness", err))
		return
	}

	serializer := NewIllnessSerializer(c, IllnessModel)
	c.JSON(http.StatusOK, gin.H{"Illness": serializer.Response()})
}

func UpdateIllness(c *gin.Context) {
	IllnessID := c.Param("id")
	IllnessIDUint, err := strconv.ParseUint(IllnessID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Illness ID"})
		return
	}
	IllnessModel, err := GetIllnessByID(uint(IllnessIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Illness", err))
		return
	}

	// Bind and update IllnessModel with new data
	modelValidator := NewIllnessModelValidatorFillWith(IllnessModel)
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Call UpdateSingleIllness function with the IllnessModel and updated data
	if err := UpdateSingleIllness(&IllnessModel, modelValidator.IllnessModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	serializer := NewIllnessSerializer(c, IllnessModel)
	c.JSON(http.StatusOK, gin.H{"Illness": serializer.Response()})
}

func DeleteIllness(c *gin.Context) {
	IllnessID := c.Param("id")
	IllnessIDUint, err := strconv.ParseUint(IllnessID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Illness ID"})
		return
	}
	IllnessModel, err := GetIllnessByID(uint(IllnessIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Illness", err))
		return
	}

	if err := DeleteSingleIllness(&IllnessModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Illness deleted successfully"})
}

func IllnesssList(c *gin.Context) {
	IllnesssModels, err := GetAllIllnesss()
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Illnesss", err))
		return
	}
	serializer := NewIllnesssSerializer(c, IllnesssModels)
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"Illnesss": response})
}

type SymptomsResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ReadSymptoms(c *gin.Context) {
	IllnessID := c.Param("id")
	IllnessIDUint, err := strconv.ParseUint(IllnessID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Illness ID"})
		return
	}

	// Create an instance of IllnessModel with the provided ID
	illness := &IllnessModel{ID: uint(IllnessIDUint)}

	// Call the GetSymptoms function on the illness instance
	symptoms, err := illness.GetSymptoms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Symptoms", err))
		return
	}

	// Create a custom response structure
	var response []SymptomsResponse
	for _, symptom := range symptoms {
		response = append(response, SymptomsResponse{
			ID:          symptom.ID,
			Name:        symptom.Name,
			Description: symptom.Description,
		})
	}

	c.JSON(http.StatusOK, gin.H{"Symptoms": response})
}
