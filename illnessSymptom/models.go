package illnesssymptom

// case class IllnessllnessSymptoms(ID: Int, illness: Int, symptom: Int)

import (
	"gorm.io/gorm"
	"healing.hospital/database"
)

// case class llnessSymptomss(ID: Int, name: String, age: Int, location: String, email: String, password: String)
type IllnessSymptomsModel struct {
	gorm.Model
	ID         uint   `gorm:"primary_key"`
	IllnessID  string `gorm:"size:2048"`
	SymptomsID string `gorm:"size:2048"`
}

// Migrate the schema to the database if needed
func AutoMigrate() {
	db := database.GetConnection()
	db.AutoMigrate(&IllnessSymptomsModel{})
}

// FindSinglellnessSymptoms finds a single llnessSymptoms based on the provided condition
func FindSinglellnessSymptoms(condition interface{}) (IllnessSymptomsModel, error) {
	db := database.GetConnection()
	var model IllnessSymptomsModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveSinglellnessSymptoms saves a single llnessSymptoms to the database
func SaveSinglellnessSymptoms(data interface{}) error {
	db := database.GetConnection()
	err := db.Save(data).Error
	return err
}

// UpdateSinglellnessSymptoms updates a llnessSymptoms with new data
func UpdateSinglellnessSymptoms(model *IllnessSymptomsModel, data interface{}) error {
	db := database.GetConnection()
	err := db.Model(model).Updates(data).Error
	return err
}

// DeleteSinglellnessSymptoms deletes a llnessSymptoms from the database
func DeleteSinglellnessSymptoms(model *IllnessSymptomsModel) error {
	db := database.GetConnection()
	err := db.Delete(model).Error
	return err
}

// GetAllllnessSymptomss gets all llnessSymptomss from the database
func GetAllllnessSymptomss() ([]IllnessSymptomsModel, error) {
	db := database.GetConnection()
	var models []IllnessSymptomsModel
	err := db.Find(&models).Error
	return models, err
}

// fix codwa
func GetllnessSymptomsByID(id uint) (IllnessSymptomsModel, error) {
	db := database.GetConnection()
	var llnessSymptoms IllnessSymptomsModel
	err := db.First(&llnessSymptoms, id).Error
	return llnessSymptoms, err
}

// get llnessSymptoms by  email
func GetllnessSymptomsByEmail(email string) (IllnessSymptomsModel, error) {
	db := database.GetConnection()
	var llnessSymptoms IllnessSymptomsModel
	err := db.Where("Email = ?", email).First(&llnessSymptoms).Error
	return llnessSymptoms, err
}
