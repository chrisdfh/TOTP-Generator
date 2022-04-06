package fnc

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/chrisdfh/totp/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PrintTOTPw(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var data []models.TOTPData
		db.Find(&data)
		w.WriteHeader(http.StatusOK)

		var tiempoRestante int64
		var otp string

		var dataOTG []models.CuadroOTG
		for k := 0; k < len(data); k++ {
			otp, tiempoRestante = Decode2FA(data[k].Secret)
			dataOTG = append(dataOTG,
				models.CuadroOTG{
					Service:        data[k].Service,
					ServiceID:      strings.ReplaceAll(data[k].Service, " ", ""),
					Otp:            otp,
					TiempoRestante: tiempoRestante,
				})
		}

		dataCuadro := models.ForTemplateOTG{
			Titulo: "Autentificación de 2 pasos",
			Cuerpo: "Códigos 2FA",
			Data:   dataOTG,
		}

		workPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		tmpl, _ := template.ParseFiles(workPath+"/templateBox.html")
		tmpl.Execute(w, dataCuadro)

		// PARA USAR UNA PLANTILLA DENTRO DE LA FUNCION:
		// plantilla := `
		// 	<!DOCTYPE html>
		// 	<html lang="en">
		// 		<head>
		// 			<meta charset="UTF-8" />
		// 			<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		// 			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		// 			<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" />
		// 			<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js"></script>
		// 			<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js"></script>
		// 			<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js"></script>
		// 			<script>
		// 				var tiempo=0
		// 			</script>
		// 			<title>{{.Titulo}}</title>
		// 		</head>
		// 		<body>
		// 			<div class="container-fluid">
		// 				<div class="row">
		// 					<div class="col-12">
		// 					<h1 class="text-center pb-4">{{.Cuerpo}}</h1>
		// 					<progress class="w-100 d-none" id="c" value="0" max="30"></progress>

		// 					<div>Actualizado en <span id="txt-seg">30</span> Segundos</div>
		// 					<div class="progress mb-2">
		// 						<div id="d" class="progress-bar bg-success" role="progressbar" style="width: 100%" ></div>
		// 					</div>
		// 					</div>
		// 				</div>
		// 				<div class="row">
		// 					<div class="col-12">
		// 						<div id="accordion" role="tablist">

		// 			{{range .Data}}
		// 							<div class="card">
		// 							<a
		// 								data-toggle="collapse"
		// 								href="#{{.ServiceID}}"
		// 								aria-expanded="true"
		// 								aria-controls="{{.ServiceID}}"
		// 							>
		// 								<div class="card-header py-2" role="tab" id="head-{{.ServiceID}}">
		// 									<h6 class="mb-0 font-weight-lighter">
		// 											{{.Service}}
		// 											<span class="float-right">
		// 												>
		// 											</span>
		// 									</h6>
		// 								</div>
		// 							</a>

		// 								<div id="{{.ServiceID}}" class="collapse" role="tabpanel" aria-labelledby="headingOne">
		// 									<div class="card-body py-2">{{.Otp}}</div>
		// 								</div>
		// 							</div>
		// 							<script>
		// 								tiempo = {{.TiempoRestante}}
		// 							</script>
		// 			{{end}}
		// 						</div>
		// 					</div>
		// 				</div>
		// 			</div>
		// 		</body>
		// 		<style>
		// 			body {
		// 				background-color: white;
		// 				max-width: 350px;
		// 				width: 100%;
		// 				margin: 0 auto;
		// 				box-shadow: 2px 2px 5px black;
		// 				padding-bottom:15px;
		// 			}
		// 			html {
		// 				background-color: grey;
		// 			}
		// 			a,a:hover{
		// 				text-decoration:none;
		// 				color:black;
		// 			}
		// 			#txt-seg{
		// 				font-weight:900;
		// 				color:darkgrey;
		// 			}
		// 		</style>
		// 		<script>
		// 				var c = document.getElementById("c")
		// 				var d = document.getElementById("d")
		// 				var t = document.getElementById("txt-seg")
		// 				c.value = tiempo
		// 				var base30
		// 				window.setInterval(function () {
		// 					c.value = c.value - 1
		// 					t.innerText = c.value
		// 					if (c.value < 5){
		// 						d.classList.remove('bg-success')
		// 						d.classList.remove('bg-warning')
		// 						d.classList.add('bg-danger')
		// 					} else if (c.value < 10){
		// 						d.classList.add('bg-warning')
		// 						d.classList.remove('bg-danger')
		// 						d.classList.remove('bg-success')
		// 					} else {
		// 						d.classList.add('bg-success')
		// 						d.classList.remove('bg-danger')
		// 						d.classList.remove('bg-warning')
		// 					}
		// 					base30 = c.value * 3.3333
		// 					newWidth = base30+'%'
		// 					d.style.width = newWidth
		// 					if (c.value == 0) {window.location.reload()}
		// 				}, 1000);
		// 		</script>
		// 	</html>
		// 	`
		// tmpl, _ := template.New("OTPBox").Parse(plantilla)
		// tmpl.ExecuteTemplate(w, "OTPBox", dataCuadro)
	}
}

func UpdatePasswdw(db *gorm.DB, security models.Security) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Printf("%s %s\n", "Cambiando contraseña por", vars["newPass"])
		UpdatePasswd(db, security, vars["newPass"])
		fmt.Fprintln(w, "<head></head><h3>Contraseña cambiada</h3>")
	}
}

func AddServicew(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		service := vars["service"]
		secret := vars["secret"]
		totp, tiempoRestante := Decode2FA(secret)

		// var data models.TOTPData
		// data.Service = service
		// data.Secret = secret
		// AddService(db, data)

		dataTemplate := models.TemplateData{
			Service:       service,
			OTP:           totp,
			Secret:        strings.ToUpper(secret),
			TimeRemaining: tiempoRestante,
			Array: []models.TOTPData{
				{
					ID:      3,
					Service: "Prueba3",
					Secret:  "Elsecreto3",
				},
				{
					ID:      5,
					Service: "Prueba5",
					Secret:  "Secreto4",
				},
			},
		}

		plantilla := `
			<head>
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<style>
			body{
				width: 301px;
				background: white;
				height: 350px;
				box-shadow: 6px 6px 23px;
				position: absolute;
				top: 25%;
				left: 50%;
				transform: translate(-50%, -50%);
				text-align: center;
				display:grid;
				place-items:center;
			}
			html{
				background:grey
			}
			</style>
			</head>
			<body>
			<div>
			<h3>AGREGADO EL SERVICIO A BD</h3>
			<p>{{.Service}}</p>
			<p>{{.Secret}}</p>
			<p>{{.OTP}}</p> 
		{{range .Array}}
		<div>
			<span>{{.Service}}</span>
			<span>{{.Secret}}</span>
		</div>
		{{end}}

			<progress id="c" value="0" max="30"></progress>
			</div>
			</body>
			<script>
				// window.location.reload();
				var tiempo = {{.TimeRemaining}}
				var c = document.getElementById("c")
				c.value = tiempo
				window.setInterval(function () {
					c.value = c.value - 0.5
					if (c.value == 0) {window.location.reload()}
				  }, 500);
			</script>
		`
		tmpl, _ := template.New("page").Parse(plantilla)
		tmpl.ExecuteTemplate(w, "page", dataTemplate)
	}
}
