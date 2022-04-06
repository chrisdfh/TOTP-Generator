package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chrisdfh/totp/pkg/fnc"
	"github.com/gorilla/mux"

	"golang.org/x/term"
)

func main() {
	web := flag.Bool("w", false, "Habilita app web")
	flag.Parse()

	workPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	
	fileNameDB := workPath+"/OTPn.db"

	// REVISA SI EXISTE LA BASE DE DATOS
	if _,err:= os.Stat(fileNameDB); err != nil{
		log.Fatal(fnc.BrightRed,"Base de Datos no existe",fnc.Reset)
	}
	db := fnc.ConnectDB(fileNameDB)

	// OBTIENE DATA DE BASE DE DATOS
	security := fnc.ReadSecurity(db)

	var passwd string
	fmt.Print("Introduzca Contraseña: ")
	passEnByte, _ := term.ReadPassword(0)
	passwd = string(passEnByte)

	//REVISAR SI LA CLAVE COINCIDE CON LA ALMACENADA EN BD
	if !fnc.IsPasswordOK(passwd, security.Passwd) {
		log.Fatal("Clave Incorrecta")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", fnc.PrintTOTPw(db))
	r.HandleFunc("/pwd-db/{newPass}", fnc.UpdatePasswdw(db, security))
	r.HandleFunc("/add/{service}/{secret}", fnc.AddServicew(db))
	http.Handle("/", r)

	if *web {
		if _,err :=os.Stat(workPath+"/templateBox.html"); err != nil {
			log.Fatalln("\nArchivo de plantila HTML no encontrado")
		}
		hostname, _ := os.Hostname()
		puerto := ":8123"
		fmt.Printf("\nSirviendo por http://%s%s\n", hostname, puerto)
		http.ListenAndServe(puerto, r)
	}
	data := fnc.ReadData(db)
	fnc.PrintTOTP(data)
}
