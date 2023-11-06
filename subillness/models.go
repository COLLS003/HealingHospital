package subillness

// case class SubIllness(ID: Int, name: String, illness: Int)

import (
	"gorm.io/gorm"
	"healing.hospital/database"
)

// case class SubIllnesss(ID: Int, name: String, age: Int, location: String, email: String, password: String)
type SubIllnessModel struct {
	gorm.Model
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:2048"`
	IllnessID string `gorm:"size:2048"`
}

// Migrate the schema to the database if needed
func AutoMigrate() {
	db := database.GetConnection()
	db.AutoMigrate(&SubIllnessModel{})
}

// FindSingleSubIllness finds a single SubIllness based on the provided condition
func FindSingleSubIllness(condition interface{}) (SubIllnessModel, error) {
	db := database.GetConnection()
	var model SubIllnessModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveSingleSubIllness saves a single SubIllness to the database
func SaveSingleSubIllness(data interface{}) error {
	db := database.GetConnection()
	err := db.Save(data).Error
	return err
}

// UpdateSingleSubIllness updates a SubIllness with new data
func UpdateSingleSubIllness(model *SubIllnessModel, data interface{}) error {
	db := database.GetConnection()
	err := db.Model(model).Updates(data).Error
	return err
}

// DeleteSingleSubIllness deletes a SubIllness from the database
func DeleteSingleSubIllness(model *SubIllnessModel) error {
	db := database.GetConnection()
	err := db.Delete(model).Error
	return err
}

// GetAllSubIllnesss gets all SubIllnesss from the database
func GetAllSubIllnesss() ([]SubIllnessModel, error) {
	db := database.GetConnection()
	var models []SubIllnessModel
	err := db.Find(&models).Error
	return models, err
}

// fix codwa
func GetSubIllnessByID(id uint) (SubIllnessModel, error) {
	db := database.GetConnection()
	var SubIllness SubIllnessModel
	err := db.First(&SubIllness, id).Error
	return SubIllness, err
}

// get SubIllness by  email
func GetSubIllnessByEmail(email string) (SubIllnessModel, error) {
	db := database.GetConnection()
	var SubIllness SubIllnessModel
	err := db.Where("Email = ?", email).First(&SubIllness).Error
	return SubIllness, err
}
