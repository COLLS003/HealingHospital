package illness

import (
	illnesssymptom "colls.labs.claire/illnessSymptom"
	"colls.labs.claire/symptoms"
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

// get symptoms ot a given illness
// GetSymptoms returns a list of symptoms for the illness with the given ID.

func (illness *IllnessModel) GetSymptoms() ([]symptoms.SymptomsModel, error) {
	db := database.GetConnection()

	// Create a slice to store the symptom details
	var symptoms1 []symptoms.SymptomsModel

	// Find the illness-symptom associations for the given illness ID
	var illnessSymptoms []illnesssymptom.IllnessSymptomsModel
	if err := db.Where("illness_id = ?", illness.ID).Find(&illnessSymptoms).Error; err != nil {
		return nil, err
	}

	// Iterate through the associations and fetch the symptom details
	for _, is := range illnessSymptoms {
		var symptom symptoms.SymptomsModel
		if err := db.Where("ID = ?", is.SymptomsID).First(&symptom).Error; err != nil {
			return nil, err
		}
		symptoms1 = append(symptoms1, symptom)
	}

	return symptoms1, nil
}
