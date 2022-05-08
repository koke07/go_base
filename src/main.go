package main

import "fmt"

func main() {
	// linea de constante declaracion
	const pi float64 = 3.1416
	const nombre = "Jorge Luis"
	fmt.Println("Nombre:", nombre)
	fmt.Println("pi:", pi)

	// declaracion de variable enteras

	base := 22
	var altura int = 182
	var peso int

	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("peso:", peso)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// ejercicio el cuadrado de un numero

	const baseCuadrado = 50
	area := baseCuadrado * baseCuadrado
	fmt.Println(" el area del cuadro es:", area)

	/*const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)*/
}
