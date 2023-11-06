package specialist

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"healing.hospital/database"
)

// case class Specialist(ID: Int, name: String, category: Int, location: String, email: String, availability: Int)

type SpecialistModel struct {
	gorm.Model
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"size:2048"`
	Category     string `gorm:"size:2048"`
	Availability string `gorm:"size:2048"`
	Location     string `gorm:"size:2048"`
	Password     string `gorm:"column:Password;not null"`
	Email        string `gorm:"column:email;unique_index"`
}

// Migrate the schema to the database if needed
func AutoMigrate() {
	db := database.GetConnection()
	db.AutoMigrate(&SpecialistModel{})
}

// FindSingleSpecialist finds a single Specialist based on the provided condition
func FindSingleSpecialist(condition interface{}) (SpecialistModel, error) {
	db := database.GetConnection()
	var model SpecialistModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveSingleSpecialist saves a single Specialist to the database
func SaveSingleSpecialist(data interface{}) error {
	db := database.GetConnection()
	err := db.Save(data).Error
	return err
}

// UpdateSingleSpecialist updates a Specialist with new data
func UpdateSingleSpecialist(model *SpecialistModel, data interface{}) error {
	db := database.GetConnection()
	err := db.Model(model).Updates(data).Error
	return err
}

// DeleteSingleSpecialist deletes a Specialist from the database
func DeleteSingleSpecialist(model *SpecialistModel) error {
	db := database.GetConnection()
	err := db.Delete(model).Error
	return err
}

// GetAllSpecialists gets all Specialists from the database
func GetAllSpecialists() ([]SpecialistModel, error) {
	db := database.GetConnection()
	var models []SpecialistModel
	err := db.Find(&models).Error
	return models, err
}

// fix codwa
func GetSpecialistByID(id uint) (SpecialistModel, error) {
	db := database.GetConnection()
	var Specialist SpecialistModel
	err := db.First(&Specialist, id).Error
	return Specialist, err
}

// get Specialist by  email
func GetSpecialistByEmail(email string) (SpecialistModel, error) {
	db := database.GetConnection()
	var Specialist SpecialistModel
	err := db.Where("Email = ?", email).First(&Specialist).Error
	return Specialist, err
}

// Specialist dedicated function
func (u *SpecialistModel) setPassword(Password string) error {
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

// if err := SpecialistModel.checkPassword("Password0"); err != nil { Password error }
func (u *SpecialistModel) checkPassword(Password string) error {
	bytePassword := []byte(Password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
