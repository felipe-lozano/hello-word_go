package main
 
import ("fmt"
       "strings")

func x_mayor_o_menor() {
	X, Y := -5, 10
	if X > 0 || Y > 0 {
		fmt.Println("ambas condiciones son verdaderas")
	} else {
		fmt.Println("al menos una condición es falsa")
	}
}

func main() {
	// x_mayor_o_menor()
	// contraseña()
	switche()
}

func contraseña() {
	var contraseña string
	// Preguntamos por la contraseña del usuario.
    fmt.Print("Por favor, ingrese su contraseña: ")
    fmt.Scanln(&contraseña)

    // Utilizamos el operador or para verificar si la contraseña cumple con los requisitos.
    if  !strings.ContainsAny(contraseña, "abcdefghijklmnopqrstuvwxyz") || !strings.ContainsAny(contraseña, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") || !strings.ContainsAny(contraseña, "0123456789") || len(contraseña) <8{
      fmt.Println("La contraseña no cumple con los requisitos.")
    } else {
      fmt.Println("La contraseña es válida.")
    }
}

func switche() {
    var operacion string
	var x int
	var y int
    fmt.Println("Elija que tipo de operacion suma? resta, multiplicacion o division")
    fmt.Scanln(&operacion)
    switch operacion {
    case "suma":
        fmt.Println("por favor ingrese el primer numero ")
		fmt.Scanln(&x)
        fmt.Println("por favor ingrese el segundo numero ")
		fmt.Scanln(&y)
		fmt.Println("El resultado de la suma es: ", x + y)
		
	case "resta":
		fmt.Println("Por favor ingrese el primer numero ")
		fmt.Scanln(&x)
		fmt.Println("por favor ingrese el segundo numero ")
		fmt.Scanln(&y)
		fmt.Println("El resultado de la resta es: ", x - y)
    case "multiplicacion":
		fmt.Println("Por favor ingrese el primer numero ")
		fmt.Scanln(&x)
		fmt.Println("por favor ingrese el segundo numero ")
		fmt.Scanln(&y)
		fmt.Println("El resultado de la multiplicacion es: ", x * y)
	case "division":
		fmt.Println("Por favor ingrese el primer numero ")
        fmt.Scanln(&x)
        fmt.Println("por favor ingrese el segundo numero ")
        fmt.Scanln(&y)
        if y!= 0 {
            fmt.Println("El resultado de la division es: ", x / y)
        } else {
            fmt.Println("No se puede dividir entre 0")
        }
	default:
        fmt.Println("El juguete es otra categoria")
    }
}

func desconectar() {
    fmt.Println("Se ha desconectado de la base de datos")
}

func leer() {
    fmt.Println("Se han leido los registros de la base de datos")
}

func actualizar() {
    fmt.Println("Se han actalizado registros de la base de datos")
}

func conectar() {
    fmt.Println("Se ha conectado a la base de datos")
}

func main() {
    conectar()
    defer desconectar()
    leer()
    actualizar()
}