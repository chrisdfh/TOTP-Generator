package fnc

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chrisdfh/totp/pkg/models"
	"github.com/xlzd/gotp"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//ReadJSON lee un archivo JSON y regresa un mapa
//del tipo map[string]string con los datos computados
func ReadJSON(filename string) map[string]string {

	// path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	path := "."

	jsonFile, err := os.Open(path + "/" + filename)
	if err != nil {
		fmt.Println(err)
		log.Fatal("No se consiguió el " + path + filename)
	}

	var data []models.TOTPData
	json.NewDecoder(jsonFile).Decode(&data)
	defer jsonFile.Close()

	dict_2fa := make(map[string]string)
	for k := 0; k < len(data); k++ {
		dict_2fa[data[k].Service] = data[k].Secret
	}

	return dict_2fa

}

//Decode2FA devuelve el código TOTP y el tiempo hasta que falta
//hasta la expiración, dado el código de respaldo
func Decode2FA(secret string) (string, int64) {
	totp := gotp.NewDefaultTOTP(strings.ReplaceAll(strings.ToUpper(secret), " ", ""))
	otp, nextTime := totp.NowWithExpiration()
	diff := time.Now().Unix() - nextTime

	return otp, -diff
}

//HashMD5 computa el código MD5 dado un string
func HashMD5(text string) string {
	hash := md5.New()
	io.WriteString(hash, text)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//IsPasswordOK compara una clave en plaintexy y otra codificada
//en MD5, devuelve true si ambos son iguales, caso contrario
//devuelve false
func IsPasswordOK(passwd string, dbPass string) bool {
	if dbPass == HashMD5(string(passwd)) {
		return true
	} else {
		return false
	}
}

func PrintTOTP(data []models.TOTPData) {
	var tiempoRestante int64
	var otp string
	for {
		fmt.Printf("%s\n%16s   %s%s\n", Bold, strings.ToUpper("Servicio"), "Codigo OTP", Reset)
		for k := 0; k < len(data); k++ {
			otp, tiempoRestante = Decode2FA(data[k].Secret)
			fmt.Printf("%16s %9s\n", data[k].Service, otp[:3]+" "+otp[3:])
			//PRIMEROS 3 CARACTERES, ESPACIO, ULTIMOS 3 -- otp[:3]+" "+otp[3:] --
			//con %-16s empuja el texto hacia la izquierda apartando 16 carácteres en total
		}
		//HACE BLOQUE COUNTDOWN
		for t := tiempoRestante; t > 0; t-- {
			var block string
			if t < 5 {
				block = strings.Repeat(BgRed+" ", int(t))
			} else if t < 10 {
				block = strings.Repeat(BgYellow+" ", int(t))
			} else {
				block = strings.Repeat(BgGreen+" ", int(t))
			}
			fmt.Printf("%s%s%s", LimpiaLinea, block, Reset)
			time.Sleep(1 * time.Second)
		}
		// cmd := exec.Command("clear") // COMANDO A EJECUTAR
		// cmd.Stdout = os.Stdout       // SALIDA DEL COMANDO, SERÁ LA SALIDA DEL SISTEMA (PANTALLA)
		// cmd.Run()                    // EJECUTA EL COMANDO
		fmt.Print(LimpiaPantalla, GotoScreen(0, 0)) //IGUAL QUE CLEAR, PERO CON CODIGOS ANSI
	}
}

func ConnectDB(filename string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		log.Fatal("ERROR CON BD")
	}
	return db
}
