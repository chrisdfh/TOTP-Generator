package models

//DE BASE DE DATOS
type Security struct {
	ID     int `gorm:"primary_key"`
	Passwd string
}

func (Security) TableName() string {
	return "security"
}

//DE BASE DE DATOS
type TOTPData struct {
	ID      int `gorm:"primary_key"`
	Service string
	Secret  string
}

func (TOTPData) TableName() string {
	return "TOTP"
}

type TemplateData struct {
	Service       string
	Secret        string
	OTP           string
	TimeRemaining int64
	Array         []TOTPData
}

type CuadroOTG struct {
	Service        string
	ServiceID      string
	Otp            string
	TiempoRestante int64
}

type ForTemplateOTG struct {
	Titulo string
	Cuerpo string
	Data   []CuadroOTG
}
