package patients

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"healing.hospital/database"
)

// case class Patients(ID: Int, name: String, age: Int, location: String, email: String, password: String)
type PatientModel struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"size:2048"`
	Location string `gorm:"size:2048"`
	Password string `gorm:"column:Password;not null"`
	Email    string `gorm:"column:email;unique_index"`
}

// Migrate the schema to the database if needed
func AutoMigrate() {
	db := database.GetConnection()
	db.AutoMigrate(&PatientModel{})
}

// FindSinglePatient finds a single Patient based on the provided condition
func FindSinglePatient(condition interface{}) (PatientModel, error) {
	db := database.GetConnection()
	var model PatientModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveSinglePatient saves a single Patient to the database
func SaveSinglePatient(data interface{}) error {
	db := database.GetConnection()
	err := db.Save(data).Error
	return err
}

// UpdateSinglePatient updates a Patient with new data
func UpdateSinglePatient(model *PatientModel, data interface{}) error {
	db := database.GetConnection()
	err := db.Model(model).Updates(data).Error
	return err
}

// DeleteSinglePatient deletes a Patient from the database
func DeleteSinglePatient(model *PatientModel) error {
	db := database.GetConnection()
	err := db.Delete(model).Error
	return err
}

// GetAllPatients gets all Patients from the database
func GetAllPatients() ([]PatientModel, error) {
	db := database.GetConnection()
	var models []PatientModel
	err := db.Find(&models).Error
	return models, err
}

// fix codwa
func GetPatientByID(id uint) (PatientModel, error) {
	db := database.GetConnection()
	var Patient PatientModel
	err := db.First(&Patient, id).Error
	return Patient, err
}

// get Patient by  email
func GetPatientByEmail(email string) (PatientModel, error) {
	db := database.GetConnection()
	var Patient PatientModel
	err := db.Where("Email = ?", email).First(&Patient).Error
	return Patient, err
}

// Patient dedicated function
func (u *PatientModel) setPassword(Password string) error {
	//check Password legnth
	if len(Password) == 0 {
		return errors.New("Password should never be empty")
	}
	bytePassword := []byte(Password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	PasswordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(PasswordHash)
	return nil
}

// if err := PatientModel.checkPassword("Password0"); err != nil { Password error }
func (u *PatientModel) checkPassword(Password string) error {
	bytePassword := []byte(Password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
