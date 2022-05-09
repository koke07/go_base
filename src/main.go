package main

import "fmt"

func main() {
	// linea de constante declaracion
	const pi float64 = 3.1416
	const nombre = "Jorge Luis"
	fmt.Println("Nombre:", nombre)
	fmt.Println("pi:", pi)

	// declaracion de variable enteras

	/*base := 22
	var altura int = 182
	var peso int

	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("peso:", peso)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)*/

	// ejercicio el cuadrado de un numero

	/*const baseCuadrado = 50
	area := baseCuadrado * baseCuadrado
	fmt.Println(" el area del cuadro es:", area)*/

	/*const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)*/

	// operators in goland

	x := 10
	y := 20
	//suma
	suma := x + y
	fmt.Println(suma)
	//resta
	resta := x - y
	fmt.Println(resta)
	//multiplicacion
	mult := x * y
	fmt.Println(mult)
	// division
	div := x / y
	fmt.Println(div)
	// modulo
	mod := x % y
	fmt.Println(mod)
	//incrementar
	x++
	fmt.Println(x)
	//decrementar
	x--
	fmt.Println(x)

	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)

	// paquete FMT
	// print ln para hacer salto de linea
	fmt.Println("Hola Mundo")
	// print f para imprimir con variables incluidas
	// %v si no se que tipo de dato es , %d para enteros , %f para flotantes , %s para strings (cadenas de texto)
	// la sugerencia es que si se sabe el tipo de dato , se agregue
	cursos := 35
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// Sprintf: No imprime nada en consola, simplemente lo guarda como un String. Ejemplo de uso:
	var message string = fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	//si quiero saber el tipo de variable de un tipo de dato es asi
	fmt.Printf("%T\n", nombre)

}
