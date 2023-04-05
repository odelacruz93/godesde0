package main

import (
	"fmt"

	ejercicios "github.com/odelacruz93/godesde0/Ejercicios"
)

func main() {
	// estado, texto := variables.ConvierteTexto(1993)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// os := runtime.GOOS
	// if os == "Linux." || os == "OS X." {
	// 	fmt.Println("El sistema Operativo es: ", os)
	// } else {
	// 	fmt.Println("El sistema Operativo no es linux sino es: ", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "Darwin":
	// 	fmt.Println("Esto es Darwin")
	// case "windows":
	// 	fmt.Println("Esto es windows")
	// default:
	// 	fmt.Printf("%s \n", os)
	//}

	numero, texto := ejercicios.ConvertirNumerico("6500")
	fmt.Println(numero)
	fmt.Println(texto)

}
