package symptoms

// case class Symptoms(ID: Int, name: String, description: String)

import (
	"gorm.io/gorm"
	"healing.hospital/database"
)

// case class Symptomss(ID: Int, name: String, age: Int, location: String, email: String, password: String)
type SymptomsModel struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:2048"`
	Description string `gorm:"size:2048"`
}

// Migrate the schema to the database if needed
func AutoMigrate() {
	db := database.GetConnection()
	db.AutoMigrate(&SymptomsModel{})
}

// FindSingleSymptoms finds a single Symptoms based on the provided condition
func FindSingleSymptoms(condition interface{}) (SymptomsModel, error) {
	db := database.GetConnection()
	var model SymptomsModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveSingleSymptoms saves a single Symptoms to the database
func SaveSingleSymptoms(data interface{}) error {
	db := database.GetConnection()
	err := db.Save(data).Error
	return err
}

// UpdateSingleSymptoms updates a Symptoms with new data
func UpdateSingleSymptoms(model *SymptomsModel, data interface{}) error {
	db := database.GetConnection()
	err := db.Model(model).Updates(data).Error
	return err
}

// DeleteSingleSymptoms deletes a Symptoms from the database
func DeleteSingleSymptoms(model *SymptomsModel) error {
	db := database.GetConnection()
	err := db.Delete(model).Error
	return err
}

// GetAllSymptomss gets all Symptomss from the database
func GetAllSymptomss() ([]SymptomsModel, error) {
	db := database.GetConnection()
	var models []SymptomsModel
	err := db.Find(&models).Error
	return models, err
}

// fix codwa
func GetSymptomsByID(id uint) (SymptomsModel, error) {
	db := database.GetConnection()
	var Symptoms SymptomsModel
	err := db.First(&Symptoms, id).Error
	return Symptoms, err
}

// get Symptoms by  email
func GetSymptomsByEmail(email string) (SymptomsModel, error) {
	db := database.GetConnection()
	var Symptoms SymptomsModel
	err := db.Where("Email = ?", email).First(&Symptoms).Error
	return Symptoms, err
}
