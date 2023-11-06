package illness

// case class Illness(ID: Int, name: String, description: String)

// case class Illness(ID: Int, name: String, description: String)

import (
	"gorm.io/gorm"
	"healing.hospital/database"
)

// case class Illnesss(ID: Int, name: String, age: Int, location: String, email: String, password: String)
type IllnessModel struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:2048"`
	Description string `gorm:"size:2048"`
}

// Migrate the schema to the database if needed
func AutoMigrate() {
	db := database.GetConnection()
	db.AutoMigrate(&IllnessModel{})
}

// FindSingleIllness finds a single Illness based on the provided condition
func FindSingleIllness(condition interface{}) (IllnessModel, error) {
	db := database.GetConnection()
	var model IllnessModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveSingleIllness saves a single Illness to the database
func SaveSingleIllness(data interface{}) error {
	db := database.GetConnection()
	err := db.Save(data).Error
	return err
}

// UpdateSingleIllness updates a Illness with new data
func UpdateSingleIllness(model *IllnessModel, data interface{}) error {
	db := database.GetConnection()
	err := db.Model(model).Updates(data).Error
	return err
}

// DeleteSingleIllness deletes a Illness from the database
func DeleteSingleIllness(model *IllnessModel) error {
	db := database.GetConnection()
	err := db.Delete(model).Error
	return err
}

// GetAllIllnesss gets all Illnesss from the database
func GetAllIllnesss() ([]IllnessModel, error) {
	db := database.GetConnection()
	var models []IllnessModel
	err := db.Find(&models).Error
	return models, err
}

// fix codwa
func GetIllnessByID(id uint) (IllnessModel, error) {
	db := database.GetConnection()
	var Illness IllnessModel
	err := db.First(&Illness, id).Error
	return Illness, err
}

// get Illness by  email
func GetIllnessByEmail(email string) (IllnessModel, error) {
	db := database.GetConnection()
	var Illness IllnessModel
	err := db.Where("Email = ?", email).First(&Illness).Error
	return Illness, err
}
