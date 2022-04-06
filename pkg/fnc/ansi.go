package fnc

import "fmt"

// Instrucciones de código ANSI para colores y
// otras funciones en pantalla
const (
	LimpiaPantalla = "\033[2J"
	LimpiaLinea    = "\033[2K\r" //2K borra la linea y \r devuelve al comienzo
	OcultaCursor   = "\033[?25l"
	MuestraCursor  = "\033[?25h"

	Reset = "\033[0m"
	Bold  = "\033[1m"

	Black = "\033[30m"
	Gray  = "\033[37m"
	White = "\033[97m"

	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"

	BrightRed    = "\033[91m"
	BrightGreen  = "\033[92m"
	BrightYellow = "\033[93m"
	BrightBlue   = "\033[94m"
	BrightPurple = "\033[95m"
	BrightCyan   = "\033[96m"

	BgReset   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

func GotoScreen(x int, y int) string {
	return fmt.Sprintf("\033[%d;%dH", x, y)
}
