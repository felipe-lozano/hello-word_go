package main

import "fmt"



func categoria() {
    var juguete string
    fmt.Println("Elige persona, animal o cosa:")
    fmt.Scanln(&juguete)
    if juguete == "persona" {
        fmt.Println("El objeto es una persona")
    } else if juguete == "cosa" {
        fmt.Println("El objeto es una cosa")
    } else if juguete == "animal" {
        fmt.Println("El objeto es un animal")
    } else {
        fmt.Println("El objeto es otra categoria")
    }

}


// func hayar_X_y_Y() {
    // var X float32
    // var Y float32
    // fmt.Println("Ingresa el valor de X:")
    // fmt.Scanln(&X)
    // fmt.Println("Ingresa el valor de Y:")
    // fmt.Scanln(&Y)
    // fmt.Println("la division entre  X y Y es: ", X/Y )
    // fmt.Println("El residuo de la division entre X y Y es: ", X%Y )
// }

func numero_par_o_impar(){
    var numero int
    fmt.Println("Ingresa un numero:")
    fmt.Scanln(&numero)
    if numero % 2 == 0 {
        fmt.Println("El numero es par")
        } else {
            fmt.Println("El numero es impar")
        }
    }


func x_mayor_o_menor () {
    X, Y := -5, 10

    if X > 0 || Y > 0 {
        fmt.Println("ambas condiciones son verdaderas")
    }   else {
        fmt.Println("al menos una condición es falsa")
    }
}

func main() {
    x_mayor_o_menor()
}
