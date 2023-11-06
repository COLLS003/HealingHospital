package subillness


import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"healing.hospital/database"
)

func Create(router *gin.RouterGroup) {
	router.POST("/create", CreateSubIllness)
	router.GET("/read/:id", ReadSingleSubIllness)
	router.PUT("/update/:id", UpdateSubIllness)
	router.DELETE("/delete/:id", DeleteSubIllness)
	router.GET("/list", SubIllnesssList)
}

func CreateSubIllness(c *gin.Context) {
	modelValidator := NewSubIllnessModelValidator()
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	if err := SaveSingleSubIllness(&modelValidator.SubIllnessModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.Set("my_SubIllness_model", modelValidator.SubIllnessModel)
	serializer := NewSubIllnessSerializer(c, modelValidator.SubIllnessModel)
	c.JSON(http.StatusCreated, gin.H{"SubIllness": serializer.Response()})

	fmt.Println("SubIllness saved ...")
}

func ReadSingleSubIllness(c *gin.Context) {
	SubIllnessID := c.Param("id")
	SubIllnessIDUint, err := strconv.ParseUint(SubIllnessID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SubIllness ID"})
		return
	}

	SubIllnessModel, err := GetSubIllnessByID(uint(SubIllnessIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("SubIllness", err))
		return
	}

	serializer := NewSubIllnessSerializer(c, SubIllnessModel)
	c.JSON(http.StatusOK, gin.H{"SubIllness": serializer.Response()})
}

func UpdateSubIllness(c *gin.Context) {
	SubIllnessID := c.Param("id")
	SubIllnessIDUint, err := strconv.ParseUint(SubIllnessID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SubIllness ID"})
		return
	}
	SubIllnessModel, err := GetSubIllnessByID(uint(SubIllnessIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("SubIllness", err))
		return
	}

	// Bind and update SubIllnessModel with new data
	modelValidator := NewSubIllnessModelValidatorFillWith(SubIllnessModel)
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Call UpdateSingleSubIllness function with the SubIllnessModel and updated data
	if err := UpdateSingleSubIllness(&SubIllnessModel, modelValidator.SubIllnessModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	serializer := NewSubIllnessSerializer(c, SubIllnessModel)
	c.JSON(http.StatusOK, gin.H{"SubIllness": serializer.Response()})
}

func DeleteSubIllness(c *gin.Context) {
	SubIllnessID := c.Param("id")
	SubIllnessIDUint, err := strconv.ParseUint(SubIllnessID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SubIllness ID"})
		return
	}
	SubIllnessModel, err := GetSubIllnessByID(uint(SubIllnessIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("SubIllness", err))
		return
	}

	if err := DeleteSingleSubIllness(&SubIllnessModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SubIllness deleted successfully"})
}

func SubIllnesssList(c *gin.Context) {
	SubIllnesssModels, err := GetAllSubIllnesss()
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("SubIllnesss", err))
		return
	}
	serializer := NewSubIllnesssSerializer(c, SubIllnesssModels)
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"SubIllnesss": response})
}
