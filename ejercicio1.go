package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func imprimir() {
	fmt.Println("texto desde la funcion imprimir")
}

func hola_string(s string) string {
	return "hola" + s
}

func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	//b := true
	a = false
	if a == true {
		fmt.Println("a y b son iguales")
	} else {
		fmt.Println("a y b no son iguales")
	}
}

// func main() {
//fmt.Println("texto desde la funcion main")
// imprimir()
// fmt.Println(hola_string(" felipe"))
// fmt.Println(sumar(3, 5))
// comparar_bool()

func arreglo() {
	var aprendices [5]string
	aprendices[0] = "felipe"
	aprendices[1] = "juan"
	aprendices[2] = "frank"
	aprendices[3] = "jose"
	aprendices[4] = "jaider"

	fmt.Println(aprendices[1])

}

// func main() {
	// arreglo()
	// tipo_datos()
	//convertir_string_to_bool()
	// convertir_bool_to_string()


func tipo_datos() {
	var texto string = "felipe"
	var numero int = 1000
	var decimal float64 = 0.00001
	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))

}

func convertir_string_to_bool() {
	var palabra string = "true"

	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Println("este es el error", err)
}

func convertir_bool_to_string() {
	var palabra_bool bool = true

	strbool := strconv.FormatBool(palabra_bool)
	fmt.Println(strbool, reflect.TypeOf(strbool))
}


func variables1() {
    var (
        nombre     string = "felipe"
        edad       int    = 19
        pensionado bool   = false
    )
    fmt.Println("Nombre: ", nombre)
    fmt.Println("Edad: ", edad)
    fmt.Println("Pensionado: ", pensionado)
}

func main() {
	// variables1()
	// valor0()
	// varianle_go()	
	// variable_go2()
	// ejercicio18()
	// equivalenciaEnPies()
	fmt.Println("La altura es:", altura, "mts")
  fmt.Println("La altura es:", equivalenciaEnPies(altura), " pies")
  fmt.Println("La nueva altura es:", altura, "mts")
}

func valor0() {
    var (
	 nombre2 string 
     edad2 int
     peso2 float64
     estudiante2 bool
	)
    fmt.Println("Nombre: ", nombre2)
    fmt.Println("Edad: ", edad2)
    fmt.Println("Peso: ", peso2)
    fmt.Println("Estudiante: ", estudiante2)
}

func varianle_go() {
    nombre3 := "Benjamin Button"
    edad3 := 19
    peso3 := 80
    estudiante3 := false
    fmt.Println("Nombre: ", nombre3)
    fmt.Println("Edad: ", edad3)
    fmt.Println("Peso: ", peso3)
    fmt.Println("Estudiante: ", estudiante3)
}

var profesion = "Deportista"

func variable_go2() {
    sueldo := 1000000
    fmt.Println("Profesion: ", profesion)
    fmt.Println("Sueldo: ", sueldo)
}


var var1 = "Este es el nivel 1"

func ejercicio18() {
    var var2 = "Este es el nivel 2"
    {
        var var3 = "Este es el nivel 3"
        fmt.Println(var3)
    }
    fmt.Println(var1)
    fmt.Println(var2)
}

func equivalenciaEnPies(altura float32) float32 {
    altura = altura * 3.28
    return altura
}

var altura float32 = 1.70

