package fnc

import (
	"github.com/chrisdfh/totp/pkg/models"
	"gorm.io/gorm"
)

// Codigo para actualizar data en security U en CRUD
// security.Passwd = "3c77b9358a5b87b75f6d762900b5f730"
// db.Save(&security)

// Código para crear data en GORM C en CRUD
// var newSec models.Security
// newSec.Passwd = "123456789"
// db.Create(&newSec)
// fmt.Println("New ID ", newSec.ID)

// Código para eliminar data en GORM D en CRUD
// var sec2 models.Security
// db.Find(&sec2, 2) //id=2
// db.Delete(&sec2)

func ReadSecurity(db *gorm.DB) models.Security {
	var security models.Security
	db.Find(&security, 1)
	return security
}

func ReadData(db *gorm.DB) []models.TOTPData {
	var data []models.TOTPData
	db.Find(&data)
	return data
}

func UpdatePasswd(db *gorm.DB, security models.Security, newPass string) {
	security.Passwd = HashMD5(newPass)
	db.Save(&security)
}

func SaveService(db *gorm.DB, data models.TOTPData) {
	db.Save(&data)
}

func DeleteService(db *gorm.DB, data models.TOTPData) {
	db.Delete(&data)
}

func AddService(db *gorm.DB, data models.TOTPData) {
	db.Create(&data)
}
